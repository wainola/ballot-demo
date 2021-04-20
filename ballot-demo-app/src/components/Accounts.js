import React, { useReducer, useEffect, useRef, useContext } from "react";
import { Context } from "../App";
import { getInfoVoter } from "../utils/helpers";
import { AccountSelector } from "./";

const reducer = (state, action) => {
  switch (action.type) {
    case "REGISTER_VOTER":
      return state;
    case "SELECT_ACCOUNT":
      return {
        ...state,
        accountSelected: action.payload,
      };
    case "FILL_USER_DATA":
      const {
        payload: { name, value },
      } = action;
      return {
        ...state,
        voterToRegister: {
          ...state.voterToRegister,
          [name]: value,
        },
      };
    case "CLEAR_USER_TO_REGISTER":
      return {
        ...state,
        voterToRegister: {},
        accountSelected: "",
      };
    case "SET_CHAIRPERSON":
      return {
        ...state,
        chairperson: action.payload,
      };
    case "SUCCESS_REGISTERING":
      return {
        ...state,
        successRegistering: true,
        voterToRegister: {},
      };
    case "SET_LAST_INFO_VOTER": {
      const {
        payload: { name, lastname, email },
      } = action;
      return {
        ...state,
        lastVoterInfo: {
          name,
          lastname,
          email,
        },
      };
    }
    default:
      return state;
  }
};

const Accounts = ({ accounts, contractInstance }) => {
  const context = useContext(Context);
  const initialState = {
    voters: {},
    infoVoter: {},
    accountSelected: "",
    voterToRegister: {},
    chairperson: "",
    getInfoVoter: false,
    successRegistering: false,
    lastVoterInfo: {},
  };

  const ref = useRef(null);

  const [state, dispatcher] = useReducer(reducer, initialState);

  useEffect(() => {
    if (accounts.length) {
      const chairperson = accounts[0];
      dispatcher({
        type: "SET_CHAIRPERSON",
        payload: chairperson,
      });

      context.setBallotData({
        ...context.ballotContext,
        chairperson,
      });
    }
  }, [accounts]);

  useEffect(() => {
    if (Object.keys(state.lastVoterInfo).length) {
      context.setBallotData({
        ...context.ballotContext,
        lastVoterInfo: state.lastVoterInfo,
      });
    }
  }, [state.lastVoterInfo]);

  const handleChange = async ({ target: { value } }) => {
    dispatcher({
      type: "SELECT_ACCOUNT",
      payload: value,
    });

    context.setBallotData({
      ...context.ballotContext,
      accountSelected: value,
    });
  };

  const handleInputChange = ({ target: { name, value } }) =>
    dispatcher({
      type: "FILL_USER_DATA",
      payload: { name, value },
    });

  const register = async (evt) => {
    evt.preventDefault();
    if (!Object.values(state.voterToRegister).includes("")) {
      // register voter
      const { voterToRegister, accountSelected, chairperson } = state;
      const { deployedInstance } = contractInstance;
      try {
        const response = await deployedInstance.registerVoters(
          accountSelected,
          voterToRegister.name,
          voterToRegister.lastname,
          voterToRegister.email,
          { from: chairperson }
        );
        ref.current.reset();
        return dispatcher({
          type: "SUCCESS_REGISTERING",
        });
      } catch (error) {
        console.log("Error", error);
      }
    }
    return dispatcher({
      type: "CLEAR_USER_TO_REGISTER",
    });
  };

  const callGetInfoVoter = async () => {
    const { deployedInstance } = contractInstance;
    const { accountSelected, chairperson } = state;
    return getInfoVoter(
      deployedInstance,
      accountSelected,
      chairperson,
      dispatcher
    );
  };

  return (
    <div>
      <div>
        <h3>Accounts</h3>
      </div>
      <div>
        {accounts.length && (
          <AccountSelector
            handleChange={handleChange}
            accounts={accounts}
            value={state.accountSelected}
            message="Select an account"
          />
        )}
        <div>
          <button onClick={callGetInfoVoter}>Get Info of last voter!</button>
          <div>
            {Object.values(state.lastVoterInfo).length !== 0 && (
              <div>
                <h4>Last voter registered!</h4>
                <p>
                  {state.lastVoterInfo.name} {state.lastVoterInfo.lastname}
                </p>
                <p>Email: {state.lastVoterInfo.email}</p>
              </div>
            )}
          </div>
        </div>
        <div>
          {state.accountSelected.length !== 0 && (
            <form action="" onSubmit={register} ref={ref}>
              <h4>
                Registering voter for account:{" "}
                {state.accountSelected.length && state.accountSelected}
              </h4>
              <input
                name="name"
                type="text"
                placeholder="name"
                onChange={handleInputChange}
              />
              <input
                name="lastname"
                type="text"
                placeholder="lastname"
                onChange={handleInputChange}
              />
              <input
                name="email"
                type="text"
                placeholder="email"
                onChange={handleInputChange}
              />
              <button type="submit">Register Voter!</button>
            </form>
          )}
        </div>
      </div>
    </div>
  );
};

export default Accounts;
