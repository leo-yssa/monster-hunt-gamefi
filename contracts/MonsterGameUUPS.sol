// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MonsterGame.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract MonsterGameUUPS is MonsterGame, UUPSUpgradeable {
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    function initialize(address tokenAddress) public override initializer {
        MonsterGame.initialize(tokenAddress);
        __UUPSUpgradeable_init();
    }
} 