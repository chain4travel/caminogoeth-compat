// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// CaminoAdminMetaData contains all meta data concerning the CaminoAdmin contract.
var CaminoAdminMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"DropRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasFee\",\"type\":\"uint256\"}],\"name\":\"GasFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"}],\"name\":\"KycStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"SetRole\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"remove\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"applyKycState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"getBlacklistState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getKycState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getRoles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setBaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"setBlacklistState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CaminoAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use CaminoAdminMetaData.ABI instead.
var CaminoAdminABI = CaminoAdminMetaData.ABI

// CaminoAdmin is an auto generated Go binding around an Ethereum contract.
type CaminoAdmin struct {
	CaminoAdminCaller     // Read-only binding to the contract
	CaminoAdminTransactor // Write-only binding to the contract
	CaminoAdminFilterer   // Log filterer for contract events
}

// CaminoAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type CaminoAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaminoAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CaminoAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaminoAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CaminoAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CaminoAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CaminoAdminSession struct {
	Contract     *CaminoAdmin      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CaminoAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CaminoAdminCallerSession struct {
	Contract *CaminoAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CaminoAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CaminoAdminTransactorSession struct {
	Contract     *CaminoAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CaminoAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type CaminoAdminRaw struct {
	Contract *CaminoAdmin // Generic contract binding to access the raw methods on
}

// CaminoAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CaminoAdminCallerRaw struct {
	Contract *CaminoAdminCaller // Generic read-only contract binding to access the raw methods on
}

// CaminoAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CaminoAdminTransactorRaw struct {
	Contract *CaminoAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaminoAdmin creates a new instance of CaminoAdmin, bound to a specific deployed contract.
func NewCaminoAdmin(address common.Address, backend bind.ContractBackend) (*CaminoAdmin, error) {
	contract, err := bindCaminoAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CaminoAdmin{CaminoAdminCaller: CaminoAdminCaller{contract: contract}, CaminoAdminTransactor: CaminoAdminTransactor{contract: contract}, CaminoAdminFilterer: CaminoAdminFilterer{contract: contract}}, nil
}

// NewCaminoAdminCaller creates a new read-only instance of CaminoAdmin, bound to a specific deployed contract.
func NewCaminoAdminCaller(address common.Address, caller bind.ContractCaller) (*CaminoAdminCaller, error) {
	contract, err := bindCaminoAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CaminoAdminCaller{contract: contract}, nil
}

// NewCaminoAdminTransactor creates a new write-only instance of CaminoAdmin, bound to a specific deployed contract.
func NewCaminoAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*CaminoAdminTransactor, error) {
	contract, err := bindCaminoAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CaminoAdminTransactor{contract: contract}, nil
}

// NewCaminoAdminFilterer creates a new log filterer instance of CaminoAdmin, bound to a specific deployed contract.
func NewCaminoAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*CaminoAdminFilterer, error) {
	contract, err := bindCaminoAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CaminoAdminFilterer{contract: contract}, nil
}

// bindCaminoAdmin binds a generic wrapper to an already deployed contract.
func bindCaminoAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CaminoAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaminoAdmin *CaminoAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CaminoAdmin.Contract.CaminoAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaminoAdmin *CaminoAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.CaminoAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaminoAdmin *CaminoAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.CaminoAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CaminoAdmin *CaminoAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CaminoAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CaminoAdmin *CaminoAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CaminoAdmin *CaminoAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.contract.Transact(opts, method, params...)
}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_CaminoAdmin *CaminoAdminCaller) GetBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CaminoAdmin.contract.Call(opts, &out, "getBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_CaminoAdmin *CaminoAdminSession) GetBaseFee() (*big.Int, error) {
	return _CaminoAdmin.Contract.GetBaseFee(&_CaminoAdmin.CallOpts)
}

// GetBaseFee is a free data retrieval call binding the contract method 0x15e812ad.
//
// Solidity: function getBaseFee() view returns(uint256)
func (_CaminoAdmin *CaminoAdminCallerSession) GetBaseFee() (*big.Int, error) {
	return _CaminoAdmin.Contract.GetBaseFee(&_CaminoAdmin.CallOpts)
}

