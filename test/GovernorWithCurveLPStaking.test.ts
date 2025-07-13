import { expect } from "chai";
import { ethers } from "hardhat";

describe("Governor with CurveLPStaking", function () {
  let lpToken: any;
  let staking: any;
  let governor: any;
  let owner: any;
  let user1: any;
  let user2: any;
  let user3: any;

  beforeEach(async () => {
    [owner, user1, user2, user3] = await ethers.getSigners();
    // Curve LP 토큰 배포 및 mint
    const LPToken = await ethers.getContractFactory("CurveLPToken");
    lpToken = await LPToken.deploy();
    await lpToken.waitForDeployment();
    await lpToken.connect(owner).mint(user1.address, 1000);
    await lpToken.connect(owner).mint(user2.address, 1000);
    await lpToken.connect(owner).mint(user3.address, 1000);
    // Curve LP Staking 배포
    const Staking = await ethers.getContractFactory("CurveLPStaking");
    staking = await Staking.deploy(await lpToken.getAddress());
    await staking.waitForDeployment();
    // approve
    await lpToken.connect(user1).approve(await staking.getAddress(), 1000);
    await lpToken.connect(user2).approve(await staking.getAddress(), 1000);
    await lpToken.connect(user3).approve(await staking.getAddress(), 1000);
    // GovernorContract 배포 (CurveLPStaking을 IVotes로)
    const Governor = await ethers.getContractFactory("GovernorContract");
    governor = await Governor.deploy(await staking.getAddress());
    await governor.waitForDeployment();
  });

  it("should allow proposal, voting, and execution with Curve LP staking as voting power", async () => {
    // user1, user3 스테이킹 (user2는 일부만)
    await staking.connect(user1).stake(500);
    await staking.connect(user2).stake(200);
    await staking.connect(user3).stake(500);

    // 제안 생성
    const [target] = [await governor.getAddress()];
    const value = 0;
    const calldata = "0x"; // 아무 동작 없음
    const description = "Test proposal with Curve LP voting power";
    const tx = await governor.connect(user1).propose([target], [value], [calldata], description);
    const receipt = await tx.wait();
    const proposalId = receipt?.logs?.find((log: any) => log.fragment?.name === "ProposalCreated")?.args?.proposalId || (await governor.hashProposal([target], [value], [calldata], ethers.keccak256(ethers.toUtf8Bytes(description))));

    // 블록 진행(votingDelay 이후)
    await ethers.provider.send("evm_mine");

    // 투표 (user1, user3 찬성, user2 반대)
    await governor.connect(user1).castVote(proposalId, 1); // For (500)
    await governor.connect(user2).castVote(proposalId, 0); // Against (200)
    await governor.connect(user3).castVote(proposalId, 1); // For (500)

    // 블록 진행(votingPeriod 이후)
    const votingPeriod = await governor.votingPeriod();
    for (let i = 0; i < votingPeriod; i++) {
      await ethers.provider.send("evm_mine");
    }

    // 실행 가능 여부 확인 (찬성 1000, 반대 200 → Succeeded)
    await expect(
      governor.execute([target], [value], [calldata], ethers.keccak256(ethers.toUtf8Bytes(description)))
    ).to.not.be.reverted;
  });
}); 