// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NodeMetaData contains all meta data concerning the Node contract.
var NodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenInstance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"startTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"finishTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getAvailableNodeStartsFrom\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContract\",\"type\":\"address\"}],\"name\":\"updateTaskContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNodeAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeMetaData.ABI instead.
var NodeABI = NodeMetaData.ABI

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_Node *NodeCaller) AvailableNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "availableNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_Node *NodeSession) AvailableNodes() (*big.Int, error) {
	return _Node.Contract.AvailableNodes(&_Node.CallOpts)
}

// AvailableNodes is a free data retrieval call binding the contract method 0x204401e7.
//
// Solidity: function availableNodes() view returns(uint256)
func (_Node *NodeCallerSession) AvailableNodes() (*big.Int, error) {
	return _Node.Contract.AvailableNodes(&_Node.CallOpts)
}

// GetAllNodeAddresses is a free data retrieval call binding the contract method 0xc8fe3a01.
//
// Solidity: function getAllNodeAddresses() view returns(address[])
func (_Node *NodeCaller) GetAllNodeAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getAllNodeAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllNodeAddresses is a free data retrieval call binding the contract method 0xc8fe3a01.
//
// Solidity: function getAllNodeAddresses() view returns(address[])
func (_Node *NodeSession) GetAllNodeAddresses() ([]common.Address, error) {
	return _Node.Contract.GetAllNodeAddresses(&_Node.CallOpts)
}

// GetAllNodeAddresses is a free data retrieval call binding the contract method 0xc8fe3a01.
//
// Solidity: function getAllNodeAddresses() view returns(address[])
func (_Node *NodeCallerSession) GetAllNodeAddresses() ([]common.Address, error) {
	return _Node.Contract.GetAllNodeAddresses(&_Node.CallOpts)
}

// GetAvailableNodeStartsFrom is a free data retrieval call binding the contract method 0x68f52d17.
//
// Solidity: function getAvailableNodeStartsFrom(uint256 i) view returns(address)
func (_Node *NodeCaller) GetAvailableNodeStartsFrom(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getAvailableNodeStartsFrom", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAvailableNodeStartsFrom is a free data retrieval call binding the contract method 0x68f52d17.
//
// Solidity: function getAvailableNodeStartsFrom(uint256 i) view returns(address)
func (_Node *NodeSession) GetAvailableNodeStartsFrom(i *big.Int) (common.Address, error) {
	return _Node.Contract.GetAvailableNodeStartsFrom(&_Node.CallOpts, i)
}

// GetAvailableNodeStartsFrom is a free data retrieval call binding the contract method 0x68f52d17.
//
// Solidity: function getAvailableNodeStartsFrom(uint256 i) view returns(address)
func (_Node *NodeCallerSession) GetAvailableNodeStartsFrom(i *big.Int) (common.Address, error) {
	return _Node.Contract.GetAvailableNodeStartsFrom(&_Node.CallOpts, i)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeCaller) GetNodeStatus(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getNodeStatus", nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeSession) GetNodeStatus(nodeAddress common.Address) (*big.Int, error) {
	return _Node.Contract.GetNodeStatus(&_Node.CallOpts, nodeAddress)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0xb65f8177.
//
// Solidity: function getNodeStatus(address nodeAddress) view returns(uint256)
func (_Node *NodeCallerSession) GetNodeStatus(nodeAddress common.Address) (*big.Int, error) {
	return _Node.Contract.GetNodeStatus(&_Node.CallOpts, nodeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Node *NodeCallerSession) Owner() (common.Address, error) {
	return _Node.Contract.Owner(&_Node.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_Node *NodeCaller) TotalNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "totalNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_Node *NodeSession) TotalNodes() (*big.Int, error) {
	return _Node.Contract.TotalNodes(&_Node.CallOpts)
}

// TotalNodes is a free data retrieval call binding the contract method 0x9592d424.
//
// Solidity: function totalNodes() view returns(uint256)
func (_Node *NodeCallerSession) TotalNodes() (*big.Int, error) {
	return _Node.Contract.TotalNodes(&_Node.CallOpts)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeTransactor) FinishTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "finishTask", nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.FinishTask(&_Node.TransactOpts, nodeAddress)
}

// FinishTask is a paid mutator transaction binding the contract method 0x3fc0f48b.
//
// Solidity: function finishTask(address nodeAddress) returns()
func (_Node *NodeTransactorSession) FinishTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.FinishTask(&_Node.TransactOpts, nodeAddress)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Node *NodeTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Node *NodeSession) Join() (*types.Transaction, error) {
	return _Node.Contract.Join(&_Node.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Node *NodeTransactorSession) Join() (*types.Transaction, error) {
	return _Node.Contract.Join(&_Node.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeSession) Pause() (*types.Transaction, error) {
	return _Node.Contract.Pause(&_Node.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Node *NodeTransactorSession) Pause() (*types.Transaction, error) {
	return _Node.Contract.Pause(&_Node.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeTransactor) Quit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "quit")
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeSession) Quit() (*types.Transaction, error) {
	return _Node.Contract.Quit(&_Node.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_Node *NodeTransactorSession) Quit() (*types.Transaction, error) {
	return _Node.Contract.Quit(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Node *NodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Node.Contract.RenounceOwnership(&_Node.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeSession) Resume() (*types.Transaction, error) {
	return _Node.Contract.Resume(&_Node.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Node *NodeTransactorSession) Resume() (*types.Transaction, error) {
	return _Node.Contract.Resume(&_Node.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeTransactor) Slash(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "slash", nodeAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeSession) Slash(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.Slash(&_Node.TransactOpts, nodeAddress)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address nodeAddress) returns()
func (_Node *NodeTransactorSession) Slash(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.Slash(&_Node.TransactOpts, nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeTransactor) StartTask(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "startTask", nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.StartTask(&_Node.TransactOpts, nodeAddress)
}

// StartTask is a paid mutator transaction binding the contract method 0x5f51c765.
//
// Solidity: function startTask(address nodeAddress) returns()
func (_Node *NodeTransactorSession) StartTask(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.StartTask(&_Node.TransactOpts, nodeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Node *NodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.TransferOwnership(&_Node.TransactOpts, newOwner)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeTransactor) UpdateTaskContractAddress(opts *bind.TransactOpts, taskContract common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "updateTaskContractAddress", taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _Node.Contract.UpdateTaskContractAddress(&_Node.TransactOpts, taskContract)
}

// UpdateTaskContractAddress is a paid mutator transaction binding the contract method 0x42145230.
//
// Solidity: function updateTaskContractAddress(address taskContract) returns()
func (_Node *NodeTransactorSession) UpdateTaskContractAddress(taskContract common.Address) (*types.Transaction, error) {
	return _Node.Contract.UpdateTaskContractAddress(&_Node.TransactOpts, taskContract)
}

// NodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Node contract.
type NodeOwnershipTransferredIterator struct {
	Event *NodeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnershipTransferred represents a OwnershipTransferred event raised by the Node contract.
type NodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnershipTransferredIterator{contract: _Node.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnershipTransferred)
				if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Node *NodeFilterer) ParseOwnershipTransferred(log types.Log) (*NodeOwnershipTransferred, error) {
	event := new(NodeOwnershipTransferred)
	if err := _Node.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
