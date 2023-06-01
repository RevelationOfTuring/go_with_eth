// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scrds

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

// SCRDSMetaData contains all meta data concerning the SCRDS contract.
var SCRDSMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preDeadline\",\"type\":\"uint256\"}],\"name\":\"DeadlineChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchanger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"preVerifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierAddressChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newDeadline\",\"type\":\"uint256\"}],\"name\":\"__SCRDS_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exchangerNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"name\":\"setDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifierAddress\",\"type\":\"address\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifierAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SCRDSABI is the input ABI used to generate the binding from.
// Deprecated: Use SCRDSMetaData.ABI instead.
var SCRDSABI = SCRDSMetaData.ABI

// SCRDS is an auto generated Go binding around an Ethereum contract.
type SCRDS struct {
	SCRDSCaller     // Read-only binding to the contract
	SCRDSTransactor // Write-only binding to the contract
	SCRDSFilterer   // Log filterer for contract events
}

// SCRDSCaller is an auto generated read-only Go binding around an Ethereum contract.
type SCRDSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRDSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SCRDSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRDSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCRDSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCRDSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCRDSSession struct {
	Contract     *SCRDS            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCRDSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCRDSCallerSession struct {
	Contract *SCRDSCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SCRDSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCRDSTransactorSession struct {
	Contract     *SCRDSTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCRDSRaw is an auto generated low-level Go binding around an Ethereum contract.
type SCRDSRaw struct {
	Contract *SCRDS // Generic contract binding to access the raw methods on
}

// SCRDSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCRDSCallerRaw struct {
	Contract *SCRDSCaller // Generic read-only contract binding to access the raw methods on
}

// SCRDSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCRDSTransactorRaw struct {
	Contract *SCRDSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSCRDS creates a new instance of SCRDS, bound to a specific deployed contract.
func NewSCRDS(address common.Address, backend bind.ContractBackend) (*SCRDS, error) {
	contract, err := bindSCRDS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCRDS{SCRDSCaller: SCRDSCaller{contract: contract}, SCRDSTransactor: SCRDSTransactor{contract: contract}, SCRDSFilterer: SCRDSFilterer{contract: contract}}, nil
}

// NewSCRDSCaller creates a new read-only instance of SCRDS, bound to a specific deployed contract.
func NewSCRDSCaller(address common.Address, caller bind.ContractCaller) (*SCRDSCaller, error) {
	contract, err := bindSCRDS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCRDSCaller{contract: contract}, nil
}

// NewSCRDSTransactor creates a new write-only instance of SCRDS, bound to a specific deployed contract.
func NewSCRDSTransactor(address common.Address, transactor bind.ContractTransactor) (*SCRDSTransactor, error) {
	contract, err := bindSCRDS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCRDSTransactor{contract: contract}, nil
}

// NewSCRDSFilterer creates a new log filterer instance of SCRDS, bound to a specific deployed contract.
func NewSCRDSFilterer(address common.Address, filterer bind.ContractFilterer) (*SCRDSFilterer, error) {
	contract, err := bindSCRDS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCRDSFilterer{contract: contract}, nil
}

// bindSCRDS binds a generic wrapper to an already deployed contract.
func bindSCRDS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SCRDSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCRDS *SCRDSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCRDS.Contract.SCRDSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCRDS *SCRDSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRDS.Contract.SCRDSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCRDS *SCRDSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCRDS.Contract.SCRDSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCRDS *SCRDSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCRDS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCRDS *SCRDSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRDS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCRDS *SCRDSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCRDS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SCRDS *SCRDSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SCRDS *SCRDSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SCRDS.Contract.Allowance(&_SCRDS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SCRDS *SCRDSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SCRDS.Contract.Allowance(&_SCRDS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SCRDS *SCRDSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SCRDS *SCRDSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SCRDS.Contract.BalanceOf(&_SCRDS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SCRDS *SCRDSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SCRDS.Contract.BalanceOf(&_SCRDS.CallOpts, account)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_SCRDS *SCRDSCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_SCRDS *SCRDSSession) Deadline() (*big.Int, error) {
	return _SCRDS.Contract.Deadline(&_SCRDS.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_SCRDS *SCRDSCallerSession) Deadline() (*big.Int, error) {
	return _SCRDS.Contract.Deadline(&_SCRDS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SCRDS *SCRDSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SCRDS *SCRDSSession) Decimals() (uint8, error) {
	return _SCRDS.Contract.Decimals(&_SCRDS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SCRDS *SCRDSCallerSession) Decimals() (uint8, error) {
	return _SCRDS.Contract.Decimals(&_SCRDS.CallOpts)
}

// ExchangerNonces is a free data retrieval call binding the contract method 0x65bb30be.
//
// Solidity: function exchangerNonces(address ) view returns(uint256)
func (_SCRDS *SCRDSCaller) ExchangerNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "exchangerNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangerNonces is a free data retrieval call binding the contract method 0x65bb30be.
//
// Solidity: function exchangerNonces(address ) view returns(uint256)
func (_SCRDS *SCRDSSession) ExchangerNonces(arg0 common.Address) (*big.Int, error) {
	return _SCRDS.Contract.ExchangerNonces(&_SCRDS.CallOpts, arg0)
}

// ExchangerNonces is a free data retrieval call binding the contract method 0x65bb30be.
//
// Solidity: function exchangerNonces(address ) view returns(uint256)
func (_SCRDS *SCRDSCallerSession) ExchangerNonces(arg0 common.Address) (*big.Int, error) {
	return _SCRDS.Contract.ExchangerNonces(&_SCRDS.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCRDS *SCRDSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCRDS *SCRDSSession) Name() (string, error) {
	return _SCRDS.Contract.Name(&_SCRDS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCRDS *SCRDSCallerSession) Name() (string, error) {
	return _SCRDS.Contract.Name(&_SCRDS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCRDS *SCRDSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCRDS *SCRDSSession) Owner() (common.Address, error) {
	return _SCRDS.Contract.Owner(&_SCRDS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCRDS *SCRDSCallerSession) Owner() (common.Address, error) {
	return _SCRDS.Contract.Owner(&_SCRDS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SCRDS *SCRDSCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SCRDS *SCRDSSession) Paused() (bool, error) {
	return _SCRDS.Contract.Paused(&_SCRDS.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SCRDS *SCRDSCallerSession) Paused() (bool, error) {
	return _SCRDS.Contract.Paused(&_SCRDS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCRDS *SCRDSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCRDS *SCRDSSession) Symbol() (string, error) {
	return _SCRDS.Contract.Symbol(&_SCRDS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCRDS *SCRDSCallerSession) Symbol() (string, error) {
	return _SCRDS.Contract.Symbol(&_SCRDS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCRDS *SCRDSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCRDS *SCRDSSession) TotalSupply() (*big.Int, error) {
	return _SCRDS.Contract.TotalSupply(&_SCRDS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCRDS *SCRDSCallerSession) TotalSupply() (*big.Int, error) {
	return _SCRDS.Contract.TotalSupply(&_SCRDS.CallOpts)
}

// VerifierAddress is a free data retrieval call binding the contract method 0x18bdffbb.
//
// Solidity: function verifierAddress() view returns(address)
func (_SCRDS *SCRDSCaller) VerifierAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCRDS.contract.Call(opts, &out, "verifierAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifierAddress is a free data retrieval call binding the contract method 0x18bdffbb.
//
// Solidity: function verifierAddress() view returns(address)
func (_SCRDS *SCRDSSession) VerifierAddress() (common.Address, error) {
	return _SCRDS.Contract.VerifierAddress(&_SCRDS.CallOpts)
}

// VerifierAddress is a free data retrieval call binding the contract method 0x18bdffbb.
//
// Solidity: function verifierAddress() view returns(address)
func (_SCRDS *SCRDSCallerSession) VerifierAddress() (common.Address, error) {
	return _SCRDS.Contract.VerifierAddress(&_SCRDS.CallOpts)
}

// SCRDSInit is a paid mutator transaction binding the contract method 0x81d8167e.
//
// Solidity: function __SCRDS_init(string name, string symbol, address newVerifierAddress, uint256 newDeadline) returns()
func (_SCRDS *SCRDSTransactor) SCRDSInit(opts *bind.TransactOpts, name string, symbol string, newVerifierAddress common.Address, newDeadline *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "__SCRDS_init", name, symbol, newVerifierAddress, newDeadline)
}

// SCRDSInit is a paid mutator transaction binding the contract method 0x81d8167e.
//
// Solidity: function __SCRDS_init(string name, string symbol, address newVerifierAddress, uint256 newDeadline) returns()
func (_SCRDS *SCRDSSession) SCRDSInit(name string, symbol string, newVerifierAddress common.Address, newDeadline *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.SCRDSInit(&_SCRDS.TransactOpts, name, symbol, newVerifierAddress, newDeadline)
}

// SCRDSInit is a paid mutator transaction binding the contract method 0x81d8167e.
//
// Solidity: function __SCRDS_init(string name, string symbol, address newVerifierAddress, uint256 newDeadline) returns()
func (_SCRDS *SCRDSTransactorSession) SCRDSInit(name string, symbol string, newVerifierAddress common.Address, newDeadline *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.SCRDSInit(&_SCRDS.TransactOpts, name, symbol, newVerifierAddress, newDeadline)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SCRDS *SCRDSSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.Approve(&_SCRDS.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.Approve(&_SCRDS.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SCRDS *SCRDSTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SCRDS *SCRDSSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.DecreaseAllowance(&_SCRDS.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SCRDS *SCRDSTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.DecreaseAllowance(&_SCRDS.TransactOpts, spender, subtractedValue)
}

// Exchange is a paid mutator transaction binding the contract method 0x064a181e.
//
// Solidity: function exchange(uint256 amount, bytes signature) returns()
func (_SCRDS *SCRDSTransactor) Exchange(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "exchange", amount, signature)
}

// Exchange is a paid mutator transaction binding the contract method 0x064a181e.
//
// Solidity: function exchange(uint256 amount, bytes signature) returns()
func (_SCRDS *SCRDSSession) Exchange(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SCRDS.Contract.Exchange(&_SCRDS.TransactOpts, amount, signature)
}

// Exchange is a paid mutator transaction binding the contract method 0x064a181e.
//
// Solidity: function exchange(uint256 amount, bytes signature) returns()
func (_SCRDS *SCRDSTransactorSession) Exchange(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SCRDS.Contract.Exchange(&_SCRDS.TransactOpts, amount, signature)
}

// FlipPause is a paid mutator transaction binding the contract method 0x385df649.
//
// Solidity: function flipPause() returns()
func (_SCRDS *SCRDSTransactor) FlipPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "flipPause")
}

// FlipPause is a paid mutator transaction binding the contract method 0x385df649.
//
// Solidity: function flipPause() returns()
func (_SCRDS *SCRDSSession) FlipPause() (*types.Transaction, error) {
	return _SCRDS.Contract.FlipPause(&_SCRDS.TransactOpts)
}

// FlipPause is a paid mutator transaction binding the contract method 0x385df649.
//
// Solidity: function flipPause() returns()
func (_SCRDS *SCRDSTransactorSession) FlipPause() (*types.Transaction, error) {
	return _SCRDS.Contract.FlipPause(&_SCRDS.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SCRDS *SCRDSTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SCRDS *SCRDSSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.IncreaseAllowance(&_SCRDS.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SCRDS *SCRDSTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.IncreaseAllowance(&_SCRDS.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCRDS *SCRDSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCRDS *SCRDSSession) RenounceOwnership() (*types.Transaction, error) {
	return _SCRDS.Contract.RenounceOwnership(&_SCRDS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCRDS *SCRDSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SCRDS.Contract.RenounceOwnership(&_SCRDS.TransactOpts)
}

// SetDeadline is a paid mutator transaction binding the contract method 0x195199f6.
//
// Solidity: function setDeadline(uint256 newTimestamp) returns()
func (_SCRDS *SCRDSTransactor) SetDeadline(opts *bind.TransactOpts, newTimestamp *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "setDeadline", newTimestamp)
}

// SetDeadline is a paid mutator transaction binding the contract method 0x195199f6.
//
// Solidity: function setDeadline(uint256 newTimestamp) returns()
func (_SCRDS *SCRDSSession) SetDeadline(newTimestamp *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.SetDeadline(&_SCRDS.TransactOpts, newTimestamp)
}

// SetDeadline is a paid mutator transaction binding the contract method 0x195199f6.
//
// Solidity: function setDeadline(uint256 newTimestamp) returns()
func (_SCRDS *SCRDSTransactorSession) SetDeadline(newTimestamp *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.SetDeadline(&_SCRDS.TransactOpts, newTimestamp)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address newVerifierAddress) returns()
func (_SCRDS *SCRDSTransactor) SetVerifierAddress(opts *bind.TransactOpts, newVerifierAddress common.Address) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "setVerifierAddress", newVerifierAddress)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address newVerifierAddress) returns()
func (_SCRDS *SCRDSSession) SetVerifierAddress(newVerifierAddress common.Address) (*types.Transaction, error) {
	return _SCRDS.Contract.SetVerifierAddress(&_SCRDS.TransactOpts, newVerifierAddress)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0x17e95526.
//
// Solidity: function setVerifierAddress(address newVerifierAddress) returns()
func (_SCRDS *SCRDSTransactorSession) SetVerifierAddress(newVerifierAddress common.Address) (*types.Transaction, error) {
	return _SCRDS.Contract.SetVerifierAddress(&_SCRDS.TransactOpts, newVerifierAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.Transfer(&_SCRDS.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.Transfer(&_SCRDS.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.TransferFrom(&_SCRDS.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SCRDS *SCRDSTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SCRDS.Contract.TransferFrom(&_SCRDS.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCRDS *SCRDSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SCRDS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCRDS *SCRDSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SCRDS.Contract.TransferOwnership(&_SCRDS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCRDS *SCRDSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SCRDS.Contract.TransferOwnership(&_SCRDS.TransactOpts, newOwner)
}

// SCRDSApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SCRDS contract.
type SCRDSApprovalIterator struct {
	Event *SCRDSApproval // Event containing the contract specifics and raw log

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
func (it *SCRDSApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSApproval)
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
		it.Event = new(SCRDSApproval)
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
func (it *SCRDSApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSApproval represents a Approval event raised by the SCRDS contract.
type SCRDSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SCRDS *SCRDSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SCRDSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SCRDSApprovalIterator{contract: _SCRDS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SCRDS *SCRDSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SCRDSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSApproval)
				if err := _SCRDS.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SCRDS *SCRDSFilterer) ParseApproval(log types.Log) (*SCRDSApproval, error) {
	event := new(SCRDSApproval)
	if err := _SCRDS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSDeadlineChangedIterator is returned from FilterDeadlineChanged and is used to iterate over the raw logs and unpacked data for DeadlineChanged events raised by the SCRDS contract.
type SCRDSDeadlineChangedIterator struct {
	Event *SCRDSDeadlineChanged // Event containing the contract specifics and raw log

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
func (it *SCRDSDeadlineChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSDeadlineChanged)
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
		it.Event = new(SCRDSDeadlineChanged)
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
func (it *SCRDSDeadlineChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSDeadlineChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSDeadlineChanged represents a DeadlineChanged event raised by the SCRDS contract.
type SCRDSDeadlineChanged struct {
	Deadline    *big.Int
	PreDeadline *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeadlineChanged is a free log retrieval operation binding the contract event 0xb50b49938e915b0ab39e99a2d0810e190d6a377647ef06ec45b1e54b892968a6.
//
// Solidity: event DeadlineChanged(uint256 deadline, uint256 preDeadline)
func (_SCRDS *SCRDSFilterer) FilterDeadlineChanged(opts *bind.FilterOpts) (*SCRDSDeadlineChangedIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "DeadlineChanged")
	if err != nil {
		return nil, err
	}
	return &SCRDSDeadlineChangedIterator{contract: _SCRDS.contract, event: "DeadlineChanged", logs: logs, sub: sub}, nil
}

// WatchDeadlineChanged is a free log subscription operation binding the contract event 0xb50b49938e915b0ab39e99a2d0810e190d6a377647ef06ec45b1e54b892968a6.
//
// Solidity: event DeadlineChanged(uint256 deadline, uint256 preDeadline)
func (_SCRDS *SCRDSFilterer) WatchDeadlineChanged(opts *bind.WatchOpts, sink chan<- *SCRDSDeadlineChanged) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "DeadlineChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSDeadlineChanged)
				if err := _SCRDS.contract.UnpackLog(event, "DeadlineChanged", log); err != nil {
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

// ParseDeadlineChanged is a log parse operation binding the contract event 0xb50b49938e915b0ab39e99a2d0810e190d6a377647ef06ec45b1e54b892968a6.
//
// Solidity: event DeadlineChanged(uint256 deadline, uint256 preDeadline)
func (_SCRDS *SCRDSFilterer) ParseDeadlineChanged(log types.Log) (*SCRDSDeadlineChanged, error) {
	event := new(SCRDSDeadlineChanged)
	if err := _SCRDS.contract.UnpackLog(event, "DeadlineChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the SCRDS contract.
type SCRDSExchangeIterator struct {
	Event *SCRDSExchange // Event containing the contract specifics and raw log

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
func (it *SCRDSExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSExchange)
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
		it.Event = new(SCRDSExchange)
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
func (it *SCRDSExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSExchange represents a Exchange event raised by the SCRDS contract.
type SCRDSExchange struct {
	Exchanger common.Address
	Amount    *big.Int
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0x26981b9aefbb0f732b0264bd34c255e831001eb50b06bc85b32cc39e14389721.
//
// Solidity: event Exchange(address exchanger, uint256 amount, uint256 nonce)
func (_SCRDS *SCRDSFilterer) FilterExchange(opts *bind.FilterOpts) (*SCRDSExchangeIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &SCRDSExchangeIterator{contract: _SCRDS.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0x26981b9aefbb0f732b0264bd34c255e831001eb50b06bc85b32cc39e14389721.
//
// Solidity: event Exchange(address exchanger, uint256 amount, uint256 nonce)
func (_SCRDS *SCRDSFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *SCRDSExchange) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSExchange)
				if err := _SCRDS.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// ParseExchange is a log parse operation binding the contract event 0x26981b9aefbb0f732b0264bd34c255e831001eb50b06bc85b32cc39e14389721.
//
// Solidity: event Exchange(address exchanger, uint256 amount, uint256 nonce)
func (_SCRDS *SCRDSFilterer) ParseExchange(log types.Log) (*SCRDSExchange, error) {
	event := new(SCRDSExchange)
	if err := _SCRDS.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SCRDS contract.
type SCRDSInitializedIterator struct {
	Event *SCRDSInitialized // Event containing the contract specifics and raw log

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
func (it *SCRDSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSInitialized)
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
		it.Event = new(SCRDSInitialized)
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
func (it *SCRDSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSInitialized represents a Initialized event raised by the SCRDS contract.
type SCRDSInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCRDS *SCRDSFilterer) FilterInitialized(opts *bind.FilterOpts) (*SCRDSInitializedIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SCRDSInitializedIterator{contract: _SCRDS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCRDS *SCRDSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SCRDSInitialized) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSInitialized)
				if err := _SCRDS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCRDS *SCRDSFilterer) ParseInitialized(log types.Log) (*SCRDSInitialized, error) {
	event := new(SCRDSInitialized)
	if err := _SCRDS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SCRDS contract.
type SCRDSOwnershipTransferredIterator struct {
	Event *SCRDSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SCRDSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSOwnershipTransferred)
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
		it.Event = new(SCRDSOwnershipTransferred)
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
func (it *SCRDSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSOwnershipTransferred represents a OwnershipTransferred event raised by the SCRDS contract.
type SCRDSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SCRDS *SCRDSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SCRDSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SCRDSOwnershipTransferredIterator{contract: _SCRDS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SCRDS *SCRDSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SCRDSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSOwnershipTransferred)
				if err := _SCRDS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SCRDS *SCRDSFilterer) ParseOwnershipTransferred(log types.Log) (*SCRDSOwnershipTransferred, error) {
	event := new(SCRDSOwnershipTransferred)
	if err := _SCRDS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SCRDS contract.
type SCRDSPausedIterator struct {
	Event *SCRDSPaused // Event containing the contract specifics and raw log

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
func (it *SCRDSPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSPaused)
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
		it.Event = new(SCRDSPaused)
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
func (it *SCRDSPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSPaused represents a Paused event raised by the SCRDS contract.
type SCRDSPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SCRDS *SCRDSFilterer) FilterPaused(opts *bind.FilterOpts) (*SCRDSPausedIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SCRDSPausedIterator{contract: _SCRDS.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SCRDS *SCRDSFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SCRDSPaused) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSPaused)
				if err := _SCRDS.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SCRDS *SCRDSFilterer) ParsePaused(log types.Log) (*SCRDSPaused, error) {
	event := new(SCRDSPaused)
	if err := _SCRDS.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SCRDS contract.
type SCRDSTransferIterator struct {
	Event *SCRDSTransfer // Event containing the contract specifics and raw log

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
func (it *SCRDSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSTransfer)
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
		it.Event = new(SCRDSTransfer)
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
func (it *SCRDSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSTransfer represents a Transfer event raised by the SCRDS contract.
type SCRDSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SCRDS *SCRDSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SCRDSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SCRDSTransferIterator{contract: _SCRDS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SCRDS *SCRDSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SCRDSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSTransfer)
				if err := _SCRDS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SCRDS *SCRDSFilterer) ParseTransfer(log types.Log) (*SCRDSTransfer, error) {
	event := new(SCRDSTransfer)
	if err := _SCRDS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SCRDS contract.
type SCRDSUnpausedIterator struct {
	Event *SCRDSUnpaused // Event containing the contract specifics and raw log

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
func (it *SCRDSUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSUnpaused)
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
		it.Event = new(SCRDSUnpaused)
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
func (it *SCRDSUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSUnpaused represents a Unpaused event raised by the SCRDS contract.
type SCRDSUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SCRDS *SCRDSFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SCRDSUnpausedIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SCRDSUnpausedIterator{contract: _SCRDS.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SCRDS *SCRDSFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SCRDSUnpaused) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSUnpaused)
				if err := _SCRDS.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SCRDS *SCRDSFilterer) ParseUnpaused(log types.Log) (*SCRDSUnpaused, error) {
	event := new(SCRDSUnpaused)
	if err := _SCRDS.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCRDSVerifierAddressChangedIterator is returned from FilterVerifierAddressChanged and is used to iterate over the raw logs and unpacked data for VerifierAddressChanged events raised by the SCRDS contract.
type SCRDSVerifierAddressChangedIterator struct {
	Event *SCRDSVerifierAddressChanged // Event containing the contract specifics and raw log

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
func (it *SCRDSVerifierAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCRDSVerifierAddressChanged)
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
		it.Event = new(SCRDSVerifierAddressChanged)
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
func (it *SCRDSVerifierAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCRDSVerifierAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCRDSVerifierAddressChanged represents a VerifierAddressChanged event raised by the SCRDS contract.
type SCRDSVerifierAddressChanged struct {
	VerifierAddress    common.Address
	PreVerifierAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerifierAddressChanged is a free log retrieval operation binding the contract event 0xd2dd85682aa0e840c5ef1783f4dcf6a6821fe7cab63f7099956f5f742934c544.
//
// Solidity: event VerifierAddressChanged(address verifierAddress, address preVerifierAddress)
func (_SCRDS *SCRDSFilterer) FilterVerifierAddressChanged(opts *bind.FilterOpts) (*SCRDSVerifierAddressChangedIterator, error) {

	logs, sub, err := _SCRDS.contract.FilterLogs(opts, "VerifierAddressChanged")
	if err != nil {
		return nil, err
	}
	return &SCRDSVerifierAddressChangedIterator{contract: _SCRDS.contract, event: "VerifierAddressChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierAddressChanged is a free log subscription operation binding the contract event 0xd2dd85682aa0e840c5ef1783f4dcf6a6821fe7cab63f7099956f5f742934c544.
//
// Solidity: event VerifierAddressChanged(address verifierAddress, address preVerifierAddress)
func (_SCRDS *SCRDSFilterer) WatchVerifierAddressChanged(opts *bind.WatchOpts, sink chan<- *SCRDSVerifierAddressChanged) (event.Subscription, error) {

	logs, sub, err := _SCRDS.contract.WatchLogs(opts, "VerifierAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCRDSVerifierAddressChanged)
				if err := _SCRDS.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
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

// ParseVerifierAddressChanged is a log parse operation binding the contract event 0xd2dd85682aa0e840c5ef1783f4dcf6a6821fe7cab63f7099956f5f742934c544.
//
// Solidity: event VerifierAddressChanged(address verifierAddress, address preVerifierAddress)
func (_SCRDS *SCRDSFilterer) ParseVerifierAddressChanged(log types.Log) (*SCRDSVerifierAddressChanged, error) {
	event := new(SCRDSVerifierAddressChanged)
	if err := _SCRDS.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
