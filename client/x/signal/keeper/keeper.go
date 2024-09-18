package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/core/store"
	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/piplabs/story/client/x/signal/types"
	"github.com/piplabs/story/lib/log"
)

// DefaultUpgradeHeightDelay is the number of blocks after a quorum has been
// reached that the chain should upgrade to the new version. Assuming a block
// interval of 12 seconds, this is 7 days.
const DefaultUpgradeHeightDelay = int64(7 * 24 * 60 * 60 / 12) // 7 days * 24 hours * 60 minutes * 60 seconds / 12 seconds per block = 50,400 blocks.

// Keeper implements the MsgServer and QueryServer interfaces.
var (
	_ types.MsgServer   = &Keeper{}
	_ types.QueryServer = Keeper{}

	// defaultSignalThreshold is 5/6 or approximately 83.33%.
	defaultSignalThreshold = sdkmath.LegacyNewDec(5).Quo(sdkmath.LegacyNewDec(6))
)

// Threshold is the fraction of voting power that is required
// to signal for a version change. It is set to 5/6 as the middle point
// between 2/3 and 3/3 providing 1/6 fault tolerance to halting the
// network during an upgrade period. It can be modified through a
// hard fork change that modified the app version.
func Threshold(_ uint64) sdkmath.LegacyDec {
	return defaultSignalThreshold
}

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	authority    string

	authKeeper    types.AccountKeeper
	stakingKeeper types.StakingKeeper
}

// NewKeeper returns a signal keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	ak types.AccountKeeper,
	stk types.StakingKeeper,
	authority string,
) *Keeper {
	// ensure that authority is a valid AccAddress
	if _, err := ak.AddressCodec().StringToBytes(authority); err != nil {
		panic("authority is not a valid acc address")
	}

	// ensure the module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(types.ModuleName + " module account has not been set")
	}

	return &Keeper{
		cdc:           cdc,
		storeService:  storeService,
		authKeeper:    ak,
		stakingKeeper: stk,
	}
}

