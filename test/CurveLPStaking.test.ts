import { expect } from "chai";
import { ethers } from "hardhat";

describe("CurveLPStaking", function () {
  let lpToken: any;
  let staking: any;
  let owner: any;
  let user1: any;
  let user2: any;

  beforeEach(async () => {
    [owner, user1, user2] = await ethers.getSigners();
    const LPToken = await ethers.getContractFactory("CurveLPToken");
    lpToken = await LPToken.deploy();
    await lpToken.waitForDeployment();
    await lpToken.connect(owner).mint(user1.address, 1000);
    await lpToken.connect(owner).mint(user2.address, 1000);
    const Staking = await ethers.getContractFactory("CurveLPStaking");
    staking = await Staking.deploy(await lpToken.getAddress());
    await staking.waitForDeployment();
    await lpToken.connect(user1).approve(await staking.getAddress(), 1000);
    await lpToken.connect(user2).approve(await staking.getAddress(), 1000);
  });

  it("should allow staking and update voting power", async () => {
    await staking.connect(user1).stake(500);
    expect(await staking.staked(user1.address)).to.equal(500);
    expect(await staking.votingPower(user1.address)).to.equal(500);
  });

  it("should allow unstaking", async () => {
    await staking.connect(user1).stake(300);
    await staking.connect(user1).unstake(200);
    expect(await staking.staked(user1.address)).to.equal(100);
    expect(await lpToken.balanceOf(user1.address)).to.equal(900);
  });

  it("should not allow unstaking more than staked", async () => {
    await staking.connect(user1).stake(100);
    await expect(staking.connect(user1).unstake(200)).to.be.revertedWith("Not enough staked");
  });

  it("should not allow staking 0", async () => {
    await expect(staking.connect(user1).stake(0)).to.be.revertedWith("Cannot stake 0");
  });
}); 