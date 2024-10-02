// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import { IPTokenSlashing } from "../src/protocol/IPTokenSlashing.sol";

// forge script script/CallUnjailScript.s.sol --fork-url http://54.215.121.164:8545 --broadcast
contract CallUnjailScript is Script {
    function run() external {
        //   IPTokenSlashing will be deployed at: 0xEEf1c4fD443965404f13BE2705766988317b3B32
        //   IP_TOKEN_STAKING 0xCCcCcC0000000000000000000000000000000001

        // address ipTokenSlashingAddr = 0xCccCCC0000000000000000000000000000000002;
        // address ipTokenSlashingAddr = 0xEEf1c4fD443965404f13BE2705766988317b3B32;
        address ipTokenSlashingAddr = 0x6c50BCba511668a6541dCc3A242d25B20698613b;
        IPTokenSlashing ipTokenSlashing = IPTokenSlashing(ipTokenSlashingAddr);
        console2.log("ipTokenSlashingAddr address:", ipTokenSlashingAddr);


        bytes memory validatorUncmpPubkey = hex"04721186a1721db08e2492b60b1b678cd5a7971c19716cce508fb29d18f9b434c76f7f0ab6d8bbd1a5f31e7ece5c33839b4519be3c1d772a16e84cb92c0dbaf320";

        uint256 unjailFee = 1 ether;

        // export PRIVATE_KEY=0x748af76e64066cf4994c434b55cad8a5972e04833b91754358eb6b441b747706
        address currentSender = vm.addr(vm.envUint("PRIVATE_KEY"));
        console2.log("Current sender address:", currentSender);
    
        uint256 privateKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(privateKey);
        ipTokenSlashing.unjail{value: unjailFee}(validatorUncmpPubkey);
        // ipTokenSlashing.unjail(validatorUncmpPubkey);
        vm.stopBroadcast();
    }
}