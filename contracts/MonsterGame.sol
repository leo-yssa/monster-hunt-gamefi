// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./MyGameToken.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MonsterGame is Ownable {
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

    constructor(address tokenAddress) Ownable(msg.sender) {
        rewardToken = MyGameToken(tokenAddress);
    }

    function registerPlayer(string memory name) public {
        require(!players[msg.sender].registered, "Already registered");
        players[msg.sender] = Player({name: name, level: 1, registered: true});
        emit PlayerRegistered(msg.sender, name);
    }

    function addMonster(string memory name, uint256 hp, uint256 reward) public onlyOwner {
        monsters.push(Monster({name: name, hp: hp, reward: reward}));
        emit MonsterAdded(monsters.length - 1, name, hp, reward);
    }

    function huntMonster(uint monsterId) public {
        require(players[msg.sender].registered, "Player not registered");
        require(monsterId < monsters.length, "Invalid monster");
        Monster memory m = monsters[monsterId];
        // 실제 게임 로직에서는 HP/전투 등 추가 가능
        rewardToken.mint(msg.sender, m.reward);
        emit MonsterHunted(msg.sender, monsterId, m.reward);
    }
}
