const FactoryBallot = artifacts.require("FactoryBallot");

module.exports = function (dpeloyer, network, accounts) {
  const chairperson = accounts[0];
  dpeloyer.deploy(FactoryBallot, { from: chairperson });
};
