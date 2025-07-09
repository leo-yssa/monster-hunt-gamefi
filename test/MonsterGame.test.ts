import { expect } from "chai";
// @ts-ignore
import { ethers, upgrades } from "hardhat";
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

    // 게임 컨트랙트 프록시 배포
    const Game = await ethers.getContractFactory("MonsterGame");
    game = await upgrades.deployProxy(Game, [await token.getAddress()], { initializer: "initialize" });
    await game.waitForDeployment();

    // 토큰 소유권 게임 컨트랙트(프록시)로 이전
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

  it("should upgrade to V2 and preserve state", async () => {
    // 1. 기존 상태 저장
    await game.connect(owner).addMonster("Goblin", 50, 10);
    await game.connect(player1).registerPlayer("Alice");

    // 2. V2로 업그레이드
    const MonsterGameV2 = await ethers.getContractFactory("MonsterGameV2");
    const gameV2 = await upgrades.upgradeProxy(await game.getAddress(), MonsterGameV2);

    // 3. 기존 상태가 유지되는지 확인
    const monster = await gameV2.monsters(0);
    expect(monster.name).to.equal("Goblin");
    const playerData = await gameV2.players(player1.address);
    expect(playerData.name).to.equal("Alice");

    // 4. V2 기능 확인
    expect(await gameV2.version()).to.equal("V2");
  });

  it("should not allow actions when paused", async () => {
    await game.connect(owner).pause();
    await expect(game.connect(player1).registerPlayer("Bob")).to.be.revertedWithCustomError(game, "EnforcedPause");
    await expect(game.connect(owner).addMonster("Orc", 100, 20)).to.be.revertedWithCustomError(game, "EnforcedPause");
    await expect(game.connect(player1).huntMonster(0)).to.be.revertedWithCustomError(game, "EnforcedPause");
    await game.connect(owner).unpause();
    // 정상 동작 확인
    await game.connect(owner).addMonster("Orc", 100, 20);
    expect((await game.monsters(0)).name).to.equal("Orc");
  });

  it("should validate input values", async () => {
    await expect(game.connect(player1).registerPlayer("")).to.be.revertedWith("Invalid name");
    await expect(game.connect(owner).addMonster("", 100, 10)).to.be.revertedWith("Invalid monster name");
    await expect(game.connect(owner).addMonster("Orc", 0, 10)).to.be.revertedWith("Invalid HP");
    await expect(game.connect(owner).addMonster("Orc", 100, 0)).to.be.revertedWith("Invalid reward");
  });
});