// GetBlacklistState is a free data retrieval call binding the contract method 0x4de525d9.
//
// Solidity: function getBlacklistState(address account, bytes4 signature) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCaller) GetBlacklistState(opts *bind.CallOpts, account common.Address, signature [4]byte) (*big.Int, error) {
	var out []interface{}
	err := _CaminoAdmin.contract.Call(opts, &out, "getBlacklistState", account, signature)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlacklistState is a free data retrieval call binding the contract method 0x4de525d9.
//
// Solidity: function getBlacklistState(address account, bytes4 signature) view returns(uint256)
func (_CaminoAdmin *CaminoAdminSession) GetBlacklistState(account common.Address, signature [4]byte) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetBlacklistState(&_CaminoAdmin.CallOpts, account, signature)
}

// GetBlacklistState is a free data retrieval call binding the contract method 0x4de525d9.
//
// Solidity: function getBlacklistState(address account, bytes4 signature) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCallerSession) GetBlacklistState(account common.Address, signature [4]byte) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetBlacklistState(&_CaminoAdmin.CallOpts, account, signature)
}

// GetKycState is a free data retrieval call binding the contract method 0xbb03fa1d.
//
// Solidity: function getKycState(address account) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCaller) GetKycState(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CaminoAdmin.contract.Call(opts, &out, "getKycState", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKycState is a free data retrieval call binding the contract method 0xbb03fa1d.
//
// Solidity: function getKycState(address account) view returns(uint256)
func (_CaminoAdmin *CaminoAdminSession) GetKycState(account common.Address) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetKycState(&_CaminoAdmin.CallOpts, account)
}

// GetKycState is a free data retrieval call binding the contract method 0xbb03fa1d.
//
// Solidity: function getKycState(address account) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCallerSession) GetKycState(account common.Address) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetKycState(&_CaminoAdmin.CallOpts, account)
}

// GetRoles is a free data retrieval call binding the contract method 0xce6ccfaf.
//
// Solidity: function getRoles(address addr) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCaller) GetRoles(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CaminoAdmin.contract.Call(opts, &out, "getRoles", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoles is a free data retrieval call binding the contract method 0xce6ccfaf.
//
// Solidity: function getRoles(address addr) view returns(uint256)
func (_CaminoAdmin *CaminoAdminSession) GetRoles(addr common.Address) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetRoles(&_CaminoAdmin.CallOpts, addr)
}

// GetRoles is a free data retrieval call binding the contract method 0xce6ccfaf.
//
// Solidity: function getRoles(address addr) view returns(uint256)
func (_CaminoAdmin *CaminoAdminCallerSession) GetRoles(addr common.Address) (*big.Int, error) {
	return _CaminoAdmin.Contract.GetRoles(&_CaminoAdmin.CallOpts, addr)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address addr, uint256 role) view returns(bool)
func (_CaminoAdmin *CaminoAdminCaller) HasRole(opts *bind.CallOpts, addr common.Address, role *big.Int) (bool, error) {
	var out []interface{}
	err := _CaminoAdmin.contract.Call(opts, &out, "hasRole", addr, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address addr, uint256 role) view returns(bool)
func (_CaminoAdmin *CaminoAdminSession) HasRole(addr common.Address, role *big.Int) (bool, error) {
	return _CaminoAdmin.Contract.HasRole(&_CaminoAdmin.CallOpts, addr, role)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address addr, uint256 role) view returns(bool)
func (_CaminoAdmin *CaminoAdminCallerSession) HasRole(addr common.Address, role *big.Int) (bool, error) {
	return _CaminoAdmin.Contract.HasRole(&_CaminoAdmin.CallOpts, addr, role)
}

// ApplyKycState is a paid mutator transaction binding the contract method 0x9a11b2e8.
//
// Solidity: function applyKycState(address account, bool remove, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminTransactor) ApplyKycState(opts *bind.TransactOpts, account common.Address, remove bool, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "applyKycState", account, remove, state)
}

// ApplyKycState is a paid mutator transaction binding the contract method 0x9a11b2e8.
//
// Solidity: function applyKycState(address account, bool remove, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminSession) ApplyKycState(account common.Address, remove bool, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.ApplyKycState(&_CaminoAdmin.TransactOpts, account, remove, state)
}

