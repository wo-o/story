// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import { Script } from "forge-std/Script.sol";
import { console2 } from "forge-std/console2.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { UBIPool } from "../src/protocol/UBIPool.sol";
import { ITransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { EIP1967Helper } from "./utils/EIP1967Helper.sol";
import { Predeploys } from "../src/libraries/Predeploys.sol";
import { TimelockController } from "@openzeppelin/contracts/governance/TimelockController.sol";

abstract contract MockNewFeatures {
    function foo() external pure returns (string memory) {
        return "bar";
    }
}

contract UBIPoolV2 is UBIPool, MockNewFeatures {
    constructor(uint32 maxUBIPercentage) UBIPool(maxUBIPercentage) {}
}

contract TestPrecompileUpgrades is Script {
    TimelockController internal timelock;
    address public newImpl = 0xfd7F50b24E5159e5213D4744A8F0a7D66822c94b;
    bytes32 public salt = keccak256(abi.encodePacked("UBIPoolUpgrade"));

    // Entry point for forge script
    function run() public {
        bool isExecution = vm.envBool("IS_EXECUTE");

        if (isExecution) {
            executeUpgrade();
        } else {
            scheduleUpgrade();
        }
    }

    // Schedule the upgrade
    function scheduleUpgrade() public {
        timelock = TimelockController(payable(0x4827c76bD61A223Ddd36D013c78F825eb0bb3Be3));

        // Read env for admin address
        uint256 upgradeKey = vm.envUint("UPGRADE_ADMIN_KEY");
        address upgrader = vm.addr(upgradeKey);
        console2.log("Upgrader address:", upgrader);
        vm.startBroadcast(upgradeKey);

        // Deploy UBIPoolV2 and store the address for later reuse
        newImpl = address(new UBIPoolV2(650));

        console2.log("IMPL:", newImpl);

        ProxyAdmin proxyAdmin = ProxyAdmin(EIP1967Helper.getAdmin(Predeploys.UBIPool));
        console2.log("ProxyAdmin address during scheduling:", address(proxyAdmin));

        // Generate calldata
        bytes memory data = abi.encodeWithSelector(
            proxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(Predeploys.UBIPool),
            newImpl,
            ""
        );
        console2.log("Calldata during scheduling:", vm.toString(data));

        // Log operation ID
        bytes32 operationId = keccak256(abi.encode(address(proxyAdmin), 0, data, bytes32(0), salt));
        console2.log("Operation ID during scheduling:", vm.toString(operationId));

        uint256 minDelay = timelock.getMinDelay(); // Use the timelock's minimum delay

        timelock.schedule(
            address(proxyAdmin),
            0,
            data,
            bytes32(0), // predecessor
            salt, // Fixed salt
            minDelay // Timelock's minimum delay
        );

        console2.log("UBIPool minDelay: ", minDelay);
        console2.log("UBIPool upgrade scheduled");

        vm.stopBroadcast();
    }

    // Execute the upgrade
    function executeUpgrade() public {
        timelock = TimelockController(payable(0x4827c76bD61A223Ddd36D013c78F825eb0bb3Be3));

        // Read env for executor address
        uint256 executorKey = vm.envUint("EXECUTOR_KEY");
        address executor = vm.addr(executorKey);
        console2.log("Executor address:", executor);
        vm.startBroadcast(executorKey);

        // Use the stored newImpl address from scheduling
        ProxyAdmin proxyAdmin = ProxyAdmin(EIP1967Helper.getAdmin(Predeploys.UBIPool));
        console2.log("ProxyAdmin address during execution:", address(proxyAdmin));

        bytes memory data = abi.encodeWithSelector(
            proxyAdmin.upgradeAndCall.selector,
            ITransparentUpgradeableProxy(Predeploys.UBIPool),
            newImpl,
            ""
        );
        console2.log("Calldata during execution:", vm.toString(data));

        // Log operation ID for execution
        bytes32 operationId = keccak256(abi.encode(address(proxyAdmin), 0, data, bytes32(0), salt));
        console2.log("Operation ID during execution:", vm.toString(operationId));

        // Check if the operation is ready
        bool isReady = timelock.isOperationReady(operationId);
        console2.log("Is operation ready:", isReady);

        if (!isReady) {
            revert("Operation is not ready for execution. Wait until the minDelay has passed.");
        }

        // Execute the UBIPool upgrade via TimelockController
        console2.log("Executing UBIPool upgrade via TimelockController");

        timelock.execute(
            address(proxyAdmin),
            0,
            data,
            bytes32(0), // predecessor
            salt // Fixed salt used earlier
        );

        console2.log("UBIPool upgrade executed by executor", executor);

        // === Verify the upgrade ===
        // Check if UBIPoolV2 upgrade was successful
        if (keccak256(abi.encode(UBIPoolV2(Predeploys.UBIPool).foo())) != keccak256(abi.encode("bar"))) {
            revert("Upgraded to wrong iface for UBIPoolV2");
        }

        // Check if UBIPool implementation address is correct
        if (EIP1967Helper.getImplementation(Predeploys.UBIPool) != newImpl) {
            revert("UBIPoolV2 not upgraded");
        }

        console2.log("UBIPoolV2 upgrade verified successfully!");

        vm.stopBroadcast();
    }
}