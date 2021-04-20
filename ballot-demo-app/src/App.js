import { useEffect, useState, useContext, createContext } from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import FactoryBallot from "./contracts/FactoryBallot.json";
import { FactoryBallot as FactoryBallotProvider, getAccounts } from "./utils";
import { Accounts, Proposals, ProposalByUser } from "./components";
import "./App.css";

export const Context = createContext(null);

function App({ web3 }) {
  const [contractInstance, setContractInstance] = useState(null);
  const [accounts, setAccounts] = useState([]);
  const [ballotContext, setBallotData] = useState(null);

  const setFactoryBallot = async () => {
    try {
      const factoryBallot = await FactoryBallotProvider.setDeployedInstance(
        FactoryBallot
      );
      setContractInstance(factoryBallot);
    } catch (error) {
      console.log("error", error);
    }
  };

  useEffect(() => {
    setFactoryBallot();
    getAccounts(web3, setAccounts);
  }, []);

  return (
    <div className="App">
      <Context.Provider value={{ ballotContext, setBallotData }}>
        <Router>
          <Switch>
            <Route exact={true} path="/proposal-by-user">
              <ProposalByUser
                web3={web3}
                accounts={accounts}
                contractInstance={contractInstance}
              />
              <Link to="/">Go back to main page</Link>
            </Route>
            <Route exact={true} path="/create-proposals">
              <Proposals
                web3={web3}
                accounts={accounts}
                contractInstance={contractInstance}
              />
              <Link to="/">Go back to main page</Link>
            </Route>
            <Route path="/">
              <h1>Ballor App</h1>
              <Accounts
                accounts={accounts}
                contractInstance={contractInstance}
              />
              <Link to="/create-proposals">Go to create proposal</Link>
              <Link to="/proposal-by-user">Go to proposal by user</Link>
            </Route>
          </Switch>
        </Router>
      </Context.Provider>
    </div>
  );
}

export default App;
