// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MonsterGame.sol";

contract MonsterGameV2 is MonsterGame {
    function version() public pure returns (string memory) {
        return "V2";
    }
} 