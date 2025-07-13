import { expect } from "chai";
import { ethers } from "hardhat";

describe("GovernorContract", function () {
  let token: any;
  let governor: any;
  let owner: any;
  let user1: any;
  let user2: any;
  let user3: any;

  beforeEach(async () => {
    [owner, user1, user2, user3] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("GovernanceToken");
    token = await Token.deploy();
    await token.waitForDeployment();
    await token.connect(owner).mint(user1.address, 1000);
    await token.connect(owner).mint(user2.address, 1000);
    await token.connect(owner).mint(user3.address, 1000);
    await token.connect(user1).delegate(user1.address);
    await token.connect(user2).delegate(user2.address);
    await token.connect(user3).delegate(user3.address);
    const Governor = await ethers.getContractFactory("GovernorContract");
    governor = await Governor.deploy(await token.getAddress());
    await governor.waitForDeployment();
  });

  it("should allow proposal, voting, and execution (dummy call)", async () => {
    // 테스트용 타겟 컨트랙트(자기 자신)와 call data
    const [target] = [await governor.getAddress()];
    const value = 0;
    const calldata = "0x"; // 아무 동작 없음
    const description = "Test proposal";

    // 제안 생성
    const tx = await governor.connect(user1).propose([target], [value], [calldata], description);
    const receipt = await tx.wait();
    const proposalId = receipt?.logs?.find((log: any) => log.fragment?.name === "ProposalCreated")?.args?.proposalId || (await governor.hashProposal([target], [value], [calldata], ethers.keccak256(ethers.toUtf8Bytes(description))));

    // 블록 진행(votingDelay 이후)
    await ethers.provider.send("evm_mine");

    // 투표
    await governor.connect(user1).castVote(proposalId, 1); // For
    await governor.connect(user2).castVote(proposalId, 0); // Against
    await governor.connect(user3).castVote(proposalId, 1); // For

    // 블록 진행(votingPeriod 이후)
    const votingPeriod = await governor.votingPeriod();
    for (let i = 0; i < votingPeriod; i++) {
      await ethers.provider.send("evm_mine");
    }

    // 실행 가능 여부 확인
    await expect(
      governor.execute([target], [value], [calldata], ethers.keccak256(ethers.toUtf8Bytes(description)))
    ).to.not.be.reverted;
  });
}); 