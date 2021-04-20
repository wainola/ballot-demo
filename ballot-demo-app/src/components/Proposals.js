import React, {
  useContext,
  useEffect,
  useState,
  useReducer,
  useRef,
} from "react";
import { Context } from "../App";
import { getBalance } from "../utils/web3";

const reducer = (state, action) => {
  switch (action.type) {
    case "SET_PROPOSAL":
      const {
        payload: { name, value },
      } = action;
      return {
        ...state,
        proposal: {
          ...state.proposal,
          [name]: value,
        },
      };
    case "ACCOUNT_BALANCE_FETCHED":
      return {
        ...state,
        balanceFromAccount: action.payload,
      };
    case "SUCCESS_REGISTER_PROPOSAL":
      return {
        ...state,
        isProposalRegister: action.payload,
      };
    case "SET_PROPOSAL_SAVED_BY_USER":
      return {
        ...state,
        proposalRecentlySaved: action.payload,
      };
    default:
      return state;
  }
};

const Proposals = ({ accounts, contractInstance, web3 }) => {
  const [voter, setVoter] = useState(null);
  const context = useContext(Context);
  const ref = useRef(null);

  const initState = {
    proposal: {},
    balanceFromAccount: "",
    isProposalRegister: false,
    proposals: [],
    proposalRecentlySaved: {},
  };

  const getProposal = async (deployedInstance, accountSelected) => {
    const {
      proposal: { title },
    } = state;

    try {
      const proposalResponse = await deployedInstance.getProposal(
        accountSelected,
        title,
        { from: accountSelected }
      );
      const {
        title: titleSaved,
        description,
        votes,
        weight,
      } = proposalResponse;

      dispatcher({
        type: "SET_PROPOSAL_SAVED_BY_USER",
        payload: {
          title: titleSaved,
          description,
          votes,
          weight: web3.utils.fromWei(weight, "ether"),
        },
      });
    } catch (error) {
      console.log("Error", error);
    }
  };

  const [state, dispatcher] = useReducer(reducer, initState);

  console.log("state", state);

  useEffect(() => {
    if (state.isProposalRegister) {
      getProposal(
        contractInstance.deployedInstance,
        context.ballotContext.accountSelected
      );
    }
  }, [state.isProposalRegister]);

  const getBalanceFromAccount = async (web3, accountSelected) => {
    try {
      const balanceFromAccount = await getBalance(web3, accountSelected);
      return dispatcher({
        type: "ACCOUNT_BALANCE_FETCHED",
        payload: balanceFromAccount,
      });
    } catch (error) {
      return console.log("Error getting balanace", error);
    }
  };

  useEffect(() => {
    if (
      context.ballotContext !== null ||
      Object.keys(context.ballotContext).length
    ) {
      const {
        ballotContext: { lastVoterInfo, accountSelected },
      } = context;

      setVoter({
        ...lastVoterInfo,
      });

      getBalanceFromAccount(web3, accountSelected);
    }
  }, [context.ballotContext]);

  const handleChange = ({ target: { name, value } }) => {
    dispatcher({
      type: "SET_PROPOSAL",
      payload: { name, value },
    });
  };

  const submit = async (evt) => {
    evt.preventDefault();
    if (Object.values(state.proposal).length) {
      const {
        proposal: { title, description, fee },
      } = state;

      const { deployedInstance } = contractInstance;

      const {
        ballotContext: { accountSelected },
      } = context;

      try {
        const {
          receipt: { status },
        } = await deployedInstance.setProposal(title, description, {
          from: accountSelected,
          value: web3.utils.toWei(fee, "ether"),
        });

        dispatcher({
          type: "SUCCESS_REGISTER_PROPOSAL",
          payload: status,
        });
      } catch (error) {
        console.log("error", error);
      }
    }
  };

  return (
    <div>
      <h2>
        {voter ? (
          <div>
            <div>
              <span>
                Hello {voter.name} {voter.lastname}
              </span>
            </div>
            <div>Account number: {context.ballotContext.accountSelected}</div>
          </div>
        ) : (
          <span>Hello stranger, there is no account related to you</span>
        )}
      </h2>
      {state.balanceFromAccount && (
        <div>Your current balance is {state.balanceFromAccount} ETH</div>
      )}
      <div>
        {voter && (
          <>
            <form onSubmit={submit} ref={ref}>
              <input
                type="text"
                name="title"
                placeholder="Proposal title"
                onChange={handleChange}
              />
              <input
                type="text"
                name="description"
                placeholder="Proposal description"
                onChange={handleChange}
              />
              <input
                type="number"
                placeholder="Amount of eth to endorse"
                name="fee"
                onChange={handleChange}
              />

              <button type="submit">Submit proposal</button>
            </form>
            <div>
              {Object.values(state.proposalRecentlySaved).length !== 0 && (
                <div>
                  <p>Title: {state.proposalRecentlySaved.title}</p>
                  <p>Dscription: {state.proposalRecentlySaved.description}</p>
                  <p>Weight: {state.proposalRecentlySaved.weight} ETH</p>
                </div>
              )}
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default Proposals;
