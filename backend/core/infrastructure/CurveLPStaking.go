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

// CurveLPStakingMetaData contains all meta data concerning the CurveLPStaking contract.
var CurveLPStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractCurveLPToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"staked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610d2b380380610d2b833981810160405281019061003291906100db565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100a88261007d565b9050919050565b6100b88161009d565b81146100c357600080fd5b50565b6000815190506100d5816100af565b92915050565b6000602082840312156100f1576100f0610078565b5b60006100ff848285016100c6565b91505092915050565b610c14806101176000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638e539e8c116100665780638e539e8c1461014957806398807d8414610179578063a694fc3a146101a9578063c07473f6146101c5578063eb9019d4146101f55761009e565b80632e17de78146100a35780633a46b1a8146100bf5780634bf5d7e9146100ef5780635fcbd2851461010d578063817b1cd21461012b575b600080fd5b6100bd60048036038101906100b8919061074d565b610225565b005b6100d960048036038101906100d491906107d8565b610406565b6040516100e69190610827565b60405180910390f35b6100f7610450565b60405161010491906108d2565b60405180910390f35b61011561048d565b6040516101229190610953565b60405180910390f35b6101336104b1565b6040516101409190610827565b60405180910390f35b610163600480360381019061015e919061074d565b6104b7565b6040516101709190610827565b60405180910390f35b610193600480360381019061018e919061096e565b6104c3565b6040516101a09190610827565b60405180910390f35b6101c360048036038101906101be919061074d565b6104db565b005b6101df60048036038101906101da919061096e565b61067f565b6040516101ec9190610827565b60405180910390f35b61020f600480360381019061020a91906107d8565b6106c8565b60405161021c9190610827565b60405180910390f35b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029e906109e7565b60405180910390fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102f69190610a36565b92505081905550806002600082825461030f9190610a36565b9250508190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610371929190610a79565b6020604051808303816000875af1158015610390573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b49190610ada565b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75826040516103fb9190610827565b60405180910390a250565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60606040518060400160405280601d81526020017f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000815250905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60006002549050919050565b60016020528060005260406000206000915090505481565b6000811161051e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051590610b53565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161057b93929190610b73565b6020604051808303816000875af115801561059a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105be9190610ada565b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461060e9190610baa565b9250508190555080600260008282546106279190610baa565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d826040516106749190610827565b60405180910390a250565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600080fd5b6000819050919050565b61072a81610717565b811461073557600080fd5b50565b60008135905061074781610721565b92915050565b60006020828403121561076357610762610712565b5b600061077184828501610738565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107a58261077a565b9050919050565b6107b58161079a565b81146107c057600080fd5b50565b6000813590506107d2816107ac565b92915050565b600080604083850312156107ef576107ee610712565b5b60006107fd858286016107c3565b925050602061080e85828601610738565b9150509250929050565b61082181610717565b82525050565b600060208201905061083c6000830184610818565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561087c578082015181840152602081019050610861565b60008484015250505050565b6000601f19601f8301169050919050565b60006108a482610842565b6108ae818561084d565b93506108be81856020860161085e565b6108c781610888565b840191505092915050565b600060208201905081810360008301526108ec8184610899565b905092915050565b6000819050919050565b600061091961091461090f8461077a565b6108f4565b61077a565b9050919050565b600061092b826108fe565b9050919050565b600061093d82610920565b9050919050565b61094d81610932565b82525050565b60006020820190506109686000830184610944565b92915050565b60006020828403121561098457610983610712565b5b6000610992848285016107c3565b91505092915050565b7f4e6f7420656e6f756768207374616b6564000000000000000000000000000000600082015250565b60006109d160118361084d565b91506109dc8261099b565b602082019050919050565b60006020820190508181036000830152610a00816109c4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610a4182610717565b9150610a4c83610717565b9250828203905081811115610a6457610a63610a07565b5b92915050565b610a738161079a565b82525050565b6000604082019050610a8e6000830185610a6a565b610a9b6020830184610818565b9392505050565b60008115159050919050565b610ab781610aa2565b8114610ac257600080fd5b50565b600081519050610ad481610aae565b92915050565b600060208284031215610af057610aef610712565b5b6000610afe84828501610ac5565b91505092915050565b7f43616e6e6f74207374616b652030000000000000000000000000000000000000600082015250565b6000610b3d600e8361084d565b9150610b4882610b07565b602082019050919050565b60006020820190508181036000830152610b6c81610b30565b9050919050565b6000606082019050610b886000830186610a6a565b610b956020830185610a6a565b610ba26040830184610818565b949350505050565b6000610bb582610717565b9150610bc083610717565b9250828201905080821115610bd857610bd7610a07565b5b9291505056fea2646970667358221220326c52734cad7786bede2a2870d1f2cf4de5023209ef173e3e324ee3e1c34acc64736f6c63430008140033",
}

// CurveLPStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveLPStakingMetaData.ABI instead.
var CurveLPStakingABI = CurveLPStakingMetaData.ABI

// CurveLPStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CurveLPStakingMetaData.Bin instead.
var CurveLPStakingBin = CurveLPStakingMetaData.Bin

// DeployCurveLPStaking deploys a new Ethereum contract, binding an instance of CurveLPStaking to it.
func DeployCurveLPStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _lpToken common.Address) (common.Address, *types.Transaction, *CurveLPStaking, error) {
	parsed, err := CurveLPStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CurveLPStakingBin), backend, _lpToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CurveLPStaking{CurveLPStakingCaller: CurveLPStakingCaller{contract: contract}, CurveLPStakingTransactor: CurveLPStakingTransactor{contract: contract}, CurveLPStakingFilterer: CurveLPStakingFilterer{contract: contract}}, nil
}

// CurveLPStaking is an auto generated Go binding around an Ethereum contract.
type CurveLPStaking struct {
	CurveLPStakingCaller     // Read-only binding to the contract
	CurveLPStakingTransactor // Write-only binding to the contract
	CurveLPStakingFilterer   // Log filterer for contract events
}

// CurveLPStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveLPStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLPStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveLPStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLPStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveLPStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLPStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveLPStakingSession struct {
	Contract     *CurveLPStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveLPStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveLPStakingCallerSession struct {
	Contract *CurveLPStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CurveLPStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveLPStakingTransactorSession struct {
	Contract     *CurveLPStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CurveLPStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveLPStakingRaw struct {
	Contract *CurveLPStaking // Generic contract binding to access the raw methods on
}

// CurveLPStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveLPStakingCallerRaw struct {
	Contract *CurveLPStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CurveLPStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveLPStakingTransactorRaw struct {
	Contract *CurveLPStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveLPStaking creates a new instance of CurveLPStaking, bound to a specific deployed contract.
func NewCurveLPStaking(address common.Address, backend bind.ContractBackend) (*CurveLPStaking, error) {
	contract, err := bindCurveLPStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveLPStaking{CurveLPStakingCaller: CurveLPStakingCaller{contract: contract}, CurveLPStakingTransactor: CurveLPStakingTransactor{contract: contract}, CurveLPStakingFilterer: CurveLPStakingFilterer{contract: contract}}, nil
}

// NewCurveLPStakingCaller creates a new read-only instance of CurveLPStaking, bound to a specific deployed contract.
func NewCurveLPStakingCaller(address common.Address, caller bind.ContractCaller) (*CurveLPStakingCaller, error) {
	contract, err := bindCurveLPStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingCaller{contract: contract}, nil
}

// NewCurveLPStakingTransactor creates a new write-only instance of CurveLPStaking, bound to a specific deployed contract.
func NewCurveLPStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveLPStakingTransactor, error) {
	contract, err := bindCurveLPStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingTransactor{contract: contract}, nil
}

// NewCurveLPStakingFilterer creates a new log filterer instance of CurveLPStaking, bound to a specific deployed contract.
func NewCurveLPStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveLPStakingFilterer, error) {
	contract, err := bindCurveLPStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingFilterer{contract: contract}, nil
}

// bindCurveLPStaking binds a generic wrapper to an already deployed contract.
func bindCurveLPStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveLPStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLPStaking *CurveLPStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLPStaking.Contract.CurveLPStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLPStaking *CurveLPStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.CurveLPStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLPStaking *CurveLPStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.CurveLPStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLPStaking *CurveLPStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLPStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLPStaking *CurveLPStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLPStaking *CurveLPStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.contract.Transact(opts, method, params...)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_CurveLPStaking *CurveLPStakingCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_CurveLPStaking *CurveLPStakingSession) CLOCKMODE() (string, error) {
	return _CurveLPStaking.Contract.CLOCKMODE(&_CurveLPStaking.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() pure returns(string)
func (_CurveLPStaking *CurveLPStakingCallerSession) CLOCKMODE() (string, error) {
	return _CurveLPStaking.Contract.CLOCKMODE(&_CurveLPStaking.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) GetPastTotalSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "getPastTotalSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) GetPastTotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetPastTotalSupply(&_CurveLPStaking.CallOpts, arg0)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) GetPastTotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetPastTotalSupply(&_CurveLPStaking.CallOpts, arg0)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "getPastVotes", account, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) GetPastVotes(account common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetPastVotes(&_CurveLPStaking.CallOpts, account, arg1)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) GetPastVotes(account common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetPastVotes(&_CurveLPStaking.CallOpts, account, arg1)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) GetVotes(opts *bind.CallOpts, account common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "getVotes", account, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) GetVotes(account common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetVotes(&_CurveLPStaking.CallOpts, account, arg1)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) GetVotes(account common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CurveLPStaking.Contract.GetVotes(&_CurveLPStaking.CallOpts, account, arg1)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_CurveLPStaking *CurveLPStakingCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_CurveLPStaking *CurveLPStakingSession) LpToken() (common.Address, error) {
	return _CurveLPStaking.Contract.LpToken(&_CurveLPStaking.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_CurveLPStaking *CurveLPStakingCallerSession) LpToken() (common.Address, error) {
	return _CurveLPStaking.Contract.LpToken(&_CurveLPStaking.CallOpts)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) Staked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "staked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) Staked(arg0 common.Address) (*big.Int, error) {
	return _CurveLPStaking.Contract.Staked(&_CurveLPStaking.CallOpts, arg0)
}

// Staked is a free data retrieval call binding the contract method 0x98807d84.
//
// Solidity: function staked(address ) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) Staked(arg0 common.Address) (*big.Int, error) {
	return _CurveLPStaking.Contract.Staked(&_CurveLPStaking.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) TotalStaked() (*big.Int, error) {
	return _CurveLPStaking.Contract.TotalStaked(&_CurveLPStaking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) TotalStaked() (*big.Int, error) {
	return _CurveLPStaking.Contract.TotalStaked(&_CurveLPStaking.CallOpts)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCaller) VotingPower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLPStaking.contract.Call(opts, &out, "votingPower", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingSession) VotingPower(user common.Address) (*big.Int, error) {
	return _CurveLPStaking.Contract.VotingPower(&_CurveLPStaking.CallOpts, user)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_CurveLPStaking *CurveLPStakingCallerSession) VotingPower(user common.Address) (*big.Int, error) {
	return _CurveLPStaking.Contract.VotingPower(&_CurveLPStaking.CallOpts, user)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.Stake(&_CurveLPStaking.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.Stake(&_CurveLPStaking.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.Unstake(&_CurveLPStaking.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_CurveLPStaking *CurveLPStakingTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _CurveLPStaking.Contract.Unstake(&_CurveLPStaking.TransactOpts, amount)
}

// CurveLPStakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CurveLPStaking contract.
type CurveLPStakingStakedIterator struct {
	Event *CurveLPStakingStaked // Event containing the contract specifics and raw log

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
func (it *CurveLPStakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLPStakingStaked)
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
		it.Event = new(CurveLPStakingStaked)
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
func (it *CurveLPStakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLPStakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLPStakingStaked represents a Staked event raised by the CurveLPStaking contract.
type CurveLPStakingStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*CurveLPStakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveLPStaking.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingStakedIterator{contract: _CurveLPStaking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CurveLPStakingStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveLPStaking.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLPStakingStaked)
				if err := _CurveLPStaking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) ParseStaked(log types.Log) (*CurveLPStakingStaked, error) {
	event := new(CurveLPStakingStaked)
	if err := _CurveLPStaking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLPStakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the CurveLPStaking contract.
type CurveLPStakingUnstakedIterator struct {
	Event *CurveLPStakingUnstaked // Event containing the contract specifics and raw log

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
func (it *CurveLPStakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLPStakingUnstaked)
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
		it.Event = new(CurveLPStakingUnstaked)
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
func (it *CurveLPStakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLPStakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLPStakingUnstaked represents a Unstaked event raised by the CurveLPStaking contract.
type CurveLPStakingUnstaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*CurveLPStakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveLPStaking.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingUnstakedIterator{contract: _CurveLPStaking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *CurveLPStakingUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveLPStaking.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLPStakingUnstaked)
				if err := _CurveLPStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed user, uint256 amount)
func (_CurveLPStaking *CurveLPStakingFilterer) ParseUnstaked(log types.Log) (*CurveLPStakingUnstaked, error) {
	event := new(CurveLPStakingUnstaked)
	if err := _CurveLPStaking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
