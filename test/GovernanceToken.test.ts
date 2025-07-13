import { expect } from "chai";
import { ethers } from "hardhat";

describe("GovernanceToken", function () {
  let token: any;
  let owner: any;
  let user1: any;
  let user2: any;

  beforeEach(async () => {
    [owner, user1, user2] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("GovernanceToken");
    token = await Token.deploy();
    await token.waitForDeployment();
  });

  it("should have correct name and symbol", async () => {
    expect(await token.name()).to.equal("GovernanceToken");
    expect(await token.symbol()).to.equal("GOV");
  });

  it("should allow owner to mint", async () => {
    await token.connect(owner).mint(user1.address, 1000);
    expect(await token.balanceOf(user1.address)).to.equal(1000);
  });

  it("should not allow non-owner to mint", async () => {
    await expect(token.connect(user1).mint(user2.address, 1000)).to.be.reverted;
  });

  it("should allow transfer and update voting power", async () => {
    await token.connect(owner).mint(user1.address, 1000);
    await token.connect(user1).delegate(user1.address);
    expect(await token.getVotes(user1.address)).to.equal(1000);
    await token.connect(user1).transfer(user2.address, 400);
    // 투표권은 transfer 후 자동 업데이트
    expect(await token.getVotes(user1.address)).to.equal(600);
    await token.connect(user2).delegate(user2.address);
    expect(await token.getVotes(user2.address)).to.equal(400);
  });
}); 