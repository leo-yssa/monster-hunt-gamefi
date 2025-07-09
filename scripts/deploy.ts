// @ts-ignore
import { ethers, upgrades } from "hardhat";

async function main() {
  // 1. 토큰 배포
  const Token = await ethers.getContractFactory("MyGameToken");
  const token = await Token.deploy();
  await token.waitForDeployment();

  const tokenAddress = await token.getAddress();
  console.log(`✅ MyGameToken deployed at: ${tokenAddress}`);

  // 2. 게임 컨트랙트 프록시 배포 (initialize 호출)
  const Game = await ethers.getContractFactory("MonsterGame");
  const game = await upgrades.deployProxy(Game, [tokenAddress], { initializer: "initialize" });
  await game.waitForDeployment();

  const gameAddress = await game.getAddress();
  console.log(`✅ MonsterGame (proxy) deployed at: ${gameAddress}`);

  // 3. 토큰 소유권 이전 (프록시 주소로)
  const tx = await token.transferOwnership(gameAddress);
  await tx.wait();
  console.log("🔑 Token ownership transferred to MonsterGame contract (proxy)");
}

main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
