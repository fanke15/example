/** @type import('hardhat/config').HardhatUserConfig */
import "@nomiclabs/hardhat-waffle"
import "@nomiclabs/hardhat-ethers"
import "hardhat-abi-exporter";
import "hardhat-gas-reporter";
import "hardhat-contract-sizer";

const mnemonic = "season turtle oblige language winner purpose call engine thunder pepper cactus base";
const privateKey = [
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386710",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386711",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386722",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386733",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386744",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386755",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386766",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386777",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386788",
  "0x6a4843f986e41899fa98984f0a559d8e870f2bb5552bfeb79b427c1199386799",
];
const LOCAL_RPC = "http://127.0.0.1:8545";

const CURRENT_RPC = LOCAL_RPC;
const DefaultNetwork = "localhost";
const GasPrice = 90e9;



module.exports = {
  solidity: "0.8.17",
  networks: {
  localhost: {
    gasPrice: GasPrice,
    accounts: privateKey,
    url: CURRENT_RPC,
    chainId: 43114,
    allowUnlimitedContractSize: true,
    timeout: 1000000,
    }
  },
  gasReporter: {
    enabled: true,
    showMethodSig: true,
    maxMethodDiff: 10,
    currency: 'USD',
    gasPrice: 127,
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false,
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
  abiExporter: {
    path: './abi',
    runOnCompile: true,
    clear: true,
    spacing: 2
  }
};

