import React, { useContext, useEffect, useRef, useState } from "react";
import { Context } from "../App";
import { AccountSelector } from "./";
import { getBalance } from "../utils";

const ProposalsByUser = ({ web3, accounts, contractInstance }) => {
  const context = useContext(Context);
  const [voter, setVoter] = useState(null);
  const [account, setAccount] = useState(null);
  const [proposals, setProposals] = useState([]);
  const [accountSelected, setAccountFromAccounts] = useState(null);
  const [voteSuccess, setVoteSuccess] = useState(false);
  const ref = useRef(null);

  const getOwnProposals = async () => {
    const { deployedInstance } = contractInstance;
    const {
      ballotContext: { accountSelected },
    } = context;
    try {
      const proposals = await deployedInstance.getOwnProposals({
        from: accountSelected,
      });
      const proposalMaped = proposals.map((proposal) => {
        const [title, description, votes, weight] = proposal;
        const weightOnEther = web3.utils.fromWei(weight, "ether");
        return { title, description, votes, weight: weightOnEther };
      });
      setProposals(proposalMaped);
    } catch (error) {
      console.log("Error", error);
    }
  };

  useEffect(() => {
    if (context.ballotContext && context.ballotContext.accountSelected) {
      const {
        ballotContext: { accountSelected, lastVoterInfo },
      } = context;
      getOwnProposals();
      setVoter(lastVoterInfo);
      setAccount(accountSelected);
    }
  }, []);

  useEffect(() => {
    if (voteSuccess) {
      getOwnProposals();
      setVoteSuccess(false);
    }
  }, [voteSuccess]);

  const execVote = async (address, accountSelected, proposalTitle) => {
    console.log("::::", address, accountSelected, proposalTitle);
    const { deployedInstance } = contractInstance;
    try {
      const {
        receipt: { status },
      } = await deployedInstance.vote(address, proposalTitle, {
        from: accountSelected,
        value: web3.utils.toWei("1", "ether"),
      });
      if (status) {
        alert("Success!");
        ref.current.selectedIndex = 0;
        setVoteSuccess(true);
      }
    } catch (error) {
      console.log("Error voting", error);
    }
  };

  const handleVote = (title) => () => execVote(account, accountSelected, title);

  const handleChange = ({ target: { value } }) => {
    if (value === account) {
      alert("You cannot use the same account to vote");
      ref.current.selectedIndex = 0;
      return;
    }
    setAccountFromAccounts(value);
  };

  return (
    <div>
      {voter !== null && Object.keys(voter).length !== 0 && (
        <>
          <div>
            <h2>
              Proposals for {voter.name} {voter.lastname}
            </h2>
            <div>Account number {account}</div>
          </div>
        </>
      )}
      {accounts.length && (
        <AccountSelector
          value={accountSelected}
          handleChange={handleChange}
          accounts={accounts}
          message="Select an account to vote"
          reference={ref}
        />
      )}
      {proposals.length !== 0 && (
        <div>
          {proposals.map((proposal, idx) => (
            <div key={idx}>
              <p>{proposal.title}</p>
              <p>
                Description: {proposal.description} - votes: {proposal.votes} -
                weight: {proposal.weight} ETH
              </p>
              <button onClick={handleVote(proposal.title)}>Vote!</button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default ProposalsByUser;
