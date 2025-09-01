// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// IBrevisMarketBidder is an auto generated low-level Go binding around an user-defined struct.
type IBrevisMarketBidder struct {
	Prover common.Address
	Fee    *big.Int
}

// IBrevisMarketFeeParams is an auto generated low-level Go binding around an user-defined struct.
type IBrevisMarketFeeParams struct {
	MaxFee   *big.Int
	MinStake *big.Int
	Deadline uint64
}

// IBrevisMarketProofRequest is an auto generated low-level Go binding around an user-defined struct.
type IBrevisMarketProofRequest struct {
	Nonce              uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
	ImgURL             string
	InputData          [][]byte
	InputURL           string
	Fee                IBrevisMarketFeeParams
}

// IStakingControllerUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakingControllerUnstakeRequest struct {
	Amount        *big.Int
	RequestTime   *big.Int
	ScaleSnapshot *big.Int
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlSession) Owner() (common.Address, error) {
	return _AccessControl.Contract.Owner(&_AccessControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccessControl *AccessControlCallerSession) Owner() (common.Address, error) {
	return _AccessControl.Contract.Owner(&_AccessControl.CallOpts)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCaller) RoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "roleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.RoleMemberCount(&_AccessControl.CallOpts, role)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControl *AccessControlCallerSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControl.Contract.RoleMemberCount(&_AccessControl.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlCaller) RoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "roleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _AccessControl.Contract.RoleMembers(&_AccessControl.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_AccessControl *AccessControlCallerSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _AccessControl.Contract.RoleMembers(&_AccessControl.CallOpts, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactor) GrantRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRoles", role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRoles(&_AccessControl.TransactOpts, role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRoles(&_AccessControl.TransactOpts, role, accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactor) RevokeRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRoles", role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRoles(&_AccessControl.TransactOpts, role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRoles(&_AccessControl.TransactOpts, role, accounts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.TransferOwnership(&_AccessControl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccessControl *AccessControlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.TransferOwnership(&_AccessControl.TransactOpts, newOwner)
}

// AccessControlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccessControl contract.
type AccessControlOwnershipTransferredIterator struct {
	Event *AccessControlOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccessControlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlOwnershipTransferred)
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
		it.Event = new(AccessControlOwnershipTransferred)
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
func (it *AccessControlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlOwnershipTransferred represents a OwnershipTransferred event raised by the AccessControl contract.
type AccessControlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccessControl *AccessControlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccessControlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlOwnershipTransferredIterator{contract: _AccessControl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccessControl *AccessControlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessControlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlOwnershipTransferred)
				if err := _AccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseOwnershipTransferred(log types.Log) (*AccessControlOwnershipTransferred, error) {
	event := new(AccessControlOwnershipTransferred)
	if err := _AccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArraysMetaData contains all meta data concerning the Arrays contract.
var ArraysMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220bf69931d809f116cce0635b9d23ecabb95193b80bc0c50a553237a65f0f06fc564736f6c63430008140033",
}

// ArraysABI is the input ABI used to generate the binding from.
// Deprecated: Use ArraysMetaData.ABI instead.
var ArraysABI = ArraysMetaData.ABI

// ArraysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ArraysMetaData.Bin instead.
var ArraysBin = ArraysMetaData.Bin

// DeployArrays deploys a new Ethereum contract, binding an instance of Arrays to it.
func DeployArrays(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Arrays, error) {
	parsed, err := ArraysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArraysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// Arrays is an auto generated Go binding around an Ethereum contract.
type Arrays struct {
	ArraysCaller     // Read-only binding to the contract
	ArraysTransactor // Write-only binding to the contract
	ArraysFilterer   // Log filterer for contract events
}

// ArraysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArraysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArraysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArraysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraysSession struct {
	Contract     *Arrays           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArraysCallerSession struct {
	Contract *ArraysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArraysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArraysTransactorSession struct {
	Contract     *ArraysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArraysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArraysRaw struct {
	Contract *Arrays // Generic contract binding to access the raw methods on
}

// ArraysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArraysCallerRaw struct {
	Contract *ArraysCaller // Generic read-only contract binding to access the raw methods on
}

// ArraysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArraysTransactorRaw struct {
	Contract *ArraysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrays creates a new instance of Arrays, bound to a specific deployed contract.
func NewArrays(address common.Address, backend bind.ContractBackend) (*Arrays, error) {
	contract, err := bindArrays(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arrays{ArraysCaller: ArraysCaller{contract: contract}, ArraysTransactor: ArraysTransactor{contract: contract}, ArraysFilterer: ArraysFilterer{contract: contract}}, nil
}

// NewArraysCaller creates a new read-only instance of Arrays, bound to a specific deployed contract.
func NewArraysCaller(address common.Address, caller bind.ContractCaller) (*ArraysCaller, error) {
	contract, err := bindArrays(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysCaller{contract: contract}, nil
}

// NewArraysTransactor creates a new write-only instance of Arrays, bound to a specific deployed contract.
func NewArraysTransactor(address common.Address, transactor bind.ContractTransactor) (*ArraysTransactor, error) {
	contract, err := bindArrays(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArraysTransactor{contract: contract}, nil
}

// NewArraysFilterer creates a new log filterer instance of Arrays, bound to a specific deployed contract.
func NewArraysFilterer(address common.Address, filterer bind.ContractFilterer) (*ArraysFilterer, error) {
	contract, err := bindArrays(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArraysFilterer{contract: contract}, nil
}

// bindArrays binds a generic wrapper to an already deployed contract.
func bindArrays(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArraysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.ArraysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.ArraysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arrays *ArraysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arrays.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arrays *ArraysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arrays *ArraysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arrays.Contract.contract.Transact(opts, method, params...)
}

// BrevisMarketMetaData contains all meta data concerning the BrevisMarket contract.
var BrevisMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"_picoVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIStakingController\",\"name\":\"_stakingController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_biddingPhaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_revealPhaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_minMaxFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"MarketBeforeDeadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MarketBidRevealMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketBiddingPhaseEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketBiddingPhaseNotEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketCannotRefundYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeadlineBeforeRevealPhaseEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeadlineMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"MarketDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"MarketDeadlineTooFar\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"MarketFeeExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidProtocolFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIBrevisMarket.ReqStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"MarketInvalidRequestStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidSlashBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidStakingController\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"MarketMaxFeeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"MarketMinStakeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketNoAssignedProverToSlash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNoProtocolFeeToWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"MarketNotExpectedProver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualStake\",\"type\":\"uint256\"}],\"name\":\"MarketProverNotEligible\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketRequestAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketRequestNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketRevealPhaseEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketRevealPhaseNotEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashWindowEnd\",\"type\":\"uint256\"}],\"name\":\"MarketSlashWindowExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"BiddingPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"FeeTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"MinMaxFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisMarket.FeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBrevisMarket.ProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"PicoVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"ProverSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"RevealPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"SlashBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"SlashWindowUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEADLINE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getBidders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winnerFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPlace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secondFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getRequest\",\"outputs\":[{\"internalType\":\"enumIBrevisMarket.ReqStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"_picoVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIStakingController\",\"name\":\"_stakingController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_biddingPhaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_revealPhaseDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_minMaxFee\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMaxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"picoVerifier\",\"outputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisMarket.FeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structIBrevisMarket.ProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumIBrevisMarket.ReqStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisMarket.FeeParams\",\"name\":\"fee\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"bidCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structIBrevisMarket.Bidder\",\"name\":\"winner\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"internalType\":\"structIBrevisMarket.Bidder\",\"name\":\"second\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setBiddingPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinFee\",\"type\":\"uint256\"}],\"name\":\"setMinMaxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"setPicoVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setRevealPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"setSlashBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"setSlashWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingController\",\"outputs\":[{\"internalType\":\"contractIStakingController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604081815234620002655760a082620034a080380380916200002482856200026a565b833981010312620002655781516001600160a01b0380821693909184900362000265576020908181015190838216809203620002655762000067858201620002a4565b60806200007760608401620002a4565b920151600080546001600160a01b03198082163390811784558a51919b9399949694939284167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08b80a36001600255811515806200025b575b620000e5575b89516131e69081620002ba8239f35b86156200024c57508960095416176009558489600b541617600b5587516372f702f360e01b81528681600481895afa9081156200024257889162000201575b50600a8054909a169116908117909855600380546001600160801b0319166001600160401b039093169290921790871b6fffffffffffffffff0000000000000000161790556004908155845163095ea7b360e01b8152908101919091526000196024820152929391929181908390604490829087905af18015620001f757620001b6575b8080808080808998620000d6565b81813d8311620001ef575b620001cd81836200026a565b81010312620001eb575180151503620001e8578080620001a8565b80fd5b5080fd5b503d620001c1565b84513d85823e3d90fd5b90508681813d83116200023a575b6200021b81836200026a565b81010312620002365751818116810362000236578662000124565b8780fd5b503d6200020f565b89513d8a823e3d90fd5b632836ef1560e01b8152600490fd5b50861515620000d0565b600080fd5b601f909101601f19168101906001600160401b038211908210176200028e57604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160401b0382168203620002655756fe608080604052600436101561001357600080fd5b600090813560e01c9081630681e6511461278f575080630a22d68c14612771578063196f0f62146127135780631b80bb3a1461267e5780632cf5d279146122945780632f2ff15d1461224e57806335659fb814612230578063390e226d14612212578063434f967c146120a95780635039cb0b14611ff9578063647846a514611fc5578063668fb6dc14611f0c578063695639b714611ee75780636dd5bd6014611e7f578063713e6a0914611e4b5780637249fbb614611c8e5780637322ae3714611c705780637d9b715814611bc05780637e88b1a0146117ac57806383d860a01461173d5780638894a097146117095780638bb9c5bf146116eb5780638da5cb5b146116b857806391d1485414611654578063955919661461159257806396f6fe911461156a5780639d86698514611417578063b6ba1ca7146113ed578063b844bb001461113d578063c0417e581461109b578063c22d31551461096a578063d2b8f2fc146108c8578063d547741f14610882578063deb9a3a21461081b578063dfc75372146107f0578063e1a45218146107d3578063e30c1fc314610764578063eaf57ad71461069e578063f2fde38b146105d7578063f415ed1414610314578063f628d88d146102bc578063fafe42011461029e5763fb1e61ca146101fa57600080fd5b3461029b57602060031936011261029b576040610100916004358152600c60205220805460018201549160028101549267ffffffffffffffff93846003840154169173ffffffffffffffffffffffffffffffffffffffff60056004860154950154956040519761026d8960ff8416612970565b8160081c16602089015260481c1660408701526060860152608085015260a084015260c083015260e0820152f35b80fd5b503461029b578060031936011261029b576020600454604051908152f35b503461029b57604060031936011261029b5773ffffffffffffffffffffffffffffffffffffffff600660406102ef612825565b936004358152600c602052200191166000526020526020604060002054604051908152f35b503461029b576020806003193601126105d357600435808352600c82526040832060ff81541660048110156105a6576002810361056d575067ffffffffffffffff600382015416928342111561052c576008820173ffffffffffffffffffffffffffffffffffffffff90818154169586156104fb576006549061039782826129ce565b42116104bb575050610419836127106103b7600288015460055490612ae2565b04809885600b5416908b6040518096819582947f3046198c000000000000000000000000000000000000000000000000000000008452600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af180156104b057610481575b508360037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007f6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6965416179055541693604051908152a380f35b8390813d83116104a9575b61049681836127e4565b810103126104a45738610428565b600080fd5b503d61048c565b6040513d8a823e3d90fd5b6044916104c7916129ce565b604051907f2ac29f300000000000000000000000000000000000000000000000000000000082524260048301526024820152fd5b602486604051907f8a0c28630000000000000000000000000000000000000000000000000000000082526004820152fd5b6040517f0c821b2900000000000000000000000000000000000000000000000000000000815242600482015267ffffffffffffffff85166024820152604490fd5b6024906105a4604051917f40d41d890000000000000000000000000000000000000000000000000000000083526004830190612970565bfd5b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b5080fd5b503461029b57602060031936011261029b576105f1612848565b73ffffffffffffffffffffffffffffffffffffffff8083541633810361064f57508116156106255761062290613143565b80f35b60046040517f12c44af2000000000000000000000000000000000000000000000000000000008152fd5b6040517fe4cae21a00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff919091166024820152604490fd5b503461029b57602060031936011261029b576106b8612910565b73ffffffffffffffffffffffffffffffffffffffff82541633810361064f5750600380546fffffffffffffffff0000000000000000604084811b919091167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff831617909255815167ffffffffffffffff91831c82168152921660208301527ffbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda9190819081015b0390a180f35b503461029b57602060031936011261029b5760043573ffffffffffffffffffffffffffffffffffffffff82541633810361064f575060407fbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f91600454908060045582519182526020820152a180f35b503461029b578060031936011261029b5760206040516127108152f35b503461029b578060031936011261029b57602067ffffffffffffffff60035460401c16604051908152f35b503461029b5761082a3661286b565b73ffffffffffffffffffffffffffffffffffffffff929192908183541633810361064f5750825b815181101561087e57806108748461086c6108799486612d5c565b511687612d9f565b612ab5565b610851565b8380f35b503461029b57604060031936011261029b5761089c612825565b73ffffffffffffffffffffffffffffffffffffffff82541633810361064f575061062290600435612e72565b503461029b57602060031936011261029b5760043573ffffffffffffffffffffffffffffffffffffffff82541633810361064f575061271081116109405760407f7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee38891600554908060055582519182526020820152a180f35b60046040517f0b642bcc000000000000000000000000000000000000000000000000000000008152fd5b503461029b576003196020813601126105d35767ffffffffffffffff600435116105d357610120906004353603011261029b576109ac610104600435016129b9565b67ffffffffffffffff4291161115611071576109cd610104600435016129b9565b62278d004201908142116110445767ffffffffffffffff82911611610ff457506109fc610104600435016129b9565b67ffffffffffffffff610a2360035482610a18818316426129ce565b9160401c16906129ce565b911610610fca576004548060c4600435013510610f8d57506004602073ffffffffffffffffffffffffffffffffffffffff600b5416604051928380927fc5f530af0000000000000000000000000000000000000000000000000000000082525afa908115610f82578291610f50575b508060e4600435013510610f135750610aaf6004356004016129b9565b7fffffffffffffffff0000000000000000000000000000000000000000000000006040519160c01b1660208201526024600435013560288201526044600435013560488201526048815280608081011067ffffffffffffffff608083011117610eb3576080810160405260208151910120808252600c6020526040822067ffffffffffffffff815460081c16610ee25773ffffffffffffffffffffffffffffffffffffffff600a54166040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015233602482015230604482015260c460043501356064820152606481528060a081011067ffffffffffffffff60a083011117610eb357610bc69160a08201604052612c89565b80547fffffff0000000000000000000000000000000000000000000000000000000000164260081b68ffffffffffffffff0016173360481b7cffffffffffffffffffffffffffffffffffffffff0000000000000000001617815560043560c4810135600183015560e48101356002830155600382019067ffffffffffffffff90610c5390610104016129b9565b167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008254161790556024600435013560048201556005604460043501359101556040516020815267ffffffffffffffff610cb1600435600401612927565b166020820152602460043501356040820152604460043501356060820152610cf9610ce6606460043501600435600401612a0a565b6101206080850152610140840191612a5a565b608460043501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdd600435360301811215610eaf57600435019067ffffffffffffffff600483013511610eaf57600482013560051b8036036024840113610eab57907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08482030160a08501526004830135815260208082019282010192602481019287915b60048101358310610e545788887f7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd1238980610e168b610de660a460043501600435600401612a0a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152612a5a565b60c4600435013560e083015260e4600435013561010083015267ffffffffffffffff610e4761010460043501612927565b166101208301520390a280f35b9091929394602080610e9c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08960019603018752610e968a60248801612a0a565b90612a5a565b97019301930191939290610d9e565b8580fd5b8480fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b602482604051907f3b2433e10000000000000000000000000000000000000000000000000000000082526004820152fd5b604490604051907fb932475a00000000000000000000000000000000000000000000000000000000825260e4600435013560048301526024820152fd5b90506020813d602011610f7a575b81610f6b602093836127e4565b810103126104a4575138610a92565b3d9150610f5e565b6040513d84823e3d90fd5b604490604051907f4c1beb6200000000000000000000000000000000000000000000000000000000825260c4600435013560048301526024820152fd5b60046040517ffc254531000000000000000000000000000000000000000000000000000000008152fd5b604490611006610104600435016129b9565b9067ffffffffffffffff604051927fb81cc27b0000000000000000000000000000000000000000000000000000000084521660048301526024820152fd5b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b60046040517f1b755fb7000000000000000000000000000000000000000000000000000000008152fd5b503461029b57602060031936011261029b5760043573ffffffffffffffffffffffffffffffffffffffff82541633810361064f575061271081116111135760407ff247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be91600754908060075582519182526020820152a180f35b60046040517fadacaded000000000000000000000000000000000000000000000000000000008152fd5b503461029b5760a060031936011261029b57611157612848565b60243573ffffffffffffffffffffffffffffffffffffffff918282168092036113e95760443567ffffffffffffffff8082168092036104a45760643590811681036104a45783156113bf578593857fffffffffffffffffffffffff0000000000000000000000000000000000000000941684600954161760095584600b548286821617600b55161792604051937f72f702f30000000000000000000000000000000000000000000000000000000085526020948581600481855afa9485156113b4578694898993849861136c575b506044959697169182600a5497881617600a557fffffffffffffffffffffffffffffffff000000000000000000000000000000006fffffffffffffffff00000000000000006003549360401b169216171760035560843560045560405197889586947f095ea7b300000000000000000000000000000000000000000000000000000000865260048601527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602486015216175af1801561136157611329575b5050815416806112f8575061062233613143565b602490604051907f62c926120000000000000000000000000000000000000000000000000000000082526004820152fd5b81813d831161135a575b61133d81836127e4565b810103126113565761134e906129ac565b5038806112e4565b8280fd5b503d611333565b6040513d86823e3d90fd5b935050939490955081813d83116113ad575b61138881836127e4565b810103126113a9575187811681036113a95793859392879190896044611225565b8680fd5b503d61137e565b6040513d89823e3d90fd5b60046040517f2836ef15000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b503461029b57602060031936011261029b5760406020916004358152600183522054604051908152f35b503461029b57602060031936011261029b576040906004358152600c60205220805460405190606082019067ffffffffffffffff9183811083821117610eb3576101a0946114db946115689461153e93604052600183015482526002830154906020830191825280600385015416906040840191825260048501549160058601549360078701549573ffffffffffffffffffffffffffffffffffffffff6114cc600a6114c560088c0161293c565b9a0161293c565b9a6040519d8e60ff8316612970565b8d6020878360081c1691015260481c1660408d01525160608c01525160808b0152511660a089015260c088015260e08701526101008601526101208501906020809173ffffffffffffffffffffffffffffffffffffffff81511684520151910152565b805173ffffffffffffffffffffffffffffffffffffffff1661016084015260200151610180830152565bf35b503461029b578060031936011261029b57602067ffffffffffffffff60035416604051908152f35b503461029b57602060031936011261029b576115ac612848565b73ffffffffffffffffffffffffffffffffffffffff908183541633810361064f5750811690811561162a57600954827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600955167f289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a808380a380f35b60046040517f72085081000000000000000000000000000000000000000000000000000000008152fd5b503461029b57604060031936011261029b576116ae60209173ffffffffffffffffffffffffffffffffffffffff604061168b612825565b926004358152600186522091169060019160005201602052604060002054151590565b6040519015158152f35b503461029b578060031936011261029b5773ffffffffffffffffffffffffffffffffffffffff6020915416604051908152f35b503461029b57602060031936011261029b5761062233600435612e72565b503461029b578060031936011261029b57602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b503461029b57602060031936011261029b5760043573ffffffffffffffffffffffffffffffffffffffff82541633810361064f575060407fc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f91600654908060065582519182526020820152a180f35b503461029b57610120806003193601126105d3576004353661012411611356576117d4612d23565b808352600c602092818452604085209182549260035467ffffffffffffffff96876118176008938261180c818316828c891c16612a99565b9160401c1690612a99565b1680421115611b89575087600384015416804211611b4657508183019373ffffffffffffffffffffffffffffffffffffffff9687865416803303611af7575060ff166004811015611aca578061056d575086600954169860048501549060058601548b3b15611ac657604051927ff39751a8000000000000000000000000000000000000000000000000000000008452600484015260248301528b82610144816101009e8f602460448401375afa8015611abb57611a7a575b50506024908a5b848110611a66575050505060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00835416178255600b8201549285600a8401541615611a5a575b61271061192d60075486612ae2565b04916119398386612af5565b9280611a46575b5050816119a7575b5050836119909161198a8460017f859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac3498600a541694845460481c16930154612af5565b91612c20565b6040519480602487378501523393a3600160025580f35b600b549054604080517fa9fc507b00000000000000000000000000000000000000000000000000000000815291881673ffffffffffffffffffffffffffffffffffffffff1660048301526024820193909352919082908716818b816044810103925af180156104b057611a1b575b80611948565b604090813d8311611a3f575b611a3181836127e4565b810103126113a95738611a15565b503d611a27565b611a519082546129ce565b90553880611940565b6009830154935061191e565b8235868201830155918301916001016118d7565b819b929b11611a8e576040529838806118d0565b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6040513d8e823e3d90fd5b8c80fd5b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b6040517f5186daa400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff919091166004820152336024820152604490fd5b6040517f5cc404fa00000000000000000000000000000000000000000000000000000000815242600482015267ffffffffffffffff919091166024820152604490fd5b604490604051907f53dfa38c0000000000000000000000000000000000000000000000000000000082524260048301526024820152fd5b503461029b57602060031936011261029b57611bda612910565b73ffffffffffffffffffffffffffffffffffffffff82541633810361064f57506003805467ffffffffffffffff9283167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000821681179092556040805193909116835260208301919091527fae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e91908190810161075e565b503461029b578060031936011261029b57602060405162278d008152f35b503461029b57602060031936011261029b57600435611cab612d23565b808252600c602052604082209081549160ff831660048110156105a6578061056d57506003549267ffffffffffffffff93611cfe8580611cf2818516828760081c16612a99565b169260401c16826129ce565b6003840154600185019616918742841015611df1575060015b15611daf575050507ff552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe1839160027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00602093161790818155611d9973ffffffffffffffffffffffffffffffffffffffff9283600a54168489549260481c1690612c20565b5460481c169354604051908152a3600160025580f35b60849350604051927ffc8988cb000000000000000000000000000000000000000000000000000000008452426004850152602484015260448301526064820152fd5b81421180611e3f575b15611e0757506001611d17565b82421180611e1d575b15611d1757506001611d17565b5073ffffffffffffffffffffffffffffffffffffffff60088701541615611e10565b50600786015415611dfa565b503461029b578060031936011261029b57602073ffffffffffffffffffffffffffffffffffffffff600b5416604051908152f35b503461029b57602060031936011261029b5760406080916004358152600c6020522073ffffffffffffffffffffffffffffffffffffffff8060088301541691600b600982015492600a8301541691015491604051938452602084015260408301526060820152f35b503461029b578060031936011261029b57604060075460085482519182526020820152f35b503461029b57602060031936011261029b57611f26612848565b73ffffffffffffffffffffffffffffffffffffffff908183541633810361064f575081811691821561162a576008548015611f9b57611f91817f052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf739460209488600855600a5416612c20565b604051908152a280f35b60046040517f24f94102000000000000000000000000000000000000000000000000000000008152fd5b503461029b578060031936011261029b57602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b503461029b576020908160031936011261029b5790600435825260018082526040832092604051808486549182815201908196845285842090845b818110612096575050508161204a9103826127e4565b6040519380850191818652518092526040850195925b82811061206d5785870386f35b835173ffffffffffffffffffffffffffffffffffffffff16875295810195928101928401612060565b8254845292870192918601918601612034565b503461029b57604060031936011261029b5760043560243590808352600c60205260408320805467ffffffffffffffff90818160081c169081156121e15760ff1660048110156121b4578061056d575061210890826003541690612a99565b1680421161217d575061211f600282015433612b02565b3384526006810160205260408420838154159155612168575b506040519182527fe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e60203393a380f35b6007016121758154612ab5565b905538612138565b604490604051907fc5fed1210000000000000000000000000000000000000000000000000000000082524260048301526024820152fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b602485604051907ff72a3af10000000000000000000000000000000000000000000000000000000082526004820152fd5b503461029b578060031936011261029b576020600654604051908152f35b503461029b578060031936011261029b576020600754604051908152f35b503461029b57604060031936011261029b57612268612825565b73ffffffffffffffffffffffffffffffffffffffff82541633810361064f575061062290600435612d9f565b503461029b57606060031936011261029b57600435602490813591818452602090600c82526040852090815467ffffffffffffffff90818160081c1690811561264e5760ff16600481101561262257806125ec57506003549082821690836122fc8383612a99565b16804211156125b65750918361180c612316938295612a99565b1680421161258057506040518481018781526044356040830152604082526060820192828410908411176125545782604052815190209033895260068501865260408920549180830361252257505050506001820154908186116124ec5750509083826123a860027f9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f95015433612b02565b600881019073ffffffffffffffffffffffffffffffffffffffff91828154161580156124df575b1561246a57600992600a830190828203612430575b505083856040516123f4816127ab565b3381520152337fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905501555b6040519384523393a380f35b8254167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905582820154600b83015538806123e4565b50600a810180549283161580156124d2575b61248a575b50505050612424565b600b92848660405161249b816127ab565b33815201527fffffffffffffffffffffffff0000000000000000000000000000000000000000339116179055015583388080612481565b50600b820154841061247c565b50600982015484106123cf565b6044918691604051927f2365c7240000000000000000000000000000000000000000000000000000000084526004840152820152fd5b7ff83ac0ed00000000000000000000000000000000000000000000000000000000845260648201526084015260449150fd5b83897f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b82604491604051917f680fb097000000000000000000000000000000000000000000000000000000008352426004840152820152fd5b85604491604051917f290cc1aa000000000000000000000000000000000000000000000000000000008352426004840152820152fd5b83906105a4604051917f40d41d890000000000000000000000000000000000000000000000000000000083526004830190612970565b83897f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b8387604051907ff72a3af10000000000000000000000000000000000000000000000000000000082526004820152fd5b503461029b576020908160031936011261029b5790604051916126a0836127c7565b61010080933690376004358152600c8252600c604082200191604051928383905b600882106126fd575050506126d5836127c7565b6040519291835b600882106126e8578585f35b828060019286518152019401910190926126dc565b82548152600192830192919091019083016126c1565b503461029b576127223661286b565b73ffffffffffffffffffffffffffffffffffffffff929192908183541633810361064f5750825b815181101561087e57806108748461276461276c9486612d5c565b511687612e72565b612749565b503461029b578060031936011261029b576020600854604051908152f35b9050346105d357816003193601126105d3576020906005548152f35b6040810190811067ffffffffffffffff821117610eb357604052565b610100810190811067ffffffffffffffff821117610eb357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610eb357604052565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036104a457565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036104a457565b9060406003198301126104a457600435916024359067ffffffffffffffff908183116104a457806023840112156104a4578260040135918211610eb3578160051b604051936020936128bf858401876127e4565b85526024848601928201019283116104a457602401905b8282106128e4575050505090565b813573ffffffffffffffffffffffffffffffffffffffff811681036104a45781529083019083016128d6565b6004359067ffffffffffffffff821682036104a457565b359067ffffffffffffffff821682036104a457565b90604051612949816127ab565b60206001829473ffffffffffffffffffffffffffffffffffffffff81541684520154910152565b90600482101561297d5752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b519081151582036104a457565b3567ffffffffffffffff811681036104a45790565b919082018092116129db57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156104a457016020813591019167ffffffffffffffff82116104a45781360383136104a457565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91909167ffffffffffffffff808094169116019182116129db57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146129db5760010190565b818102929181159184041417156129db57565b919082039182116129db57565b9073ffffffffffffffffffffffffffffffffffffffff9081600b541691604080805180957fad732cc10000000000000000000000000000000000000000000000000000000082528180612b7c888b600484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03915afa938415612c15576000908195612bd9575b5015612b9e575050505050565b6064955051937f8c60a9e300000000000000000000000000000000000000000000000000000000855216600484015260248301526044820152fd5b94508185813d8311612c0e575b612bf081836127e4565b8101031261029b57506020612c04856129ac565b9401519338612b91565b503d612be6565b50513d6000823e3d90fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff929092166024830152604480830193909352918152612c8791612c826064836127e4565b612c89565b565b906000602091828151910182855af115612d17576000513d612d0e575073ffffffffffffffffffffffffffffffffffffffff81163b155b612cc75750565b60249073ffffffffffffffffffffffffffffffffffffffff604051917f5274afe7000000000000000000000000000000000000000000000000000000008352166004820152fd5b60011415612cc0565b6040513d6000823e3d90fd5b6002805414612d325760028055565b60046040517f3ee5aeb5000000000000000000000000000000000000000000000000000000008152fd5b8051821015612d705760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081815260016020526040812092612ddf73ffffffffffffffffffffffffffffffffffffffff8216809560019160005201602052604060002054151590565b612e20575080827f2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f392526001602052612e1b8460408320612f59565b5080a3565b6040517f6b75226d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff919091166004820152602481019290925250604490fd5b600081815260016020526040812092612eb273ffffffffffffffffffffffffffffffffffffffff8216809560019160005201602052604060002054151590565b15612eef575080827f155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a5292526001602052612e1b8460408320612fea565b6040517f5c6b51c900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff919091166004820152602481019290925250604490fd5b8054821015612d705760005260206000200190600090565b6000828152600182016020526040902054612fe35780549068010000000000000000821015610eb35782612fcc612f97846001809601855584612f41565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905580549260005201602052604060002055600190565b5050600090565b9060018201906000928184528260205260408420549081151560001461313c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9182810181811161310f578254908482019182116130e2578181036130ad575b50505080548015613080578201916130638383612f41565b909182549160031b1b191690555582526020526040812055600190565b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526031600452fd5b6130cd6130bd612f979386612f41565b90549060031b1c92839286612f41565b9055865284602052604086205538808061304b565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5050505090565b6000549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a356fea2646970667358221220d4334f144cda553af305c25860ced50923168b346ca8ba2221ccecf71bc5402164736f6c63430008140033",
}

// BrevisMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use BrevisMarketMetaData.ABI instead.
var BrevisMarketABI = BrevisMarketMetaData.ABI

// BrevisMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BrevisMarketMetaData.Bin instead.
var BrevisMarketBin = BrevisMarketMetaData.Bin

// DeployBrevisMarket deploys a new Ethereum contract, binding an instance of BrevisMarket to it.
func DeployBrevisMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _picoVerifier common.Address, _stakingController common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64, _minMaxFee *big.Int) (common.Address, *types.Transaction, *BrevisMarket, error) {
	parsed, err := BrevisMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BrevisMarketBin), backend, _picoVerifier, _stakingController, _biddingPhaseDuration, _revealPhaseDuration, _minMaxFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrevisMarket{BrevisMarketCaller: BrevisMarketCaller{contract: contract}, BrevisMarketTransactor: BrevisMarketTransactor{contract: contract}, BrevisMarketFilterer: BrevisMarketFilterer{contract: contract}}, nil
}

// BrevisMarket is an auto generated Go binding around an Ethereum contract.
type BrevisMarket struct {
	BrevisMarketCaller     // Read-only binding to the contract
	BrevisMarketTransactor // Write-only binding to the contract
	BrevisMarketFilterer   // Log filterer for contract events
}

// BrevisMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrevisMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrevisMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrevisMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrevisMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrevisMarketSession struct {
	Contract     *BrevisMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrevisMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrevisMarketCallerSession struct {
	Contract *BrevisMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BrevisMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrevisMarketTransactorSession struct {
	Contract     *BrevisMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BrevisMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrevisMarketRaw struct {
	Contract *BrevisMarket // Generic contract binding to access the raw methods on
}

// BrevisMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrevisMarketCallerRaw struct {
	Contract *BrevisMarketCaller // Generic read-only contract binding to access the raw methods on
}

// BrevisMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrevisMarketTransactorRaw struct {
	Contract *BrevisMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrevisMarket creates a new instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarket(address common.Address, backend bind.ContractBackend) (*BrevisMarket, error) {
	contract, err := bindBrevisMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrevisMarket{BrevisMarketCaller: BrevisMarketCaller{contract: contract}, BrevisMarketTransactor: BrevisMarketTransactor{contract: contract}, BrevisMarketFilterer: BrevisMarketFilterer{contract: contract}}, nil
}

// NewBrevisMarketCaller creates a new read-only instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketCaller(address common.Address, caller bind.ContractCaller) (*BrevisMarketCaller, error) {
	contract, err := bindBrevisMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketCaller{contract: contract}, nil
}

// NewBrevisMarketTransactor creates a new write-only instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*BrevisMarketTransactor, error) {
	contract, err := bindBrevisMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketTransactor{contract: contract}, nil
}

// NewBrevisMarketFilterer creates a new log filterer instance of BrevisMarket, bound to a specific deployed contract.
func NewBrevisMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*BrevisMarketFilterer, error) {
	contract, err := bindBrevisMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketFilterer{contract: contract}, nil
}

// bindBrevisMarket binds a generic wrapper to an already deployed contract.
func bindBrevisMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrevisMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisMarket *BrevisMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisMarket.Contract.BrevisMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisMarket *BrevisMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisMarket.Contract.BrevisMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisMarket *BrevisMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisMarket.Contract.BrevisMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrevisMarket *BrevisMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrevisMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrevisMarket *BrevisMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrevisMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrevisMarket *BrevisMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrevisMarket.Contract.contract.Transact(opts, method, params...)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) BPSDENOMINATOR() (*big.Int, error) {
	return _BrevisMarket.Contract.BPSDENOMINATOR(&_BrevisMarket.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _BrevisMarket.Contract.BPSDENOMINATOR(&_BrevisMarket.CallOpts)
}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) MAXDEADLINEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "MAX_DEADLINE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) MAXDEADLINEDURATION() (*big.Int, error) {
	return _BrevisMarket.Contract.MAXDEADLINEDURATION(&_BrevisMarket.CallOpts)
}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) MAXDEADLINEDURATION() (*big.Int, error) {
	return _BrevisMarket.Contract.MAXDEADLINEDURATION(&_BrevisMarket.CallOpts)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCaller) BiddingPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "biddingPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketSession) BiddingPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.BiddingPhaseDuration(&_BrevisMarket.CallOpts)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCallerSession) BiddingPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.BiddingPhaseDuration(&_BrevisMarket.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_BrevisMarket *BrevisMarketSession) FeeToken() (common.Address, error) {
	return _BrevisMarket.Contract.FeeToken(&_BrevisMarket.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) FeeToken() (common.Address, error) {
	return _BrevisMarket.Contract.FeeToken(&_BrevisMarket.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_BrevisMarket *BrevisMarketCaller) GetBidHash(opts *bind.CallOpts, reqid [32]byte, prover common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getBidHash", reqid, prover)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_BrevisMarket *BrevisMarketSession) GetBidHash(reqid [32]byte, prover common.Address) ([32]byte, error) {
	return _BrevisMarket.Contract.GetBidHash(&_BrevisMarket.CallOpts, reqid, prover)
}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_BrevisMarket *BrevisMarketCallerSession) GetBidHash(reqid [32]byte, prover common.Address) ([32]byte, error) {
	return _BrevisMarket.Contract.GetBidHash(&_BrevisMarket.CallOpts, reqid, prover)
}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_BrevisMarket *BrevisMarketCaller) GetBidders(opts *bind.CallOpts, reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getBidders", reqid)

	outstruct := new(struct {
		Winner      common.Address
		WinnerFee   *big.Int
		SecondPlace common.Address
		SecondFee   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Winner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WinnerFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondPlace = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SecondFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_BrevisMarket *BrevisMarketSession) GetBidders(reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	return _BrevisMarket.Contract.GetBidders(&_BrevisMarket.CallOpts, reqid)
}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_BrevisMarket *BrevisMarketCallerSession) GetBidders(reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	return _BrevisMarket.Contract.GetBidders(&_BrevisMarket.CallOpts, reqid)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_BrevisMarket *BrevisMarketCaller) GetProof(opts *bind.CallOpts, reqid [32]byte) ([8]*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getProof", reqid)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_BrevisMarket *BrevisMarketSession) GetProof(reqid [32]byte) ([8]*big.Int, error) {
	return _BrevisMarket.Contract.GetProof(&_BrevisMarket.CallOpts, reqid)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_BrevisMarket *BrevisMarketCallerSession) GetProof(reqid [32]byte) ([8]*big.Int, error) {
	return _BrevisMarket.Contract.GetProof(&_BrevisMarket.CallOpts, reqid)
}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_BrevisMarket *BrevisMarketCaller) GetProtocolFeeInfo(opts *bind.CallOpts) (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getProtocolFeeInfo")

	outstruct := new(struct {
		FeeBps  *big.Int
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeBps = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_BrevisMarket *BrevisMarketSession) GetProtocolFeeInfo() (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	return _BrevisMarket.Contract.GetProtocolFeeInfo(&_BrevisMarket.CallOpts)
}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_BrevisMarket *BrevisMarketCallerSession) GetProtocolFeeInfo() (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	return _BrevisMarket.Contract.GetProtocolFeeInfo(&_BrevisMarket.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_BrevisMarket *BrevisMarketCaller) GetRequest(opts *bind.CallOpts, reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "getRequest", reqid)

	outstruct := new(struct {
		Status             uint8
		Timestamp          uint64
		Sender             common.Address
		MaxFee             *big.Int
		MinStake           *big.Int
		Deadline           uint64
		Vk                 [32]byte
		PublicValuesDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinStake = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Vk = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.PublicValuesDigest = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_BrevisMarket *BrevisMarketSession) GetRequest(reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	return _BrevisMarket.Contract.GetRequest(&_BrevisMarket.CallOpts, reqid)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_BrevisMarket *BrevisMarketCallerSession) GetRequest(reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	return _BrevisMarket.Contract.GetRequest(&_BrevisMarket.CallOpts, reqid)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.HasRole(&_BrevisMarket.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BrevisMarket *BrevisMarketCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BrevisMarket.Contract.HasRole(&_BrevisMarket.CallOpts, role, account)
}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) MinMaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "minMaxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) MinMaxFee() (*big.Int, error) {
	return _BrevisMarket.Contract.MinMaxFee(&_BrevisMarket.CallOpts)
}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) MinMaxFee() (*big.Int, error) {
	return _BrevisMarket.Contract.MinMaxFee(&_BrevisMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketSession) Owner() (common.Address, error) {
	return _BrevisMarket.Contract.Owner(&_BrevisMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) Owner() (common.Address, error) {
	return _BrevisMarket.Contract.Owner(&_BrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) PicoVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "picoVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketSession) PicoVerifier() (common.Address, error) {
	return _BrevisMarket.Contract.PicoVerifier(&_BrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) PicoVerifier() (common.Address, error) {
	return _BrevisMarket.Contract.PicoVerifier(&_BrevisMarket.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) ProtocolFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "protocolFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) ProtocolFeeBalance() (*big.Int, error) {
	return _BrevisMarket.Contract.ProtocolFeeBalance(&_BrevisMarket.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) ProtocolFeeBalance() (*big.Int, error) {
	return _BrevisMarket.Contract.ProtocolFeeBalance(&_BrevisMarket.CallOpts)
}

// ProtocolFeeBps is a free data retrieval call binding the contract method 0x35659fb8.
//
// Solidity: function protocolFeeBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) ProtocolFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "protocolFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeBps is a free data retrieval call binding the contract method 0x35659fb8.
//
// Solidity: function protocolFeeBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) ProtocolFeeBps() (*big.Int, error) {
	return _BrevisMarket.Contract.ProtocolFeeBps(&_BrevisMarket.CallOpts)
}

// ProtocolFeeBps is a free data retrieval call binding the contract method 0x35659fb8.
//
// Solidity: function protocolFeeBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) ProtocolFeeBps() (*big.Int, error) {
	return _BrevisMarket.Contract.ProtocolFeeBps(&_BrevisMarket.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, address sender, (uint256,uint256,uint64) fee, bytes32 vk, bytes32 publicValuesDigest, uint256 bidCount, (address,uint256) winner, (address,uint256) second)
func (_BrevisMarket *BrevisMarketCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	Fee                IBrevisMarketFeeParams
	Vk                 [32]byte
	PublicValuesDigest [32]byte
	BidCount           *big.Int
	Winner             IBrevisMarketBidder
	Second             IBrevisMarketBidder
}, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Status             uint8
		Timestamp          uint64
		Sender             common.Address
		Fee                IBrevisMarketFeeParams
		Vk                 [32]byte
		PublicValuesDigest [32]byte
		BidCount           *big.Int
		Winner             IBrevisMarketBidder
		Second             IBrevisMarketBidder
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(IBrevisMarketFeeParams)).(*IBrevisMarketFeeParams)
	outstruct.Vk = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.PublicValuesDigest = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.BidCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[7], new(IBrevisMarketBidder)).(*IBrevisMarketBidder)
	outstruct.Second = *abi.ConvertType(out[8], new(IBrevisMarketBidder)).(*IBrevisMarketBidder)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, address sender, (uint256,uint256,uint64) fee, bytes32 vk, bytes32 publicValuesDigest, uint256 bidCount, (address,uint256) winner, (address,uint256) second)
func (_BrevisMarket *BrevisMarketSession) Requests(arg0 [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	Fee                IBrevisMarketFeeParams
	Vk                 [32]byte
	PublicValuesDigest [32]byte
	BidCount           *big.Int
	Winner             IBrevisMarketBidder
	Second             IBrevisMarketBidder
}, error) {
	return _BrevisMarket.Contract.Requests(&_BrevisMarket.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(uint8 status, uint64 timestamp, address sender, (uint256,uint256,uint64) fee, bytes32 vk, bytes32 publicValuesDigest, uint256 bidCount, (address,uint256) winner, (address,uint256) second)
func (_BrevisMarket *BrevisMarketCallerSession) Requests(arg0 [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	Fee                IBrevisMarketFeeParams
	Vk                 [32]byte
	PublicValuesDigest [32]byte
	BidCount           *big.Int
	Winner             IBrevisMarketBidder
	Second             IBrevisMarketBidder
}, error) {
	return _BrevisMarket.Contract.Requests(&_BrevisMarket.CallOpts, arg0)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCaller) RevealPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "revealPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketSession) RevealPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.RevealPhaseDuration(&_BrevisMarket.CallOpts)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64)
func (_BrevisMarket *BrevisMarketCallerSession) RevealPhaseDuration() (uint64, error) {
	return _BrevisMarket.Contract.RevealPhaseDuration(&_BrevisMarket.CallOpts)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) RoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "roleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BrevisMarket.Contract.RoleMemberCount(&_BrevisMarket.CallOpts, role)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BrevisMarket.Contract.RoleMemberCount(&_BrevisMarket.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketCaller) RoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "roleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _BrevisMarket.Contract.RoleMembers(&_BrevisMarket.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_BrevisMarket *BrevisMarketCallerSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _BrevisMarket.Contract.RoleMembers(&_BrevisMarket.CallOpts, role)
}

// SlashBps is a free data retrieval call binding the contract method 0x0681e651.
//
// Solidity: function slashBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) SlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "slashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashBps is a free data retrieval call binding the contract method 0x0681e651.
//
// Solidity: function slashBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) SlashBps() (*big.Int, error) {
	return _BrevisMarket.Contract.SlashBps(&_BrevisMarket.CallOpts)
}

// SlashBps is a free data retrieval call binding the contract method 0x0681e651.
//
// Solidity: function slashBps() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) SlashBps() (*big.Int, error) {
	return _BrevisMarket.Contract.SlashBps(&_BrevisMarket.CallOpts)
}

// SlashWindow is a free data retrieval call binding the contract method 0x390e226d.
//
// Solidity: function slashWindow() view returns(uint256)
func (_BrevisMarket *BrevisMarketCaller) SlashWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "slashWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashWindow is a free data retrieval call binding the contract method 0x390e226d.
//
// Solidity: function slashWindow() view returns(uint256)
func (_BrevisMarket *BrevisMarketSession) SlashWindow() (*big.Int, error) {
	return _BrevisMarket.Contract.SlashWindow(&_BrevisMarket.CallOpts)
}

// SlashWindow is a free data retrieval call binding the contract method 0x390e226d.
//
// Solidity: function slashWindow() view returns(uint256)
func (_BrevisMarket *BrevisMarketCallerSession) SlashWindow() (*big.Int, error) {
	return _BrevisMarket.Contract.SlashWindow(&_BrevisMarket.CallOpts)
}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address)
func (_BrevisMarket *BrevisMarketCaller) StakingController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrevisMarket.contract.Call(opts, &out, "stakingController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address)
func (_BrevisMarket *BrevisMarketSession) StakingController() (common.Address, error) {
	return _BrevisMarket.Contract.StakingController(&_BrevisMarket.CallOpts)
}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address)
func (_BrevisMarket *BrevisMarketCallerSession) StakingController() (common.Address, error) {
	return _BrevisMarket.Contract.StakingController(&_BrevisMarket.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketTransactor) Bid(opts *bind.TransactOpts, reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "bid", reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Bid(&_BrevisMarket.TransactOpts, reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Bid(&_BrevisMarket.TransactOpts, reqid, bidHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRole(&_BrevisMarket.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRole(&_BrevisMarket.TransactOpts, role, account)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactor) GrantRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "grantRoles", role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.GrantRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// Init is a paid mutator transaction binding the contract method 0xb844bb00.
//
// Solidity: function init(address _picoVerifier, address _stakingController, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration, uint256 _minMaxFee) returns()
func (_BrevisMarket *BrevisMarketTransactor) Init(opts *bind.TransactOpts, _picoVerifier common.Address, _stakingController common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64, _minMaxFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "init", _picoVerifier, _stakingController, _biddingPhaseDuration, _revealPhaseDuration, _minMaxFee)
}

// Init is a paid mutator transaction binding the contract method 0xb844bb00.
//
// Solidity: function init(address _picoVerifier, address _stakingController, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration, uint256 _minMaxFee) returns()
func (_BrevisMarket *BrevisMarketSession) Init(_picoVerifier common.Address, _stakingController common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64, _minMaxFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Init(&_BrevisMarket.TransactOpts, _picoVerifier, _stakingController, _biddingPhaseDuration, _revealPhaseDuration, _minMaxFee)
}

// Init is a paid mutator transaction binding the contract method 0xb844bb00.
//
// Solidity: function init(address _picoVerifier, address _stakingController, uint64 _biddingPhaseDuration, uint64 _revealPhaseDuration, uint256 _minMaxFee) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Init(_picoVerifier common.Address, _stakingController common.Address, _biddingPhaseDuration uint64, _revealPhaseDuration uint64, _minMaxFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Init(&_BrevisMarket.TransactOpts, _picoVerifier, _stakingController, _biddingPhaseDuration, _revealPhaseDuration, _minMaxFee)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactor) Refund(opts *bind.TransactOpts, reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "refund", reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Refund(&_BrevisMarket.TransactOpts, reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Refund(&_BrevisMarket.TransactOpts, reqid)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RenounceRole(&_BrevisMarket.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RenounceRole(&_BrevisMarket.TransactOpts, role)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_BrevisMarket *BrevisMarketTransactor) RequestProof(opts *bind.TransactOpts, req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "requestProof", req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_BrevisMarket *BrevisMarketSession) RequestProof(req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RequestProof(&_BrevisMarket.TransactOpts, req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RequestProof(req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RequestProof(&_BrevisMarket.TransactOpts, req)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketTransactor) Reveal(opts *bind.TransactOpts, reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "reveal", reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Reveal(&_BrevisMarket.TransactOpts, reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Reveal(&_BrevisMarket.TransactOpts, reqid, fee, nonce)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRole(&_BrevisMarket.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRole(&_BrevisMarket.TransactOpts, role, account)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactor) RevokeRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "revokeRoles", role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.RevokeRoles(&_BrevisMarket.TransactOpts, role, accounts)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetBiddingPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setBiddingPhaseDuration", newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetBiddingPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetBiddingPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetMinMaxFee(opts *bind.TransactOpts, newMinFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setMinMaxFee", newMinFee)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_BrevisMarket *BrevisMarketSession) SetMinMaxFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetMinMaxFee(&_BrevisMarket.TransactOpts, newMinFee)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetMinMaxFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetMinMaxFee(&_BrevisMarket.TransactOpts, newMinFee)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetPicoVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setPicoVerifier", newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetPicoVerifier(&_BrevisMarket.TransactOpts, newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetPicoVerifier(&_BrevisMarket.TransactOpts, newVerifier)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetProtocolFeeBps(opts *bind.TransactOpts, newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setProtocolFeeBps", newBps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketSession) SetProtocolFeeBps(newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetProtocolFeeBps(&_BrevisMarket.TransactOpts, newBps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetProtocolFeeBps(newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetProtocolFeeBps(&_BrevisMarket.TransactOpts, newBps)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetRevealPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setRevealPhaseDuration", newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetRevealPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetRevealPhaseDuration(&_BrevisMarket.TransactOpts, newDuration)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetSlashBps(opts *bind.TransactOpts, newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setSlashBps", newBps)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketSession) SetSlashBps(newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetSlashBps(&_BrevisMarket.TransactOpts, newBps)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetSlashBps(newBps *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetSlashBps(&_BrevisMarket.TransactOpts, newBps)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_BrevisMarket *BrevisMarketTransactor) SetSlashWindow(opts *bind.TransactOpts, newWindow *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "setSlashWindow", newWindow)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_BrevisMarket *BrevisMarketSession) SetSlashWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetSlashWindow(&_BrevisMarket.TransactOpts, newWindow)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SetSlashWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SetSlashWindow(&_BrevisMarket.TransactOpts, newWindow)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactor) Slash(opts *bind.TransactOpts, reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "slash", reqid)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketSession) Slash(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Slash(&_BrevisMarket.TransactOpts, reqid)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) Slash(reqid [32]byte) (*types.Transaction, error) {
	return _BrevisMarket.Contract.Slash(&_BrevisMarket.TransactOpts, reqid)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketTransactor) SubmitProof(opts *bind.TransactOpts, reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "submitProof", reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SubmitProof(&_BrevisMarket.TransactOpts, reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _BrevisMarket.Contract.SubmitProof(&_BrevisMarket.TransactOpts, reqid, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.TransferOwnership(&_BrevisMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.TransferOwnership(&_BrevisMarket.TransactOpts, newOwner)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_BrevisMarket *BrevisMarketTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BrevisMarket.contract.Transact(opts, "withdrawProtocolFee", to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_BrevisMarket *BrevisMarketSession) WithdrawProtocolFee(to common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.WithdrawProtocolFee(&_BrevisMarket.TransactOpts, to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_BrevisMarket *BrevisMarketTransactorSession) WithdrawProtocolFee(to common.Address) (*types.Transaction, error) {
	return _BrevisMarket.Contract.WithdrawProtocolFee(&_BrevisMarket.TransactOpts, to)
}

// BrevisMarketBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the BrevisMarket contract.
type BrevisMarketBidRevealedIterator struct {
	Event *BrevisMarketBidRevealed // Event containing the contract specifics and raw log

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
func (it *BrevisMarketBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketBidRevealed)
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
		it.Event = new(BrevisMarketBidRevealed)
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
func (it *BrevisMarketBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketBidRevealed represents a BidRevealed event raised by the BrevisMarket contract.
type BrevisMarketBidRevealed struct {
	Reqid  [32]byte
	Prover common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) FilterBidRevealed(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketBidRevealedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketBidRevealedIterator{contract: _BrevisMarket.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *BrevisMarketBidRevealed, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketBidRevealed)
				if err := _BrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
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

// ParseBidRevealed is a log parse operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_BrevisMarket *BrevisMarketFilterer) ParseBidRevealed(log types.Log) (*BrevisMarketBidRevealed, error) {
	event := new(BrevisMarketBidRevealed)
	if err := _BrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketBiddingPhaseDurationUpdatedIterator is returned from FilterBiddingPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for BiddingPhaseDurationUpdated events raised by the BrevisMarket contract.
type BrevisMarketBiddingPhaseDurationUpdatedIterator struct {
	Event *BrevisMarketBiddingPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketBiddingPhaseDurationUpdated)
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
		it.Event = new(BrevisMarketBiddingPhaseDurationUpdated)
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
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketBiddingPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketBiddingPhaseDurationUpdated represents a BiddingPhaseDurationUpdated event raised by the BrevisMarket contract.
type BrevisMarketBiddingPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBiddingPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) FilterBiddingPhaseDurationUpdated(opts *bind.FilterOpts) (*BrevisMarketBiddingPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketBiddingPhaseDurationUpdatedIterator{contract: _BrevisMarket.contract, event: "BiddingPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchBiddingPhaseDurationUpdated is a free log subscription operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) WatchBiddingPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketBiddingPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketBiddingPhaseDurationUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
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

// ParseBiddingPhaseDurationUpdated is a log parse operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) ParseBiddingPhaseDurationUpdated(log types.Log) (*BrevisMarketBiddingPhaseDurationUpdated, error) {
	event := new(BrevisMarketBiddingPhaseDurationUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketFeeTokenUpdatedIterator is returned from FilterFeeTokenUpdated and is used to iterate over the raw logs and unpacked data for FeeTokenUpdated events raised by the BrevisMarket contract.
type BrevisMarketFeeTokenUpdatedIterator struct {
	Event *BrevisMarketFeeTokenUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketFeeTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketFeeTokenUpdated)
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
		it.Event = new(BrevisMarketFeeTokenUpdated)
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
func (it *BrevisMarketFeeTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketFeeTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketFeeTokenUpdated represents a FeeTokenUpdated event raised by the BrevisMarket contract.
type BrevisMarketFeeTokenUpdated struct {
	OldToken common.Address
	NewToken common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeTokenUpdated is a free log retrieval operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_BrevisMarket *BrevisMarketFilterer) FilterFeeTokenUpdated(opts *bind.FilterOpts, oldToken []common.Address, newToken []common.Address) (*BrevisMarketFeeTokenUpdatedIterator, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "FeeTokenUpdated", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketFeeTokenUpdatedIterator{contract: _BrevisMarket.contract, event: "FeeTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeTokenUpdated is a free log subscription operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_BrevisMarket *BrevisMarketFilterer) WatchFeeTokenUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketFeeTokenUpdated, oldToken []common.Address, newToken []common.Address) (event.Subscription, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "FeeTokenUpdated", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketFeeTokenUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "FeeTokenUpdated", log); err != nil {
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

// ParseFeeTokenUpdated is a log parse operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_BrevisMarket *BrevisMarketFilterer) ParseFeeTokenUpdated(log types.Log) (*BrevisMarketFeeTokenUpdated, error) {
	event := new(BrevisMarketFeeTokenUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "FeeTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketMinMaxFeeUpdatedIterator is returned from FilterMinMaxFeeUpdated and is used to iterate over the raw logs and unpacked data for MinMaxFeeUpdated events raised by the BrevisMarket contract.
type BrevisMarketMinMaxFeeUpdatedIterator struct {
	Event *BrevisMarketMinMaxFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketMinMaxFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketMinMaxFeeUpdated)
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
		it.Event = new(BrevisMarketMinMaxFeeUpdated)
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
func (it *BrevisMarketMinMaxFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketMinMaxFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketMinMaxFeeUpdated represents a MinMaxFeeUpdated event raised by the BrevisMarket contract.
type BrevisMarketMinMaxFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinMaxFeeUpdated is a free log retrieval operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_BrevisMarket *BrevisMarketFilterer) FilterMinMaxFeeUpdated(opts *bind.FilterOpts) (*BrevisMarketMinMaxFeeUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "MinMaxFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketMinMaxFeeUpdatedIterator{contract: _BrevisMarket.contract, event: "MinMaxFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinMaxFeeUpdated is a free log subscription operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_BrevisMarket *BrevisMarketFilterer) WatchMinMaxFeeUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketMinMaxFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "MinMaxFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketMinMaxFeeUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "MinMaxFeeUpdated", log); err != nil {
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

// ParseMinMaxFeeUpdated is a log parse operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_BrevisMarket *BrevisMarketFilterer) ParseMinMaxFeeUpdated(log types.Log) (*BrevisMarketMinMaxFeeUpdated, error) {
	event := new(BrevisMarketMinMaxFeeUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "MinMaxFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the BrevisMarket contract.
type BrevisMarketNewBidIterator struct {
	Event *BrevisMarketNewBid // Event containing the contract specifics and raw log

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
func (it *BrevisMarketNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketNewBid)
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
		it.Event = new(BrevisMarketNewBid)
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
func (it *BrevisMarketNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketNewBid represents a NewBid event raised by the BrevisMarket contract.
type BrevisMarketNewBid struct {
	Reqid   [32]byte
	Prover  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) FilterNewBid(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketNewBidIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketNewBidIterator{contract: _BrevisMarket.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *BrevisMarketNewBid, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketNewBid)
				if err := _BrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_BrevisMarket *BrevisMarketFilterer) ParseNewBid(log types.Log) (*BrevisMarketNewBid, error) {
	event := new(BrevisMarketNewBid)
	if err := _BrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the BrevisMarket contract.
type BrevisMarketNewRequestIterator struct {
	Event *BrevisMarketNewRequest // Event containing the contract specifics and raw log

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
func (it *BrevisMarketNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketNewRequest)
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
		it.Event = new(BrevisMarketNewRequest)
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
func (it *BrevisMarketNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketNewRequest represents a NewRequest event raised by the BrevisMarket contract.
type BrevisMarketNewRequest struct {
	Reqid [32]byte
	Req   IBrevisMarketProofRequest
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) FilterNewRequest(opts *bind.FilterOpts, reqid [][32]byte) (*BrevisMarketNewRequestIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketNewRequestIterator{contract: _BrevisMarket.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *BrevisMarketNewRequest, reqid [][32]byte) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketNewRequest)
				if err := _BrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_BrevisMarket *BrevisMarketFilterer) ParseNewRequest(log types.Log) (*BrevisMarketNewRequest, error) {
	event := new(BrevisMarketNewRequest)
	if err := _BrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BrevisMarket contract.
type BrevisMarketOwnershipTransferredIterator struct {
	Event *BrevisMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BrevisMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketOwnershipTransferred)
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
		it.Event = new(BrevisMarketOwnershipTransferred)
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
func (it *BrevisMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketOwnershipTransferred represents a OwnershipTransferred event raised by the BrevisMarket contract.
type BrevisMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisMarket *BrevisMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BrevisMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketOwnershipTransferredIterator{contract: _BrevisMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BrevisMarket *BrevisMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BrevisMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketOwnershipTransferred)
				if err := _BrevisMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BrevisMarket *BrevisMarketFilterer) ParseOwnershipTransferred(log types.Log) (*BrevisMarketOwnershipTransferred, error) {
	event := new(BrevisMarketOwnershipTransferred)
	if err := _BrevisMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketPicoVerifierUpdatedIterator is returned from FilterPicoVerifierUpdated and is used to iterate over the raw logs and unpacked data for PicoVerifierUpdated events raised by the BrevisMarket contract.
type BrevisMarketPicoVerifierUpdatedIterator struct {
	Event *BrevisMarketPicoVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketPicoVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketPicoVerifierUpdated)
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
		it.Event = new(BrevisMarketPicoVerifierUpdated)
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
func (it *BrevisMarketPicoVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketPicoVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketPicoVerifierUpdated represents a PicoVerifierUpdated event raised by the BrevisMarket contract.
type BrevisMarketPicoVerifierUpdated struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPicoVerifierUpdated is a free log retrieval operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) FilterPicoVerifierUpdated(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*BrevisMarketPicoVerifierUpdatedIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketPicoVerifierUpdatedIterator{contract: _BrevisMarket.contract, event: "PicoVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchPicoVerifierUpdated is a free log subscription operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) WatchPicoVerifierUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketPicoVerifierUpdated, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketPicoVerifierUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
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

// ParsePicoVerifierUpdated is a log parse operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_BrevisMarket *BrevisMarketFilterer) ParsePicoVerifierUpdated(log types.Log) (*BrevisMarketPicoVerifierUpdated, error) {
	event := new(BrevisMarketPicoVerifierUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the BrevisMarket contract.
type BrevisMarketProofSubmittedIterator struct {
	Event *BrevisMarketProofSubmitted // Event containing the contract specifics and raw log

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
func (it *BrevisMarketProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketProofSubmitted)
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
		it.Event = new(BrevisMarketProofSubmitted)
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
func (it *BrevisMarketProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketProofSubmitted represents a ProofSubmitted event raised by the BrevisMarket contract.
type BrevisMarketProofSubmitted struct {
	Reqid     [32]byte
	Prover    common.Address
	Proof     [8]*big.Int
	ActualFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_BrevisMarket *BrevisMarketFilterer) FilterProofSubmitted(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketProofSubmittedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketProofSubmittedIterator{contract: _BrevisMarket.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_BrevisMarket *BrevisMarketFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *BrevisMarketProofSubmitted, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketProofSubmitted)
				if err := _BrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_BrevisMarket *BrevisMarketFilterer) ParseProofSubmitted(log types.Log) (*BrevisMarketProofSubmitted, error) {
	event := new(BrevisMarketProofSubmitted)
	if err := _BrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketProtocolFeeBpsUpdatedIterator is returned from FilterProtocolFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeBpsUpdated events raised by the BrevisMarket contract.
type BrevisMarketProtocolFeeBpsUpdatedIterator struct {
	Event *BrevisMarketProtocolFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketProtocolFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketProtocolFeeBpsUpdated)
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
		it.Event = new(BrevisMarketProtocolFeeBpsUpdated)
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
func (it *BrevisMarketProtocolFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketProtocolFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketProtocolFeeBpsUpdated represents a ProtocolFeeBpsUpdated event raised by the BrevisMarket contract.
type BrevisMarketProtocolFeeBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeBpsUpdated is a free log retrieval operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) FilterProtocolFeeBpsUpdated(opts *bind.FilterOpts) (*BrevisMarketProtocolFeeBpsUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "ProtocolFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketProtocolFeeBpsUpdatedIterator{contract: _BrevisMarket.contract, event: "ProtocolFeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeBpsUpdated is a free log subscription operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) WatchProtocolFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketProtocolFeeBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "ProtocolFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketProtocolFeeBpsUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "ProtocolFeeBpsUpdated", log); err != nil {
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

// ParseProtocolFeeBpsUpdated is a log parse operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) ParseProtocolFeeBpsUpdated(log types.Log) (*BrevisMarketProtocolFeeBpsUpdated, error) {
	event := new(BrevisMarketProtocolFeeBpsUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "ProtocolFeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketProtocolFeeWithdrawnIterator is returned from FilterProtocolFeeWithdrawn and is used to iterate over the raw logs and unpacked data for ProtocolFeeWithdrawn events raised by the BrevisMarket contract.
type BrevisMarketProtocolFeeWithdrawnIterator struct {
	Event *BrevisMarketProtocolFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *BrevisMarketProtocolFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketProtocolFeeWithdrawn)
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
		it.Event = new(BrevisMarketProtocolFeeWithdrawn)
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
func (it *BrevisMarketProtocolFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketProtocolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketProtocolFeeWithdrawn represents a ProtocolFeeWithdrawn event raised by the BrevisMarket contract.
type BrevisMarketProtocolFeeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeWithdrawn is a free log retrieval operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) FilterProtocolFeeWithdrawn(opts *bind.FilterOpts, to []common.Address) (*BrevisMarketProtocolFeeWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "ProtocolFeeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketProtocolFeeWithdrawnIterator{contract: _BrevisMarket.contract, event: "ProtocolFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeWithdrawn is a free log subscription operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) WatchProtocolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BrevisMarketProtocolFeeWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "ProtocolFeeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketProtocolFeeWithdrawn)
				if err := _BrevisMarket.contract.UnpackLog(event, "ProtocolFeeWithdrawn", log); err != nil {
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

// ParseProtocolFeeWithdrawn is a log parse operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) ParseProtocolFeeWithdrawn(log types.Log) (*BrevisMarketProtocolFeeWithdrawn, error) {
	event := new(BrevisMarketProtocolFeeWithdrawn)
	if err := _BrevisMarket.contract.UnpackLog(event, "ProtocolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketProverSlashedIterator is returned from FilterProverSlashed and is used to iterate over the raw logs and unpacked data for ProverSlashed events raised by the BrevisMarket contract.
type BrevisMarketProverSlashedIterator struct {
	Event *BrevisMarketProverSlashed // Event containing the contract specifics and raw log

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
func (it *BrevisMarketProverSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketProverSlashed)
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
		it.Event = new(BrevisMarketProverSlashed)
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
func (it *BrevisMarketProverSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketProverSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketProverSlashed represents a ProverSlashed event raised by the BrevisMarket contract.
type BrevisMarketProverSlashed struct {
	Reqid       [32]byte
	Prover      common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProverSlashed is a free log retrieval operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_BrevisMarket *BrevisMarketFilterer) FilterProverSlashed(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*BrevisMarketProverSlashedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "ProverSlashed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketProverSlashedIterator{contract: _BrevisMarket.contract, event: "ProverSlashed", logs: logs, sub: sub}, nil
}

// WatchProverSlashed is a free log subscription operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_BrevisMarket *BrevisMarketFilterer) WatchProverSlashed(opts *bind.WatchOpts, sink chan<- *BrevisMarketProverSlashed, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "ProverSlashed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketProverSlashed)
				if err := _BrevisMarket.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
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

// ParseProverSlashed is a log parse operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_BrevisMarket *BrevisMarketFilterer) ParseProverSlashed(log types.Log) (*BrevisMarketProverSlashed, error) {
	event := new(BrevisMarketProverSlashed)
	if err := _BrevisMarket.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the BrevisMarket contract.
type BrevisMarketRefundedIterator struct {
	Event *BrevisMarketRefunded // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRefunded)
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
		it.Event = new(BrevisMarketRefunded)
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
func (it *BrevisMarketRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRefunded represents a Refunded event raised by the BrevisMarket contract.
type BrevisMarketRefunded struct {
	Reqid     [32]byte
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) FilterRefunded(opts *bind.FilterOpts, reqid [][32]byte, requester []common.Address) (*BrevisMarketRefundedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRefundedIterator{contract: _BrevisMarket.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *BrevisMarketRefunded, reqid [][32]byte, requester []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRefunded)
				if err := _BrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_BrevisMarket *BrevisMarketFilterer) ParseRefunded(log types.Log) (*BrevisMarketRefunded, error) {
	event := new(BrevisMarketRefunded)
	if err := _BrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRevealPhaseDurationUpdatedIterator is returned from FilterRevealPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for RevealPhaseDurationUpdated events raised by the BrevisMarket contract.
type BrevisMarketRevealPhaseDurationUpdatedIterator struct {
	Event *BrevisMarketRevealPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRevealPhaseDurationUpdated)
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
		it.Event = new(BrevisMarketRevealPhaseDurationUpdated)
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
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRevealPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRevealPhaseDurationUpdated represents a RevealPhaseDurationUpdated event raised by the BrevisMarket contract.
type BrevisMarketRevealPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) FilterRevealPhaseDurationUpdated(opts *bind.FilterOpts) (*BrevisMarketRevealPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRevealPhaseDurationUpdatedIterator{contract: _BrevisMarket.contract, event: "RevealPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRevealPhaseDurationUpdated is a free log subscription operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) WatchRevealPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketRevealPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRevealPhaseDurationUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
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

// ParseRevealPhaseDurationUpdated is a log parse operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_BrevisMarket *BrevisMarketFilterer) ParseRevealPhaseDurationUpdated(log types.Log) (*BrevisMarketRevealPhaseDurationUpdated, error) {
	event := new(BrevisMarketRevealPhaseDurationUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BrevisMarket contract.
type BrevisMarketRoleGrantedIterator struct {
	Event *BrevisMarketRoleGranted // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRoleGranted)
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
		it.Event = new(BrevisMarketRoleGranted)
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
func (it *BrevisMarketRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRoleGranted represents a RoleGranted event raised by the BrevisMarket contract.
type BrevisMarketRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*BrevisMarketRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRoleGrantedIterator{contract: _BrevisMarket.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BrevisMarketRoleGranted, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRoleGranted)
				if err := _BrevisMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) ParseRoleGranted(log types.Log) (*BrevisMarketRoleGranted, error) {
	event := new(BrevisMarketRoleGranted)
	if err := _BrevisMarket.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BrevisMarket contract.
type BrevisMarketRoleRevokedIterator struct {
	Event *BrevisMarketRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BrevisMarketRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketRoleRevoked)
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
		it.Event = new(BrevisMarketRoleRevoked)
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
func (it *BrevisMarketRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketRoleRevoked represents a RoleRevoked event raised by the BrevisMarket contract.
type BrevisMarketRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*BrevisMarketRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BrevisMarketRoleRevokedIterator{contract: _BrevisMarket.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BrevisMarketRoleRevoked, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketRoleRevoked)
				if err := _BrevisMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_BrevisMarket *BrevisMarketFilterer) ParseRoleRevoked(log types.Log) (*BrevisMarketRoleRevoked, error) {
	event := new(BrevisMarketRoleRevoked)
	if err := _BrevisMarket.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketSlashBpsUpdatedIterator is returned from FilterSlashBpsUpdated and is used to iterate over the raw logs and unpacked data for SlashBpsUpdated events raised by the BrevisMarket contract.
type BrevisMarketSlashBpsUpdatedIterator struct {
	Event *BrevisMarketSlashBpsUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketSlashBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketSlashBpsUpdated)
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
		it.Event = new(BrevisMarketSlashBpsUpdated)
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
func (it *BrevisMarketSlashBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketSlashBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketSlashBpsUpdated represents a SlashBpsUpdated event raised by the BrevisMarket contract.
type BrevisMarketSlashBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashBpsUpdated is a free log retrieval operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) FilterSlashBpsUpdated(opts *bind.FilterOpts) (*BrevisMarketSlashBpsUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "SlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketSlashBpsUpdatedIterator{contract: _BrevisMarket.contract, event: "SlashBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashBpsUpdated is a free log subscription operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) WatchSlashBpsUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketSlashBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "SlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketSlashBpsUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "SlashBpsUpdated", log); err != nil {
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

// ParseSlashBpsUpdated is a log parse operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_BrevisMarket *BrevisMarketFilterer) ParseSlashBpsUpdated(log types.Log) (*BrevisMarketSlashBpsUpdated, error) {
	event := new(BrevisMarketSlashBpsUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "SlashBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrevisMarketSlashWindowUpdatedIterator is returned from FilterSlashWindowUpdated and is used to iterate over the raw logs and unpacked data for SlashWindowUpdated events raised by the BrevisMarket contract.
type BrevisMarketSlashWindowUpdatedIterator struct {
	Event *BrevisMarketSlashWindowUpdated // Event containing the contract specifics and raw log

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
func (it *BrevisMarketSlashWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrevisMarketSlashWindowUpdated)
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
		it.Event = new(BrevisMarketSlashWindowUpdated)
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
func (it *BrevisMarketSlashWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrevisMarketSlashWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrevisMarketSlashWindowUpdated represents a SlashWindowUpdated event raised by the BrevisMarket contract.
type BrevisMarketSlashWindowUpdated struct {
	OldWindow *big.Int
	NewWindow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashWindowUpdated is a free log retrieval operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_BrevisMarket *BrevisMarketFilterer) FilterSlashWindowUpdated(opts *bind.FilterOpts) (*BrevisMarketSlashWindowUpdatedIterator, error) {

	logs, sub, err := _BrevisMarket.contract.FilterLogs(opts, "SlashWindowUpdated")
	if err != nil {
		return nil, err
	}
	return &BrevisMarketSlashWindowUpdatedIterator{contract: _BrevisMarket.contract, event: "SlashWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashWindowUpdated is a free log subscription operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_BrevisMarket *BrevisMarketFilterer) WatchSlashWindowUpdated(opts *bind.WatchOpts, sink chan<- *BrevisMarketSlashWindowUpdated) (event.Subscription, error) {

	logs, sub, err := _BrevisMarket.contract.WatchLogs(opts, "SlashWindowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrevisMarketSlashWindowUpdated)
				if err := _BrevisMarket.contract.UnpackLog(event, "SlashWindowUpdated", log); err != nil {
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

// ParseSlashWindowUpdated is a log parse operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_BrevisMarket *BrevisMarketFilterer) ParseSlashWindowUpdated(log types.Log) (*BrevisMarketSlashWindowUpdated, error) {
	event := new(BrevisMarketSlashWindowUpdated)
	if err := _BrevisMarket.contract.UnpackLog(event, "SlashWindowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComparatorsMetaData contains all meta data concerning the Comparators contract.
var ComparatorsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220256b03ff945900a1ae20d2f332058e96fa9a52fa621ba3b202757b1b27dace1564736f6c63430008140033",
}

// ComparatorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ComparatorsMetaData.ABI instead.
var ComparatorsABI = ComparatorsMetaData.ABI

// ComparatorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ComparatorsMetaData.Bin instead.
var ComparatorsBin = ComparatorsMetaData.Bin

// DeployComparators deploys a new Ethereum contract, binding an instance of Comparators to it.
func DeployComparators(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Comparators, error) {
	parsed, err := ComparatorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ComparatorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Comparators{ComparatorsCaller: ComparatorsCaller{contract: contract}, ComparatorsTransactor: ComparatorsTransactor{contract: contract}, ComparatorsFilterer: ComparatorsFilterer{contract: contract}}, nil
}

// Comparators is an auto generated Go binding around an Ethereum contract.
type Comparators struct {
	ComparatorsCaller     // Read-only binding to the contract
	ComparatorsTransactor // Write-only binding to the contract
	ComparatorsFilterer   // Log filterer for contract events
}

// ComparatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComparatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComparatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComparatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComparatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComparatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComparatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComparatorsSession struct {
	Contract     *Comparators      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComparatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComparatorsCallerSession struct {
	Contract *ComparatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ComparatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComparatorsTransactorSession struct {
	Contract     *ComparatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComparatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComparatorsRaw struct {
	Contract *Comparators // Generic contract binding to access the raw methods on
}

// ComparatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComparatorsCallerRaw struct {
	Contract *ComparatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ComparatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComparatorsTransactorRaw struct {
	Contract *ComparatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComparators creates a new instance of Comparators, bound to a specific deployed contract.
func NewComparators(address common.Address, backend bind.ContractBackend) (*Comparators, error) {
	contract, err := bindComparators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comparators{ComparatorsCaller: ComparatorsCaller{contract: contract}, ComparatorsTransactor: ComparatorsTransactor{contract: contract}, ComparatorsFilterer: ComparatorsFilterer{contract: contract}}, nil
}

// NewComparatorsCaller creates a new read-only instance of Comparators, bound to a specific deployed contract.
func NewComparatorsCaller(address common.Address, caller bind.ContractCaller) (*ComparatorsCaller, error) {
	contract, err := bindComparators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComparatorsCaller{contract: contract}, nil
}

// NewComparatorsTransactor creates a new write-only instance of Comparators, bound to a specific deployed contract.
func NewComparatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ComparatorsTransactor, error) {
	contract, err := bindComparators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComparatorsTransactor{contract: contract}, nil
}

// NewComparatorsFilterer creates a new log filterer instance of Comparators, bound to a specific deployed contract.
func NewComparatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ComparatorsFilterer, error) {
	contract, err := bindComparators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComparatorsFilterer{contract: contract}, nil
}

// bindComparators binds a generic wrapper to an already deployed contract.
func bindComparators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComparatorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comparators *ComparatorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comparators.Contract.ComparatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comparators *ComparatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comparators.Contract.ComparatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comparators *ComparatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comparators.Contract.ComparatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comparators *ComparatorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comparators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comparators *ComparatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comparators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comparators *ComparatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comparators.Contract.contract.Transact(opts, method, params...)
}

// EnumerableMapMetaData contains all meta data concerning the EnumerableMap contract.
var EnumerableMapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"}]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212201f61d607e7a0afe07a36f48f49bbfa74f700393a2f6acfc1c3c7ad375dda909264736f6c63430008140033",
}

// EnumerableMapABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableMapMetaData.ABI instead.
var EnumerableMapABI = EnumerableMapMetaData.ABI

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableMapMetaData.Bin instead.
var EnumerableMapBin = EnumerableMapMetaData.Bin

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220f0c921e37a39e124815cbeb94225b12c5ee9a4b810f49dcea32ebaf281e6aea064736f6c63430008140033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountAlreadyHasRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlAccountDoesNotHaveRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedRole\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"roleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControl *IAccessControlCaller) RoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "roleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControl *IAccessControlSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControl.Contract.RoleMemberCount(&_IAccessControl.CallOpts, role)
}

// RoleMemberCount is a free data retrieval call binding the contract method 0xb6ba1ca7.
//
// Solidity: function roleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControl *IAccessControlCallerSession) RoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControl.Contract.RoleMemberCount(&_IAccessControl.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_IAccessControl *IAccessControlCaller) RoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "roleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_IAccessControl *IAccessControlSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _IAccessControl.Contract.RoleMembers(&_IAccessControl.CallOpts, role)
}

// RoleMembers is a free data retrieval call binding the contract method 0x5039cb0b.
//
// Solidity: function roleMembers(bytes32 role) view returns(address[] accounts)
func (_IAccessControl *IAccessControlCallerSession) RoleMembers(role [32]byte) ([]common.Address, error) {
	return _IAccessControl.Contract.RoleMembers(&_IAccessControl.CallOpts, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRoles", role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRoles(&_IAccessControl.TransactOpts, role, accounts)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xdeb9a3a2.
//
// Solidity: function grantRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRoles(&_IAccessControl.TransactOpts, role, accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRoles(opts *bind.TransactOpts, role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRoles", role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRoles(&_IAccessControl.TransactOpts, role, accounts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x196f0f62.
//
// Solidity: function revokeRoles(bytes32 role, address[] accounts) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRoles(role [32]byte, accounts []common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRoles(&_IAccessControl.TransactOpts, role, accounts)
}

// IAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControl contract.
type IAccessControlRoleGrantedIterator struct {
	Event *IAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleGranted)
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
		it.Event = new(IAccessControlRoleGranted)
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
func (it *IAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleGranted represents a RoleGranted event raised by the IAccessControl contract.
type IAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*IAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleGrantedIterator{contract: _IAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleGranted, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleGranted)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2ae6a113c0ed5b78a53413ffbb7679881f11145ccfba4fb92e863dfcd5a1d2f3.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) ParseRoleGranted(log types.Log) (*IAccessControlRoleGranted, error) {
	event := new(IAccessControlRoleGranted)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControl contract.
type IAccessControlRoleRevokedIterator struct {
	Event *IAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleRevoked)
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
		it.Event = new(IAccessControlRoleRevoked)
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
func (it *IAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleRevoked represents a RoleRevoked event raised by the IAccessControl contract.
type IAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address) (*IAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleRevokedIterator{contract: _IAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleRevoked, role [][32]byte, account []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleRevoked)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0x155aaafb6329a2098580462df33ec4b7441b19729b9601c5fc17ae1cf99a8a52.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account)
func (_IAccessControl *IAccessControlFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlRoleRevoked, error) {
	event := new(IAccessControlRoleRevoked)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketMetaData contains all meta data concerning the IBrevisMarket contract.
var IBrevisMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"MarketBeforeDeadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"MarketBidRevealMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketBiddingPhaseEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketBiddingPhaseNotEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketCannotRefundYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeadlineBeforeRevealPhaseEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketDeadlineMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"MarketDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAllowed\",\"type\":\"uint256\"}],\"name\":\"MarketDeadlineTooFar\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"MarketFeeExceedsMaximum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidProtocolFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumIBrevisMarket.ReqStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"MarketInvalidRequestStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidSlashBps\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketInvalidStakingController\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"MarketMaxFeeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"MarketMinStakeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketNoAssignedProverToSlash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNoProtocolFeeToWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"MarketNotExpectedProver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualStake\",\"type\":\"uint256\"}],\"name\":\"MarketProverNotEligible\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketRequestAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"MarketRequestNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketRevealPhaseEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEndTime\",\"type\":\"uint256\"}],\"name\":\"MarketRevealPhaseNotEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashWindowEnd\",\"type\":\"uint256\"}],\"name\":\"MarketSlashWindowExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"BiddingPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"FeeTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"MinMaxFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisMarket.FeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBrevisMarket.ProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"PicoVerifierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"}],\"name\":\"ProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"}],\"name\":\"ProverSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDuration\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"RevealPhaseDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"SlashBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"SlashWindowUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_DEADLINE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getBidders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"winnerFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"secondPlace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secondFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"getRequest\",\"outputs\":[{\"internalType\":\"enumIBrevisMarket.ReqStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMaxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"picoVerifier\",\"outputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"verifier\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"vk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesDigest\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"imgURL\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"inputData\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"inputURL\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"}],\"internalType\":\"structIBrevisMarket.FeeParams\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structIBrevisMarket.ProofRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealPhaseDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setBiddingPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinFee\",\"type\":\"uint256\"}],\"name\":\"setMinMaxFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPicoVerifier\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"setPicoVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"setProtocolFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDuration\",\"type\":\"uint64\"}],\"name\":\"setRevealPhaseDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBps\",\"type\":\"uint256\"}],\"name\":\"setSlashBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"setSlashWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingController\",\"outputs\":[{\"internalType\":\"contractIStakingController\",\"name\":\"controller\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBrevisMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use IBrevisMarketMetaData.ABI instead.
var IBrevisMarketABI = IBrevisMarketMetaData.ABI

// IBrevisMarket is an auto generated Go binding around an Ethereum contract.
type IBrevisMarket struct {
	IBrevisMarketCaller     // Read-only binding to the contract
	IBrevisMarketTransactor // Write-only binding to the contract
	IBrevisMarketFilterer   // Log filterer for contract events
}

// IBrevisMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBrevisMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBrevisMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBrevisMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBrevisMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBrevisMarketSession struct {
	Contract     *IBrevisMarket    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBrevisMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBrevisMarketCallerSession struct {
	Contract *IBrevisMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBrevisMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBrevisMarketTransactorSession struct {
	Contract     *IBrevisMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBrevisMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBrevisMarketRaw struct {
	Contract *IBrevisMarket // Generic contract binding to access the raw methods on
}

// IBrevisMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBrevisMarketCallerRaw struct {
	Contract *IBrevisMarketCaller // Generic read-only contract binding to access the raw methods on
}

// IBrevisMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBrevisMarketTransactorRaw struct {
	Contract *IBrevisMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBrevisMarket creates a new instance of IBrevisMarket, bound to a specific deployed contract.
func NewIBrevisMarket(address common.Address, backend bind.ContractBackend) (*IBrevisMarket, error) {
	contract, err := bindIBrevisMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarket{IBrevisMarketCaller: IBrevisMarketCaller{contract: contract}, IBrevisMarketTransactor: IBrevisMarketTransactor{contract: contract}, IBrevisMarketFilterer: IBrevisMarketFilterer{contract: contract}}, nil
}

// NewIBrevisMarketCaller creates a new read-only instance of IBrevisMarket, bound to a specific deployed contract.
func NewIBrevisMarketCaller(address common.Address, caller bind.ContractCaller) (*IBrevisMarketCaller, error) {
	contract, err := bindIBrevisMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketCaller{contract: contract}, nil
}

// NewIBrevisMarketTransactor creates a new write-only instance of IBrevisMarket, bound to a specific deployed contract.
func NewIBrevisMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*IBrevisMarketTransactor, error) {
	contract, err := bindIBrevisMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketTransactor{contract: contract}, nil
}

// NewIBrevisMarketFilterer creates a new log filterer instance of IBrevisMarket, bound to a specific deployed contract.
func NewIBrevisMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*IBrevisMarketFilterer, error) {
	contract, err := bindIBrevisMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketFilterer{contract: contract}, nil
}

// bindIBrevisMarket binds a generic wrapper to an already deployed contract.
func bindIBrevisMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBrevisMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisMarket *IBrevisMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisMarket.Contract.IBrevisMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisMarket *IBrevisMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.IBrevisMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisMarket *IBrevisMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.IBrevisMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBrevisMarket *IBrevisMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBrevisMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBrevisMarket *IBrevisMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBrevisMarket *IBrevisMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.contract.Transact(opts, method, params...)
}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256 duration)
func (_IBrevisMarket *IBrevisMarketCaller) MAXDEADLINEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "MAX_DEADLINE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256 duration)
func (_IBrevisMarket *IBrevisMarketSession) MAXDEADLINEDURATION() (*big.Int, error) {
	return _IBrevisMarket.Contract.MAXDEADLINEDURATION(&_IBrevisMarket.CallOpts)
}

// MAXDEADLINEDURATION is a free data retrieval call binding the contract method 0x7322ae37.
//
// Solidity: function MAX_DEADLINE_DURATION() view returns(uint256 duration)
func (_IBrevisMarket *IBrevisMarketCallerSession) MAXDEADLINEDURATION() (*big.Int, error) {
	return _IBrevisMarket.Contract.MAXDEADLINEDURATION(&_IBrevisMarket.CallOpts)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketCaller) BiddingPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "biddingPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketSession) BiddingPhaseDuration() (uint64, error) {
	return _IBrevisMarket.Contract.BiddingPhaseDuration(&_IBrevisMarket.CallOpts)
}

// BiddingPhaseDuration is a free data retrieval call binding the contract method 0x96f6fe91.
//
// Solidity: function biddingPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketCallerSession) BiddingPhaseDuration() (uint64, error) {
	return _IBrevisMarket.Contract.BiddingPhaseDuration(&_IBrevisMarket.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address token)
func (_IBrevisMarket *IBrevisMarketCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address token)
func (_IBrevisMarket *IBrevisMarketSession) FeeToken() (common.Address, error) {
	return _IBrevisMarket.Contract.FeeToken(&_IBrevisMarket.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address token)
func (_IBrevisMarket *IBrevisMarketCallerSession) FeeToken() (common.Address, error) {
	return _IBrevisMarket.Contract.FeeToken(&_IBrevisMarket.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketCaller) GetBidHash(opts *bind.CallOpts, reqid [32]byte, prover common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "getBidHash", reqid, prover)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketSession) GetBidHash(reqid [32]byte, prover common.Address) ([32]byte, error) {
	return _IBrevisMarket.Contract.GetBidHash(&_IBrevisMarket.CallOpts, reqid, prover)
}

// GetBidHash is a free data retrieval call binding the contract method 0xf628d88d.
//
// Solidity: function getBidHash(bytes32 reqid, address prover) view returns(bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketCallerSession) GetBidHash(reqid [32]byte, prover common.Address) ([32]byte, error) {
	return _IBrevisMarket.Contract.GetBidHash(&_IBrevisMarket.CallOpts, reqid, prover)
}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_IBrevisMarket *IBrevisMarketCaller) GetBidders(opts *bind.CallOpts, reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "getBidders", reqid)

	outstruct := new(struct {
		Winner      common.Address
		WinnerFee   *big.Int
		SecondPlace common.Address
		SecondFee   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Winner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WinnerFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondPlace = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SecondFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_IBrevisMarket *IBrevisMarketSession) GetBidders(reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	return _IBrevisMarket.Contract.GetBidders(&_IBrevisMarket.CallOpts, reqid)
}

// GetBidders is a free data retrieval call binding the contract method 0x6dd5bd60.
//
// Solidity: function getBidders(bytes32 reqid) view returns(address winner, uint256 winnerFee, address secondPlace, uint256 secondFee)
func (_IBrevisMarket *IBrevisMarketCallerSession) GetBidders(reqid [32]byte) (struct {
	Winner      common.Address
	WinnerFee   *big.Int
	SecondPlace common.Address
	SecondFee   *big.Int
}, error) {
	return _IBrevisMarket.Contract.GetBidders(&_IBrevisMarket.CallOpts, reqid)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_IBrevisMarket *IBrevisMarketCaller) GetProof(opts *bind.CallOpts, reqid [32]byte) ([8]*big.Int, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "getProof", reqid)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_IBrevisMarket *IBrevisMarketSession) GetProof(reqid [32]byte) ([8]*big.Int, error) {
	return _IBrevisMarket.Contract.GetProof(&_IBrevisMarket.CallOpts, reqid)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 reqid) view returns(uint256[8] proof)
func (_IBrevisMarket *IBrevisMarketCallerSession) GetProof(reqid [32]byte) ([8]*big.Int, error) {
	return _IBrevisMarket.Contract.GetProof(&_IBrevisMarket.CallOpts, reqid)
}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_IBrevisMarket *IBrevisMarketCaller) GetProtocolFeeInfo(opts *bind.CallOpts) (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "getProtocolFeeInfo")

	outstruct := new(struct {
		FeeBps  *big.Int
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeBps = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_IBrevisMarket *IBrevisMarketSession) GetProtocolFeeInfo() (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	return _IBrevisMarket.Contract.GetProtocolFeeInfo(&_IBrevisMarket.CallOpts)
}

// GetProtocolFeeInfo is a free data retrieval call binding the contract method 0x695639b7.
//
// Solidity: function getProtocolFeeInfo() view returns(uint256 feeBps, uint256 balance)
func (_IBrevisMarket *IBrevisMarketCallerSession) GetProtocolFeeInfo() (struct {
	FeeBps  *big.Int
	Balance *big.Int
}, error) {
	return _IBrevisMarket.Contract.GetProtocolFeeInfo(&_IBrevisMarket.CallOpts)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_IBrevisMarket *IBrevisMarketCaller) GetRequest(opts *bind.CallOpts, reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "getRequest", reqid)

	outstruct := new(struct {
		Status             uint8
		Timestamp          uint64
		Sender             common.Address
		MaxFee             *big.Int
		MinStake           *big.Int
		Deadline           uint64
		Vk                 [32]byte
		PublicValuesDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MaxFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinStake = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Vk = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.PublicValuesDigest = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_IBrevisMarket *IBrevisMarketSession) GetRequest(reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	return _IBrevisMarket.Contract.GetRequest(&_IBrevisMarket.CallOpts, reqid)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 reqid) view returns(uint8 status, uint64 timestamp, address sender, uint256 maxFee, uint256 minStake, uint64 deadline, bytes32 vk, bytes32 publicValuesDigest)
func (_IBrevisMarket *IBrevisMarketCallerSession) GetRequest(reqid [32]byte) (struct {
	Status             uint8
	Timestamp          uint64
	Sender             common.Address
	MaxFee             *big.Int
	MinStake           *big.Int
	Deadline           uint64
	Vk                 [32]byte
	PublicValuesDigest [32]byte
}, error) {
	return _IBrevisMarket.Contract.GetRequest(&_IBrevisMarket.CallOpts, reqid)
}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256 fee)
func (_IBrevisMarket *IBrevisMarketCaller) MinMaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "minMaxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256 fee)
func (_IBrevisMarket *IBrevisMarketSession) MinMaxFee() (*big.Int, error) {
	return _IBrevisMarket.Contract.MinMaxFee(&_IBrevisMarket.CallOpts)
}

// MinMaxFee is a free data retrieval call binding the contract method 0xfafe4201.
//
// Solidity: function minMaxFee() view returns(uint256 fee)
func (_IBrevisMarket *IBrevisMarketCallerSession) MinMaxFee() (*big.Int, error) {
	return _IBrevisMarket.Contract.MinMaxFee(&_IBrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address verifier)
func (_IBrevisMarket *IBrevisMarketCaller) PicoVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "picoVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address verifier)
func (_IBrevisMarket *IBrevisMarketSession) PicoVerifier() (common.Address, error) {
	return _IBrevisMarket.Contract.PicoVerifier(&_IBrevisMarket.CallOpts)
}

// PicoVerifier is a free data retrieval call binding the contract method 0x8894a097.
//
// Solidity: function picoVerifier() view returns(address verifier)
func (_IBrevisMarket *IBrevisMarketCallerSession) PicoVerifier() (common.Address, error) {
	return _IBrevisMarket.Contract.PicoVerifier(&_IBrevisMarket.CallOpts)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketCaller) RevealPhaseDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "revealPhaseDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketSession) RevealPhaseDuration() (uint64, error) {
	return _IBrevisMarket.Contract.RevealPhaseDuration(&_IBrevisMarket.CallOpts)
}

// RevealPhaseDuration is a free data retrieval call binding the contract method 0xdfc75372.
//
// Solidity: function revealPhaseDuration() view returns(uint64 duration)
func (_IBrevisMarket *IBrevisMarketCallerSession) RevealPhaseDuration() (uint64, error) {
	return _IBrevisMarket.Contract.RevealPhaseDuration(&_IBrevisMarket.CallOpts)
}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address controller)
func (_IBrevisMarket *IBrevisMarketCaller) StakingController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBrevisMarket.contract.Call(opts, &out, "stakingController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address controller)
func (_IBrevisMarket *IBrevisMarketSession) StakingController() (common.Address, error) {
	return _IBrevisMarket.Contract.StakingController(&_IBrevisMarket.CallOpts)
}

// StakingController is a free data retrieval call binding the contract method 0x713e6a09.
//
// Solidity: function stakingController() view returns(address controller)
func (_IBrevisMarket *IBrevisMarketCallerSession) StakingController() (common.Address, error) {
	return _IBrevisMarket.Contract.StakingController(&_IBrevisMarket.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) Bid(opts *bind.TransactOpts, reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "bid", reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_IBrevisMarket *IBrevisMarketSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Bid(&_IBrevisMarket.TransactOpts, reqid, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0x434f967c.
//
// Solidity: function bid(bytes32 reqid, bytes32 bidHash) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) Bid(reqid [32]byte, bidHash [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Bid(&_IBrevisMarket.TransactOpts, reqid, bidHash)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) Refund(opts *bind.TransactOpts, reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "refund", reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Refund(&_IBrevisMarket.TransactOpts, reqid)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) Refund(reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Refund(&_IBrevisMarket.TransactOpts, reqid)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) RequestProof(opts *bind.TransactOpts, req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "requestProof", req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_IBrevisMarket *IBrevisMarketSession) RequestProof(req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.RequestProof(&_IBrevisMarket.TransactOpts, req)
}

// RequestProof is a paid mutator transaction binding the contract method 0xc22d3155.
//
// Solidity: function requestProof((uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) RequestProof(req IBrevisMarketProofRequest) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.RequestProof(&_IBrevisMarket.TransactOpts, req)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) Reveal(opts *bind.TransactOpts, reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "reveal", reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_IBrevisMarket *IBrevisMarketSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Reveal(&_IBrevisMarket.TransactOpts, reqid, fee, nonce)
}

// Reveal is a paid mutator transaction binding the contract method 0x2cf5d279.
//
// Solidity: function reveal(bytes32 reqid, uint256 fee, uint256 nonce) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) Reveal(reqid [32]byte, fee *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Reveal(&_IBrevisMarket.TransactOpts, reqid, fee, nonce)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetBiddingPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setBiddingPhaseDuration", newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetBiddingPhaseDuration(&_IBrevisMarket.TransactOpts, newDuration)
}

// SetBiddingPhaseDuration is a paid mutator transaction binding the contract method 0x7d9b7158.
//
// Solidity: function setBiddingPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetBiddingPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetBiddingPhaseDuration(&_IBrevisMarket.TransactOpts, newDuration)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetMinMaxFee(opts *bind.TransactOpts, newMinFee *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setMinMaxFee", newMinFee)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetMinMaxFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetMinMaxFee(&_IBrevisMarket.TransactOpts, newMinFee)
}

// SetMinMaxFee is a paid mutator transaction binding the contract method 0xe30c1fc3.
//
// Solidity: function setMinMaxFee(uint256 newMinFee) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetMinMaxFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetMinMaxFee(&_IBrevisMarket.TransactOpts, newMinFee)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetPicoVerifier(opts *bind.TransactOpts, newVerifier common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setPicoVerifier", newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetPicoVerifier(&_IBrevisMarket.TransactOpts, newVerifier)
}

// SetPicoVerifier is a paid mutator transaction binding the contract method 0x95591966.
//
// Solidity: function setPicoVerifier(address newVerifier) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetPicoVerifier(newVerifier common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetPicoVerifier(&_IBrevisMarket.TransactOpts, newVerifier)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetProtocolFeeBps(opts *bind.TransactOpts, newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setProtocolFeeBps", newBps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetProtocolFeeBps(newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetProtocolFeeBps(&_IBrevisMarket.TransactOpts, newBps)
}

// SetProtocolFeeBps is a paid mutator transaction binding the contract method 0xc0417e58.
//
// Solidity: function setProtocolFeeBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetProtocolFeeBps(newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetProtocolFeeBps(&_IBrevisMarket.TransactOpts, newBps)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetRevealPhaseDuration(opts *bind.TransactOpts, newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setRevealPhaseDuration", newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetRevealPhaseDuration(&_IBrevisMarket.TransactOpts, newDuration)
}

// SetRevealPhaseDuration is a paid mutator transaction binding the contract method 0xeaf57ad7.
//
// Solidity: function setRevealPhaseDuration(uint64 newDuration) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetRevealPhaseDuration(newDuration uint64) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetRevealPhaseDuration(&_IBrevisMarket.TransactOpts, newDuration)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetSlashBps(opts *bind.TransactOpts, newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setSlashBps", newBps)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetSlashBps(newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetSlashBps(&_IBrevisMarket.TransactOpts, newBps)
}

// SetSlashBps is a paid mutator transaction binding the contract method 0xd2b8f2fc.
//
// Solidity: function setSlashBps(uint256 newBps) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetSlashBps(newBps *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetSlashBps(&_IBrevisMarket.TransactOpts, newBps)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SetSlashWindow(opts *bind.TransactOpts, newWindow *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "setSlashWindow", newWindow)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_IBrevisMarket *IBrevisMarketSession) SetSlashWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetSlashWindow(&_IBrevisMarket.TransactOpts, newWindow)
}

// SetSlashWindow is a paid mutator transaction binding the contract method 0x83d860a0.
//
// Solidity: function setSlashWindow(uint256 newWindow) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SetSlashWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SetSlashWindow(&_IBrevisMarket.TransactOpts, newWindow)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) Slash(opts *bind.TransactOpts, reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "slash", reqid)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketSession) Slash(reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Slash(&_IBrevisMarket.TransactOpts, reqid)
}

// Slash is a paid mutator transaction binding the contract method 0xf415ed14.
//
// Solidity: function slash(bytes32 reqid) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) Slash(reqid [32]byte) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.Slash(&_IBrevisMarket.TransactOpts, reqid)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) SubmitProof(opts *bind.TransactOpts, reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "submitProof", reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_IBrevisMarket *IBrevisMarketSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SubmitProof(&_IBrevisMarket.TransactOpts, reqid, proof)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x7e88b1a0.
//
// Solidity: function submitProof(bytes32 reqid, uint256[8] proof) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) SubmitProof(reqid [32]byte, proof [8]*big.Int) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.SubmitProof(&_IBrevisMarket.TransactOpts, reqid, proof)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_IBrevisMarket *IBrevisMarketTransactor) WithdrawProtocolFee(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.contract.Transact(opts, "withdrawProtocolFee", to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_IBrevisMarket *IBrevisMarketSession) WithdrawProtocolFee(to common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.WithdrawProtocolFee(&_IBrevisMarket.TransactOpts, to)
}

// WithdrawProtocolFee is a paid mutator transaction binding the contract method 0x668fb6dc.
//
// Solidity: function withdrawProtocolFee(address to) returns()
func (_IBrevisMarket *IBrevisMarketTransactorSession) WithdrawProtocolFee(to common.Address) (*types.Transaction, error) {
	return _IBrevisMarket.Contract.WithdrawProtocolFee(&_IBrevisMarket.TransactOpts, to)
}

// IBrevisMarketBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the IBrevisMarket contract.
type IBrevisMarketBidRevealedIterator struct {
	Event *IBrevisMarketBidRevealed // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketBidRevealed)
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
		it.Event = new(IBrevisMarketBidRevealed)
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
func (it *IBrevisMarketBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketBidRevealed represents a BidRevealed event raised by the IBrevisMarket contract.
type IBrevisMarketBidRevealed struct {
	Reqid  [32]byte
	Prover common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterBidRevealed(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*IBrevisMarketBidRevealedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketBidRevealedIterator{contract: _IBrevisMarket.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *IBrevisMarketBidRevealed, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "BidRevealed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketBidRevealed)
				if err := _IBrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
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

// ParseBidRevealed is a log parse operation binding the contract event 0x9c253a2ca87d91c2658cfbfcc3d6081283acc3cf95c2961667a858b531ec491f.
//
// Solidity: event BidRevealed(bytes32 indexed reqid, address indexed prover, uint256 fee)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseBidRevealed(log types.Log) (*IBrevisMarketBidRevealed, error) {
	event := new(IBrevisMarketBidRevealed)
	if err := _IBrevisMarket.contract.UnpackLog(event, "BidRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketBiddingPhaseDurationUpdatedIterator is returned from FilterBiddingPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for BiddingPhaseDurationUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketBiddingPhaseDurationUpdatedIterator struct {
	Event *IBrevisMarketBiddingPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketBiddingPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketBiddingPhaseDurationUpdated)
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
		it.Event = new(IBrevisMarketBiddingPhaseDurationUpdated)
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
func (it *IBrevisMarketBiddingPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketBiddingPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketBiddingPhaseDurationUpdated represents a BiddingPhaseDurationUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketBiddingPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBiddingPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterBiddingPhaseDurationUpdated(opts *bind.FilterOpts) (*IBrevisMarketBiddingPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketBiddingPhaseDurationUpdatedIterator{contract: _IBrevisMarket.contract, event: "BiddingPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchBiddingPhaseDurationUpdated is a free log subscription operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchBiddingPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketBiddingPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "BiddingPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketBiddingPhaseDurationUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
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

// ParseBiddingPhaseDurationUpdated is a log parse operation binding the contract event 0xae2a85369654bba2b48f3887a0b5c87ccef1022b39a5e25fd30d599edc93ad6e.
//
// Solidity: event BiddingPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseBiddingPhaseDurationUpdated(log types.Log) (*IBrevisMarketBiddingPhaseDurationUpdated, error) {
	event := new(IBrevisMarketBiddingPhaseDurationUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "BiddingPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketFeeTokenUpdatedIterator is returned from FilterFeeTokenUpdated and is used to iterate over the raw logs and unpacked data for FeeTokenUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketFeeTokenUpdatedIterator struct {
	Event *IBrevisMarketFeeTokenUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketFeeTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketFeeTokenUpdated)
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
		it.Event = new(IBrevisMarketFeeTokenUpdated)
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
func (it *IBrevisMarketFeeTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketFeeTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketFeeTokenUpdated represents a FeeTokenUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketFeeTokenUpdated struct {
	OldToken common.Address
	NewToken common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeTokenUpdated is a free log retrieval operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterFeeTokenUpdated(opts *bind.FilterOpts, oldToken []common.Address, newToken []common.Address) (*IBrevisMarketFeeTokenUpdatedIterator, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "FeeTokenUpdated", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketFeeTokenUpdatedIterator{contract: _IBrevisMarket.contract, event: "FeeTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeTokenUpdated is a free log subscription operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchFeeTokenUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketFeeTokenUpdated, oldToken []common.Address, newToken []common.Address) (event.Subscription, error) {

	var oldTokenRule []interface{}
	for _, oldTokenItem := range oldToken {
		oldTokenRule = append(oldTokenRule, oldTokenItem)
	}
	var newTokenRule []interface{}
	for _, newTokenItem := range newToken {
		newTokenRule = append(newTokenRule, newTokenItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "FeeTokenUpdated", oldTokenRule, newTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketFeeTokenUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "FeeTokenUpdated", log); err != nil {
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

// ParseFeeTokenUpdated is a log parse operation binding the contract event 0x91a03e1d689caf891fe531c01e290f7b718f9c6a3af6726d6d837d2b7bd82e67.
//
// Solidity: event FeeTokenUpdated(address indexed oldToken, address indexed newToken)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseFeeTokenUpdated(log types.Log) (*IBrevisMarketFeeTokenUpdated, error) {
	event := new(IBrevisMarketFeeTokenUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "FeeTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketMinMaxFeeUpdatedIterator is returned from FilterMinMaxFeeUpdated and is used to iterate over the raw logs and unpacked data for MinMaxFeeUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketMinMaxFeeUpdatedIterator struct {
	Event *IBrevisMarketMinMaxFeeUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketMinMaxFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketMinMaxFeeUpdated)
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
		it.Event = new(IBrevisMarketMinMaxFeeUpdated)
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
func (it *IBrevisMarketMinMaxFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketMinMaxFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketMinMaxFeeUpdated represents a MinMaxFeeUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketMinMaxFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinMaxFeeUpdated is a free log retrieval operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterMinMaxFeeUpdated(opts *bind.FilterOpts) (*IBrevisMarketMinMaxFeeUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "MinMaxFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketMinMaxFeeUpdatedIterator{contract: _IBrevisMarket.contract, event: "MinMaxFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinMaxFeeUpdated is a free log subscription operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchMinMaxFeeUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketMinMaxFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "MinMaxFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketMinMaxFeeUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "MinMaxFeeUpdated", log); err != nil {
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

// ParseMinMaxFeeUpdated is a log parse operation binding the contract event 0xbed878bea5faee24f1f99a52bda105fcd767c49edb25248985fa0838e5a92f9f.
//
// Solidity: event MinMaxFeeUpdated(uint256 oldFee, uint256 newFee)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseMinMaxFeeUpdated(log types.Log) (*IBrevisMarketMinMaxFeeUpdated, error) {
	event := new(IBrevisMarketMinMaxFeeUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "MinMaxFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the IBrevisMarket contract.
type IBrevisMarketNewBidIterator struct {
	Event *IBrevisMarketNewBid // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketNewBid)
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
		it.Event = new(IBrevisMarketNewBid)
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
func (it *IBrevisMarketNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketNewBid represents a NewBid event raised by the IBrevisMarket contract.
type IBrevisMarketNewBid struct {
	Reqid   [32]byte
	Prover  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterNewBid(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*IBrevisMarketNewBidIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketNewBidIterator{contract: _IBrevisMarket.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *IBrevisMarketNewBid, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "NewBid", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketNewBid)
				if err := _IBrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0xe3b42b8d48a26fbc202ccd924bc9c016974833ac32de0a612cf9eb6160e5f45e.
//
// Solidity: event NewBid(bytes32 indexed reqid, address indexed prover, bytes32 bidHash)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseNewBid(log types.Log) (*IBrevisMarketNewBid, error) {
	event := new(IBrevisMarketNewBid)
	if err := _IBrevisMarket.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the IBrevisMarket contract.
type IBrevisMarketNewRequestIterator struct {
	Event *IBrevisMarketNewRequest // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketNewRequest)
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
		it.Event = new(IBrevisMarketNewRequest)
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
func (it *IBrevisMarketNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketNewRequest represents a NewRequest event raised by the IBrevisMarket contract.
type IBrevisMarketNewRequest struct {
	Reqid [32]byte
	Req   IBrevisMarketProofRequest
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterNewRequest(opts *bind.FilterOpts, reqid [][32]byte) (*IBrevisMarketNewRequestIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketNewRequestIterator{contract: _IBrevisMarket.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *IBrevisMarketNewRequest, reqid [][32]byte) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "NewRequest", reqidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketNewRequest)
				if err := _IBrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0x7efcf76179636bbad459f07114e7c620a02919c12199daf057c0c63f9debd123.
//
// Solidity: event NewRequest(bytes32 indexed reqid, (uint64,bytes32,bytes32,string,bytes[],string,(uint256,uint256,uint64)) req)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseNewRequest(log types.Log) (*IBrevisMarketNewRequest, error) {
	event := new(IBrevisMarketNewRequest)
	if err := _IBrevisMarket.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketPicoVerifierUpdatedIterator is returned from FilterPicoVerifierUpdated and is used to iterate over the raw logs and unpacked data for PicoVerifierUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketPicoVerifierUpdatedIterator struct {
	Event *IBrevisMarketPicoVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketPicoVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketPicoVerifierUpdated)
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
		it.Event = new(IBrevisMarketPicoVerifierUpdated)
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
func (it *IBrevisMarketPicoVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketPicoVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketPicoVerifierUpdated represents a PicoVerifierUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketPicoVerifierUpdated struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPicoVerifierUpdated is a free log retrieval operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterPicoVerifierUpdated(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*IBrevisMarketPicoVerifierUpdatedIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketPicoVerifierUpdatedIterator{contract: _IBrevisMarket.contract, event: "PicoVerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchPicoVerifierUpdated is a free log subscription operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchPicoVerifierUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketPicoVerifierUpdated, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "PicoVerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketPicoVerifierUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
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

// ParsePicoVerifierUpdated is a log parse operation binding the contract event 0x289e68b8f41113af2cd98bcf4fd4369937169415c5c43937f1d9a7aea3270a80.
//
// Solidity: event PicoVerifierUpdated(address indexed oldVerifier, address indexed newVerifier)
func (_IBrevisMarket *IBrevisMarketFilterer) ParsePicoVerifierUpdated(log types.Log) (*IBrevisMarketPicoVerifierUpdated, error) {
	event := new(IBrevisMarketPicoVerifierUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "PicoVerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketProofSubmittedIterator is returned from FilterProofSubmitted and is used to iterate over the raw logs and unpacked data for ProofSubmitted events raised by the IBrevisMarket contract.
type IBrevisMarketProofSubmittedIterator struct {
	Event *IBrevisMarketProofSubmitted // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketProofSubmitted)
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
		it.Event = new(IBrevisMarketProofSubmitted)
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
func (it *IBrevisMarketProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketProofSubmitted represents a ProofSubmitted event raised by the IBrevisMarket contract.
type IBrevisMarketProofSubmitted struct {
	Reqid     [32]byte
	Prover    common.Address
	Proof     [8]*big.Int
	ActualFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitted is a free log retrieval operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterProofSubmitted(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*IBrevisMarketProofSubmittedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketProofSubmittedIterator{contract: _IBrevisMarket.contract, event: "ProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchProofSubmitted is a free log subscription operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchProofSubmitted(opts *bind.WatchOpts, sink chan<- *IBrevisMarketProofSubmitted, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "ProofSubmitted", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketProofSubmitted)
				if err := _IBrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
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

// ParseProofSubmitted is a log parse operation binding the contract event 0x859954befd38b6b306ee0105f3d1e765aa65e922a7190c8d9f9a3003b60fac34.
//
// Solidity: event ProofSubmitted(bytes32 indexed reqid, address indexed prover, uint256[8] proof, uint256 actualFee)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseProofSubmitted(log types.Log) (*IBrevisMarketProofSubmitted, error) {
	event := new(IBrevisMarketProofSubmitted)
	if err := _IBrevisMarket.contract.UnpackLog(event, "ProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketProtocolFeeBpsUpdatedIterator is returned from FilterProtocolFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeBpsUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketProtocolFeeBpsUpdatedIterator struct {
	Event *IBrevisMarketProtocolFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketProtocolFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketProtocolFeeBpsUpdated)
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
		it.Event = new(IBrevisMarketProtocolFeeBpsUpdated)
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
func (it *IBrevisMarketProtocolFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketProtocolFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketProtocolFeeBpsUpdated represents a ProtocolFeeBpsUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketProtocolFeeBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeBpsUpdated is a free log retrieval operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterProtocolFeeBpsUpdated(opts *bind.FilterOpts) (*IBrevisMarketProtocolFeeBpsUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "ProtocolFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketProtocolFeeBpsUpdatedIterator{contract: _IBrevisMarket.contract, event: "ProtocolFeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeBpsUpdated is a free log subscription operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchProtocolFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketProtocolFeeBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "ProtocolFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketProtocolFeeBpsUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "ProtocolFeeBpsUpdated", log); err != nil {
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

// ParseProtocolFeeBpsUpdated is a log parse operation binding the contract event 0xf247e7445763d3e9a01aee4b64de35f02273d6a1d7264694bb9298109319c8be.
//
// Solidity: event ProtocolFeeBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseProtocolFeeBpsUpdated(log types.Log) (*IBrevisMarketProtocolFeeBpsUpdated, error) {
	event := new(IBrevisMarketProtocolFeeBpsUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "ProtocolFeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketProtocolFeeWithdrawnIterator is returned from FilterProtocolFeeWithdrawn and is used to iterate over the raw logs and unpacked data for ProtocolFeeWithdrawn events raised by the IBrevisMarket contract.
type IBrevisMarketProtocolFeeWithdrawnIterator struct {
	Event *IBrevisMarketProtocolFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketProtocolFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketProtocolFeeWithdrawn)
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
		it.Event = new(IBrevisMarketProtocolFeeWithdrawn)
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
func (it *IBrevisMarketProtocolFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketProtocolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketProtocolFeeWithdrawn represents a ProtocolFeeWithdrawn event raised by the IBrevisMarket contract.
type IBrevisMarketProtocolFeeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeWithdrawn is a free log retrieval operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterProtocolFeeWithdrawn(opts *bind.FilterOpts, to []common.Address) (*IBrevisMarketProtocolFeeWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "ProtocolFeeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketProtocolFeeWithdrawnIterator{contract: _IBrevisMarket.contract, event: "ProtocolFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeWithdrawn is a free log subscription operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchProtocolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *IBrevisMarketProtocolFeeWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "ProtocolFeeWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketProtocolFeeWithdrawn)
				if err := _IBrevisMarket.contract.UnpackLog(event, "ProtocolFeeWithdrawn", log); err != nil {
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

// ParseProtocolFeeWithdrawn is a log parse operation binding the contract event 0x052c2c1904fab85ddafadaeeae6731433f2cba1bcb770a300d25a40d989acf73.
//
// Solidity: event ProtocolFeeWithdrawn(address indexed to, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseProtocolFeeWithdrawn(log types.Log) (*IBrevisMarketProtocolFeeWithdrawn, error) {
	event := new(IBrevisMarketProtocolFeeWithdrawn)
	if err := _IBrevisMarket.contract.UnpackLog(event, "ProtocolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketProverSlashedIterator is returned from FilterProverSlashed and is used to iterate over the raw logs and unpacked data for ProverSlashed events raised by the IBrevisMarket contract.
type IBrevisMarketProverSlashedIterator struct {
	Event *IBrevisMarketProverSlashed // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketProverSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketProverSlashed)
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
		it.Event = new(IBrevisMarketProverSlashed)
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
func (it *IBrevisMarketProverSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketProverSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketProverSlashed represents a ProverSlashed event raised by the IBrevisMarket contract.
type IBrevisMarketProverSlashed struct {
	Reqid       [32]byte
	Prover      common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProverSlashed is a free log retrieval operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterProverSlashed(opts *bind.FilterOpts, reqid [][32]byte, prover []common.Address) (*IBrevisMarketProverSlashedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "ProverSlashed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketProverSlashedIterator{contract: _IBrevisMarket.contract, event: "ProverSlashed", logs: logs, sub: sub}, nil
}

// WatchProverSlashed is a free log subscription operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchProverSlashed(opts *bind.WatchOpts, sink chan<- *IBrevisMarketProverSlashed, reqid [][32]byte, prover []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "ProverSlashed", reqidRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketProverSlashed)
				if err := _IBrevisMarket.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
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

// ParseProverSlashed is a log parse operation binding the contract event 0x6371a7077bd5be10ad8186ac347e7acd8fed70c411b5f223f01fcc9190f42ac6.
//
// Solidity: event ProverSlashed(bytes32 indexed reqid, address indexed prover, uint256 slashAmount)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseProverSlashed(log types.Log) (*IBrevisMarketProverSlashed, error) {
	event := new(IBrevisMarketProverSlashed)
	if err := _IBrevisMarket.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the IBrevisMarket contract.
type IBrevisMarketRefundedIterator struct {
	Event *IBrevisMarketRefunded // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketRefunded)
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
		it.Event = new(IBrevisMarketRefunded)
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
func (it *IBrevisMarketRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketRefunded represents a Refunded event raised by the IBrevisMarket contract.
type IBrevisMarketRefunded struct {
	Reqid     [32]byte
	Requester common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterRefunded(opts *bind.FilterOpts, reqid [][32]byte, requester []common.Address) (*IBrevisMarketRefundedIterator, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketRefundedIterator{contract: _IBrevisMarket.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *IBrevisMarketRefunded, reqid [][32]byte, requester []common.Address) (event.Subscription, error) {

	var reqidRule []interface{}
	for _, reqidItem := range reqid {
		reqidRule = append(reqidRule, reqidItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "Refunded", reqidRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketRefunded)
				if err := _IBrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xf552ca82e113ac3c539c3d617f29fcd19c172a0c75dad017555c9e109f7fe183.
//
// Solidity: event Refunded(bytes32 indexed reqid, address indexed requester, uint256 amount)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseRefunded(log types.Log) (*IBrevisMarketRefunded, error) {
	event := new(IBrevisMarketRefunded)
	if err := _IBrevisMarket.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketRevealPhaseDurationUpdatedIterator is returned from FilterRevealPhaseDurationUpdated and is used to iterate over the raw logs and unpacked data for RevealPhaseDurationUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketRevealPhaseDurationUpdatedIterator struct {
	Event *IBrevisMarketRevealPhaseDurationUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketRevealPhaseDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketRevealPhaseDurationUpdated)
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
		it.Event = new(IBrevisMarketRevealPhaseDurationUpdated)
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
func (it *IBrevisMarketRevealPhaseDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketRevealPhaseDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketRevealPhaseDurationUpdated represents a RevealPhaseDurationUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketRevealPhaseDurationUpdated struct {
	OldDuration uint64
	NewDuration uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealPhaseDurationUpdated is a free log retrieval operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterRevealPhaseDurationUpdated(opts *bind.FilterOpts) (*IBrevisMarketRevealPhaseDurationUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketRevealPhaseDurationUpdatedIterator{contract: _IBrevisMarket.contract, event: "RevealPhaseDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRevealPhaseDurationUpdated is a free log subscription operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchRevealPhaseDurationUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketRevealPhaseDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "RevealPhaseDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketRevealPhaseDurationUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
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

// ParseRevealPhaseDurationUpdated is a log parse operation binding the contract event 0xfbc83668acd4181d31302d42027aaf88afe3917743b0f5c6935fec9a7583afda.
//
// Solidity: event RevealPhaseDurationUpdated(uint64 oldDuration, uint64 newDuration)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseRevealPhaseDurationUpdated(log types.Log) (*IBrevisMarketRevealPhaseDurationUpdated, error) {
	event := new(IBrevisMarketRevealPhaseDurationUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "RevealPhaseDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketSlashBpsUpdatedIterator is returned from FilterSlashBpsUpdated and is used to iterate over the raw logs and unpacked data for SlashBpsUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketSlashBpsUpdatedIterator struct {
	Event *IBrevisMarketSlashBpsUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketSlashBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketSlashBpsUpdated)
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
		it.Event = new(IBrevisMarketSlashBpsUpdated)
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
func (it *IBrevisMarketSlashBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketSlashBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketSlashBpsUpdated represents a SlashBpsUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketSlashBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashBpsUpdated is a free log retrieval operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterSlashBpsUpdated(opts *bind.FilterOpts) (*IBrevisMarketSlashBpsUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "SlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketSlashBpsUpdatedIterator{contract: _IBrevisMarket.contract, event: "SlashBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashBpsUpdated is a free log subscription operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchSlashBpsUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketSlashBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "SlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketSlashBpsUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "SlashBpsUpdated", log); err != nil {
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

// ParseSlashBpsUpdated is a log parse operation binding the contract event 0x7a9645161550a715662c2036d7ff1b7089f139727f6aa3be8d26a8c55a1ee388.
//
// Solidity: event SlashBpsUpdated(uint256 oldBps, uint256 newBps)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseSlashBpsUpdated(log types.Log) (*IBrevisMarketSlashBpsUpdated, error) {
	event := new(IBrevisMarketSlashBpsUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "SlashBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBrevisMarketSlashWindowUpdatedIterator is returned from FilterSlashWindowUpdated and is used to iterate over the raw logs and unpacked data for SlashWindowUpdated events raised by the IBrevisMarket contract.
type IBrevisMarketSlashWindowUpdatedIterator struct {
	Event *IBrevisMarketSlashWindowUpdated // Event containing the contract specifics and raw log

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
func (it *IBrevisMarketSlashWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBrevisMarketSlashWindowUpdated)
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
		it.Event = new(IBrevisMarketSlashWindowUpdated)
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
func (it *IBrevisMarketSlashWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBrevisMarketSlashWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBrevisMarketSlashWindowUpdated represents a SlashWindowUpdated event raised by the IBrevisMarket contract.
type IBrevisMarketSlashWindowUpdated struct {
	OldWindow *big.Int
	NewWindow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashWindowUpdated is a free log retrieval operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_IBrevisMarket *IBrevisMarketFilterer) FilterSlashWindowUpdated(opts *bind.FilterOpts) (*IBrevisMarketSlashWindowUpdatedIterator, error) {

	logs, sub, err := _IBrevisMarket.contract.FilterLogs(opts, "SlashWindowUpdated")
	if err != nil {
		return nil, err
	}
	return &IBrevisMarketSlashWindowUpdatedIterator{contract: _IBrevisMarket.contract, event: "SlashWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashWindowUpdated is a free log subscription operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_IBrevisMarket *IBrevisMarketFilterer) WatchSlashWindowUpdated(opts *bind.WatchOpts, sink chan<- *IBrevisMarketSlashWindowUpdated) (event.Subscription, error) {

	logs, sub, err := _IBrevisMarket.contract.WatchLogs(opts, "SlashWindowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBrevisMarketSlashWindowUpdated)
				if err := _IBrevisMarket.contract.UnpackLog(event, "SlashWindowUpdated", log); err != nil {
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

// ParseSlashWindowUpdated is a log parse operation binding the contract event 0xc7bb9641822e3c3318570dd58e519090e9189272549564b4644822309e8e6e9f.
//
// Solidity: event SlashWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_IBrevisMarket *IBrevisMarketFilterer) ParseSlashWindowUpdated(log types.Log) (*IBrevisMarketSlashWindowUpdated, error) {
	event := new(IBrevisMarketSlashWindowUpdated)
	if err := _IBrevisMarket.contract.UnpackLog(event, "SlashWindowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1363MetaData contains all meta data concerning the IERC1363 contract.
var IERC1363MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferFromAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFromAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC1363ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1363MetaData.ABI instead.
var IERC1363ABI = IERC1363MetaData.ABI

// IERC1363 is an auto generated Go binding around an Ethereum contract.
type IERC1363 struct {
	IERC1363Caller     // Read-only binding to the contract
	IERC1363Transactor // Write-only binding to the contract
	IERC1363Filterer   // Log filterer for contract events
}

// IERC1363Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1363Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1363Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1363Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1363Session struct {
	Contract     *IERC1363         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1363CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1363CallerSession struct {
	Contract *IERC1363Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1363TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1363TransactorSession struct {
	Contract     *IERC1363Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1363Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1363Raw struct {
	Contract *IERC1363 // Generic contract binding to access the raw methods on
}

// IERC1363CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1363CallerRaw struct {
	Contract *IERC1363Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1363TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1363TransactorRaw struct {
	Contract *IERC1363Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1363 creates a new instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363(address common.Address, backend bind.ContractBackend) (*IERC1363, error) {
	contract, err := bindIERC1363(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1363{IERC1363Caller: IERC1363Caller{contract: contract}, IERC1363Transactor: IERC1363Transactor{contract: contract}, IERC1363Filterer: IERC1363Filterer{contract: contract}}, nil
}

// NewIERC1363Caller creates a new read-only instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Caller(address common.Address, caller bind.ContractCaller) (*IERC1363Caller, error) {
	contract, err := bindIERC1363(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1363Caller{contract: contract}, nil
}

// NewIERC1363Transactor creates a new write-only instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1363Transactor, error) {
	contract, err := bindIERC1363(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1363Transactor{contract: contract}, nil
}

// NewIERC1363Filterer creates a new log filterer instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1363Filterer, error) {
	contract, err := bindIERC1363(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1363Filterer{contract: contract}, nil
}

// bindIERC1363 binds a generic wrapper to an already deployed contract.
func bindIERC1363(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1363MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1363 *IERC1363Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1363.Contract.IERC1363Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1363 *IERC1363Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1363.Contract.IERC1363Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1363 *IERC1363Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1363.Contract.IERC1363Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1363 *IERC1363CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1363.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1363 *IERC1363TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1363.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1363 *IERC1363TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1363.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC1363.Contract.Allowance(&_IERC1363.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC1363.Contract.Allowance(&_IERC1363.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC1363.Contract.BalanceOf(&_IERC1363.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC1363.Contract.BalanceOf(&_IERC1363.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1363.Contract.SupportsInterface(&_IERC1363.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1363.Contract.SupportsInterface(&_IERC1363.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363Session) TotalSupply() (*big.Int, error) {
	return _IERC1363.Contract.TotalSupply(&_IERC1363.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC1363.Contract.TotalSupply(&_IERC1363.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Approve(&_IERC1363.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Approve(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approveAndCall", spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) ApproveAndCall(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) ApproveAndCall(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) ApproveAndCall0(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approveAndCall0", spender, value, data)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) ApproveAndCall0(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall0(&_IERC1363.TransactOpts, spender, value, data)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) ApproveAndCall0(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall0(&_IERC1363.TransactOpts, spender, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Transfer(&_IERC1363.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Transfer(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferAndCall", to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferAndCall(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferAndCall(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferAndCall0(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferAndCall0", to, value, data)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) TransferAndCall0(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall0(&_IERC1363.TransactOpts, to, value, data)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferAndCall0(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall0(&_IERC1363.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFrom(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFrom(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFromAndCall(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFromAndCall", from, to, value, data)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFromAndCall(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall(&_IERC1363.TransactOpts, from, to, value, data)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFromAndCall(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall(&_IERC1363.TransactOpts, from, to, value, data)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFromAndCall0(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFromAndCall0", from, to, value)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFromAndCall0(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall0(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFromAndCall0(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall0(&_IERC1363.TransactOpts, from, to, value)
}

// IERC1363ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC1363 contract.
type IERC1363ApprovalIterator struct {
	Event *IERC1363Approval // Event containing the contract specifics and raw log

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
func (it *IERC1363ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1363Approval)
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
		it.Event = new(IERC1363Approval)
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
func (it *IERC1363ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1363ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1363Approval represents a Approval event raised by the IERC1363 contract.
type IERC1363Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC1363 *IERC1363Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC1363ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC1363.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC1363ApprovalIterator{contract: _IERC1363.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC1363 *IERC1363Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC1363Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC1363.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1363Approval)
				if err := _IERC1363.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC1363 *IERC1363Filterer) ParseApproval(log types.Log) (*IERC1363Approval, error) {
	event := new(IERC1363Approval)
	if err := _IERC1363.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1363TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC1363 contract.
type IERC1363TransferIterator struct {
	Event *IERC1363Transfer // Event containing the contract specifics and raw log

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
func (it *IERC1363TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1363Transfer)
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
		it.Event = new(IERC1363Transfer)
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
func (it *IERC1363TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1363TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1363Transfer represents a Transfer event raised by the IERC1363 contract.
type IERC1363Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC1363 *IERC1363Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC1363TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1363.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1363TransferIterator{contract: _IERC1363.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC1363 *IERC1363Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC1363Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1363.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1363Transfer)
				if err := _IERC1363.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC1363 *IERC1363Filterer) ParseTransfer(log types.Log) (*IERC1363Transfer, error) {
	event := new(IERC1363Transfer)
	if err := _IERC1363.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnableMetaData contains all meta data concerning the IOwnable contract.
var IOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use IOwnableMetaData.ABI instead.
var IOwnableABI = IOwnableMetaData.ABI

// IOwnable is an auto generated Go binding around an Ethereum contract.
type IOwnable struct {
	IOwnableCaller     // Read-only binding to the contract
	IOwnableTransactor // Write-only binding to the contract
	IOwnableFilterer   // Log filterer for contract events
}

// IOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnableSession struct {
	Contract     *IOwnable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnableCallerSession struct {
	Contract *IOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnableTransactorSession struct {
	Contract     *IOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnableRaw struct {
	Contract *IOwnable // Generic contract binding to access the raw methods on
}

// IOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnableCallerRaw struct {
	Contract *IOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnableTransactorRaw struct {
	Contract *IOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwnable creates a new instance of IOwnable, bound to a specific deployed contract.
func NewIOwnable(address common.Address, backend bind.ContractBackend) (*IOwnable, error) {
	contract, err := bindIOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwnable{IOwnableCaller: IOwnableCaller{contract: contract}, IOwnableTransactor: IOwnableTransactor{contract: contract}, IOwnableFilterer: IOwnableFilterer{contract: contract}}, nil
}

// NewIOwnableCaller creates a new read-only instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableCaller(address common.Address, caller bind.ContractCaller) (*IOwnableCaller, error) {
	contract, err := bindIOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableCaller{contract: contract}, nil
}

// NewIOwnableTransactor creates a new write-only instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnableTransactor, error) {
	contract, err := bindIOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableTransactor{contract: contract}, nil
}

// NewIOwnableFilterer creates a new log filterer instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnableFilterer, error) {
	contract, err := bindIOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnableFilterer{contract: contract}, nil
}

// bindIOwnable binds a generic wrapper to an already deployed contract.
func bindIOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnable *IOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnable.Contract.IOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnable *IOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnable.Contract.IOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnable *IOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnable.Contract.IOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnable *IOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnable *IOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnable *IOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableSession) Owner() (common.Address, error) {
	return _IOwnable.Contract.Owner(&_IOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableCallerSession) Owner() (common.Address, error) {
	return _IOwnable.Contract.Owner(&_IOwnable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOwnable *IOwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IOwnable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOwnable *IOwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IOwnable.Contract.TransferOwnership(&_IOwnable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOwnable *IOwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IOwnable.Contract.TransferOwnership(&_IOwnable.TransactOpts, newOwner)
}

// IOwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IOwnable contract.
type IOwnableOwnershipTransferredIterator struct {
	Event *IOwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IOwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnableOwnershipTransferred)
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
		it.Event = new(IOwnableOwnershipTransferred)
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
func (it *IOwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnableOwnershipTransferred represents a OwnershipTransferred event raised by the IOwnable contract.
type IOwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IOwnable *IOwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IOwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IOwnable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IOwnableOwnershipTransferredIterator{contract: _IOwnable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IOwnable *IOwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IOwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IOwnable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnableOwnershipTransferred)
				if err := _IOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IOwnable *IOwnableFilterer) ParseOwnershipTransferred(log types.Log) (*IOwnableOwnershipTransferred, error) {
	event := new(IOwnableOwnershipTransferred)
	if err := _IOwnable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPicoVerifierMetaData contains all meta data concerning the IPicoVerifier contract.
var IPicoVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"riscvVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicValuesHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"verifyPicoProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPicoVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IPicoVerifierMetaData.ABI instead.
var IPicoVerifierABI = IPicoVerifierMetaData.ABI

// IPicoVerifier is an auto generated Go binding around an Ethereum contract.
type IPicoVerifier struct {
	IPicoVerifierCaller     // Read-only binding to the contract
	IPicoVerifierTransactor // Write-only binding to the contract
	IPicoVerifierFilterer   // Log filterer for contract events
}

// IPicoVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPicoVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPicoVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPicoVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPicoVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPicoVerifierSession struct {
	Contract     *IPicoVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPicoVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPicoVerifierCallerSession struct {
	Contract *IPicoVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPicoVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPicoVerifierTransactorSession struct {
	Contract     *IPicoVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPicoVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPicoVerifierRaw struct {
	Contract *IPicoVerifier // Generic contract binding to access the raw methods on
}

// IPicoVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPicoVerifierCallerRaw struct {
	Contract *IPicoVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IPicoVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPicoVerifierTransactorRaw struct {
	Contract *IPicoVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPicoVerifier creates a new instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifier(address common.Address, backend bind.ContractBackend) (*IPicoVerifier, error) {
	contract, err := bindIPicoVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifier{IPicoVerifierCaller: IPicoVerifierCaller{contract: contract}, IPicoVerifierTransactor: IPicoVerifierTransactor{contract: contract}, IPicoVerifierFilterer: IPicoVerifierFilterer{contract: contract}}, nil
}

// NewIPicoVerifierCaller creates a new read-only instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierCaller(address common.Address, caller bind.ContractCaller) (*IPicoVerifierCaller, error) {
	contract, err := bindIPicoVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierCaller{contract: contract}, nil
}

// NewIPicoVerifierTransactor creates a new write-only instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IPicoVerifierTransactor, error) {
	contract, err := bindIPicoVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierTransactor{contract: contract}, nil
}

// NewIPicoVerifierFilterer creates a new log filterer instance of IPicoVerifier, bound to a specific deployed contract.
func NewIPicoVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IPicoVerifierFilterer, error) {
	contract, err := bindIPicoVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPicoVerifierFilterer{contract: contract}, nil
}

// bindIPicoVerifier binds a generic wrapper to an already deployed contract.
func bindIPicoVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPicoVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPicoVerifier *IPicoVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPicoVerifier.Contract.IPicoVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPicoVerifier *IPicoVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.IPicoVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPicoVerifier *IPicoVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.IPicoVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPicoVerifier *IPicoVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPicoVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPicoVerifier *IPicoVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPicoVerifier *IPicoVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPicoVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCaller) VerifyPicoProof(opts *bind.CallOpts, riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _IPicoVerifier.contract.Call(opts, &out, "verifyPicoProof", riscvVkey, publicValues, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof(&_IPicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof is a free data retrieval call binding the contract method 0xee87de37.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes publicValues, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCallerSession) VerifyPicoProof(riscvVkey [32]byte, publicValues []byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof(&_IPicoVerifier.CallOpts, riscvVkey, publicValues, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCaller) VerifyPicoProof0(opts *bind.CallOpts, riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	var out []interface{}
	err := _IPicoVerifier.contract.Call(opts, &out, "verifyPicoProof0", riscvVkey, publicValuesHash, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof0(&_IPicoVerifier.CallOpts, riscvVkey, publicValuesHash, proof)
}

// VerifyPicoProof0 is a free data retrieval call binding the contract method 0xf39751a8.
//
// Solidity: function verifyPicoProof(bytes32 riscvVkey, bytes32 publicValuesHash, uint256[8] proof) view returns()
func (_IPicoVerifier *IPicoVerifierCallerSession) VerifyPicoProof0(riscvVkey [32]byte, publicValuesHash [32]byte, proof [8]*big.Int) error {
	return _IPicoVerifier.Contract.VerifyPicoProof0(&_IPicoVerifier.CallOpts, riscvVkey, publicValuesHash, proof)
}

// IStakingControllerMetaData contains all meta data concerning the IStakingController contract.
var IStakingControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ControllerCannotRetireProverWithAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerCannotRetireProverWithPendingUnstakes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerInsufficientShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerInsufficientTreasury\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerInvalidArg\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerMinSelfStakeNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerNoShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerNoUnstakeRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerOnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerOnlyProver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerOnlyVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerProverAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerProverNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerProverNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerShareAccountingMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerSlashTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerTooManyPendingUnstakes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerUnstakeNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ControllerZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldRate\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newRate\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefault\",\"type\":\"bool\"}],\"name\":\"CommissionRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MaxSlashBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinSelfStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"defaultCommissionRate\",\"type\":\"uint64\"}],\"name\":\"ProverInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"ProverRetired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"ProverSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIStakingController.ProverState\",\"name\":\"oldState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumIStakingController.ProverState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"ProverStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"required\",\"type\":\"bool\"}],\"name\":\"RequireAuthorizationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toStakers\",\"type\":\"uint256\"}],\"name\":\"RewardsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnstakeRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakers\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"beforeShareTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"completeUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"deactivateProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRecover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveProverCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveProvers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"provers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllProvers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"provers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"getCommissionRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"rate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getCommissionRates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"sources\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"rates\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getPendingUnstakes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaleSnapshot\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingController.UnstakeRequest[]\",\"name\":\"requests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProverCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverInfo\",\"outputs\":[{\"internalType\":\"enumIStakingController.ProverState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"defaultCommissionRate\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pendingCommission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numStakers\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverSlashingScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverState\",\"outputs\":[{\"internalType\":\"enumIStakingController.ProverState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverTotalUnstaking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getProverVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getStakersWithPendingUnstakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"getStakersWithPendingUnstakesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveProverVaultAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalVaultAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getUnstakingInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"readyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"defaultCommissionRate\",\"type\":\"uint64\"}],\"name\":\"initializeProver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumVaultAssets\",\"type\":\"uint256\"}],\"name\":\"isProverEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"eligible\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentVaultAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"jailProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSlashBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"onShareTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"reactivateProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"requestUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"name\":\"resetCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"retireProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"newRate\",\"type\":\"uint64\"}],\"name\":\"setCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMaxSlashBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinSelfStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"required\",\"type\":\"bool\"}],\"name\":\"setRequireAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"setUnstakeDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashByAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakerHasPendingUnstakes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasPending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStakingControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingControllerMetaData.ABI instead.
var IStakingControllerABI = IStakingControllerMetaData.ABI

// IStakingController is an auto generated Go binding around an Ethereum contract.
type IStakingController struct {
	IStakingControllerCaller     // Read-only binding to the contract
	IStakingControllerTransactor // Write-only binding to the contract
	IStakingControllerFilterer   // Log filterer for contract events
}

// IStakingControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingControllerSession struct {
	Contract     *IStakingController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingControllerCallerSession struct {
	Contract *IStakingControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IStakingControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingControllerTransactorSession struct {
	Contract     *IStakingControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IStakingControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingControllerRaw struct {
	Contract *IStakingController // Generic contract binding to access the raw methods on
}

// IStakingControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingControllerCallerRaw struct {
	Contract *IStakingControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingControllerTransactorRaw struct {
	Contract *IStakingControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingController creates a new instance of IStakingController, bound to a specific deployed contract.
func NewIStakingController(address common.Address, backend bind.ContractBackend) (*IStakingController, error) {
	contract, err := bindIStakingController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingController{IStakingControllerCaller: IStakingControllerCaller{contract: contract}, IStakingControllerTransactor: IStakingControllerTransactor{contract: contract}, IStakingControllerFilterer: IStakingControllerFilterer{contract: contract}}, nil
}

// NewIStakingControllerCaller creates a new read-only instance of IStakingController, bound to a specific deployed contract.
func NewIStakingControllerCaller(address common.Address, caller bind.ContractCaller) (*IStakingControllerCaller, error) {
	contract, err := bindIStakingController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerCaller{contract: contract}, nil
}

// NewIStakingControllerTransactor creates a new write-only instance of IStakingController, bound to a specific deployed contract.
func NewIStakingControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingControllerTransactor, error) {
	contract, err := bindIStakingController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerTransactor{contract: contract}, nil
}

// NewIStakingControllerFilterer creates a new log filterer instance of IStakingController, bound to a specific deployed contract.
func NewIStakingControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingControllerFilterer, error) {
	contract, err := bindIStakingController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerFilterer{contract: contract}, nil
}

// bindIStakingController binds a generic wrapper to an already deployed contract.
func bindIStakingController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingController *IStakingControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingController.Contract.IStakingControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingController *IStakingControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingController.Contract.IStakingControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingController *IStakingControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingController.Contract.IStakingControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingController *IStakingControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingController *IStakingControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingController *IStakingControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingController.Contract.contract.Transact(opts, method, params...)
}

// BeforeShareTransfer is a free data retrieval call binding the contract method 0x603442d1.
//
// Solidity: function beforeShareTransfer(address prover, address from, uint256 shares) view returns()
func (_IStakingController *IStakingControllerCaller) BeforeShareTransfer(opts *bind.CallOpts, prover common.Address, from common.Address, shares *big.Int) error {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "beforeShareTransfer", prover, from, shares)

	if err != nil {
		return err
	}

	return err

}

// BeforeShareTransfer is a free data retrieval call binding the contract method 0x603442d1.
//
// Solidity: function beforeShareTransfer(address prover, address from, uint256 shares) view returns()
func (_IStakingController *IStakingControllerSession) BeforeShareTransfer(prover common.Address, from common.Address, shares *big.Int) error {
	return _IStakingController.Contract.BeforeShareTransfer(&_IStakingController.CallOpts, prover, from, shares)
}

// BeforeShareTransfer is a free data retrieval call binding the contract method 0x603442d1.
//
// Solidity: function beforeShareTransfer(address prover, address from, uint256 shares) view returns()
func (_IStakingController *IStakingControllerCallerSession) BeforeShareTransfer(prover common.Address, from common.Address, shares *big.Int) error {
	return _IStakingController.Contract.BeforeShareTransfer(&_IStakingController.CallOpts, prover, from, shares)
}

// GetActiveProverCount is a free data retrieval call binding the contract method 0x492723b4.
//
// Solidity: function getActiveProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerCaller) GetActiveProverCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getActiveProverCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveProverCount is a free data retrieval call binding the contract method 0x492723b4.
//
// Solidity: function getActiveProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerSession) GetActiveProverCount() (*big.Int, error) {
	return _IStakingController.Contract.GetActiveProverCount(&_IStakingController.CallOpts)
}

// GetActiveProverCount is a free data retrieval call binding the contract method 0x492723b4.
//
// Solidity: function getActiveProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerCallerSession) GetActiveProverCount() (*big.Int, error) {
	return _IStakingController.Contract.GetActiveProverCount(&_IStakingController.CallOpts)
}

// GetActiveProvers is a free data retrieval call binding the contract method 0x39686c9c.
//
// Solidity: function getActiveProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerCaller) GetActiveProvers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getActiveProvers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveProvers is a free data retrieval call binding the contract method 0x39686c9c.
//
// Solidity: function getActiveProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerSession) GetActiveProvers() ([]common.Address, error) {
	return _IStakingController.Contract.GetActiveProvers(&_IStakingController.CallOpts)
}

// GetActiveProvers is a free data retrieval call binding the contract method 0x39686c9c.
//
// Solidity: function getActiveProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerCallerSession) GetActiveProvers() ([]common.Address, error) {
	return _IStakingController.Contract.GetActiveProvers(&_IStakingController.CallOpts)
}

// GetAllProvers is a free data retrieval call binding the contract method 0x48ceaa59.
//
// Solidity: function getAllProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerCaller) GetAllProvers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getAllProvers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllProvers is a free data retrieval call binding the contract method 0x48ceaa59.
//
// Solidity: function getAllProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerSession) GetAllProvers() ([]common.Address, error) {
	return _IStakingController.Contract.GetAllProvers(&_IStakingController.CallOpts)
}

// GetAllProvers is a free data retrieval call binding the contract method 0x48ceaa59.
//
// Solidity: function getAllProvers() view returns(address[] provers)
func (_IStakingController *IStakingControllerCallerSession) GetAllProvers() ([]common.Address, error) {
	return _IStakingController.Contract.GetAllProvers(&_IStakingController.CallOpts)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xc10702a5.
//
// Solidity: function getCommissionRate(address prover, address source) view returns(uint64 rate)
func (_IStakingController *IStakingControllerCaller) GetCommissionRate(opts *bind.CallOpts, prover common.Address, source common.Address) (uint64, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getCommissionRate", prover, source)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCommissionRate is a free data retrieval call binding the contract method 0xc10702a5.
//
// Solidity: function getCommissionRate(address prover, address source) view returns(uint64 rate)
func (_IStakingController *IStakingControllerSession) GetCommissionRate(prover common.Address, source common.Address) (uint64, error) {
	return _IStakingController.Contract.GetCommissionRate(&_IStakingController.CallOpts, prover, source)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xc10702a5.
//
// Solidity: function getCommissionRate(address prover, address source) view returns(uint64 rate)
func (_IStakingController *IStakingControllerCallerSession) GetCommissionRate(prover common.Address, source common.Address) (uint64, error) {
	return _IStakingController.Contract.GetCommissionRate(&_IStakingController.CallOpts, prover, source)
}

// GetCommissionRates is a free data retrieval call binding the contract method 0x7890b781.
//
// Solidity: function getCommissionRates(address prover) view returns(address[] sources, uint64[] rates)
func (_IStakingController *IStakingControllerCaller) GetCommissionRates(opts *bind.CallOpts, prover common.Address) (struct {
	Sources []common.Address
	Rates   []uint64
}, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getCommissionRates", prover)

	outstruct := new(struct {
		Sources []common.Address
		Rates   []uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sources = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Rates = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

// GetCommissionRates is a free data retrieval call binding the contract method 0x7890b781.
//
// Solidity: function getCommissionRates(address prover) view returns(address[] sources, uint64[] rates)
func (_IStakingController *IStakingControllerSession) GetCommissionRates(prover common.Address) (struct {
	Sources []common.Address
	Rates   []uint64
}, error) {
	return _IStakingController.Contract.GetCommissionRates(&_IStakingController.CallOpts, prover)
}

// GetCommissionRates is a free data retrieval call binding the contract method 0x7890b781.
//
// Solidity: function getCommissionRates(address prover) view returns(address[] sources, uint64[] rates)
func (_IStakingController *IStakingControllerCallerSession) GetCommissionRates(prover common.Address) (struct {
	Sources []common.Address
	Rates   []uint64
}, error) {
	return _IStakingController.Contract.GetCommissionRates(&_IStakingController.CallOpts, prover)
}

// GetPendingUnstakes is a free data retrieval call binding the contract method 0x8f76d7aa.
//
// Solidity: function getPendingUnstakes(address prover, address staker) view returns((uint256,uint256,uint256)[] requests)
func (_IStakingController *IStakingControllerCaller) GetPendingUnstakes(opts *bind.CallOpts, prover common.Address, staker common.Address) ([]IStakingControllerUnstakeRequest, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getPendingUnstakes", prover, staker)

	if err != nil {
		return *new([]IStakingControllerUnstakeRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingControllerUnstakeRequest)).(*[]IStakingControllerUnstakeRequest)

	return out0, err

}

// GetPendingUnstakes is a free data retrieval call binding the contract method 0x8f76d7aa.
//
// Solidity: function getPendingUnstakes(address prover, address staker) view returns((uint256,uint256,uint256)[] requests)
func (_IStakingController *IStakingControllerSession) GetPendingUnstakes(prover common.Address, staker common.Address) ([]IStakingControllerUnstakeRequest, error) {
	return _IStakingController.Contract.GetPendingUnstakes(&_IStakingController.CallOpts, prover, staker)
}

// GetPendingUnstakes is a free data retrieval call binding the contract method 0x8f76d7aa.
//
// Solidity: function getPendingUnstakes(address prover, address staker) view returns((uint256,uint256,uint256)[] requests)
func (_IStakingController *IStakingControllerCallerSession) GetPendingUnstakes(prover common.Address, staker common.Address) ([]IStakingControllerUnstakeRequest, error) {
	return _IStakingController.Contract.GetPendingUnstakes(&_IStakingController.CallOpts, prover, staker)
}

// GetProverCount is a free data retrieval call binding the contract method 0xf329a6bd.
//
// Solidity: function getProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerCaller) GetProverCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProverCount is a free data retrieval call binding the contract method 0xf329a6bd.
//
// Solidity: function getProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerSession) GetProverCount() (*big.Int, error) {
	return _IStakingController.Contract.GetProverCount(&_IStakingController.CallOpts)
}

// GetProverCount is a free data retrieval call binding the contract method 0xf329a6bd.
//
// Solidity: function getProverCount() view returns(uint256 count)
func (_IStakingController *IStakingControllerCallerSession) GetProverCount() (*big.Int, error) {
	return _IStakingController.Contract.GetProverCount(&_IStakingController.CallOpts)
}

// GetProverInfo is a free data retrieval call binding the contract method 0x259c17e4.
//
// Solidity: function getProverInfo(address prover) view returns(uint8 state, address vault, uint64 defaultCommissionRate, uint256 pendingCommission, uint256 numStakers)
func (_IStakingController *IStakingControllerCaller) GetProverInfo(opts *bind.CallOpts, prover common.Address) (struct {
	State                 uint8
	Vault                 common.Address
	DefaultCommissionRate uint64
	PendingCommission     *big.Int
	NumStakers            *big.Int
}, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverInfo", prover)

	outstruct := new(struct {
		State                 uint8
		Vault                 common.Address
		DefaultCommissionRate uint64
		PendingCommission     *big.Int
		NumStakers            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Vault = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DefaultCommissionRate = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PendingCommission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NumStakers = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProverInfo is a free data retrieval call binding the contract method 0x259c17e4.
//
// Solidity: function getProverInfo(address prover) view returns(uint8 state, address vault, uint64 defaultCommissionRate, uint256 pendingCommission, uint256 numStakers)
func (_IStakingController *IStakingControllerSession) GetProverInfo(prover common.Address) (struct {
	State                 uint8
	Vault                 common.Address
	DefaultCommissionRate uint64
	PendingCommission     *big.Int
	NumStakers            *big.Int
}, error) {
	return _IStakingController.Contract.GetProverInfo(&_IStakingController.CallOpts, prover)
}

// GetProverInfo is a free data retrieval call binding the contract method 0x259c17e4.
//
// Solidity: function getProverInfo(address prover) view returns(uint8 state, address vault, uint64 defaultCommissionRate, uint256 pendingCommission, uint256 numStakers)
func (_IStakingController *IStakingControllerCallerSession) GetProverInfo(prover common.Address) (struct {
	State                 uint8
	Vault                 common.Address
	DefaultCommissionRate uint64
	PendingCommission     *big.Int
	NumStakers            *big.Int
}, error) {
	return _IStakingController.Contract.GetProverInfo(&_IStakingController.CallOpts, prover)
}

// GetProverSlashingScale is a free data retrieval call binding the contract method 0x34c0d357.
//
// Solidity: function getProverSlashingScale(address prover) view returns(uint256 scale)
func (_IStakingController *IStakingControllerCaller) GetProverSlashingScale(opts *bind.CallOpts, prover common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverSlashingScale", prover)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProverSlashingScale is a free data retrieval call binding the contract method 0x34c0d357.
//
// Solidity: function getProverSlashingScale(address prover) view returns(uint256 scale)
func (_IStakingController *IStakingControllerSession) GetProverSlashingScale(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverSlashingScale(&_IStakingController.CallOpts, prover)
}

// GetProverSlashingScale is a free data retrieval call binding the contract method 0x34c0d357.
//
// Solidity: function getProverSlashingScale(address prover) view returns(uint256 scale)
func (_IStakingController *IStakingControllerCallerSession) GetProverSlashingScale(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverSlashingScale(&_IStakingController.CallOpts, prover)
}

// GetProverStakers is a free data retrieval call binding the contract method 0x574df063.
//
// Solidity: function getProverStakers(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerCaller) GetProverStakers(opts *bind.CallOpts, prover common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverStakers", prover)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProverStakers is a free data retrieval call binding the contract method 0x574df063.
//
// Solidity: function getProverStakers(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerSession) GetProverStakers(prover common.Address) ([]common.Address, error) {
	return _IStakingController.Contract.GetProverStakers(&_IStakingController.CallOpts, prover)
}

// GetProverStakers is a free data retrieval call binding the contract method 0x574df063.
//
// Solidity: function getProverStakers(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerCallerSession) GetProverStakers(prover common.Address) ([]common.Address, error) {
	return _IStakingController.Contract.GetProverStakers(&_IStakingController.CallOpts, prover)
}

// GetProverState is a free data retrieval call binding the contract method 0xb7a87463.
//
// Solidity: function getProverState(address prover) view returns(uint8 state)
func (_IStakingController *IStakingControllerCaller) GetProverState(opts *bind.CallOpts, prover common.Address) (uint8, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverState", prover)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProverState is a free data retrieval call binding the contract method 0xb7a87463.
//
// Solidity: function getProverState(address prover) view returns(uint8 state)
func (_IStakingController *IStakingControllerSession) GetProverState(prover common.Address) (uint8, error) {
	return _IStakingController.Contract.GetProverState(&_IStakingController.CallOpts, prover)
}

// GetProverState is a free data retrieval call binding the contract method 0xb7a87463.
//
// Solidity: function getProverState(address prover) view returns(uint8 state)
func (_IStakingController *IStakingControllerCallerSession) GetProverState(prover common.Address) (uint8, error) {
	return _IStakingController.Contract.GetProverState(&_IStakingController.CallOpts, prover)
}

// GetProverTotalAssets is a free data retrieval call binding the contract method 0xbb21d52a.
//
// Solidity: function getProverTotalAssets(address prover) view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCaller) GetProverTotalAssets(opts *bind.CallOpts, prover common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverTotalAssets", prover)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProverTotalAssets is a free data retrieval call binding the contract method 0xbb21d52a.
//
// Solidity: function getProverTotalAssets(address prover) view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerSession) GetProverTotalAssets(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverTotalAssets(&_IStakingController.CallOpts, prover)
}

// GetProverTotalAssets is a free data retrieval call binding the contract method 0xbb21d52a.
//
// Solidity: function getProverTotalAssets(address prover) view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCallerSession) GetProverTotalAssets(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverTotalAssets(&_IStakingController.CallOpts, prover)
}

// GetProverTotalUnstaking is a free data retrieval call binding the contract method 0x820b44a9.
//
// Solidity: function getProverTotalUnstaking(address prover) view returns(uint256 totalAmount)
func (_IStakingController *IStakingControllerCaller) GetProverTotalUnstaking(opts *bind.CallOpts, prover common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverTotalUnstaking", prover)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProverTotalUnstaking is a free data retrieval call binding the contract method 0x820b44a9.
//
// Solidity: function getProverTotalUnstaking(address prover) view returns(uint256 totalAmount)
func (_IStakingController *IStakingControllerSession) GetProverTotalUnstaking(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverTotalUnstaking(&_IStakingController.CallOpts, prover)
}

// GetProverTotalUnstaking is a free data retrieval call binding the contract method 0x820b44a9.
//
// Solidity: function getProverTotalUnstaking(address prover) view returns(uint256 totalAmount)
func (_IStakingController *IStakingControllerCallerSession) GetProverTotalUnstaking(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetProverTotalUnstaking(&_IStakingController.CallOpts, prover)
}

// GetProverVault is a free data retrieval call binding the contract method 0xf41f646a.
//
// Solidity: function getProverVault(address prover) view returns(address vault)
func (_IStakingController *IStakingControllerCaller) GetProverVault(opts *bind.CallOpts, prover common.Address) (common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getProverVault", prover)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProverVault is a free data retrieval call binding the contract method 0xf41f646a.
//
// Solidity: function getProverVault(address prover) view returns(address vault)
func (_IStakingController *IStakingControllerSession) GetProverVault(prover common.Address) (common.Address, error) {
	return _IStakingController.Contract.GetProverVault(&_IStakingController.CallOpts, prover)
}

// GetProverVault is a free data retrieval call binding the contract method 0xf41f646a.
//
// Solidity: function getProverVault(address prover) view returns(address vault)
func (_IStakingController *IStakingControllerCallerSession) GetProverVault(prover common.Address) (common.Address, error) {
	return _IStakingController.Contract.GetProverVault(&_IStakingController.CallOpts, prover)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0xd77c8f14.
//
// Solidity: function getStakeInfo(address prover, address staker) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCaller) GetStakeInfo(opts *bind.CallOpts, prover common.Address, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getStakeInfo", prover, staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeInfo is a free data retrieval call binding the contract method 0xd77c8f14.
//
// Solidity: function getStakeInfo(address prover, address staker) view returns(uint256 shares)
func (_IStakingController *IStakingControllerSession) GetStakeInfo(prover common.Address, staker common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetStakeInfo(&_IStakingController.CallOpts, prover, staker)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0xd77c8f14.
//
// Solidity: function getStakeInfo(address prover, address staker) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCallerSession) GetStakeInfo(prover common.Address, staker common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetStakeInfo(&_IStakingController.CallOpts, prover, staker)
}

// GetStakersWithPendingUnstakes is a free data retrieval call binding the contract method 0xff850d2a.
//
// Solidity: function getStakersWithPendingUnstakes(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerCaller) GetStakersWithPendingUnstakes(opts *bind.CallOpts, prover common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getStakersWithPendingUnstakes", prover)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakersWithPendingUnstakes is a free data retrieval call binding the contract method 0xff850d2a.
//
// Solidity: function getStakersWithPendingUnstakes(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerSession) GetStakersWithPendingUnstakes(prover common.Address) ([]common.Address, error) {
	return _IStakingController.Contract.GetStakersWithPendingUnstakes(&_IStakingController.CallOpts, prover)
}

// GetStakersWithPendingUnstakes is a free data retrieval call binding the contract method 0xff850d2a.
//
// Solidity: function getStakersWithPendingUnstakes(address prover) view returns(address[] stakers)
func (_IStakingController *IStakingControllerCallerSession) GetStakersWithPendingUnstakes(prover common.Address) ([]common.Address, error) {
	return _IStakingController.Contract.GetStakersWithPendingUnstakes(&_IStakingController.CallOpts, prover)
}

// GetStakersWithPendingUnstakesCount is a free data retrieval call binding the contract method 0x78e5c914.
//
// Solidity: function getStakersWithPendingUnstakesCount(address prover) view returns(uint256 count)
func (_IStakingController *IStakingControllerCaller) GetStakersWithPendingUnstakesCount(opts *bind.CallOpts, prover common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getStakersWithPendingUnstakesCount", prover)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakersWithPendingUnstakesCount is a free data retrieval call binding the contract method 0x78e5c914.
//
// Solidity: function getStakersWithPendingUnstakesCount(address prover) view returns(uint256 count)
func (_IStakingController *IStakingControllerSession) GetStakersWithPendingUnstakesCount(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetStakersWithPendingUnstakesCount(&_IStakingController.CallOpts, prover)
}

// GetStakersWithPendingUnstakesCount is a free data retrieval call binding the contract method 0x78e5c914.
//
// Solidity: function getStakersWithPendingUnstakesCount(address prover) view returns(uint256 count)
func (_IStakingController *IStakingControllerCallerSession) GetStakersWithPendingUnstakesCount(prover common.Address) (*big.Int, error) {
	return _IStakingController.Contract.GetStakersWithPendingUnstakesCount(&_IStakingController.CallOpts, prover)
}

// GetTotalActiveProverVaultAssets is a free data retrieval call binding the contract method 0x614ebc8d.
//
// Solidity: function getTotalActiveProverVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCaller) GetTotalActiveProverVaultAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getTotalActiveProverVaultAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveProverVaultAssets is a free data retrieval call binding the contract method 0x614ebc8d.
//
// Solidity: function getTotalActiveProverVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerSession) GetTotalActiveProverVaultAssets() (*big.Int, error) {
	return _IStakingController.Contract.GetTotalActiveProverVaultAssets(&_IStakingController.CallOpts)
}

// GetTotalActiveProverVaultAssets is a free data retrieval call binding the contract method 0x614ebc8d.
//
// Solidity: function getTotalActiveProverVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCallerSession) GetTotalActiveProverVaultAssets() (*big.Int, error) {
	return _IStakingController.Contract.GetTotalActiveProverVaultAssets(&_IStakingController.CallOpts)
}

// GetTotalVaultAssets is a free data retrieval call binding the contract method 0x7a95f5fe.
//
// Solidity: function getTotalVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCaller) GetTotalVaultAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getTotalVaultAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVaultAssets is a free data retrieval call binding the contract method 0x7a95f5fe.
//
// Solidity: function getTotalVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerSession) GetTotalVaultAssets() (*big.Int, error) {
	return _IStakingController.Contract.GetTotalVaultAssets(&_IStakingController.CallOpts)
}

// GetTotalVaultAssets is a free data retrieval call binding the contract method 0x7a95f5fe.
//
// Solidity: function getTotalVaultAssets() view returns(uint256 totalAssets)
func (_IStakingController *IStakingControllerCallerSession) GetTotalVaultAssets() (*big.Int, error) {
	return _IStakingController.Contract.GetTotalVaultAssets(&_IStakingController.CallOpts)
}

// GetUnstakingInfo is a free data retrieval call binding the contract method 0xd4fab088.
//
// Solidity: function getUnstakingInfo(address prover, address staker) view returns(uint256 totalAmount, uint256 readyAmount)
func (_IStakingController *IStakingControllerCaller) GetUnstakingInfo(opts *bind.CallOpts, prover common.Address, staker common.Address) (struct {
	TotalAmount *big.Int
	ReadyAmount *big.Int
}, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "getUnstakingInfo", prover, staker)

	outstruct := new(struct {
		TotalAmount *big.Int
		ReadyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReadyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnstakingInfo is a free data retrieval call binding the contract method 0xd4fab088.
//
// Solidity: function getUnstakingInfo(address prover, address staker) view returns(uint256 totalAmount, uint256 readyAmount)
func (_IStakingController *IStakingControllerSession) GetUnstakingInfo(prover common.Address, staker common.Address) (struct {
	TotalAmount *big.Int
	ReadyAmount *big.Int
}, error) {
	return _IStakingController.Contract.GetUnstakingInfo(&_IStakingController.CallOpts, prover, staker)
}

// GetUnstakingInfo is a free data retrieval call binding the contract method 0xd4fab088.
//
// Solidity: function getUnstakingInfo(address prover, address staker) view returns(uint256 totalAmount, uint256 readyAmount)
func (_IStakingController *IStakingControllerCallerSession) GetUnstakingInfo(prover common.Address, staker common.Address) (struct {
	TotalAmount *big.Int
	ReadyAmount *big.Int
}, error) {
	return _IStakingController.Contract.GetUnstakingInfo(&_IStakingController.CallOpts, prover, staker)
}

// IsProverEligible is a free data retrieval call binding the contract method 0xad732cc1.
//
// Solidity: function isProverEligible(address prover, uint256 minimumVaultAssets) view returns(bool eligible, uint256 currentVaultAssets)
func (_IStakingController *IStakingControllerCaller) IsProverEligible(opts *bind.CallOpts, prover common.Address, minimumVaultAssets *big.Int) (struct {
	Eligible           bool
	CurrentVaultAssets *big.Int
}, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "isProverEligible", prover, minimumVaultAssets)

	outstruct := new(struct {
		Eligible           bool
		CurrentVaultAssets *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Eligible = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CurrentVaultAssets = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IsProverEligible is a free data retrieval call binding the contract method 0xad732cc1.
//
// Solidity: function isProverEligible(address prover, uint256 minimumVaultAssets) view returns(bool eligible, uint256 currentVaultAssets)
func (_IStakingController *IStakingControllerSession) IsProverEligible(prover common.Address, minimumVaultAssets *big.Int) (struct {
	Eligible           bool
	CurrentVaultAssets *big.Int
}, error) {
	return _IStakingController.Contract.IsProverEligible(&_IStakingController.CallOpts, prover, minimumVaultAssets)
}

// IsProverEligible is a free data retrieval call binding the contract method 0xad732cc1.
//
// Solidity: function isProverEligible(address prover, uint256 minimumVaultAssets) view returns(bool eligible, uint256 currentVaultAssets)
func (_IStakingController *IStakingControllerCallerSession) IsProverEligible(prover common.Address, minimumVaultAssets *big.Int) (struct {
	Eligible           bool
	CurrentVaultAssets *big.Int
}, error) {
	return _IStakingController.Contract.IsProverEligible(&_IStakingController.CallOpts, prover, minimumVaultAssets)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x76016870.
//
// Solidity: function maxDeposit(address prover, address receiver) view returns(uint256 amount)
func (_IStakingController *IStakingControllerCaller) MaxDeposit(opts *bind.CallOpts, prover common.Address, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "maxDeposit", prover, receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x76016870.
//
// Solidity: function maxDeposit(address prover, address receiver) view returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) MaxDeposit(prover common.Address, receiver common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxDeposit(&_IStakingController.CallOpts, prover, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x76016870.
//
// Solidity: function maxDeposit(address prover, address receiver) view returns(uint256 amount)
func (_IStakingController *IStakingControllerCallerSession) MaxDeposit(prover common.Address, receiver common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxDeposit(&_IStakingController.CallOpts, prover, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xf2e586db.
//
// Solidity: function maxMint(address prover, address receiver) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCaller) MaxMint(opts *bind.CallOpts, prover common.Address, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "maxMint", prover, receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xf2e586db.
//
// Solidity: function maxMint(address prover, address receiver) view returns(uint256 shares)
func (_IStakingController *IStakingControllerSession) MaxMint(prover common.Address, receiver common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxMint(&_IStakingController.CallOpts, prover, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xf2e586db.
//
// Solidity: function maxMint(address prover, address receiver) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCallerSession) MaxMint(prover common.Address, receiver common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxMint(&_IStakingController.CallOpts, prover, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0x95b734fb.
//
// Solidity: function maxRedeem(address prover, address owner) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCaller) MaxRedeem(opts *bind.CallOpts, prover common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "maxRedeem", prover, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0x95b734fb.
//
// Solidity: function maxRedeem(address prover, address owner) view returns(uint256 shares)
func (_IStakingController *IStakingControllerSession) MaxRedeem(prover common.Address, owner common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxRedeem(&_IStakingController.CallOpts, prover, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0x95b734fb.
//
// Solidity: function maxRedeem(address prover, address owner) view returns(uint256 shares)
func (_IStakingController *IStakingControllerCallerSession) MaxRedeem(prover common.Address, owner common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxRedeem(&_IStakingController.CallOpts, prover, owner)
}

// MaxSlashBps is a free data retrieval call binding the contract method 0xab6e48db.
//
// Solidity: function maxSlashBps() view returns(uint256 bps)
func (_IStakingController *IStakingControllerCaller) MaxSlashBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "maxSlashBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSlashBps is a free data retrieval call binding the contract method 0xab6e48db.
//
// Solidity: function maxSlashBps() view returns(uint256 bps)
func (_IStakingController *IStakingControllerSession) MaxSlashBps() (*big.Int, error) {
	return _IStakingController.Contract.MaxSlashBps(&_IStakingController.CallOpts)
}

// MaxSlashBps is a free data retrieval call binding the contract method 0xab6e48db.
//
// Solidity: function maxSlashBps() view returns(uint256 bps)
func (_IStakingController *IStakingControllerCallerSession) MaxSlashBps() (*big.Int, error) {
	return _IStakingController.Contract.MaxSlashBps(&_IStakingController.CallOpts)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0x8c22e9b0.
//
// Solidity: function maxWithdraw(address prover, address owner) view returns(uint256 amount)
func (_IStakingController *IStakingControllerCaller) MaxWithdraw(opts *bind.CallOpts, prover common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "maxWithdraw", prover, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0x8c22e9b0.
//
// Solidity: function maxWithdraw(address prover, address owner) view returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) MaxWithdraw(prover common.Address, owner common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxWithdraw(&_IStakingController.CallOpts, prover, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0x8c22e9b0.
//
// Solidity: function maxWithdraw(address prover, address owner) view returns(uint256 amount)
func (_IStakingController *IStakingControllerCallerSession) MaxWithdraw(prover common.Address, owner common.Address) (*big.Int, error) {
	return _IStakingController.Contract.MaxWithdraw(&_IStakingController.CallOpts, prover, owner)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256 amount)
func (_IStakingController *IStakingControllerCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "minSelfStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) MinSelfStake() (*big.Int, error) {
	return _IStakingController.Contract.MinSelfStake(&_IStakingController.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256 amount)
func (_IStakingController *IStakingControllerCallerSession) MinSelfStake() (*big.Int, error) {
	return _IStakingController.Contract.MinSelfStake(&_IStakingController.CallOpts)
}

// StakerHasPendingUnstakes is a free data retrieval call binding the contract method 0xd7c5c24c.
//
// Solidity: function stakerHasPendingUnstakes(address prover, address staker) view returns(bool hasPending)
func (_IStakingController *IStakingControllerCaller) StakerHasPendingUnstakes(opts *bind.CallOpts, prover common.Address, staker common.Address) (bool, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "stakerHasPendingUnstakes", prover, staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakerHasPendingUnstakes is a free data retrieval call binding the contract method 0xd7c5c24c.
//
// Solidity: function stakerHasPendingUnstakes(address prover, address staker) view returns(bool hasPending)
func (_IStakingController *IStakingControllerSession) StakerHasPendingUnstakes(prover common.Address, staker common.Address) (bool, error) {
	return _IStakingController.Contract.StakerHasPendingUnstakes(&_IStakingController.CallOpts, prover, staker)
}

// StakerHasPendingUnstakes is a free data retrieval call binding the contract method 0xd7c5c24c.
//
// Solidity: function stakerHasPendingUnstakes(address prover, address staker) view returns(bool hasPending)
func (_IStakingController *IStakingControllerCallerSession) StakerHasPendingUnstakes(prover common.Address, staker common.Address) (bool, error) {
	return _IStakingController.Contract.StakerHasPendingUnstakes(&_IStakingController.CallOpts, prover, staker)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address token)
func (_IStakingController *IStakingControllerCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStakingController.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address token)
func (_IStakingController *IStakingControllerSession) StakingToken() (common.Address, error) {
	return _IStakingController.Contract.StakingToken(&_IStakingController.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address token)
func (_IStakingController *IStakingControllerCallerSession) StakingToken() (common.Address, error) {
	return _IStakingController.Contract.StakingToken(&_IStakingController.CallOpts)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address prover, uint256 amount) returns(uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerTransactor) AddRewards(opts *bind.TransactOpts, prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "addRewards", prover, amount)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address prover, uint256 amount) returns(uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerSession) AddRewards(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.AddRewards(&_IStakingController.TransactOpts, prover, amount)
}

// AddRewards is a paid mutator transaction binding the contract method 0xa9fc507b.
//
// Solidity: function addRewards(address prover, uint256 amount) returns(uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerTransactorSession) AddRewards(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.AddRewards(&_IStakingController.TransactOpts, prover, amount)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactor) ClaimCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "claimCommission")
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) ClaimCommission() (*types.Transaction, error) {
	return _IStakingController.Contract.ClaimCommission(&_IStakingController.TransactOpts)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactorSession) ClaimCommission() (*types.Transaction, error) {
	return _IStakingController.Contract.ClaimCommission(&_IStakingController.TransactOpts)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x08a1d322.
//
// Solidity: function completeUnstake(address prover) returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactor) CompleteUnstake(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "completeUnstake", prover)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x08a1d322.
//
// Solidity: function completeUnstake(address prover) returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) CompleteUnstake(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.CompleteUnstake(&_IStakingController.TransactOpts, prover)
}

// CompleteUnstake is a paid mutator transaction binding the contract method 0x08a1d322.
//
// Solidity: function completeUnstake(address prover) returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactorSession) CompleteUnstake(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.CompleteUnstake(&_IStakingController.TransactOpts, prover)
}

// DeactivateProver is a paid mutator transaction binding the contract method 0xd61d3ef3.
//
// Solidity: function deactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactor) DeactivateProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "deactivateProver", prover)
}

// DeactivateProver is a paid mutator transaction binding the contract method 0xd61d3ef3.
//
// Solidity: function deactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerSession) DeactivateProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.DeactivateProver(&_IStakingController.TransactOpts, prover)
}

// DeactivateProver is a paid mutator transaction binding the contract method 0xd61d3ef3.
//
// Solidity: function deactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactorSession) DeactivateProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.DeactivateProver(&_IStakingController.TransactOpts, prover)
}

// EmergencyRecover is a paid mutator transaction binding the contract method 0x9639011e.
//
// Solidity: function emergencyRecover(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerTransactor) EmergencyRecover(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "emergencyRecover", to, amount)
}

// EmergencyRecover is a paid mutator transaction binding the contract method 0x9639011e.
//
// Solidity: function emergencyRecover(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerSession) EmergencyRecover(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.EmergencyRecover(&_IStakingController.TransactOpts, to, amount)
}

// EmergencyRecover is a paid mutator transaction binding the contract method 0x9639011e.
//
// Solidity: function emergencyRecover(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerTransactorSession) EmergencyRecover(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.EmergencyRecover(&_IStakingController.TransactOpts, to, amount)
}

// InitializeProver is a paid mutator transaction binding the contract method 0xf97e87ed.
//
// Solidity: function initializeProver(uint64 defaultCommissionRate) returns(address vault)
func (_IStakingController *IStakingControllerTransactor) InitializeProver(opts *bind.TransactOpts, defaultCommissionRate uint64) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "initializeProver", defaultCommissionRate)
}

// InitializeProver is a paid mutator transaction binding the contract method 0xf97e87ed.
//
// Solidity: function initializeProver(uint64 defaultCommissionRate) returns(address vault)
func (_IStakingController *IStakingControllerSession) InitializeProver(defaultCommissionRate uint64) (*types.Transaction, error) {
	return _IStakingController.Contract.InitializeProver(&_IStakingController.TransactOpts, defaultCommissionRate)
}

// InitializeProver is a paid mutator transaction binding the contract method 0xf97e87ed.
//
// Solidity: function initializeProver(uint64 defaultCommissionRate) returns(address vault)
func (_IStakingController *IStakingControllerTransactorSession) InitializeProver(defaultCommissionRate uint64) (*types.Transaction, error) {
	return _IStakingController.Contract.InitializeProver(&_IStakingController.TransactOpts, defaultCommissionRate)
}

// JailProver is a paid mutator transaction binding the contract method 0x8ecdc11f.
//
// Solidity: function jailProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactor) JailProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "jailProver", prover)
}

// JailProver is a paid mutator transaction binding the contract method 0x8ecdc11f.
//
// Solidity: function jailProver(address prover) returns()
func (_IStakingController *IStakingControllerSession) JailProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.JailProver(&_IStakingController.TransactOpts, prover)
}

// JailProver is a paid mutator transaction binding the contract method 0x8ecdc11f.
//
// Solidity: function jailProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactorSession) JailProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.JailProver(&_IStakingController.TransactOpts, prover)
}

// OnShareTransfer is a paid mutator transaction binding the contract method 0x86a99082.
//
// Solidity: function onShareTransfer(address prover, address from, address to, uint256 shares) returns()
func (_IStakingController *IStakingControllerTransactor) OnShareTransfer(opts *bind.TransactOpts, prover common.Address, from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "onShareTransfer", prover, from, to, shares)
}

// OnShareTransfer is a paid mutator transaction binding the contract method 0x86a99082.
//
// Solidity: function onShareTransfer(address prover, address from, address to, uint256 shares) returns()
func (_IStakingController *IStakingControllerSession) OnShareTransfer(prover common.Address, from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.OnShareTransfer(&_IStakingController.TransactOpts, prover, from, to, shares)
}

// OnShareTransfer is a paid mutator transaction binding the contract method 0x86a99082.
//
// Solidity: function onShareTransfer(address prover, address from, address to, uint256 shares) returns()
func (_IStakingController *IStakingControllerTransactorSession) OnShareTransfer(prover common.Address, from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.OnShareTransfer(&_IStakingController.TransactOpts, prover, from, to, shares)
}

// ReactivateProver is a paid mutator transaction binding the contract method 0x866ce885.
//
// Solidity: function reactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactor) ReactivateProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "reactivateProver", prover)
}

// ReactivateProver is a paid mutator transaction binding the contract method 0x866ce885.
//
// Solidity: function reactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerSession) ReactivateProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.ReactivateProver(&_IStakingController.TransactOpts, prover)
}

// ReactivateProver is a paid mutator transaction binding the contract method 0x866ce885.
//
// Solidity: function reactivateProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactorSession) ReactivateProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.ReactivateProver(&_IStakingController.TransactOpts, prover)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x710ba631.
//
// Solidity: function requestUnstake(address prover, uint256 shares) returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactor) RequestUnstake(opts *bind.TransactOpts, prover common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "requestUnstake", prover, shares)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x710ba631.
//
// Solidity: function requestUnstake(address prover, uint256 shares) returns(uint256 amount)
func (_IStakingController *IStakingControllerSession) RequestUnstake(prover common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.RequestUnstake(&_IStakingController.TransactOpts, prover, shares)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x710ba631.
//
// Solidity: function requestUnstake(address prover, uint256 shares) returns(uint256 amount)
func (_IStakingController *IStakingControllerTransactorSession) RequestUnstake(prover common.Address, shares *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.RequestUnstake(&_IStakingController.TransactOpts, prover, shares)
}

// ResetCommissionRate is a paid mutator transaction binding the contract method 0x64176bab.
//
// Solidity: function resetCommissionRate(address source) returns()
func (_IStakingController *IStakingControllerTransactor) ResetCommissionRate(opts *bind.TransactOpts, source common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "resetCommissionRate", source)
}

// ResetCommissionRate is a paid mutator transaction binding the contract method 0x64176bab.
//
// Solidity: function resetCommissionRate(address source) returns()
func (_IStakingController *IStakingControllerSession) ResetCommissionRate(source common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.ResetCommissionRate(&_IStakingController.TransactOpts, source)
}

// ResetCommissionRate is a paid mutator transaction binding the contract method 0x64176bab.
//
// Solidity: function resetCommissionRate(address source) returns()
func (_IStakingController *IStakingControllerTransactorSession) ResetCommissionRate(source common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.ResetCommissionRate(&_IStakingController.TransactOpts, source)
}

// RetireProver is a paid mutator transaction binding the contract method 0x8b699ee1.
//
// Solidity: function retireProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactor) RetireProver(opts *bind.TransactOpts, prover common.Address) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "retireProver", prover)
}

// RetireProver is a paid mutator transaction binding the contract method 0x8b699ee1.
//
// Solidity: function retireProver(address prover) returns()
func (_IStakingController *IStakingControllerSession) RetireProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.RetireProver(&_IStakingController.TransactOpts, prover)
}

// RetireProver is a paid mutator transaction binding the contract method 0x8b699ee1.
//
// Solidity: function retireProver(address prover) returns()
func (_IStakingController *IStakingControllerTransactorSession) RetireProver(prover common.Address) (*types.Transaction, error) {
	return _IStakingController.Contract.RetireProver(&_IStakingController.TransactOpts, prover)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x988a894e.
//
// Solidity: function setCommissionRate(address source, uint64 newRate) returns()
func (_IStakingController *IStakingControllerTransactor) SetCommissionRate(opts *bind.TransactOpts, source common.Address, newRate uint64) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "setCommissionRate", source, newRate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x988a894e.
//
// Solidity: function setCommissionRate(address source, uint64 newRate) returns()
func (_IStakingController *IStakingControllerSession) SetCommissionRate(source common.Address, newRate uint64) (*types.Transaction, error) {
	return _IStakingController.Contract.SetCommissionRate(&_IStakingController.TransactOpts, source, newRate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x988a894e.
//
// Solidity: function setCommissionRate(address source, uint64 newRate) returns()
func (_IStakingController *IStakingControllerTransactorSession) SetCommissionRate(source common.Address, newRate uint64) (*types.Transaction, error) {
	return _IStakingController.Contract.SetCommissionRate(&_IStakingController.TransactOpts, source, newRate)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 value) returns()
func (_IStakingController *IStakingControllerTransactor) SetMaxSlashBps(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "setMaxSlashBps", value)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 value) returns()
func (_IStakingController *IStakingControllerSession) SetMaxSlashBps(value *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetMaxSlashBps(&_IStakingController.TransactOpts, value)
}

// SetMaxSlashBps is a paid mutator transaction binding the contract method 0x3dcc1ddd.
//
// Solidity: function setMaxSlashBps(uint256 value) returns()
func (_IStakingController *IStakingControllerTransactorSession) SetMaxSlashBps(value *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetMaxSlashBps(&_IStakingController.TransactOpts, value)
}

// SetMinSelfStake is a paid mutator transaction binding the contract method 0xd75629cf.
//
// Solidity: function setMinSelfStake(uint256 value) returns()
func (_IStakingController *IStakingControllerTransactor) SetMinSelfStake(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "setMinSelfStake", value)
}

// SetMinSelfStake is a paid mutator transaction binding the contract method 0xd75629cf.
//
// Solidity: function setMinSelfStake(uint256 value) returns()
func (_IStakingController *IStakingControllerSession) SetMinSelfStake(value *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetMinSelfStake(&_IStakingController.TransactOpts, value)
}

// SetMinSelfStake is a paid mutator transaction binding the contract method 0xd75629cf.
//
// Solidity: function setMinSelfStake(uint256 value) returns()
func (_IStakingController *IStakingControllerTransactorSession) SetMinSelfStake(value *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetMinSelfStake(&_IStakingController.TransactOpts, value)
}

// SetRequireAuthorization is a paid mutator transaction binding the contract method 0xa9aadaad.
//
// Solidity: function setRequireAuthorization(bool required) returns()
func (_IStakingController *IStakingControllerTransactor) SetRequireAuthorization(opts *bind.TransactOpts, required bool) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "setRequireAuthorization", required)
}

// SetRequireAuthorization is a paid mutator transaction binding the contract method 0xa9aadaad.
//
// Solidity: function setRequireAuthorization(bool required) returns()
func (_IStakingController *IStakingControllerSession) SetRequireAuthorization(required bool) (*types.Transaction, error) {
	return _IStakingController.Contract.SetRequireAuthorization(&_IStakingController.TransactOpts, required)
}

// SetRequireAuthorization is a paid mutator transaction binding the contract method 0xa9aadaad.
//
// Solidity: function setRequireAuthorization(bool required) returns()
func (_IStakingController *IStakingControllerTransactorSession) SetRequireAuthorization(required bool) (*types.Transaction, error) {
	return _IStakingController.Contract.SetRequireAuthorization(&_IStakingController.TransactOpts, required)
}

// SetUnstakeDelay is a paid mutator transaction binding the contract method 0x30b75e04.
//
// Solidity: function setUnstakeDelay(uint256 newDelay) returns()
func (_IStakingController *IStakingControllerTransactor) SetUnstakeDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "setUnstakeDelay", newDelay)
}

// SetUnstakeDelay is a paid mutator transaction binding the contract method 0x30b75e04.
//
// Solidity: function setUnstakeDelay(uint256 newDelay) returns()
func (_IStakingController *IStakingControllerSession) SetUnstakeDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetUnstakeDelay(&_IStakingController.TransactOpts, newDelay)
}

// SetUnstakeDelay is a paid mutator transaction binding the contract method 0x30b75e04.
//
// Solidity: function setUnstakeDelay(uint256 newDelay) returns()
func (_IStakingController *IStakingControllerTransactorSession) SetUnstakeDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SetUnstakeDelay(&_IStakingController.TransactOpts, newDelay)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address prover, uint256 bps) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerTransactor) Slash(opts *bind.TransactOpts, prover common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "slash", prover, bps)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address prover, uint256 bps) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerSession) Slash(prover common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.Slash(&_IStakingController.TransactOpts, prover, bps)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address prover, uint256 bps) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerTransactorSession) Slash(prover common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.Slash(&_IStakingController.TransactOpts, prover, bps)
}

// SlashByAmount is a paid mutator transaction binding the contract method 0x3046198c.
//
// Solidity: function slashByAmount(address prover, uint256 amount) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerTransactor) SlashByAmount(opts *bind.TransactOpts, prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "slashByAmount", prover, amount)
}

// SlashByAmount is a paid mutator transaction binding the contract method 0x3046198c.
//
// Solidity: function slashByAmount(address prover, uint256 amount) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerSession) SlashByAmount(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SlashByAmount(&_IStakingController.TransactOpts, prover, amount)
}

// SlashByAmount is a paid mutator transaction binding the contract method 0x3046198c.
//
// Solidity: function slashByAmount(address prover, uint256 amount) returns(uint256 slashedAmount)
func (_IStakingController *IStakingControllerTransactorSession) SlashByAmount(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.SlashByAmount(&_IStakingController.TransactOpts, prover, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address prover, uint256 amount) returns(uint256 shares)
func (_IStakingController *IStakingControllerTransactor) Stake(opts *bind.TransactOpts, prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "stake", prover, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address prover, uint256 amount) returns(uint256 shares)
func (_IStakingController *IStakingControllerSession) Stake(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.Stake(&_IStakingController.TransactOpts, prover, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address prover, uint256 amount) returns(uint256 shares)
func (_IStakingController *IStakingControllerTransactorSession) Stake(prover common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.Stake(&_IStakingController.TransactOpts, prover, amount)
}

// WithdrawTreasury is a paid mutator transaction binding the contract method 0x0d86419a.
//
// Solidity: function withdrawTreasury(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerTransactor) WithdrawTreasury(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.contract.Transact(opts, "withdrawTreasury", to, amount)
}

// WithdrawTreasury is a paid mutator transaction binding the contract method 0x0d86419a.
//
// Solidity: function withdrawTreasury(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerSession) WithdrawTreasury(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.WithdrawTreasury(&_IStakingController.TransactOpts, to, amount)
}

// WithdrawTreasury is a paid mutator transaction binding the contract method 0x0d86419a.
//
// Solidity: function withdrawTreasury(address to, uint256 amount) returns()
func (_IStakingController *IStakingControllerTransactorSession) WithdrawTreasury(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStakingController.Contract.WithdrawTreasury(&_IStakingController.TransactOpts, to, amount)
}

// IStakingControllerCommissionClaimedIterator is returned from FilterCommissionClaimed and is used to iterate over the raw logs and unpacked data for CommissionClaimed events raised by the IStakingController contract.
type IStakingControllerCommissionClaimedIterator struct {
	Event *IStakingControllerCommissionClaimed // Event containing the contract specifics and raw log

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
func (it *IStakingControllerCommissionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerCommissionClaimed)
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
		it.Event = new(IStakingControllerCommissionClaimed)
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
func (it *IStakingControllerCommissionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerCommissionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerCommissionClaimed represents a CommissionClaimed event raised by the IStakingController contract.
type IStakingControllerCommissionClaimed struct {
	Prover common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCommissionClaimed is a free log retrieval operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed prover, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterCommissionClaimed(opts *bind.FilterOpts, prover []common.Address) (*IStakingControllerCommissionClaimedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "CommissionClaimed", proverRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerCommissionClaimedIterator{contract: _IStakingController.contract, event: "CommissionClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionClaimed is a free log subscription operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed prover, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchCommissionClaimed(opts *bind.WatchOpts, sink chan<- *IStakingControllerCommissionClaimed, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "CommissionClaimed", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerCommissionClaimed)
				if err := _IStakingController.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
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

// ParseCommissionClaimed is a log parse operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed prover, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseCommissionClaimed(log types.Log) (*IStakingControllerCommissionClaimed, error) {
	event := new(IStakingControllerCommissionClaimed)
	if err := _IStakingController.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerCommissionRateUpdatedIterator is returned from FilterCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for CommissionRateUpdated events raised by the IStakingController contract.
type IStakingControllerCommissionRateUpdatedIterator struct {
	Event *IStakingControllerCommissionRateUpdated // Event containing the contract specifics and raw log

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
func (it *IStakingControllerCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerCommissionRateUpdated)
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
		it.Event = new(IStakingControllerCommissionRateUpdated)
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
func (it *IStakingControllerCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerCommissionRateUpdated represents a CommissionRateUpdated event raised by the IStakingController contract.
type IStakingControllerCommissionRateUpdated struct {
	Prover     common.Address
	Source     common.Address
	OldRate    uint64
	NewRate    uint64
	UseDefault bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdated is a free log retrieval operation binding the contract event 0x3e0ead68f9d0779831d43703ee338effa29c4fb4fe5094bc13b7f94ee5551322.
//
// Solidity: event CommissionRateUpdated(address indexed prover, address indexed source, uint64 oldRate, uint64 newRate, bool useDefault)
func (_IStakingController *IStakingControllerFilterer) FilterCommissionRateUpdated(opts *bind.FilterOpts, prover []common.Address, source []common.Address) (*IStakingControllerCommissionRateUpdatedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "CommissionRateUpdated", proverRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerCommissionRateUpdatedIterator{contract: _IStakingController.contract, event: "CommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdated is a free log subscription operation binding the contract event 0x3e0ead68f9d0779831d43703ee338effa29c4fb4fe5094bc13b7f94ee5551322.
//
// Solidity: event CommissionRateUpdated(address indexed prover, address indexed source, uint64 oldRate, uint64 newRate, bool useDefault)
func (_IStakingController *IStakingControllerFilterer) WatchCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *IStakingControllerCommissionRateUpdated, prover []common.Address, source []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "CommissionRateUpdated", proverRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerCommissionRateUpdated)
				if err := _IStakingController.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
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

// ParseCommissionRateUpdated is a log parse operation binding the contract event 0x3e0ead68f9d0779831d43703ee338effa29c4fb4fe5094bc13b7f94ee5551322.
//
// Solidity: event CommissionRateUpdated(address indexed prover, address indexed source, uint64 oldRate, uint64 newRate, bool useDefault)
func (_IStakingController *IStakingControllerFilterer) ParseCommissionRateUpdated(log types.Log) (*IStakingControllerCommissionRateUpdated, error) {
	event := new(IStakingControllerCommissionRateUpdated)
	if err := _IStakingController.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerEmergencyRecoveredIterator is returned from FilterEmergencyRecovered and is used to iterate over the raw logs and unpacked data for EmergencyRecovered events raised by the IStakingController contract.
type IStakingControllerEmergencyRecoveredIterator struct {
	Event *IStakingControllerEmergencyRecovered // Event containing the contract specifics and raw log

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
func (it *IStakingControllerEmergencyRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerEmergencyRecovered)
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
		it.Event = new(IStakingControllerEmergencyRecovered)
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
func (it *IStakingControllerEmergencyRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerEmergencyRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerEmergencyRecovered represents a EmergencyRecovered event raised by the IStakingController contract.
type IStakingControllerEmergencyRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyRecovered is a free log retrieval operation binding the contract event 0x1115307ac794dd199ebee09e669ca430434628a9d4c1cf126f4fe8d2202ce702.
//
// Solidity: event EmergencyRecovered(address to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterEmergencyRecovered(opts *bind.FilterOpts) (*IStakingControllerEmergencyRecoveredIterator, error) {

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "EmergencyRecovered")
	if err != nil {
		return nil, err
	}
	return &IStakingControllerEmergencyRecoveredIterator{contract: _IStakingController.contract, event: "EmergencyRecovered", logs: logs, sub: sub}, nil
}

// WatchEmergencyRecovered is a free log subscription operation binding the contract event 0x1115307ac794dd199ebee09e669ca430434628a9d4c1cf126f4fe8d2202ce702.
//
// Solidity: event EmergencyRecovered(address to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchEmergencyRecovered(opts *bind.WatchOpts, sink chan<- *IStakingControllerEmergencyRecovered) (event.Subscription, error) {

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "EmergencyRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerEmergencyRecovered)
				if err := _IStakingController.contract.UnpackLog(event, "EmergencyRecovered", log); err != nil {
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

// ParseEmergencyRecovered is a log parse operation binding the contract event 0x1115307ac794dd199ebee09e669ca430434628a9d4c1cf126f4fe8d2202ce702.
//
// Solidity: event EmergencyRecovered(address to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseEmergencyRecovered(log types.Log) (*IStakingControllerEmergencyRecovered, error) {
	event := new(IStakingControllerEmergencyRecovered)
	if err := _IStakingController.contract.UnpackLog(event, "EmergencyRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerMaxSlashBpsUpdatedIterator is returned from FilterMaxSlashBpsUpdated and is used to iterate over the raw logs and unpacked data for MaxSlashBpsUpdated events raised by the IStakingController contract.
type IStakingControllerMaxSlashBpsUpdatedIterator struct {
	Event *IStakingControllerMaxSlashBpsUpdated // Event containing the contract specifics and raw log

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
func (it *IStakingControllerMaxSlashBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerMaxSlashBpsUpdated)
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
		it.Event = new(IStakingControllerMaxSlashBpsUpdated)
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
func (it *IStakingControllerMaxSlashBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerMaxSlashBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerMaxSlashBpsUpdated represents a MaxSlashBpsUpdated event raised by the IStakingController contract.
type IStakingControllerMaxSlashBpsUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMaxSlashBpsUpdated is a free log retrieval operation binding the contract event 0x67c28ae08cb0596db1556dd74583a7f5de2f54688e7dc389176a79f2c16627db.
//
// Solidity: event MaxSlashBpsUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) FilterMaxSlashBpsUpdated(opts *bind.FilterOpts) (*IStakingControllerMaxSlashBpsUpdatedIterator, error) {

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "MaxSlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &IStakingControllerMaxSlashBpsUpdatedIterator{contract: _IStakingController.contract, event: "MaxSlashBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSlashBpsUpdated is a free log subscription operation binding the contract event 0x67c28ae08cb0596db1556dd74583a7f5de2f54688e7dc389176a79f2c16627db.
//
// Solidity: event MaxSlashBpsUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) WatchMaxSlashBpsUpdated(opts *bind.WatchOpts, sink chan<- *IStakingControllerMaxSlashBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "MaxSlashBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerMaxSlashBpsUpdated)
				if err := _IStakingController.contract.UnpackLog(event, "MaxSlashBpsUpdated", log); err != nil {
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

// ParseMaxSlashBpsUpdated is a log parse operation binding the contract event 0x67c28ae08cb0596db1556dd74583a7f5de2f54688e7dc389176a79f2c16627db.
//
// Solidity: event MaxSlashBpsUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) ParseMaxSlashBpsUpdated(log types.Log) (*IStakingControllerMaxSlashBpsUpdated, error) {
	event := new(IStakingControllerMaxSlashBpsUpdated)
	if err := _IStakingController.contract.UnpackLog(event, "MaxSlashBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerMinSelfStakeUpdatedIterator is returned from FilterMinSelfStakeUpdated and is used to iterate over the raw logs and unpacked data for MinSelfStakeUpdated events raised by the IStakingController contract.
type IStakingControllerMinSelfStakeUpdatedIterator struct {
	Event *IStakingControllerMinSelfStakeUpdated // Event containing the contract specifics and raw log

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
func (it *IStakingControllerMinSelfStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerMinSelfStakeUpdated)
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
		it.Event = new(IStakingControllerMinSelfStakeUpdated)
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
func (it *IStakingControllerMinSelfStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerMinSelfStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerMinSelfStakeUpdated represents a MinSelfStakeUpdated event raised by the IStakingController contract.
type IStakingControllerMinSelfStakeUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinSelfStakeUpdated is a free log retrieval operation binding the contract event 0x712b3a39aec1f14a0d99c01c87ef87e200d8f9f958a7a62b969b1e4333dc8e57.
//
// Solidity: event MinSelfStakeUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) FilterMinSelfStakeUpdated(opts *bind.FilterOpts) (*IStakingControllerMinSelfStakeUpdatedIterator, error) {

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "MinSelfStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &IStakingControllerMinSelfStakeUpdatedIterator{contract: _IStakingController.contract, event: "MinSelfStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinSelfStakeUpdated is a free log subscription operation binding the contract event 0x712b3a39aec1f14a0d99c01c87ef87e200d8f9f958a7a62b969b1e4333dc8e57.
//
// Solidity: event MinSelfStakeUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) WatchMinSelfStakeUpdated(opts *bind.WatchOpts, sink chan<- *IStakingControllerMinSelfStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "MinSelfStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerMinSelfStakeUpdated)
				if err := _IStakingController.contract.UnpackLog(event, "MinSelfStakeUpdated", log); err != nil {
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

// ParseMinSelfStakeUpdated is a log parse operation binding the contract event 0x712b3a39aec1f14a0d99c01c87ef87e200d8f9f958a7a62b969b1e4333dc8e57.
//
// Solidity: event MinSelfStakeUpdated(uint256 oldValue, uint256 newValue)
func (_IStakingController *IStakingControllerFilterer) ParseMinSelfStakeUpdated(log types.Log) (*IStakingControllerMinSelfStakeUpdated, error) {
	event := new(IStakingControllerMinSelfStakeUpdated)
	if err := _IStakingController.contract.UnpackLog(event, "MinSelfStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerProverInitializedIterator is returned from FilterProverInitialized and is used to iterate over the raw logs and unpacked data for ProverInitialized events raised by the IStakingController contract.
type IStakingControllerProverInitializedIterator struct {
	Event *IStakingControllerProverInitialized // Event containing the contract specifics and raw log

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
func (it *IStakingControllerProverInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerProverInitialized)
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
		it.Event = new(IStakingControllerProverInitialized)
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
func (it *IStakingControllerProverInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerProverInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerProverInitialized represents a ProverInitialized event raised by the IStakingController contract.
type IStakingControllerProverInitialized struct {
	Prover                common.Address
	Vault                 common.Address
	DefaultCommissionRate uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProverInitialized is a free log retrieval operation binding the contract event 0xffc56da2bca5bf05ebfea889d9557b9fae7af3c57574d637faf1d4b010605465.
//
// Solidity: event ProverInitialized(address indexed prover, address indexed vault, uint64 defaultCommissionRate)
func (_IStakingController *IStakingControllerFilterer) FilterProverInitialized(opts *bind.FilterOpts, prover []common.Address, vault []common.Address) (*IStakingControllerProverInitializedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "ProverInitialized", proverRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerProverInitializedIterator{contract: _IStakingController.contract, event: "ProverInitialized", logs: logs, sub: sub}, nil
}

// WatchProverInitialized is a free log subscription operation binding the contract event 0xffc56da2bca5bf05ebfea889d9557b9fae7af3c57574d637faf1d4b010605465.
//
// Solidity: event ProverInitialized(address indexed prover, address indexed vault, uint64 defaultCommissionRate)
func (_IStakingController *IStakingControllerFilterer) WatchProverInitialized(opts *bind.WatchOpts, sink chan<- *IStakingControllerProverInitialized, prover []common.Address, vault []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "ProverInitialized", proverRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerProverInitialized)
				if err := _IStakingController.contract.UnpackLog(event, "ProverInitialized", log); err != nil {
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

// ParseProverInitialized is a log parse operation binding the contract event 0xffc56da2bca5bf05ebfea889d9557b9fae7af3c57574d637faf1d4b010605465.
//
// Solidity: event ProverInitialized(address indexed prover, address indexed vault, uint64 defaultCommissionRate)
func (_IStakingController *IStakingControllerFilterer) ParseProverInitialized(log types.Log) (*IStakingControllerProverInitialized, error) {
	event := new(IStakingControllerProverInitialized)
	if err := _IStakingController.contract.UnpackLog(event, "ProverInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerProverRetiredIterator is returned from FilterProverRetired and is used to iterate over the raw logs and unpacked data for ProverRetired events raised by the IStakingController contract.
type IStakingControllerProverRetiredIterator struct {
	Event *IStakingControllerProverRetired // Event containing the contract specifics and raw log

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
func (it *IStakingControllerProverRetiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerProverRetired)
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
		it.Event = new(IStakingControllerProverRetired)
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
func (it *IStakingControllerProverRetiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerProverRetiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerProverRetired represents a ProverRetired event raised by the IStakingController contract.
type IStakingControllerProverRetired struct {
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverRetired is a free log retrieval operation binding the contract event 0xea690d60b1cd11702fa15558b25ae6d66448d2ac3b8ac33710aeb3ab37e406e3.
//
// Solidity: event ProverRetired(address indexed prover)
func (_IStakingController *IStakingControllerFilterer) FilterProverRetired(opts *bind.FilterOpts, prover []common.Address) (*IStakingControllerProverRetiredIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "ProverRetired", proverRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerProverRetiredIterator{contract: _IStakingController.contract, event: "ProverRetired", logs: logs, sub: sub}, nil
}

// WatchProverRetired is a free log subscription operation binding the contract event 0xea690d60b1cd11702fa15558b25ae6d66448d2ac3b8ac33710aeb3ab37e406e3.
//
// Solidity: event ProverRetired(address indexed prover)
func (_IStakingController *IStakingControllerFilterer) WatchProverRetired(opts *bind.WatchOpts, sink chan<- *IStakingControllerProverRetired, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "ProverRetired", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerProverRetired)
				if err := _IStakingController.contract.UnpackLog(event, "ProverRetired", log); err != nil {
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

// ParseProverRetired is a log parse operation binding the contract event 0xea690d60b1cd11702fa15558b25ae6d66448d2ac3b8ac33710aeb3ab37e406e3.
//
// Solidity: event ProverRetired(address indexed prover)
func (_IStakingController *IStakingControllerFilterer) ParseProverRetired(log types.Log) (*IStakingControllerProverRetired, error) {
	event := new(IStakingControllerProverRetired)
	if err := _IStakingController.contract.UnpackLog(event, "ProverRetired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerProverSlashedIterator is returned from FilterProverSlashed and is used to iterate over the raw logs and unpacked data for ProverSlashed events raised by the IStakingController contract.
type IStakingControllerProverSlashedIterator struct {
	Event *IStakingControllerProverSlashed // Event containing the contract specifics and raw log

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
func (it *IStakingControllerProverSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerProverSlashed)
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
		it.Event = new(IStakingControllerProverSlashed)
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
func (it *IStakingControllerProverSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerProverSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerProverSlashed represents a ProverSlashed event raised by the IStakingController contract.
type IStakingControllerProverSlashed struct {
	Prover common.Address
	Amount *big.Int
	Bps    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProverSlashed is a free log retrieval operation binding the contract event 0x1e07165745a3fd60e5c2c529df491fccbd1d5d22032b69a728ea9bae9bbccef4.
//
// Solidity: event ProverSlashed(address indexed prover, uint256 amount, uint256 bps)
func (_IStakingController *IStakingControllerFilterer) FilterProverSlashed(opts *bind.FilterOpts, prover []common.Address) (*IStakingControllerProverSlashedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "ProverSlashed", proverRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerProverSlashedIterator{contract: _IStakingController.contract, event: "ProverSlashed", logs: logs, sub: sub}, nil
}

// WatchProverSlashed is a free log subscription operation binding the contract event 0x1e07165745a3fd60e5c2c529df491fccbd1d5d22032b69a728ea9bae9bbccef4.
//
// Solidity: event ProverSlashed(address indexed prover, uint256 amount, uint256 bps)
func (_IStakingController *IStakingControllerFilterer) WatchProverSlashed(opts *bind.WatchOpts, sink chan<- *IStakingControllerProverSlashed, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "ProverSlashed", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerProverSlashed)
				if err := _IStakingController.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
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

// ParseProverSlashed is a log parse operation binding the contract event 0x1e07165745a3fd60e5c2c529df491fccbd1d5d22032b69a728ea9bae9bbccef4.
//
// Solidity: event ProverSlashed(address indexed prover, uint256 amount, uint256 bps)
func (_IStakingController *IStakingControllerFilterer) ParseProverSlashed(log types.Log) (*IStakingControllerProverSlashed, error) {
	event := new(IStakingControllerProverSlashed)
	if err := _IStakingController.contract.UnpackLog(event, "ProverSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerProverStateChangedIterator is returned from FilterProverStateChanged and is used to iterate over the raw logs and unpacked data for ProverStateChanged events raised by the IStakingController contract.
type IStakingControllerProverStateChangedIterator struct {
	Event *IStakingControllerProverStateChanged // Event containing the contract specifics and raw log

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
func (it *IStakingControllerProverStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerProverStateChanged)
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
		it.Event = new(IStakingControllerProverStateChanged)
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
func (it *IStakingControllerProverStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerProverStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerProverStateChanged represents a ProverStateChanged event raised by the IStakingController contract.
type IStakingControllerProverStateChanged struct {
	Prover   common.Address
	OldState uint8
	NewState uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProverStateChanged is a free log retrieval operation binding the contract event 0xf62d7c2be4600ef704b62e9ca7b5f210c7de47b6104775c2d013e51514a89591.
//
// Solidity: event ProverStateChanged(address indexed prover, uint8 oldState, uint8 newState)
func (_IStakingController *IStakingControllerFilterer) FilterProverStateChanged(opts *bind.FilterOpts, prover []common.Address) (*IStakingControllerProverStateChangedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "ProverStateChanged", proverRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerProverStateChangedIterator{contract: _IStakingController.contract, event: "ProverStateChanged", logs: logs, sub: sub}, nil
}

// WatchProverStateChanged is a free log subscription operation binding the contract event 0xf62d7c2be4600ef704b62e9ca7b5f210c7de47b6104775c2d013e51514a89591.
//
// Solidity: event ProverStateChanged(address indexed prover, uint8 oldState, uint8 newState)
func (_IStakingController *IStakingControllerFilterer) WatchProverStateChanged(opts *bind.WatchOpts, sink chan<- *IStakingControllerProverStateChanged, prover []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "ProverStateChanged", proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerProverStateChanged)
				if err := _IStakingController.contract.UnpackLog(event, "ProverStateChanged", log); err != nil {
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

// ParseProverStateChanged is a log parse operation binding the contract event 0xf62d7c2be4600ef704b62e9ca7b5f210c7de47b6104775c2d013e51514a89591.
//
// Solidity: event ProverStateChanged(address indexed prover, uint8 oldState, uint8 newState)
func (_IStakingController *IStakingControllerFilterer) ParseProverStateChanged(log types.Log) (*IStakingControllerProverStateChanged, error) {
	event := new(IStakingControllerProverStateChanged)
	if err := _IStakingController.contract.UnpackLog(event, "ProverStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerRequireAuthorizationUpdatedIterator is returned from FilterRequireAuthorizationUpdated and is used to iterate over the raw logs and unpacked data for RequireAuthorizationUpdated events raised by the IStakingController contract.
type IStakingControllerRequireAuthorizationUpdatedIterator struct {
	Event *IStakingControllerRequireAuthorizationUpdated // Event containing the contract specifics and raw log

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
func (it *IStakingControllerRequireAuthorizationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerRequireAuthorizationUpdated)
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
		it.Event = new(IStakingControllerRequireAuthorizationUpdated)
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
func (it *IStakingControllerRequireAuthorizationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerRequireAuthorizationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerRequireAuthorizationUpdated represents a RequireAuthorizationUpdated event raised by the IStakingController contract.
type IStakingControllerRequireAuthorizationUpdated struct {
	Required bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequireAuthorizationUpdated is a free log retrieval operation binding the contract event 0x8131635808ab3dcde7054fa5cf23a415ee65c2230297d5b413e033e5cff08a32.
//
// Solidity: event RequireAuthorizationUpdated(bool required)
func (_IStakingController *IStakingControllerFilterer) FilterRequireAuthorizationUpdated(opts *bind.FilterOpts) (*IStakingControllerRequireAuthorizationUpdatedIterator, error) {

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "RequireAuthorizationUpdated")
	if err != nil {
		return nil, err
	}
	return &IStakingControllerRequireAuthorizationUpdatedIterator{contract: _IStakingController.contract, event: "RequireAuthorizationUpdated", logs: logs, sub: sub}, nil
}

// WatchRequireAuthorizationUpdated is a free log subscription operation binding the contract event 0x8131635808ab3dcde7054fa5cf23a415ee65c2230297d5b413e033e5cff08a32.
//
// Solidity: event RequireAuthorizationUpdated(bool required)
func (_IStakingController *IStakingControllerFilterer) WatchRequireAuthorizationUpdated(opts *bind.WatchOpts, sink chan<- *IStakingControllerRequireAuthorizationUpdated) (event.Subscription, error) {

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "RequireAuthorizationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerRequireAuthorizationUpdated)
				if err := _IStakingController.contract.UnpackLog(event, "RequireAuthorizationUpdated", log); err != nil {
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

// ParseRequireAuthorizationUpdated is a log parse operation binding the contract event 0x8131635808ab3dcde7054fa5cf23a415ee65c2230297d5b413e033e5cff08a32.
//
// Solidity: event RequireAuthorizationUpdated(bool required)
func (_IStakingController *IStakingControllerFilterer) ParseRequireAuthorizationUpdated(log types.Log) (*IStakingControllerRequireAuthorizationUpdated, error) {
	event := new(IStakingControllerRequireAuthorizationUpdated)
	if err := _IStakingController.contract.UnpackLog(event, "RequireAuthorizationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerRewardsAddedIterator is returned from FilterRewardsAdded and is used to iterate over the raw logs and unpacked data for RewardsAdded events raised by the IStakingController contract.
type IStakingControllerRewardsAddedIterator struct {
	Event *IStakingControllerRewardsAdded // Event containing the contract specifics and raw log

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
func (it *IStakingControllerRewardsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerRewardsAdded)
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
		it.Event = new(IStakingControllerRewardsAdded)
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
func (it *IStakingControllerRewardsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerRewardsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerRewardsAdded represents a RewardsAdded event raised by the IStakingController contract.
type IStakingControllerRewardsAdded struct {
	Prover     common.Address
	Source     common.Address
	Amount     *big.Int
	Commission *big.Int
	ToStakers  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardsAdded is a free log retrieval operation binding the contract event 0x1b0006daff5bb99d1383ffc4326984a27347192b77ddb40fe18ec9720d0929a8.
//
// Solidity: event RewardsAdded(address indexed prover, address indexed source, uint256 amount, uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerFilterer) FilterRewardsAdded(opts *bind.FilterOpts, prover []common.Address, source []common.Address) (*IStakingControllerRewardsAddedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "RewardsAdded", proverRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerRewardsAddedIterator{contract: _IStakingController.contract, event: "RewardsAdded", logs: logs, sub: sub}, nil
}

// WatchRewardsAdded is a free log subscription operation binding the contract event 0x1b0006daff5bb99d1383ffc4326984a27347192b77ddb40fe18ec9720d0929a8.
//
// Solidity: event RewardsAdded(address indexed prover, address indexed source, uint256 amount, uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerFilterer) WatchRewardsAdded(opts *bind.WatchOpts, sink chan<- *IStakingControllerRewardsAdded, prover []common.Address, source []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "RewardsAdded", proverRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerRewardsAdded)
				if err := _IStakingController.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
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

// ParseRewardsAdded is a log parse operation binding the contract event 0x1b0006daff5bb99d1383ffc4326984a27347192b77ddb40fe18ec9720d0929a8.
//
// Solidity: event RewardsAdded(address indexed prover, address indexed source, uint256 amount, uint256 commission, uint256 toStakers)
func (_IStakingController *IStakingControllerFilterer) ParseRewardsAdded(log types.Log) (*IStakingControllerRewardsAdded, error) {
	event := new(IStakingControllerRewardsAdded)
	if err := _IStakingController.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the IStakingController contract.
type IStakingControllerStakedIterator struct {
	Event *IStakingControllerStaked // Event containing the contract specifics and raw log

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
func (it *IStakingControllerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerStaked)
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
		it.Event = new(IStakingControllerStaked)
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
func (it *IStakingControllerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerStaked represents a Staked event raised by the IStakingController contract.
type IStakingControllerStaked struct {
	Prover common.Address
	Staker common.Address
	Shares *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterStaked(opts *bind.FilterOpts, prover []common.Address, staker []common.Address) (*IStakingControllerStakedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "Staked", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerStakedIterator{contract: _IStakingController.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *IStakingControllerStaked, prover []common.Address, staker []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "Staked", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerStaked)
				if err := _IStakingController.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseStaked(log types.Log) (*IStakingControllerStaked, error) {
	event := new(IStakingControllerStaked)
	if err := _IStakingController.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerTreasuryWithdrawnIterator is returned from FilterTreasuryWithdrawn and is used to iterate over the raw logs and unpacked data for TreasuryWithdrawn events raised by the IStakingController contract.
type IStakingControllerTreasuryWithdrawnIterator struct {
	Event *IStakingControllerTreasuryWithdrawn // Event containing the contract specifics and raw log

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
func (it *IStakingControllerTreasuryWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerTreasuryWithdrawn)
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
		it.Event = new(IStakingControllerTreasuryWithdrawn)
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
func (it *IStakingControllerTreasuryWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerTreasuryWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerTreasuryWithdrawn represents a TreasuryWithdrawn event raised by the IStakingController contract.
type IStakingControllerTreasuryWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryWithdrawn is a free log retrieval operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterTreasuryWithdrawn(opts *bind.FilterOpts, to []common.Address) (*IStakingControllerTreasuryWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "TreasuryWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerTreasuryWithdrawnIterator{contract: _IStakingController.contract, event: "TreasuryWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTreasuryWithdrawn is a free log subscription operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchTreasuryWithdrawn(opts *bind.WatchOpts, sink chan<- *IStakingControllerTreasuryWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "TreasuryWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerTreasuryWithdrawn)
				if err := _IStakingController.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
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

// ParseTreasuryWithdrawn is a log parse operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseTreasuryWithdrawn(log types.Log) (*IStakingControllerTreasuryWithdrawn, error) {
	event := new(IStakingControllerTreasuryWithdrawn)
	if err := _IStakingController.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerUnstakeCompletedIterator is returned from FilterUnstakeCompleted and is used to iterate over the raw logs and unpacked data for UnstakeCompleted events raised by the IStakingController contract.
type IStakingControllerUnstakeCompletedIterator struct {
	Event *IStakingControllerUnstakeCompleted // Event containing the contract specifics and raw log

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
func (it *IStakingControllerUnstakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerUnstakeCompleted)
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
		it.Event = new(IStakingControllerUnstakeCompleted)
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
func (it *IStakingControllerUnstakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerUnstakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerUnstakeCompleted represents a UnstakeCompleted event raised by the IStakingController contract.
type IStakingControllerUnstakeCompleted struct {
	Prover common.Address
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstakeCompleted is a free log retrieval operation binding the contract event 0xf3bf64f3e95494bffbfa1a414348cfe31a4abdc20fb1ef2f225a6b65f02de852.
//
// Solidity: event UnstakeCompleted(address indexed prover, address indexed staker, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterUnstakeCompleted(opts *bind.FilterOpts, prover []common.Address, staker []common.Address) (*IStakingControllerUnstakeCompletedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "UnstakeCompleted", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerUnstakeCompletedIterator{contract: _IStakingController.contract, event: "UnstakeCompleted", logs: logs, sub: sub}, nil
}

// WatchUnstakeCompleted is a free log subscription operation binding the contract event 0xf3bf64f3e95494bffbfa1a414348cfe31a4abdc20fb1ef2f225a6b65f02de852.
//
// Solidity: event UnstakeCompleted(address indexed prover, address indexed staker, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchUnstakeCompleted(opts *bind.WatchOpts, sink chan<- *IStakingControllerUnstakeCompleted, prover []common.Address, staker []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "UnstakeCompleted", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerUnstakeCompleted)
				if err := _IStakingController.contract.UnpackLog(event, "UnstakeCompleted", log); err != nil {
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

// ParseUnstakeCompleted is a log parse operation binding the contract event 0xf3bf64f3e95494bffbfa1a414348cfe31a4abdc20fb1ef2f225a6b65f02de852.
//
// Solidity: event UnstakeCompleted(address indexed prover, address indexed staker, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseUnstakeCompleted(log types.Log) (*IStakingControllerUnstakeCompleted, error) {
	event := new(IStakingControllerUnstakeCompleted)
	if err := _IStakingController.contract.UnpackLog(event, "UnstakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerUnstakeDelayUpdatedIterator is returned from FilterUnstakeDelayUpdated and is used to iterate over the raw logs and unpacked data for UnstakeDelayUpdated events raised by the IStakingController contract.
type IStakingControllerUnstakeDelayUpdatedIterator struct {
	Event *IStakingControllerUnstakeDelayUpdated // Event containing the contract specifics and raw log

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
func (it *IStakingControllerUnstakeDelayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerUnstakeDelayUpdated)
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
		it.Event = new(IStakingControllerUnstakeDelayUpdated)
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
func (it *IStakingControllerUnstakeDelayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerUnstakeDelayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerUnstakeDelayUpdated represents a UnstakeDelayUpdated event raised by the IStakingController contract.
type IStakingControllerUnstakeDelayUpdated struct {
	OldDelay *big.Int
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayUpdated is a free log retrieval operation binding the contract event 0x308abc9a9ccc3f3ffbaefa845a2e756ce4d367abd5a3078eb950285e6476f2d9.
//
// Solidity: event UnstakeDelayUpdated(uint256 oldDelay, uint256 newDelay)
func (_IStakingController *IStakingControllerFilterer) FilterUnstakeDelayUpdated(opts *bind.FilterOpts) (*IStakingControllerUnstakeDelayUpdatedIterator, error) {

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "UnstakeDelayUpdated")
	if err != nil {
		return nil, err
	}
	return &IStakingControllerUnstakeDelayUpdatedIterator{contract: _IStakingController.contract, event: "UnstakeDelayUpdated", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayUpdated is a free log subscription operation binding the contract event 0x308abc9a9ccc3f3ffbaefa845a2e756ce4d367abd5a3078eb950285e6476f2d9.
//
// Solidity: event UnstakeDelayUpdated(uint256 oldDelay, uint256 newDelay)
func (_IStakingController *IStakingControllerFilterer) WatchUnstakeDelayUpdated(opts *bind.WatchOpts, sink chan<- *IStakingControllerUnstakeDelayUpdated) (event.Subscription, error) {

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "UnstakeDelayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerUnstakeDelayUpdated)
				if err := _IStakingController.contract.UnpackLog(event, "UnstakeDelayUpdated", log); err != nil {
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

// ParseUnstakeDelayUpdated is a log parse operation binding the contract event 0x308abc9a9ccc3f3ffbaefa845a2e756ce4d367abd5a3078eb950285e6476f2d9.
//
// Solidity: event UnstakeDelayUpdated(uint256 oldDelay, uint256 newDelay)
func (_IStakingController *IStakingControllerFilterer) ParseUnstakeDelayUpdated(log types.Log) (*IStakingControllerUnstakeDelayUpdated, error) {
	event := new(IStakingControllerUnstakeDelayUpdated)
	if err := _IStakingController.contract.UnpackLog(event, "UnstakeDelayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingControllerUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the IStakingController contract.
type IStakingControllerUnstakeRequestedIterator struct {
	Event *IStakingControllerUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *IStakingControllerUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingControllerUnstakeRequested)
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
		it.Event = new(IStakingControllerUnstakeRequested)
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
func (it *IStakingControllerUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingControllerUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingControllerUnstakeRequested represents a UnstakeRequested event raised by the IStakingController contract.
type IStakingControllerUnstakeRequested struct {
	Prover common.Address
	Staker common.Address
	Shares *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0xfe07ce9fff39f8420b3de5fbc6909ce08f809e2572b62f9df35c25f56d610bb0.
//
// Solidity: event UnstakeRequested(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, prover []common.Address, staker []common.Address) (*IStakingControllerUnstakeRequestedIterator, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.FilterLogs(opts, "UnstakeRequested", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &IStakingControllerUnstakeRequestedIterator{contract: _IStakingController.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0xfe07ce9fff39f8420b3de5fbc6909ce08f809e2572b62f9df35c25f56d610bb0.
//
// Solidity: event UnstakeRequested(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *IStakingControllerUnstakeRequested, prover []common.Address, staker []common.Address) (event.Subscription, error) {

	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _IStakingController.contract.WatchLogs(opts, "UnstakeRequested", proverRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingControllerUnstakeRequested)
				if err := _IStakingController.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0xfe07ce9fff39f8420b3de5fbc6909ce08f809e2572b62f9df35c25f56d610bb0.
//
// Solidity: event UnstakeRequested(address indexed prover, address indexed staker, uint256 shares, uint256 amount)
func (_IStakingController *IStakingControllerFilterer) ParseUnstakeRequested(log types.Log) (*IStakingControllerUnstakeRequested, error) {
	event := new(IStakingControllerUnstakeRequested)
	if err := _IStakingController.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212205c8f54f4577c9f3f9a4a5fe11d44eaa99426caf037efabc6301e700a0d58262564736f6c63430008140033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PanicMetaData contains all meta data concerning the Panic contract.
var PanicMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212209e7da5834b772e5f034bdd762a8601b740a7311b10cfadea8ab5c756856c71ed64736f6c63430008140033",
}

// PanicABI is the input ABI used to generate the binding from.
// Deprecated: Use PanicMetaData.ABI instead.
var PanicABI = PanicMetaData.ABI

// PanicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PanicMetaData.Bin instead.
var PanicBin = PanicMetaData.Bin

// DeployPanic deploys a new Ethereum contract, binding an instance of Panic to it.
func DeployPanic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Panic, error) {
	parsed, err := PanicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PanicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Panic{PanicCaller: PanicCaller{contract: contract}, PanicTransactor: PanicTransactor{contract: contract}, PanicFilterer: PanicFilterer{contract: contract}}, nil
}

// Panic is an auto generated Go binding around an Ethereum contract.
type Panic struct {
	PanicCaller     // Read-only binding to the contract
	PanicTransactor // Write-only binding to the contract
	PanicFilterer   // Log filterer for contract events
}

// PanicCaller is an auto generated read-only Go binding around an Ethereum contract.
type PanicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PanicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PanicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PanicSession struct {
	Contract     *Panic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PanicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PanicCallerSession struct {
	Contract *PanicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PanicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PanicTransactorSession struct {
	Contract     *PanicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PanicRaw is an auto generated low-level Go binding around an Ethereum contract.
type PanicRaw struct {
	Contract *Panic // Generic contract binding to access the raw methods on
}

// PanicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PanicCallerRaw struct {
	Contract *PanicCaller // Generic read-only contract binding to access the raw methods on
}

// PanicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PanicTransactorRaw struct {
	Contract *PanicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPanic creates a new instance of Panic, bound to a specific deployed contract.
func NewPanic(address common.Address, backend bind.ContractBackend) (*Panic, error) {
	contract, err := bindPanic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Panic{PanicCaller: PanicCaller{contract: contract}, PanicTransactor: PanicTransactor{contract: contract}, PanicFilterer: PanicFilterer{contract: contract}}, nil
}

// NewPanicCaller creates a new read-only instance of Panic, bound to a specific deployed contract.
func NewPanicCaller(address common.Address, caller bind.ContractCaller) (*PanicCaller, error) {
	contract, err := bindPanic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PanicCaller{contract: contract}, nil
}

// NewPanicTransactor creates a new write-only instance of Panic, bound to a specific deployed contract.
func NewPanicTransactor(address common.Address, transactor bind.ContractTransactor) (*PanicTransactor, error) {
	contract, err := bindPanic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PanicTransactor{contract: contract}, nil
}

// NewPanicFilterer creates a new log filterer instance of Panic, bound to a specific deployed contract.
func NewPanicFilterer(address common.Address, filterer bind.ContractFilterer) (*PanicFilterer, error) {
	contract, err := bindPanic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PanicFilterer{contract: contract}, nil
}

// bindPanic binds a generic wrapper to an already deployed contract.
func bindPanic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PanicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Panic *PanicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Panic.Contract.PanicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Panic *PanicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Panic.Contract.PanicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Panic *PanicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Panic.Contract.PanicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Panic *PanicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Panic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Panic *PanicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Panic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Panic *PanicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Panic.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"}]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220ace471c10bad8d20333dcb6f92680f4fc75dfe59d1da38b87129f3f88b152e6264736f6c63430008140033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"SafeERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"}]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea264697066735822122074936108e5a70f550d0c0fd0b0c6e2f4e947af45ab55e670bd8ec16f20f32c8164736f6c63430008140033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SlotDerivationMetaData contains all meta data concerning the SlotDerivation contract.
var SlotDerivationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212204a5dd24399ac431ba58dec2fae8eb0a8cb1d081c524a82b07a4d726df400a0a264736f6c63430008140033",
}

// SlotDerivationABI is the input ABI used to generate the binding from.
// Deprecated: Use SlotDerivationMetaData.ABI instead.
var SlotDerivationABI = SlotDerivationMetaData.ABI

// SlotDerivationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlotDerivationMetaData.Bin instead.
var SlotDerivationBin = SlotDerivationMetaData.Bin

// DeploySlotDerivation deploys a new Ethereum contract, binding an instance of SlotDerivation to it.
func DeploySlotDerivation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlotDerivation, error) {
	parsed, err := SlotDerivationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlotDerivationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlotDerivation{SlotDerivationCaller: SlotDerivationCaller{contract: contract}, SlotDerivationTransactor: SlotDerivationTransactor{contract: contract}, SlotDerivationFilterer: SlotDerivationFilterer{contract: contract}}, nil
}

// SlotDerivation is an auto generated Go binding around an Ethereum contract.
type SlotDerivation struct {
	SlotDerivationCaller     // Read-only binding to the contract
	SlotDerivationTransactor // Write-only binding to the contract
	SlotDerivationFilterer   // Log filterer for contract events
}

// SlotDerivationCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlotDerivationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotDerivationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlotDerivationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotDerivationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlotDerivationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlotDerivationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlotDerivationSession struct {
	Contract     *SlotDerivation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlotDerivationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlotDerivationCallerSession struct {
	Contract *SlotDerivationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SlotDerivationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlotDerivationTransactorSession struct {
	Contract     *SlotDerivationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SlotDerivationRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlotDerivationRaw struct {
	Contract *SlotDerivation // Generic contract binding to access the raw methods on
}

// SlotDerivationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlotDerivationCallerRaw struct {
	Contract *SlotDerivationCaller // Generic read-only contract binding to access the raw methods on
}

// SlotDerivationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlotDerivationTransactorRaw struct {
	Contract *SlotDerivationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlotDerivation creates a new instance of SlotDerivation, bound to a specific deployed contract.
func NewSlotDerivation(address common.Address, backend bind.ContractBackend) (*SlotDerivation, error) {
	contract, err := bindSlotDerivation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlotDerivation{SlotDerivationCaller: SlotDerivationCaller{contract: contract}, SlotDerivationTransactor: SlotDerivationTransactor{contract: contract}, SlotDerivationFilterer: SlotDerivationFilterer{contract: contract}}, nil
}

// NewSlotDerivationCaller creates a new read-only instance of SlotDerivation, bound to a specific deployed contract.
func NewSlotDerivationCaller(address common.Address, caller bind.ContractCaller) (*SlotDerivationCaller, error) {
	contract, err := bindSlotDerivation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlotDerivationCaller{contract: contract}, nil
}

// NewSlotDerivationTransactor creates a new write-only instance of SlotDerivation, bound to a specific deployed contract.
func NewSlotDerivationTransactor(address common.Address, transactor bind.ContractTransactor) (*SlotDerivationTransactor, error) {
	contract, err := bindSlotDerivation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlotDerivationTransactor{contract: contract}, nil
}

// NewSlotDerivationFilterer creates a new log filterer instance of SlotDerivation, bound to a specific deployed contract.
func NewSlotDerivationFilterer(address common.Address, filterer bind.ContractFilterer) (*SlotDerivationFilterer, error) {
	contract, err := bindSlotDerivation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlotDerivationFilterer{contract: contract}, nil
}

// bindSlotDerivation binds a generic wrapper to an already deployed contract.
func bindSlotDerivation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlotDerivationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotDerivation *SlotDerivationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotDerivation.Contract.SlotDerivationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotDerivation *SlotDerivationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotDerivation.Contract.SlotDerivationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotDerivation *SlotDerivationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotDerivation.Contract.SlotDerivationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlotDerivation *SlotDerivationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlotDerivation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlotDerivation *SlotDerivationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlotDerivation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlotDerivation *SlotDerivationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlotDerivation.Contract.contract.Transact(opts, method, params...)
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212205b303505b0771586e7e5f68185834ea051f7bd3eeb3c62f6ef94ae3cd29b05bb64736f6c63430008140033",
}

// StorageSlotABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotMetaData.ABI instead.
var StorageSlotABI = StorageSlotMetaData.ABI

// StorageSlotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotMetaData.Bin instead.
var StorageSlotBin = StorageSlotMetaData.Bin

// DeployStorageSlot deploys a new Ethereum contract, binding an instance of StorageSlot to it.
func DeployStorageSlot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlot, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// StorageSlot is an auto generated Go binding around an Ethereum contract.
type StorageSlot struct {
	StorageSlotCaller     // Read-only binding to the contract
	StorageSlotTransactor // Write-only binding to the contract
	StorageSlotFilterer   // Log filterer for contract events
}

// StorageSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotSession struct {
	Contract     *StorageSlot      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageSlotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotCallerSession struct {
	Contract *StorageSlotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StorageSlotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotTransactorSession struct {
	Contract     *StorageSlotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StorageSlotRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotRaw struct {
	Contract *StorageSlot // Generic contract binding to access the raw methods on
}

// StorageSlotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotCallerRaw struct {
	Contract *StorageSlotCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotTransactorRaw struct {
	Contract *StorageSlotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlot creates a new instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlot(address common.Address, backend bind.ContractBackend) (*StorageSlot, error) {
	contract, err := bindStorageSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// NewStorageSlotCaller creates a new read-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotCaller, error) {
	contract, err := bindStorageSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotCaller{contract: contract}, nil
}

// NewStorageSlotTransactor creates a new write-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotTransactor, error) {
	contract, err := bindStorageSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotTransactor{contract: contract}, nil
}

// NewStorageSlotFilterer creates a new log filterer instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotFilterer, error) {
	contract, err := bindStorageSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotFilterer{contract: contract}, nil
}

// bindStorageSlot binds a generic wrapper to an already deployed contract.
func bindStorageSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.StorageSlotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transact(opts, method, params...)
}
