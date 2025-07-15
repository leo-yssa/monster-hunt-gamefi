// @ts-ignore
import { ethers, upgrades } from "hardhat";

async function main() {
  // 1. 토큰 배포
  const Token = await ethers.getContractFactory("MyGameToken");
  const token = await Token.deploy();
  await token.waitForDeployment();
  const tokenAddress = await token.getAddress();
  console.log(`✅ MyGameToken deployed at: ${tokenAddress}`);

  // 2-1. Transparent 프록시 배포
  const GameTransparent = await ethers.getContractFactory("MonsterGameTransparent");
  const gameTransparent = await upgrades.deployProxy(GameTransparent, [tokenAddress], { initializer: "initialize" });
  await gameTransparent.waitForDeployment();
  const gameTransparentAddress = await gameTransparent.getAddress();
  console.log(`✅ MonsterGameTransparent (Transparent Proxy) deployed at: ${gameTransparentAddress}`);

  // 2-2. UUPS 프록시 배포
  const GameUUPS = await ethers.getContractFactory("MonsterGameUUPS");
  const gameUUPS = await upgrades.deployProxy(GameUUPS, [tokenAddress], { initializer: "initialize", kind: "uups" });
  await gameUUPS.waitForDeployment();
  const gameUUPSAddress = await gameUUPS.getAddress();
  console.log(`✅ MonsterGameUUPS (UUPS Proxy) deployed at: ${gameUUPSAddress}`);

  // 2-3. Beacon 프록시 배포
  const GameBeaconImpl = await ethers.getContractFactory("MonsterGameBeaconImpl");
  const beacon = await upgrades.deployBeacon(GameBeaconImpl);
  await beacon.waitForDeployment();
  const beaconAddress = await beacon.getAddress();
  const gameBeacon = await upgrades.deployBeaconProxy(beacon, GameBeaconImpl, [tokenAddress], { initializer: "initialize" });
  await gameBeacon.waitForDeployment();
  const gameBeaconAddress = await gameBeacon.getAddress();
  console.log(`✅ MonsterGameBeaconImpl (Beacon Proxy) deployed at: ${gameBeaconAddress}`);
  console.log(`   └ Beacon address: ${beaconAddress}`);

  // 3. 토큰 소유권 이전 (각 프록시 주소로)
  await token.transferOwnership(gameTransparentAddress);
  await token.transferOwnership(gameUUPSAddress);
  await token.transferOwnership(gameBeaconAddress);
  console.log("🔑 Token ownership transferred to all MonsterGame proxies");
}

main().catch((error) => {
  console.error("❌ Deployment failed:", error);
  process.exitCode = 1;
});
