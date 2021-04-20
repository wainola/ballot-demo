const FactoryBallot = artifacts.require("FactoryBallot");

contract("FactoryBallot", async (accounts) => {
  it("should register voters and proposals", async () => {
    const instance = await FactoryBallot.deployed({ from: accounts[0] });

    const registered = accounts[1];

    const r = await instance.registerVoters(
      registered,
      "name 1",
      "lastname 1",
      "mail@emaail.com",
      { from: accounts[0] }
    );

    const registerVoter = await instance.getInfoVoter(registered, {
      from: accounts[0],
    });

    const registerProposal = await instance.setProposal(
      "proposal 1",
      "description 1",
      {
        from: registered,
        value: web3.utils.toWei("1", "ether"),
      }
    );

    const proposalSaved = await instance.getProposal(registered, "proposal 1", {
      from: registered,
    });

    console.log("proposal:", proposalSaved);
  });
});
