// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;
/* solhint-disable no-console */
/* solhint-disable max-line-length */

import { Script } from "forge-std/Script.sol";
import { console2 } from "forge-std/console2.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import { CREATE3 } from "solady/src/utils/CREATE3.sol";
import { IPTokenSlashing } from "../src/protocol/IPTokenSlashing.sol";

/**
 * @title DeployIPTokenSlashing
 * @dev A script to deploy IPTokenSlashing for Illiad
 */

// == Logs ==
//   deployer 0x13919a0d8603c35DAC923f92D7E4e1D55e993898
//   IPTokenSlashing will be deployed at: 0xEEf1c4fD443965404f13BE2705766988317b3B32
//   IP_TOKEN_STAKING 0xCCcCcC0000000000000000000000000000000001
//   IPTokenSlashing deployed at: 0xEEf1c4fD443965404f13BE2705766988317b3B32
contract DeployIPTokenSlashing is Script {
    // To run the script:
    // mininet v4 "http://18.144.99.223:8545

    // dev
    // const RPC_URL = "http://18.144.89.70:8545";

    // export ADMIN_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8 
    // export IPTOKENSTAKING_DEPLOYER_KEY=0x45813784c38bb79693e587efc8b754ca12e04552358855cc02f069da9c11b79c
    
    // - Dry run
    // forge script script/DeployIPTokenSlashing.s.sol --fork-url http://18.144.99.223:8545
    //
    // - Deploy (OK for devnet)
    // forge script script/DeployIPTokenSlashing.s.sol --fork-url http://54.215.121.164:8545 --broadcast
    //
    // - Deploy and Verify (for testnet)
    // forge script script/DeployIPTokenSlashing.s.sol --fork-url https://testnet.storyrpc.io --broadcast --verify --verifier blockscout --verifier-url https://testnet.storyscan.xyz/api\?
    function run() public {
        // Read env for admin address
        address protocolAccessManagerAddr = vm.envAddress("ADMIN_ADDRESS");
        require(protocolAccessManagerAddr != address(0), "address not set");
        // Read env for deployer private key
        uint256 deployerKey = vm.envUint("IPTOKENSTAKING_DEPLOYER_KEY");
        address deployer = vm.addr(deployerKey);
        console2.log("deployer", deployer);
        vm.startBroadcast(deployerKey);

        address ipTokenStaking = 0xCCcCcC0000000000000000000000000000000001;

        address impl = address(new IPTokenSlashing(ipTokenStaking));
        bytes memory initializationData = abi.encodeCall(
            IPTokenSlashing.initialize,
            (
                protocolAccessManagerAddr,
                1 ether // unjailFee
            )
        );
        bytes memory creationCode =
            abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(impl, initializationData));
        bytes32 salt = keccak256(abi.encode("STORY", type(IPTokenSlashing).name));
        address predicted = CREATE3.predictDeterministicAddress(salt);
        console2.log("IPTokenSlashing will be deployed at:", predicted);

        IPTokenSlashing ipTokenSlashing = IPTokenSlashing(this.deployDeterministic(creationCode, salt));

        console2.log("IP_TOKEN_STAKING", address(ipTokenSlashing.IP_TOKEN_STAKING()));
        if (address(ipTokenSlashing) != predicted) {
            revert("IPTokenSlashing mismatch");
        }
        console2.log("IPTokenSlashing deployed at:", address(ipTokenSlashing));

        vm.stopBroadcast();
    }

    function deployDeterministic(bytes calldata creationCode, bytes32 salt) public returns (address) {
        return CREATE3.deployDeterministic(0, creationCode, salt);
    }
}
