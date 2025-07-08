import { expect } from "chai";
import { ethers } from "hardhat";
import "@nomicfoundation/hardhat-chai-matchers";

describe("MonsterGame", function () {
  let token: any;
  let game: any;
  let owner: any;
  let player1: any;
  let player2: any;

  beforeEach(async () => {
    [owner, player1, player2] = await ethers.getSigners();

    // 토큰 배포
    const Token = await ethers.getContractFactory("MyGameToken");
    token = await Token.deploy();
    await token.waitForDeployment();

    // 게임 컨트랙트 배포
    const Game = await ethers.getContractFactory("MonsterGame");
    game = await Game.deploy(await token.getAddress());
    await game.waitForDeployment();

    // 토큰 소유권 게임 컨트랙트로 이전
    await token.transferOwnership(await game.getAddress());
  });

  it("should register a player and give default values", async () => {
    await game.connect(player1).registerPlayer("Alice");

    const playerData = await game.players(player1.address);
    expect(playerData.level).to.equal(1);
    expect(playerData.name).to.equal("Alice");
  });

  it("should add a monster and allow player to hunt", async () => {
    await game.connect(owner).addMonster("Goblin", 50, 10);
    const monster = await game.monsters(0);
    expect(monster.name).to.equal("Goblin");

    await game.connect(player1).registerPlayer("Alice");
    await game.connect(player1).huntMonster(0);

    const balance = await token.balanceOf(player1.address);
    expect(balance).to.equal(10); // Goblin 보상 10
  });

  it("should fail if non-owner tries to add a monster", async () => {
    await expect(
      game.connect(player1).addMonster("Dragon", 200, 100)
    ).to.be.revertedWithCustomError(game, "OwnableUnauthorizedAccount");
  });
});