// ApplyKycState is a paid mutator transaction binding the contract method 0x9a11b2e8.
//
// Solidity: function applyKycState(address account, bool remove, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) ApplyKycState(account common.Address, remove bool, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.ApplyKycState(&_CaminoAdmin.TransactOpts, account, remove, state)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminTransactor) GrantRole(opts *bind.TransactOpts, addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "grantRole", addr, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminSession) GrantRole(addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.GrantRole(&_CaminoAdmin.TransactOpts, addr, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) GrantRole(addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.GrantRole(&_CaminoAdmin.TransactOpts, addr, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminTransactor) RevokeRole(opts *bind.TransactOpts, addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "revokeRole", addr, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminSession) RevokeRole(addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.RevokeRole(&_CaminoAdmin.TransactOpts, addr, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address addr, uint256 role) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) RevokeRole(addr common.Address, role *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.RevokeRole(&_CaminoAdmin.TransactOpts, addr, role)
}

// SetBaseFee is a paid mutator transaction binding the contract method 0x46860698.
//
// Solidity: function setBaseFee(uint256 newFee) returns()
func (_CaminoAdmin *CaminoAdminTransactor) SetBaseFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "setBaseFee", newFee)
}

// SetBaseFee is a paid mutator transaction binding the contract method 0x46860698.
//
// Solidity: function setBaseFee(uint256 newFee) returns()
func (_CaminoAdmin *CaminoAdminSession) SetBaseFee(newFee *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.SetBaseFee(&_CaminoAdmin.TransactOpts, newFee)
}

// SetBaseFee is a paid mutator transaction binding the contract method 0x46860698.
//
// Solidity: function setBaseFee(uint256 newFee) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) SetBaseFee(newFee *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.SetBaseFee(&_CaminoAdmin.TransactOpts, newFee)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdff0b26.
//
// Solidity: function setBlacklistState(address account, bytes4 signature, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminTransactor) SetBlacklistState(opts *bind.TransactOpts, account common.Address, signature [4]byte, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "setBlacklistState", account, signature, state)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdff0b26.
//
// Solidity: function setBlacklistState(address account, bytes4 signature, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminSession) SetBlacklistState(account common.Address, signature [4]byte, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.SetBlacklistState(&_CaminoAdmin.TransactOpts, account, signature, state)
}

// SetBlacklistState is a paid mutator transaction binding the contract method 0xfdff0b26.
//
// Solidity: function setBlacklistState(address account, bytes4 signature, uint256 state) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) SetBlacklistState(account common.Address, signature [4]byte, state *big.Int) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.SetBlacklistState(&_CaminoAdmin.TransactOpts, account, signature, state)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_CaminoAdmin *CaminoAdminTransactor) Upgrade(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _CaminoAdmin.contract.Transact(opts, "upgrade", newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_CaminoAdmin *CaminoAdminSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.Upgrade(&_CaminoAdmin.TransactOpts, newImplementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address newImplementation) returns()
func (_CaminoAdmin *CaminoAdminTransactorSession) Upgrade(newImplementation common.Address) (*types.Transaction, error) {
	return _CaminoAdmin.Contract.Upgrade(&_CaminoAdmin.TransactOpts, newImplementation)
}

