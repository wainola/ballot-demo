# Ballot Demo App

## Setup

For this app to work you will need to install truffle suite globally

```bash
npm i -g truffle
```

Afterwards you should go to the Truffle website and download [Ganache](https://www.trufflesuite.com/ganache). This will provide you with a local ethereum network to test this app or your projects.

Clone this repo and run `npm i`. You will also need to clone this [folder](https://github.com/wainola/BIA/tree/master/dappV1/Ballot-Dapp/ballot-demo) or just clone the entire project. This folder contains the contracts that you will need to compile and migrate to Ganache.

The commands to use with the [contracts folder](https://github.com/wainola/BIA/tree/master/dappV1/Ballot-Dapp/ballot-demo) are the following (provided that you installed truffle globally)

- `truffle compile` will compile the contracts
- `truffle migrate --reset` will migrate the contracts to the network and also pase the **JSON-RPC** format to the frontend directory

You can check the configuration for the network and the compile directory in the `truffle-config.js` file

## TODO

- there is a bug in the vote section but I don't have the time rigth now to fix it
- improve on code abstracation
- add some styles, perhaps tailwind.css
- add xstate
- add more methods for the users accounts
- interop other chains?
