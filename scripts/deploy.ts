import { ethers } from "hardhat";

async function main() {
  // 1. 토큰 배포
  const Token = await ethers.getContractFactory("MyGameToken");
  const token = await Token.deploy();
  await token.waitForDeployment();

  const tokenAddress = await token.getAddress();
  console.log(`✅ MyGameToken deployed at: ${tokenAddress}`);

  // 2. 게임 컨트랙트 배포 (토큰 주소 전달)
  const Game = await ethers.getContractFactory("MonsterGame");
  const game = await Game.deploy(tokenAddress);
  await game.waitForDeployment();

  const gameAddress = await game.getAddress();
  console.log(`✅ MonsterGame deployed at: ${gameAddress}`);

  // 3. 토큰 소유권 이전
  const tx = await token.transferOwnership(gameAddress);
  await tx.wait();
  console.log("🔑 Token ownership transferred to MonsterGame contract");
}

main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