// CaminoAdminDropRoleIterator is returned from FilterDropRole and is used to iterate over the raw logs and unpacked data for DropRole events raised by the CaminoAdmin contract.
type CaminoAdminDropRoleIterator struct {
	Event *CaminoAdminDropRole // Event containing the contract specifics and raw log

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
func (it *CaminoAdminDropRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaminoAdminDropRole)
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
		it.Event = new(CaminoAdminDropRole)
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
func (it *CaminoAdminDropRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaminoAdminDropRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaminoAdminDropRole represents a DropRole event raised by the CaminoAdmin contract.
type CaminoAdminDropRole struct {
	Addr common.Address
	Role *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDropRole is a free log retrieval operation binding the contract event 0xcfa5316bd1be4ceb62f363b0a162f322c33ba870641138cd8600dd4fa603fc3b.
//
// Solidity: event DropRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) FilterDropRole(opts *bind.FilterOpts) (*CaminoAdminDropRoleIterator, error) {

	logs, sub, err := _CaminoAdmin.contract.FilterLogs(opts, "DropRole")
	if err != nil {
		return nil, err
	}
	return &CaminoAdminDropRoleIterator{contract: _CaminoAdmin.contract, event: "DropRole", logs: logs, sub: sub}, nil
}

// WatchDropRole is a free log subscription operation binding the contract event 0xcfa5316bd1be4ceb62f363b0a162f322c33ba870641138cd8600dd4fa603fc3b.
//
// Solidity: event DropRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) WatchDropRole(opts *bind.WatchOpts, sink chan<- *CaminoAdminDropRole) (event.Subscription, error) {

	logs, sub, err := _CaminoAdmin.contract.WatchLogs(opts, "DropRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaminoAdminDropRole)
				if err := _CaminoAdmin.contract.UnpackLog(event, "DropRole", log); err != nil {
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

// ParseDropRole is a log parse operation binding the contract event 0xcfa5316bd1be4ceb62f363b0a162f322c33ba870641138cd8600dd4fa603fc3b.
//
// Solidity: event DropRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) ParseDropRole(log types.Log) (*CaminoAdminDropRole, error) {
	event := new(CaminoAdminDropRole)
	if err := _CaminoAdmin.contract.UnpackLog(event, "DropRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaminoAdminGasFeeSetIterator is returned from FilterGasFeeSet and is used to iterate over the raw logs and unpacked data for GasFeeSet events raised by the CaminoAdmin contract.
type CaminoAdminGasFeeSetIterator struct {
	Event *CaminoAdminGasFeeSet // Event containing the contract specifics and raw log

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
func (it *CaminoAdminGasFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaminoAdminGasFeeSet)
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
		it.Event = new(CaminoAdminGasFeeSet)
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
func (it *CaminoAdminGasFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaminoAdminGasFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaminoAdminGasFeeSet represents a GasFeeSet event raised by the CaminoAdmin contract.
type CaminoAdminGasFeeSet struct {
	NewGasFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasFeeSet is a free log retrieval operation binding the contract event 0x1c053aa9b674900648619554980ac10e913d661c372ca30455e1e4ec0ce44071.
//
// Solidity: event GasFeeSet(uint256 newGasFee)
func (_CaminoAdmin *CaminoAdminFilterer) FilterGasFeeSet(opts *bind.FilterOpts) (*CaminoAdminGasFeeSetIterator, error) {

	logs, sub, err := _CaminoAdmin.contract.FilterLogs(opts, "GasFeeSet")
	if err != nil {
		return nil, err
	}
	return &CaminoAdminGasFeeSetIterator{contract: _CaminoAdmin.contract, event: "GasFeeSet", logs: logs, sub: sub}, nil
}

// WatchGasFeeSet is a free log subscription operation binding the contract event 0x1c053aa9b674900648619554980ac10e913d661c372ca30455e1e4ec0ce44071.
//
// Solidity: event GasFeeSet(uint256 newGasFee)
func (_CaminoAdmin *CaminoAdminFilterer) WatchGasFeeSet(opts *bind.WatchOpts, sink chan<- *CaminoAdminGasFeeSet) (event.Subscription, error) {

	logs, sub, err := _CaminoAdmin.contract.WatchLogs(opts, "GasFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaminoAdminGasFeeSet)
				if err := _CaminoAdmin.contract.UnpackLog(event, "GasFeeSet", log); err != nil {
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

// ParseGasFeeSet is a log parse operation binding the contract event 0x1c053aa9b674900648619554980ac10e913d661c372ca30455e1e4ec0ce44071.
//
// Solidity: event GasFeeSet(uint256 newGasFee)
func (_CaminoAdmin *CaminoAdminFilterer) ParseGasFeeSet(log types.Log) (*CaminoAdminGasFeeSet, error) {
	event := new(CaminoAdminGasFeeSet)
	if err := _CaminoAdmin.contract.UnpackLog(event, "GasFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaminoAdminKycStateChangedIterator is returned from FilterKycStateChanged and is used to iterate over the raw logs and unpacked data for KycStateChanged events raised by the CaminoAdmin contract.
type CaminoAdminKycStateChangedIterator struct {
	Event *CaminoAdminKycStateChanged // Event containing the contract specifics and raw log

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
func (it *CaminoAdminKycStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaminoAdminKycStateChanged)
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
		it.Event = new(CaminoAdminKycStateChanged)
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
func (it *CaminoAdminKycStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaminoAdminKycStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaminoAdminKycStateChanged represents a KycStateChanged event raised by the CaminoAdmin contract.
type CaminoAdminKycStateChanged struct {
	Account  common.Address
	OldState *big.Int
	NewState *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterKycStateChanged is a free log retrieval operation binding the contract event 0xf64784c1c207eed151b4adc53adde03b1c3a4ecad6b8de3a65539d464b3e1add.
//
// Solidity: event KycStateChanged(address indexed account, uint256 oldState, uint256 newState)
func (_CaminoAdmin *CaminoAdminFilterer) FilterKycStateChanged(opts *bind.FilterOpts, account []common.Address) (*CaminoAdminKycStateChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CaminoAdmin.contract.FilterLogs(opts, "KycStateChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &CaminoAdminKycStateChangedIterator{contract: _CaminoAdmin.contract, event: "KycStateChanged", logs: logs, sub: sub}, nil
}

// WatchKycStateChanged is a free log subscription operation binding the contract event 0xf64784c1c207eed151b4adc53adde03b1c3a4ecad6b8de3a65539d464b3e1add.
//
// Solidity: event KycStateChanged(address indexed account, uint256 oldState, uint256 newState)
func (_CaminoAdmin *CaminoAdminFilterer) WatchKycStateChanged(opts *bind.WatchOpts, sink chan<- *CaminoAdminKycStateChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CaminoAdmin.contract.WatchLogs(opts, "KycStateChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaminoAdminKycStateChanged)
				if err := _CaminoAdmin.contract.UnpackLog(event, "KycStateChanged", log); err != nil {
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

// ParseKycStateChanged is a log parse operation binding the contract event 0xf64784c1c207eed151b4adc53adde03b1c3a4ecad6b8de3a65539d464b3e1add.
//
// Solidity: event KycStateChanged(address indexed account, uint256 oldState, uint256 newState)
func (_CaminoAdmin *CaminoAdminFilterer) ParseKycStateChanged(log types.Log) (*CaminoAdminKycStateChanged, error) {
	event := new(CaminoAdminKycStateChanged)
	if err := _CaminoAdmin.contract.UnpackLog(event, "KycStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CaminoAdminSetRoleIterator is returned from FilterSetRole and is used to iterate over the raw logs and unpacked data for SetRole events raised by the CaminoAdmin contract.
type CaminoAdminSetRoleIterator struct {
	Event *CaminoAdminSetRole // Event containing the contract specifics and raw log

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
func (it *CaminoAdminSetRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CaminoAdminSetRole)
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
		it.Event = new(CaminoAdminSetRole)
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
func (it *CaminoAdminSetRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CaminoAdminSetRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CaminoAdminSetRole represents a SetRole event raised by the CaminoAdmin contract.
type CaminoAdminSetRole struct {
	Addr common.Address
	Role *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetRole is a free log retrieval operation binding the contract event 0x385a9c70004a48177c93b74796d77d5ebf7e1248f9e2369624514da454cd01b0.
//
// Solidity: event SetRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) FilterSetRole(opts *bind.FilterOpts) (*CaminoAdminSetRoleIterator, error) {

	logs, sub, err := _CaminoAdmin.contract.FilterLogs(opts, "SetRole")
	if err != nil {
		return nil, err
	}
	return &CaminoAdminSetRoleIterator{contract: _CaminoAdmin.contract, event: "SetRole", logs: logs, sub: sub}, nil
}

// WatchSetRole is a free log subscription operation binding the contract event 0x385a9c70004a48177c93b74796d77d5ebf7e1248f9e2369624514da454cd01b0.
//
// Solidity: event SetRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) WatchSetRole(opts *bind.WatchOpts, sink chan<- *CaminoAdminSetRole) (event.Subscription, error) {

	logs, sub, err := _CaminoAdmin.contract.WatchLogs(opts, "SetRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CaminoAdminSetRole)
				if err := _CaminoAdmin.contract.UnpackLog(event, "SetRole", log); err != nil {
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

// ParseSetRole is a log parse operation binding the contract event 0x385a9c70004a48177c93b74796d77d5ebf7e1248f9e2369624514da454cd01b0.
//
// Solidity: event SetRole(address addr, uint256 role)
func (_CaminoAdmin *CaminoAdminFilterer) ParseSetRole(log types.Log) (*CaminoAdminSetRole, error) {
	event := new(CaminoAdminSetRole)
	if err := _CaminoAdmin.contract.UnpackLog(event, "SetRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
