// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./CurveLPToken.sol";

// IVotes 인터페이스 호환 함수 추가
interface IVotes {
    function getVotes(address account, uint256 timepoint) external view returns (uint256);
    function getPastVotes(address account, uint256 timepoint) external view returns (uint256);
    function getPastTotalSupply(uint256 timepoint) external view returns (uint256);
    function CLOCK_MODE() external view returns (string memory);
}

contract CurveLPStaking is IVotes {
    CurveLPToken public lpToken;
    mapping(address => uint256) public staked;
    uint256 public totalStaked;

    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);

    constructor(address _lpToken) {
        lpToken = CurveLPToken(_lpToken);
    }

    function stake(uint256 amount) external {
        require(amount > 0, "Cannot stake 0");
        lpToken.transferFrom(msg.sender, address(this), amount);
        staked[msg.sender] += amount;
        totalStaked += amount;
        emit Staked(msg.sender, amount);
    }

    function unstake(uint256 amount) external {
        require(staked[msg.sender] >= amount, "Not enough staked");
        staked[msg.sender] -= amount;
        totalStaked -= amount;
        lpToken.transfer(msg.sender, amount);
        emit Unstaked(msg.sender, amount);
    }

    function votingPower(address user) external view returns (uint256) {
        return staked[user];
    }

    // IVotes 인터페이스 구현 (최신 블록만 지원, 단순 버전)
    function getVotes(address account, uint256 /*timepoint*/) external view override returns (uint256) {
        return staked[account];
    }
    function getPastVotes(address account, uint256 /*timepoint*/) external view override returns (uint256) {
        return staked[account];
    }
    function getPastTotalSupply(uint256 /*timepoint*/) external view override returns (uint256) {
        return totalStaked;
    }
    function CLOCK_MODE() external pure override returns (string memory) {
        return "mode=blocknumber&from=default";
    }
} 