// SignalVersion is a method required by the MsgServer interface.
func (k Keeper) SignalVersion(ctx context.Context, req *types.MsgSignalVersion) (*types.MsgSignalVersionResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if k.IsUpgradePending(sdkCtx) {
		return &types.MsgSignalVersionResponse{}, types.ErrUpgradePending.Wrapf("can not signal version")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	// The signaled version can not be less than the current version.
	currentVersion := sdkCtx.BlockHeader().Version.App
	if req.Version < currentVersion {
		return nil, types.ErrInvalidSignalVersion.Wrapf("signaled version %d, current version %d", req.Version, currentVersion)
	}

	if _, err := k.stakingKeeper.GetValidator(sdkCtx, valAddr); err != nil {
		return nil, stakingtypes.ErrNoValidatorFound
	}

	k.SetValidatorVersion(sdkCtx, valAddr, req.Version)
	return &types.MsgSignalVersionResponse{}, nil
}

// TryUpgrade is a method required by the MsgServer interface. It tallies the
// voting power that has voted on each version. If one version has reached a
// quorum, an upgrade is persisted to the store. The upgrade is used by the
// application later when it is time to upgrade to that version.
func (k *Keeper) TryUpgrade(ctx context.Context, _ *types.MsgTryUpgrade) (*types.MsgTryUpgradeResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if k.IsUpgradePending(sdkCtx) {
		return &types.MsgTryUpgradeResponse{}, types.ErrUpgradePending.Wrapf("can not try upgrade")
	}

	threshold := k.GetVotingPowerThreshold(sdkCtx)
	hasQuorum, version := k.TallyVotingPower(sdkCtx, threshold.Int64())
	if hasQuorum {
		if version <= sdkCtx.BlockHeader().Version.App {
			return &types.MsgTryUpgradeResponse{}, types.ErrInvalidUpgradeVersion.Wrapf("can not upgrade to version %v because it is less than or equal to current version %v", version, sdkCtx.BlockHeader().Version.App)
		}
		upgrade := types.Upgrade{
			AppVersion:    version,
			UpgradeHeight: sdkCtx.BlockHeader().Height + DefaultUpgradeHeightDelay,
		}
		k.setUpgrade(sdkCtx, upgrade)
	}
	return &types.MsgTryUpgradeResponse{}, nil
}

// VersionTally enables a client to query for the tally of voting power has
// signaled for a particular version.
func (k Keeper) VersionTally(ctx context.Context, req *types.QueryVersionTallyRequest) (*types.QueryVersionTallyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	totalVotingPower, err := k.stakingKeeper.GetLastTotalPower(ctx)
	if err != nil {
		return &types.QueryVersionTallyResponse{}, err
	}
	currentVotingPower := sdkmath.NewInt(0)

	stores := k.storeService.OpenKVStore(ctx)
	iterator, err := stores.Iterator(types.FirstSignalKey, nil)
	if err != nil {
		return &types.QueryVersionTallyResponse{}, err
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		valAddress := sdk.ValAddress(iterator.Key())
		power, err := k.stakingKeeper.GetLastValidatorPower(sdkCtx, valAddress)
		if err != nil {
			log.Error(ctx, "Failed to get last validator power", err)
			continue
		}
		version := VersionFromBytes(iterator.Value())
		if version == req.Version {
			currentVotingPower = currentVotingPower.AddRaw(power)
		}
	}

	threshold := k.GetVotingPowerThreshold(sdkCtx)

	return &types.QueryVersionTallyResponse{
		VotingPower:      currentVotingPower.Uint64(),
		ThresholdPower:   threshold.Uint64(),
		TotalVotingPower: totalVotingPower.Uint64(),
	}, nil
}

// SetValidatorVersion saves a signaled version for a validator.
func (k Keeper) SetValidatorVersion(ctx sdk.Context, valAddress sdk.ValAddress, version uint64) {
	stores := k.storeService.OpenKVStore(ctx)
	if err := stores.Set(valAddress, VersionToBytes(version)); err != nil {
		panic(err)
	}
}

// DeleteValidatorVersion deletes a signaled version for a validator.
func (k Keeper) DeleteValidatorVersion(ctx sdk.Context, valAddress sdk.ValAddress) {
	stores := k.storeService.OpenKVStore(ctx)
	if err := stores.Delete(valAddress); err != nil {
		panic(err)
	}
}

// TallyVotingPower tallies the voting power for each version and returns true
// and the version if any version has reached the quorum in voting power.
// Returns false and 0 otherwise.
func (k Keeper) TallyVotingPower(ctx sdk.Context, threshold int64) (bool, uint64) {
	stores := k.storeService.OpenKVStore(ctx)
	iterator, err := stores.Iterator(types.FirstSignalKey, nil)
	if err != nil {
		return false, 0
	}
	defer iterator.Close()

	versionToPower := make(map[uint64]int64)
	for ; iterator.Valid(); iterator.Next() {
		valAddress := sdk.ValAddress(iterator.Key())
		// check that the validator is still part of the bonded set
		val, err := k.stakingKeeper.GetValidator(ctx, valAddress)
		if err != nil {
			// if it no longer exists, delete the version
			k.DeleteValidatorVersion(ctx, valAddress)
		}
		// if the validator is not bonded, skip it's voting power
		if err != nil || !val.IsBonded() {
			continue
		}
		power, err := k.stakingKeeper.GetLastValidatorPower(ctx, valAddress)
		if err != nil {
			log.Error(ctx, "Failed to get last validator power", err)
			continue
		}
		version := VersionFromBytes(iterator.Value())
		if _, ok := versionToPower[version]; !ok {
			versionToPower[version] = power
		} else {
			versionToPower[version] += power
		}
		if versionToPower[version] >= threshold {
			return true, version
		}
	}

	return false, 0
}

// GetVotingPowerThreshold returns the voting power threshold required to
// upgrade to a new version.
func (k Keeper) GetVotingPowerThreshold(ctx sdk.Context) sdkmath.Int {
	totalVotingPower, err := k.stakingKeeper.GetLastTotalPower(ctx)
	if err != nil {
		panic(err)
	}
	thresholdFraction := Threshold(ctx.BlockHeader().Version.App)
	return thresholdFraction.MulInt(totalVotingPower).Ceil().TruncateInt()
}

// ShouldUpgrade returns whether the signaling mechanism has concluded that the
// network is ready to upgrade and the version to upgrade to. It returns false
// and 0 if no version has reached quorum.
func (k *Keeper) ShouldUpgrade(ctx sdk.Context) (isQuorumVersion bool, version uint64) {
	upgrade, ok := k.getUpgrade(ctx)
	if !ok {
		return false, 0
	}

	hasUpgradeHeightBeenReached := ctx.BlockHeight() >= upgrade.UpgradeHeight
	if hasUpgradeHeightBeenReached {
		return true, upgrade.AppVersion
	}
	return false, 0
}

// ResetTally resets the tally after a version change. It iterates over the
// store and deletes all versions. It also resets the quorumVersion and
// upgradeHeight to 0.
func (k *Keeper) ResetTally(ctx sdk.Context) {
	stores := k.storeService.OpenKVStore(ctx)
	iterator, err := stores.Iterator(nil, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	// delete the value in the upgrade key and all signals.
	for ; iterator.Valid(); iterator.Next() {
		if err := stores.Delete(iterator.Key()); err != nil {
			panic(err)
		}
	}
}

func VersionToBytes(version uint64) []byte {
	return binary.BigEndian.AppendUint64(nil, version)
}

func VersionFromBytes(version []byte) uint64 {
	return binary.BigEndian.Uint64(version)
}

// GetUpgrade returns the current upgrade information.
func (k Keeper) GetUpgrade(ctx context.Context, _ *types.QueryGetUpgradeRequest) (*types.QueryGetUpgradeResponse, error) {
	upgrade, ok := k.getUpgrade(sdk.UnwrapSDKContext(ctx))
	if !ok {
		return &types.QueryGetUpgradeResponse{}, nil
	}
	return &types.QueryGetUpgradeResponse{Upgrade: &upgrade}, nil
}

// IsUpgradePending returns true if an app version has reached quorum and the
// chain should upgrade to the app version at the upgrade height. While the
// keeper has an upgrade pending the SignalVersion and TryUpgrade messages will
// be rejected.
func (k *Keeper) IsUpgradePending(ctx sdk.Context) bool {
	_, ok := k.getUpgrade(ctx)
	return ok
}

// getUpgrade returns the current upgrade information from the store.
// If an upgrade is found, it returns the upgrade object and true.
// If no upgrade is found, it returns an empty upgrade object and false.
func (k *Keeper) getUpgrade(ctx sdk.Context) (upgrade types.Upgrade, ok bool) {
	stores := k.storeService.OpenKVStore(ctx)
	bz, err := stores.Get(types.UpgradeKey)
	if err != nil {
		return types.Upgrade{}, false
	}

	err = k.cdc.Unmarshal(bz, &upgrade)
	if err != nil {
		return upgrade, false
	}

	return upgrade, true
}

// setUpgrade sets the upgrade in the store.
func (k *Keeper) setUpgrade(ctx sdk.Context, upgrade types.Upgrade) {
	stores := k.storeService.OpenKVStore(ctx)
	value := k.cdc.MustMarshal(&upgrade)

	if err := stores.Set(types.UpgradeKey, value); err != nil {
		panic(err)
	}
}
