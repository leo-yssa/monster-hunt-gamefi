// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package infrastructure

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterHunted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PlayerRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"addMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"}],\"name\":\"huntMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsters\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerPlayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractMyGameToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_Contract *ContractCaller) Monsters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "monsters", arg0)

	outstruct := new(struct {
		Name   string
		Hp     *big.Int
		Reward *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Hp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_Contract *ContractSession) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _Contract.Contract.Monsters(&_Contract.CallOpts, arg0)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_Contract *ContractCallerSession) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _Contract.Contract.Monsters(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_Contract *ContractCaller) Players(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "players", arg0)

	outstruct := new(struct {
		Name       string
		Level      *big.Int
		Registered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Level = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_Contract *ContractSession) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _Contract.Contract.Players(&_Contract.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_Contract *ContractCallerSession) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _Contract.Contract.Players(&_Contract.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Contract *ContractCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Contract *ContractSession) RewardToken() (common.Address, error) {
	return _Contract.Contract.RewardToken(&_Contract.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Contract *ContractCallerSession) RewardToken() (common.Address, error) {
	return _Contract.Contract.RewardToken(&_Contract.CallOpts)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_Contract *ContractTransactor) AddMonster(opts *bind.TransactOpts, name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addMonster", name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_Contract *ContractSession) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddMonster(&_Contract.TransactOpts, name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_Contract *ContractTransactorSession) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddMonster(&_Contract.TransactOpts, name, hp, reward)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_Contract *ContractTransactor) HuntMonster(opts *bind.TransactOpts, monsterId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "huntMonster", monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_Contract *ContractSession) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.HuntMonster(&_Contract.TransactOpts, monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_Contract *ContractTransactorSession) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.HuntMonster(&_Contract.TransactOpts, monsterId)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_Contract *ContractTransactor) RegisterPlayer(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerPlayer", name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_Contract *ContractSession) RegisterPlayer(name string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterPlayer(&_Contract.TransactOpts, name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_Contract *ContractTransactorSession) RegisterPlayer(name string) (*types.Transaction, error) {
	return _Contract.Contract.RegisterPlayer(&_Contract.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractMonsterAddedIterator is returned from FilterMonsterAdded and is used to iterate over the raw logs and unpacked data for MonsterAdded events raised by the Contract contract.
type ContractMonsterAddedIterator struct {
	Event *ContractMonsterAdded // Event containing the contract specifics and raw log

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
func (it *ContractMonsterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMonsterAdded)
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
		it.Event = new(ContractMonsterAdded)
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
func (it *ContractMonsterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMonsterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMonsterAdded represents a MonsterAdded event raised by the Contract contract.
type ContractMonsterAdded struct {
	MonsterId *big.Int
	Name      string
	Hp        *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterAdded is a free log retrieval operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_Contract *ContractFilterer) FilterMonsterAdded(opts *bind.FilterOpts, monsterId []*big.Int) (*ContractMonsterAddedIterator, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractMonsterAddedIterator{contract: _Contract.contract, event: "MonsterAdded", logs: logs, sub: sub}, nil
}

// WatchMonsterAdded is a free log subscription operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_Contract *ContractFilterer) WatchMonsterAdded(opts *bind.WatchOpts, sink chan<- *ContractMonsterAdded, monsterId []*big.Int) (event.Subscription, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMonsterAdded)
				if err := _Contract.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
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

// ParseMonsterAdded is a log parse operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_Contract *ContractFilterer) ParseMonsterAdded(log types.Log) (*ContractMonsterAdded, error) {
	event := new(ContractMonsterAdded)
	if err := _Contract.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMonsterHuntedIterator is returned from FilterMonsterHunted and is used to iterate over the raw logs and unpacked data for MonsterHunted events raised by the Contract contract.
type ContractMonsterHuntedIterator struct {
	Event *ContractMonsterHunted // Event containing the contract specifics and raw log

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
func (it *ContractMonsterHuntedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMonsterHunted)
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
		it.Event = new(ContractMonsterHunted)
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
func (it *ContractMonsterHuntedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMonsterHuntedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMonsterHunted represents a MonsterHunted event raised by the Contract contract.
type ContractMonsterHunted struct {
	Player    common.Address
	MonsterId *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterHunted is a free log retrieval operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_Contract *ContractFilterer) FilterMonsterHunted(opts *bind.FilterOpts, player []common.Address, monsterId []*big.Int) (*ContractMonsterHuntedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractMonsterHuntedIterator{contract: _Contract.contract, event: "MonsterHunted", logs: logs, sub: sub}, nil
}

// WatchMonsterHunted is a free log subscription operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_Contract *ContractFilterer) WatchMonsterHunted(opts *bind.WatchOpts, sink chan<- *ContractMonsterHunted, player []common.Address, monsterId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMonsterHunted)
				if err := _Contract.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
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

// ParseMonsterHunted is a log parse operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_Contract *ContractFilterer) ParseMonsterHunted(log types.Log) (*ContractMonsterHunted, error) {
	event := new(ContractMonsterHunted)
	if err := _Contract.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPlayerRegisteredIterator is returned from FilterPlayerRegistered and is used to iterate over the raw logs and unpacked data for PlayerRegistered events raised by the Contract contract.
type ContractPlayerRegisteredIterator struct {
	Event *ContractPlayerRegistered // Event containing the contract specifics and raw log

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
func (it *ContractPlayerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPlayerRegistered)
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
		it.Event = new(ContractPlayerRegistered)
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
func (it *ContractPlayerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPlayerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPlayerRegistered represents a PlayerRegistered event raised by the Contract contract.
type ContractPlayerRegistered struct {
	Player common.Address
	Name   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlayerRegistered is a free log retrieval operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_Contract *ContractFilterer) FilterPlayerRegistered(opts *bind.FilterOpts, player []common.Address) (*ContractPlayerRegisteredIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPlayerRegisteredIterator{contract: _Contract.contract, event: "PlayerRegistered", logs: logs, sub: sub}, nil
}

// WatchPlayerRegistered is a free log subscription operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_Contract *ContractFilterer) WatchPlayerRegistered(opts *bind.WatchOpts, sink chan<- *ContractPlayerRegistered, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPlayerRegistered)
				if err := _Contract.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
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

// ParsePlayerRegistered is a log parse operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_Contract *ContractFilterer) ParsePlayerRegistered(log types.Log) (*ContractPlayerRegistered, error) {
	event := new(ContractPlayerRegistered)
	if err := _Contract.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
