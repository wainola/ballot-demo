// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FactoryBallotProposal is an auto generated low-level Go binding around an user-defined struct.
type FactoryBallotProposal struct {
	Title       string
	Description string
	Votes       *big.Int
	Weight      *big.Int
}

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getInfoVoter\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lastname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structFactoryBallot.Proposal[]\",\"name\":\"ownProposals\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proposalTitle\",\"type\":\"string\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structFactoryBallot.Proposal\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proposalTitle\",\"type\":\"string\"}],\"name\":\"getVotesForProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lastname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"name\":\"registerVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"setProposal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proposalTitle\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lastname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Main *MainCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Main *MainSession) Chairperson() (common.Address, error) {
	return _Main.Contract.Chairperson(&_Main.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Main *MainCallerSession) Chairperson() (common.Address, error) {
	return _Main.Contract.Chairperson(&_Main.CallOpts)
}

// GetInfoVoter is a free data retrieval call binding the contract method 0x6d18cca9.
//
// Solidity: function getInfoVoter(address voter) view returns(string name, string lastname, string email)
func (_Main *MainCaller) GetInfoVoter(opts *bind.CallOpts, voter common.Address) (struct {
	Name     string
	Lastname string
	Email    string
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getInfoVoter", voter)

	outstruct := new(struct {
		Name     string
		Lastname string
		Email    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Lastname = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetInfoVoter is a free data retrieval call binding the contract method 0x6d18cca9.
//
// Solidity: function getInfoVoter(address voter) view returns(string name, string lastname, string email)
func (_Main *MainSession) GetInfoVoter(voter common.Address) (struct {
	Name     string
	Lastname string
	Email    string
}, error) {
	return _Main.Contract.GetInfoVoter(&_Main.CallOpts, voter)
}

// GetInfoVoter is a free data retrieval call binding the contract method 0x6d18cca9.
//
// Solidity: function getInfoVoter(address voter) view returns(string name, string lastname, string email)
func (_Main *MainCallerSession) GetInfoVoter(voter common.Address) (struct {
	Name     string
	Lastname string
	Email    string
}, error) {
	return _Main.Contract.GetInfoVoter(&_Main.CallOpts, voter)
}

// GetOwnProposals is a free data retrieval call binding the contract method 0xf007cde2.
//
// Solidity: function getOwnProposals() view returns((string,string,uint256,uint256)[] ownProposals)
func (_Main *MainCaller) GetOwnProposals(opts *bind.CallOpts) ([]FactoryBallotProposal, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOwnProposals")

	if err != nil {
		return *new([]FactoryBallotProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]FactoryBallotProposal)).(*[]FactoryBallotProposal)

	return out0, err

}

// GetOwnProposals is a free data retrieval call binding the contract method 0xf007cde2.
//
// Solidity: function getOwnProposals() view returns((string,string,uint256,uint256)[] ownProposals)
func (_Main *MainSession) GetOwnProposals() ([]FactoryBallotProposal, error) {
	return _Main.Contract.GetOwnProposals(&_Main.CallOpts)
}

// GetOwnProposals is a free data retrieval call binding the contract method 0xf007cde2.
//
// Solidity: function getOwnProposals() view returns((string,string,uint256,uint256)[] ownProposals)
func (_Main *MainCallerSession) GetOwnProposals() ([]FactoryBallotProposal, error) {
	return _Main.Contract.GetOwnProposals(&_Main.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xbcceaac9.
//
// Solidity: function getProposal(address addr, string proposalTitle) view returns((string,string,uint256,uint256) result)
func (_Main *MainCaller) GetProposal(opts *bind.CallOpts, addr common.Address, proposalTitle string) (FactoryBallotProposal, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getProposal", addr, proposalTitle)

	if err != nil {
		return *new(FactoryBallotProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(FactoryBallotProposal)).(*FactoryBallotProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xbcceaac9.
//
// Solidity: function getProposal(address addr, string proposalTitle) view returns((string,string,uint256,uint256) result)
func (_Main *MainSession) GetProposal(addr common.Address, proposalTitle string) (FactoryBallotProposal, error) {
	return _Main.Contract.GetProposal(&_Main.CallOpts, addr, proposalTitle)
}

// GetProposal is a free data retrieval call binding the contract method 0xbcceaac9.
//
// Solidity: function getProposal(address addr, string proposalTitle) view returns((string,string,uint256,uint256) result)
func (_Main *MainCallerSession) GetProposal(addr common.Address, proposalTitle string) (FactoryBallotProposal, error) {
	return _Main.Contract.GetProposal(&_Main.CallOpts, addr, proposalTitle)
}

// GetVotesForProposal is a free data retrieval call binding the contract method 0xaacc937f.
//
// Solidity: function getVotesForProposal(address addr, string proposalTitle) view returns(uint256 voteCount)
func (_Main *MainCaller) GetVotesForProposal(opts *bind.CallOpts, addr common.Address, proposalTitle string) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getVotesForProposal", addr, proposalTitle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesForProposal is a free data retrieval call binding the contract method 0xaacc937f.
//
// Solidity: function getVotesForProposal(address addr, string proposalTitle) view returns(uint256 voteCount)
func (_Main *MainSession) GetVotesForProposal(addr common.Address, proposalTitle string) (*big.Int, error) {
	return _Main.Contract.GetVotesForProposal(&_Main.CallOpts, addr, proposalTitle)
}

// GetVotesForProposal is a free data retrieval call binding the contract method 0xaacc937f.
//
// Solidity: function getVotesForProposal(address addr, string proposalTitle) view returns(uint256 voteCount)
func (_Main *MainCallerSession) GetVotesForProposal(addr common.Address, proposalTitle string) (*big.Int, error) {
	return _Main.Contract.GetVotesForProposal(&_Main.CallOpts, addr, proposalTitle)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(string name, string lastname, string email, uint256 weight)
func (_Main *MainCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name     string
	Lastname string
	Email    string
	Weight   *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Name     string
		Lastname string
		Email    string
		Weight   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Lastname = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(string name, string lastname, string email, uint256 weight)
func (_Main *MainSession) Voters(arg0 common.Address) (struct {
	Name     string
	Lastname string
	Email    string
	Weight   *big.Int
}, error) {
	return _Main.Contract.Voters(&_Main.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(string name, string lastname, string email, uint256 weight)
func (_Main *MainCallerSession) Voters(arg0 common.Address) (struct {
	Name     string
	Lastname string
	Email    string
	Weight   *big.Int
}, error) {
	return _Main.Contract.Voters(&_Main.CallOpts, arg0)
}

// RegisterVoters is a paid mutator transaction binding the contract method 0x1767bdfd.
//
// Solidity: function registerVoters(address voter, string name, string lastname, string email) returns()
func (_Main *MainTransactor) RegisterVoters(opts *bind.TransactOpts, voter common.Address, name string, lastname string, email string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "registerVoters", voter, name, lastname, email)
}

// RegisterVoters is a paid mutator transaction binding the contract method 0x1767bdfd.
//
// Solidity: function registerVoters(address voter, string name, string lastname, string email) returns()
func (_Main *MainSession) RegisterVoters(voter common.Address, name string, lastname string, email string) (*types.Transaction, error) {
	return _Main.Contract.RegisterVoters(&_Main.TransactOpts, voter, name, lastname, email)
}

// RegisterVoters is a paid mutator transaction binding the contract method 0x1767bdfd.
//
// Solidity: function registerVoters(address voter, string name, string lastname, string email) returns()
func (_Main *MainTransactorSession) RegisterVoters(voter common.Address, name string, lastname string, email string) (*types.Transaction, error) {
	return _Main.Contract.RegisterVoters(&_Main.TransactOpts, voter, name, lastname, email)
}

// SetProposal is a paid mutator transaction binding the contract method 0x4092131d.
//
// Solidity: function setProposal(string title, string description) payable returns()
func (_Main *MainTransactor) SetProposal(opts *bind.TransactOpts, title string, description string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setProposal", title, description)
}

// SetProposal is a paid mutator transaction binding the contract method 0x4092131d.
//
// Solidity: function setProposal(string title, string description) payable returns()
func (_Main *MainSession) SetProposal(title string, description string) (*types.Transaction, error) {
	return _Main.Contract.SetProposal(&_Main.TransactOpts, title, description)
}

// SetProposal is a paid mutator transaction binding the contract method 0x4092131d.
//
// Solidity: function setProposal(string title, string description) payable returns()
func (_Main *MainTransactorSession) SetProposal(title string, description string) (*types.Transaction, error) {
	return _Main.Contract.SetProposal(&_Main.TransactOpts, title, description)
}

// Vote is a paid mutator transaction binding the contract method 0x1bf533a4.
//
// Solidity: function vote(address addr, string proposalTitle) payable returns()
func (_Main *MainTransactor) Vote(opts *bind.TransactOpts, addr common.Address, proposalTitle string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "vote", addr, proposalTitle)
}

// Vote is a paid mutator transaction binding the contract method 0x1bf533a4.
//
// Solidity: function vote(address addr, string proposalTitle) payable returns()
func (_Main *MainSession) Vote(addr common.Address, proposalTitle string) (*types.Transaction, error) {
	return _Main.Contract.Vote(&_Main.TransactOpts, addr, proposalTitle)
}

// Vote is a paid mutator transaction binding the contract method 0x1bf533a4.
//
// Solidity: function vote(address addr, string proposalTitle) payable returns()
func (_Main *MainTransactorSession) Vote(addr common.Address, proposalTitle string) (*types.Transaction, error) {
	return _Main.Contract.Vote(&_Main.TransactOpts, addr, proposalTitle)
}
