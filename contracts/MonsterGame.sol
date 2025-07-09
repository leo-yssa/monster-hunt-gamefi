// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MyGameToken.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

contract MonsterGame is Initializable, OwnableUpgradeable, PausableUpgradeable, ReentrancyGuardUpgradeable {
    MyGameToken public rewardToken;

    struct Player {
        string name;
        uint256 level;
        bool registered;
    }
    mapping(address => Player) public players;

    struct Monster {
        string name;
        uint256 hp;
        uint256 reward;
    }
    Monster[] public monsters;

    event PlayerRegistered(address indexed player, string name);
    event MonsterAdded(uint indexed monsterId, string name, uint256 hp, uint256 reward);
    event MonsterHunted(address indexed player, uint indexed monsterId, uint256 reward);

    function initialize(address tokenAddress) public initializer {
        __Ownable_init(msg.sender);
        __Pausable_init();
        __ReentrancyGuard_init();
        rewardToken = MyGameToken(tokenAddress);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function registerPlayer(string memory name) public whenNotPaused {
        require(!players[msg.sender].registered, "Already registered");
        require(bytes(name).length > 0 && bytes(name).length <= 32, "Invalid name");
        players[msg.sender] = Player({name: name, level: 1, registered: true});
        emit PlayerRegistered(msg.sender, name);
    }

    function addMonster(string memory name, uint256 hp, uint256 reward) public onlyOwner whenNotPaused {
        require(bytes(name).length > 0 && bytes(name).length <= 32, "Invalid monster name");
        require(hp > 0 && hp <= 10000, "Invalid HP");
        require(reward > 0 && reward <= 1_000_000 ether, "Invalid reward");
        monsters.push(Monster({name: name, hp: hp, reward: reward}));
        emit MonsterAdded(monsters.length - 1, name, hp, reward);
    }

    function huntMonster(uint monsterId) public whenNotPaused nonReentrant {
        require(players[msg.sender].registered, "Player not registered");
        require(monsterId < monsters.length, "Invalid monster");
        Monster memory m = monsters[monsterId];
        // 실제 게임 로직에서는 HP/전투 등 추가 가능
        rewardToken.mint(msg.sender, m.reward);
        emit MonsterHunted(msg.sender, monsterId, m.reward);
    }
}
