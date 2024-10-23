// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.23;

import { Script } from "forge-std/Script.sol";
import { UBIPool } from "../src/protocol/UBIPool.sol";
import { console2 } from "forge-std/console2.sol";

abstract contract MockNewFeatures {
    function foo() external pure returns (string memory) {
        return "bar";
    }
}
contract UBIPoolV2 is UBIPool, MockNewFeatures {
    constructor(uint32 maxUBIPercentage) UBIPool(maxUBIPercentage) {}
}

contract ConfirmUpgrade is Script {
    
    function run() public {
        address upgradedContractAddress = 0xCccCCC0000000000000000000000000000000002;


        UBIPoolV2 upgradedContract = UBIPoolV2(upgradedContractAddress);

        string memory result = upgradedContract.foo();


        require(keccak256(abi.encodePacked(result)) == keccak256(abi.encodePacked("bar")), "Upgrade failed: foo() does not return 'bar'");
        
        console2.log("Upgrade successful: foo() returned:", result);
    }
}