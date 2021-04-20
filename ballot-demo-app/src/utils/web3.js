import Web3 from "web3";

export const webProvider = new Web3.providers.HttpProvider(
  "http://localhost:7545"
);

const initWeb3 = () => {
  window.web3 = new Web3(webProvider);
};

export const getAccounts = (web3, accountUpdater) => {
  return web3.eth.getAccounts((error, accounts) => {
    if (error) {
      return console.error("ERROR_GETTING_ACCOUNTS", error);
    }
    return accountUpdater(accounts);
  });
};

export const getBalance = async (web3, accountAddress) => {
  try {
    const balance = await web3.eth.getBalance(accountAddress);
    const balanceFromWeiToEth = web3.utils.fromWei(balance, "ether");
    return balanceFromWeiToEth;
  } catch (error) {
    return error;
  }
};

export default initWeb3;
