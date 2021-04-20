import { webProvider } from "./web3";
import "@truffle/contract";

function FactoryBallot(deployedInstance) {
  this.deployedInstance = deployedInstance;
}

FactoryBallot.setDeployedInstance = async function (contract) {
  let ballotContract;
  if ("TruffleContract" in window) {
    ballotContract = window.TruffleContract(contract);
    ballotContract.setProvider(webProvider);
  }

  try {
    const deployedInstance = await ballotContract.deployed();
    return new FactoryBallot(deployedInstance);
  } catch (error) {
    return error;
  }
};

const factoryBallot = FactoryBallot;

export default factoryBallot;
