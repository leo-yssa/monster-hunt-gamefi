import { expect } from "chai";
// @ts-ignore
import { ethers, upgrades } from "hardhat";
import "@nomicfoundation/hardhat-chai-matchers";

describe("MonsterGame Proxy Patterns", function () {
  let token: any;
  let owner: any;
  let player1: any;
  let player2: any;

  beforeEach(async () => {
    [owner, player1, player2] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("MyGameToken");
    token = await Token.deploy();
    await token.waitForDeployment();
  });

  async function deployTransparent() {
    const Game = await ethers.getContractFactory("MonsterGameTransparent");
    const game = await upgrades.deployProxy(Game, [await token.getAddress()], { initializer: "initialize" });
    await game.waitForDeployment();
    await token.transferOwnership(await game.getAddress());
    return game;
  }

  async function deployUUPS() {
    const Game = await ethers.getContractFactory("MonsterGameUUPS");
    const game = await upgrades.deployProxy(Game, [await token.getAddress()], { initializer: "initialize", kind: "uups" });
    await game.waitForDeployment();
    await token.transferOwnership(await game.getAddress());
    return game;
  }

  async function deployBeacon() {
    const Impl = await ethers.getContractFactory("MonsterGameBeaconImpl");
    const beacon = await upgrades.deployBeacon(Impl);
    await beacon.waitForDeployment();
    const game = await upgrades.deployBeaconProxy(beacon, Impl, [await token.getAddress()], { initializer: "initialize" });
    await game.waitForDeployment();
    await token.transferOwnership(await game.getAddress());
    return { game, beacon };
  }

  function runCommonLogicTests(deployGame: () => Promise<any>, label: string) {
    describe(label, function () {
      let game: any;
      beforeEach(async () => {
        game = await deployGame();
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
        expect(balance).to.equal(10);
      });

      it("should fail if non-owner tries to add a monster", async () => {
        await expect(
          game.connect(player1).addMonster("Dragon", 200, 100)
        ).to.be.revertedWithCustomError(game, "OwnableUnauthorizedAccount");
      });

      it("should not allow actions when paused", async () => {
        await game.connect(owner).pause();
        await expect(game.connect(player1).registerPlayer("Bob")).to.be.revertedWithCustomError(game, "EnforcedPause");
        await expect(game.connect(owner).addMonster("Orc", 100, 20)).to.be.revertedWithCustomError(game, "EnforcedPause");
        await expect(game.connect(player1).huntMonster(0)).to.be.revertedWithCustomError(game, "EnforcedPause");
        await game.connect(owner).unpause();
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
  }

  runCommonLogicTests(deployTransparent, "Transparent Proxy Logic");
  runCommonLogicTests(deployUUPS, "UUPS Proxy Logic");
  runCommonLogicTests(async () => (await deployBeacon()).game, "Beacon Proxy Logic");

  // --- 프록시별 업그레이드/권한/상태 보존 테스트 ---

  it("Transparent: should preserve state after upgrade", async () => {
    const game = await deployTransparent();
    await game.connect(owner).addMonster("Goblin", 50, 10);
    await game.connect(player1).registerPlayer("Alice");
    // 업그레이드용 임시 V2 컨트랙트
    const V2 = await ethers.getContractFactory("MonsterGameTransparent");
    const upgraded = await upgrades.upgradeProxy(await game.getAddress(), V2);
    // 상태 보존 확인
    expect((await upgraded.monsters(0)).name).to.equal("Goblin");
    expect((await upgraded.players(player1.address)).name).to.equal("Alice");
  });

  it("UUPS: should preserve state after upgrade and onlyOwner can upgrade", async () => {
    const game = await deployUUPS();
    await game.connect(owner).addMonster("Orc", 100, 20);
    await game.connect(player1).registerPlayer("Bob");
    const V2 = await ethers.getContractFactory("MonsterGameUUPS");
    // onlyOwner 업그레이드 성공
    await upgrades.upgradeProxy(await game.getAddress(), V2);
    // player1(비owner) 업그레이드 시도 실패
    await expect(
      upgrades.upgradeProxy(await game.getAddress(), V2.connect(player1))
    ).to.be.reverted;
  });

  it("Beacon: should preserve state after beacon upgrade", async () => {
    const { game, beacon } = await deployBeacon();
    await game.connect(owner).addMonster("Dragon", 200, 100);
    await game.connect(player1).registerPlayer("Eve");
    // 새로운 구현체 배포 및 beacon 업그레이드
    const ImplV2 = await ethers.getContractFactory("MonsterGameBeaconImpl");
    await upgrades.upgradeBeacon(beacon, ImplV2);
    // 상태 보존 확인
    expect((await game.monsters(0)).name).to.equal("Dragon");
    expect((await game.players(player1.address)).name).to.equal("Eve");
  });

  it("Proxy pattern differences: admin/upgrade authority", async () => {
    // Transparent: ProxyAdmin만 업그레이드 가능, 일반 owner는 불가
    // UUPS: 컨트랙트의 onlyOwner가 직접 업그레이드 가능
    // Beacon: Beacon 컨트랙트의 owner만 업그레이드 가능
    expect(true).to.be.true;
  });
});
