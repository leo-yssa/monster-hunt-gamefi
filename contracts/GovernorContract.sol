// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/governance/Governor.sol";
import "@openzeppelin/contracts/governance/extensions/GovernorSettings.sol";
import "@openzeppelin/contracts/governance/extensions/GovernorCountingSimple.sol";
import "@openzeppelin/contracts/governance/extensions/GovernorVotes.sol";

contract GovernorContract is Governor, GovernorSettings, GovernorCountingSimple, GovernorVotes {
    constructor(IVotes _votingPowerSource)
        Governor("GovernorContract")
        GovernorSettings(1 /* 1 block */, 5 /* 5 blocks */, 0)
        GovernorVotes(_votingPowerSource)
    {}

    // The following functions are overrides required by Solidity.
    function votingDelay() public view override(Governor, GovernorSettings) returns (uint256) {
        return super.votingDelay();
    }

    function votingPeriod() public view override(Governor, GovernorSettings) returns (uint256) {
        return super.votingPeriod();
    }

    function quorum(uint256 blockNumber) public pure override returns (uint256) {
        return 0; // 쿼럼 없음(기본값)
    }

    function proposalThreshold() public view override(Governor, GovernorSettings) returns (uint256) {
        return super.proposalThreshold();
    }
} 