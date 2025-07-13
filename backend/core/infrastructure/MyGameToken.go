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

// MyGameTokenMetaData contains all meta data concerning the MyGameToken contract.
var MyGameTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50336040518060400160405280600c81526020017f4d6f6e73746572546f6b656e00000000000000000000000000000000000000008152506040518060400160405280600381526020017f4d47540000000000000000000000000000000000000000000000000000000000815250816003908162000090919062000472565b508060049081620000a2919062000472565b505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200011a5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016200011191906200059e565b60405180910390fd5b6200012b816200013260201b60201c565b50620005bb565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200027a57607f821691505b60208210810362000290576200028f62000232565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002fa7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002bb565b620003068683620002bb565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003536200034d62000347846200031e565b62000328565b6200031e565b9050919050565b6000819050919050565b6200036f8362000332565b620003876200037e826200035a565b848454620002c8565b825550505050565b600090565b6200039e6200038f565b620003ab81848462000364565b505050565b5b81811015620003d357620003c760008262000394565b600181019050620003b1565b5050565b601f8211156200042257620003ec8162000296565b620003f784620002ab565b8101602085101562000407578190505b6200041f6200041685620002ab565b830182620003b0565b50505b505050565b600082821c905092915050565b6000620004476000198460080262000427565b1980831691505092915050565b600062000462838362000434565b9150826002028217905092915050565b6200047d82620001f8565b67ffffffffffffffff81111562000499576200049862000203565b5b620004a5825462000261565b620004b2828285620003d7565b600060209050601f831160018114620004ea5760008415620004d5578287015190505b620004e1858262000454565b86555062000551565b601f198416620004fa8662000296565b60005b828110156200052457848901518255600182019150602085019450602081019050620004fd565b8683101562000544578489015162000540601f89168262000434565b8355505b6001600288020188555050505b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005868262000559565b9050919050565b620005988162000579565b82525050565b6000602082019050620005b560008301846200058d565b92915050565b61119b80620005cb6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806370a082311161008c57806395d89b411161006657806395d89b4114610202578063a9059cbb14610220578063dd62ed3e14610250578063f2fde38b14610280576100cf565b806370a08231146101aa578063715018a6146101da5780638da5cb5b146101e4576100cf565b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461012257806323b872dd14610140578063313ce5671461017057806340c10f191461018e575b600080fd5b6100dc61029c565b6040516100e99190610def565b60405180910390f35b61010c60048036038101906101079190610eaa565b61032e565b6040516101199190610f05565b60405180910390f35b61012a610351565b6040516101379190610f2f565b60405180910390f35b61015a60048036038101906101559190610f4a565b61035b565b6040516101679190610f05565b60405180910390f35b61017861038a565b6040516101859190610fb9565b60405180910390f35b6101a860048036038101906101a39190610eaa565b610393565b005b6101c460048036038101906101bf9190610fd4565b6103a9565b6040516101d19190610f2f565b60405180910390f35b6101e26103f1565b005b6101ec610405565b6040516101f99190611010565b60405180910390f35b61020a61042f565b6040516102179190610def565b60405180910390f35b61023a60048036038101906102359190610eaa565b6104c1565b6040516102479190610f05565b60405180910390f35b61026a6004803603810190610265919061102b565b6104e4565b6040516102779190610f2f565b60405180910390f35b61029a60048036038101906102959190610fd4565b61056b565b005b6060600380546102ab9061109a565b80601f01602080910402602001604051908101604052809291908181526020018280546102d79061109a565b80156103245780601f106102f957610100808354040283529160200191610324565b820191906000526020600020905b81548152906001019060200180831161030757829003601f168201915b5050505050905090565b6000806103396105f1565b90506103468185856105f9565b600191505092915050565b6000600254905090565b6000806103666105f1565b905061037385828561060b565b61037e8585856106a0565b60019150509392505050565b60006012905090565b61039b610794565b6103a5828261081b565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6103f9610794565b610403600061089d565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461043e9061109a565b80601f016020809104026020016040519081016040528092919081815260200182805461046a9061109a565b80156104b75780601f1061048c576101008083540402835291602001916104b7565b820191906000526020600020905b81548152906001019060200180831161049a57829003601f168201915b5050505050905090565b6000806104cc6105f1565b90506104d98185856106a0565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610573610794565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105e55760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016105dc9190611010565b60405180910390fd5b6105ee8161089d565b50565b600033905090565b6106068383836001610963565b505050565b600061061784846104e4565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561069a578181101561068a578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610681939291906110cb565b60405180910390fd5b61069984848484036000610963565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107125760006040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107099190611010565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107845760006040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161077b9190611010565b60405180910390fd5b61078f838383610b3a565b505050565b61079c6105f1565b73ffffffffffffffffffffffffffffffffffffffff166107ba610405565b73ffffffffffffffffffffffffffffffffffffffff1614610819576107dd6105f1565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016108109190611010565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361088d5760006040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108849190611010565b60405180910390fd5b61089960008383610b3a565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036109d55760006040517fe602df050000000000000000000000000000000000000000000000000000000081526004016109cc9190611010565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a475760006040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610a3e9190611010565b60405180910390fd5b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508015610b34578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610b2b9190610f2f565b60405180910390a35b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b8c578060026000828254610b809190611131565b92505081905550610c5f565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610c18578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610c0f939291906110cb565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ca85780600260008282540392505081905550610cf5565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d529190610f2f565b60405180910390a3505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d99578082015181840152602081019050610d7e565b60008484015250505050565b6000601f19601f8301169050919050565b6000610dc182610d5f565b610dcb8185610d6a565b9350610ddb818560208601610d7b565b610de481610da5565b840191505092915050565b60006020820190508181036000830152610e098184610db6565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e4182610e16565b9050919050565b610e5181610e36565b8114610e5c57600080fd5b50565b600081359050610e6e81610e48565b92915050565b6000819050919050565b610e8781610e74565b8114610e9257600080fd5b50565b600081359050610ea481610e7e565b92915050565b60008060408385031215610ec157610ec0610e11565b5b6000610ecf85828601610e5f565b9250506020610ee085828601610e95565b9150509250929050565b60008115159050919050565b610eff81610eea565b82525050565b6000602082019050610f1a6000830184610ef6565b92915050565b610f2981610e74565b82525050565b6000602082019050610f446000830184610f20565b92915050565b600080600060608486031215610f6357610f62610e11565b5b6000610f7186828701610e5f565b9350506020610f8286828701610e5f565b9250506040610f9386828701610e95565b9150509250925092565b600060ff82169050919050565b610fb381610f9d565b82525050565b6000602082019050610fce6000830184610faa565b92915050565b600060208284031215610fea57610fe9610e11565b5b6000610ff884828501610e5f565b91505092915050565b61100a81610e36565b82525050565b60006020820190506110256000830184611001565b92915050565b6000806040838503121561104257611041610e11565b5b600061105085828601610e5f565b925050602061106185828601610e5f565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806110b257607f821691505b6020821081036110c5576110c461106b565b5b50919050565b60006060820190506110e06000830186611001565b6110ed6020830185610f20565b6110fa6040830184610f20565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061113c82610e74565b915061114783610e74565b925082820190508082111561115f5761115e611102565b5b9291505056fea2646970667358221220445c8e06ef4962ae883442bbf4f8bb534b1f96e98112718eb44ccfcc12b8cc6264736f6c63430008140033",
}

// MyGameTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MyGameTokenMetaData.ABI instead.
var MyGameTokenABI = MyGameTokenMetaData.ABI

// MyGameTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyGameTokenMetaData.Bin instead.
var MyGameTokenBin = MyGameTokenMetaData.Bin

// DeployMyGameToken deploys a new Ethereum contract, binding an instance of MyGameToken to it.
func DeployMyGameToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyGameToken, error) {
	parsed, err := MyGameTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyGameTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyGameToken{MyGameTokenCaller: MyGameTokenCaller{contract: contract}, MyGameTokenTransactor: MyGameTokenTransactor{contract: contract}, MyGameTokenFilterer: MyGameTokenFilterer{contract: contract}}, nil
}

// MyGameToken is an auto generated Go binding around an Ethereum contract.
type MyGameToken struct {
	MyGameTokenCaller     // Read-only binding to the contract
	MyGameTokenTransactor // Write-only binding to the contract
	MyGameTokenFilterer   // Log filterer for contract events
}

// MyGameTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyGameTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGameTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyGameTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGameTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyGameTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGameTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyGameTokenSession struct {
	Contract     *MyGameToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyGameTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyGameTokenCallerSession struct {
	Contract *MyGameTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MyGameTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyGameTokenTransactorSession struct {
	Contract     *MyGameTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MyGameTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyGameTokenRaw struct {
	Contract *MyGameToken // Generic contract binding to access the raw methods on
}

// MyGameTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyGameTokenCallerRaw struct {
	Contract *MyGameTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyGameTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyGameTokenTransactorRaw struct {
	Contract *MyGameTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyGameToken creates a new instance of MyGameToken, bound to a specific deployed contract.
func NewMyGameToken(address common.Address, backend bind.ContractBackend) (*MyGameToken, error) {
	contract, err := bindMyGameToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyGameToken{MyGameTokenCaller: MyGameTokenCaller{contract: contract}, MyGameTokenTransactor: MyGameTokenTransactor{contract: contract}, MyGameTokenFilterer: MyGameTokenFilterer{contract: contract}}, nil
}

// NewMyGameTokenCaller creates a new read-only instance of MyGameToken, bound to a specific deployed contract.
func NewMyGameTokenCaller(address common.Address, caller bind.ContractCaller) (*MyGameTokenCaller, error) {
	contract, err := bindMyGameToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenCaller{contract: contract}, nil
}

// NewMyGameTokenTransactor creates a new write-only instance of MyGameToken, bound to a specific deployed contract.
func NewMyGameTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyGameTokenTransactor, error) {
	contract, err := bindMyGameToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenTransactor{contract: contract}, nil
}

// NewMyGameTokenFilterer creates a new log filterer instance of MyGameToken, bound to a specific deployed contract.
func NewMyGameTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyGameTokenFilterer, error) {
	contract, err := bindMyGameToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenFilterer{contract: contract}, nil
}

// bindMyGameToken binds a generic wrapper to an already deployed contract.
func bindMyGameToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyGameTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyGameToken *MyGameTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyGameToken.Contract.MyGameTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyGameToken *MyGameTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGameToken.Contract.MyGameTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyGameToken *MyGameTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyGameToken.Contract.MyGameTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyGameToken *MyGameTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyGameToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyGameToken *MyGameTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGameToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyGameToken *MyGameTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyGameToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyGameToken *MyGameTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyGameToken *MyGameTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyGameToken.Contract.Allowance(&_MyGameToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyGameToken *MyGameTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyGameToken.Contract.Allowance(&_MyGameToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyGameToken *MyGameTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyGameToken *MyGameTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyGameToken.Contract.BalanceOf(&_MyGameToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyGameToken *MyGameTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyGameToken.Contract.BalanceOf(&_MyGameToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyGameToken *MyGameTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyGameToken *MyGameTokenSession) Decimals() (uint8, error) {
	return _MyGameToken.Contract.Decimals(&_MyGameToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyGameToken *MyGameTokenCallerSession) Decimals() (uint8, error) {
	return _MyGameToken.Contract.Decimals(&_MyGameToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGameToken *MyGameTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGameToken *MyGameTokenSession) Name() (string, error) {
	return _MyGameToken.Contract.Name(&_MyGameToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGameToken *MyGameTokenCallerSession) Name() (string, error) {
	return _MyGameToken.Contract.Name(&_MyGameToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyGameToken *MyGameTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyGameToken *MyGameTokenSession) Owner() (common.Address, error) {
	return _MyGameToken.Contract.Owner(&_MyGameToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyGameToken *MyGameTokenCallerSession) Owner() (common.Address, error) {
	return _MyGameToken.Contract.Owner(&_MyGameToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyGameToken *MyGameTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyGameToken *MyGameTokenSession) Symbol() (string, error) {
	return _MyGameToken.Contract.Symbol(&_MyGameToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyGameToken *MyGameTokenCallerSession) Symbol() (string, error) {
	return _MyGameToken.Contract.Symbol(&_MyGameToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyGameToken *MyGameTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGameToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyGameToken *MyGameTokenSession) TotalSupply() (*big.Int, error) {
	return _MyGameToken.Contract.TotalSupply(&_MyGameToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyGameToken *MyGameTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MyGameToken.Contract.TotalSupply(&_MyGameToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Approve(&_MyGameToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Approve(&_MyGameToken.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyGameToken *MyGameTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyGameToken *MyGameTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Mint(&_MyGameToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyGameToken *MyGameTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Mint(&_MyGameToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyGameToken *MyGameTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyGameToken *MyGameTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyGameToken.Contract.RenounceOwnership(&_MyGameToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyGameToken *MyGameTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyGameToken.Contract.RenounceOwnership(&_MyGameToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Transfer(&_MyGameToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.Transfer(&_MyGameToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.TransferFrom(&_MyGameToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyGameToken *MyGameTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyGameToken.Contract.TransferFrom(&_MyGameToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyGameToken *MyGameTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyGameToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyGameToken *MyGameTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyGameToken.Contract.TransferOwnership(&_MyGameToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyGameToken *MyGameTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyGameToken.Contract.TransferOwnership(&_MyGameToken.TransactOpts, newOwner)
}

// MyGameTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyGameToken contract.
type MyGameTokenApprovalIterator struct {
	Event *MyGameTokenApproval // Event containing the contract specifics and raw log

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
func (it *MyGameTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGameTokenApproval)
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
		it.Event = new(MyGameTokenApproval)
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
func (it *MyGameTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGameTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGameTokenApproval represents a Approval event raised by the MyGameToken contract.
type MyGameTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyGameToken *MyGameTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyGameTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyGameToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenApprovalIterator{contract: _MyGameToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyGameToken *MyGameTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyGameTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyGameToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGameTokenApproval)
				if err := _MyGameToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MyGameToken *MyGameTokenFilterer) ParseApproval(log types.Log) (*MyGameTokenApproval, error) {
	event := new(MyGameTokenApproval)
	if err := _MyGameToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGameTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyGameToken contract.
type MyGameTokenOwnershipTransferredIterator struct {
	Event *MyGameTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MyGameTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGameTokenOwnershipTransferred)
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
		it.Event = new(MyGameTokenOwnershipTransferred)
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
func (it *MyGameTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGameTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGameTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MyGameToken contract.
type MyGameTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyGameToken *MyGameTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyGameTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyGameToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenOwnershipTransferredIterator{contract: _MyGameToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyGameToken *MyGameTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyGameTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyGameToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGameTokenOwnershipTransferred)
				if err := _MyGameToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MyGameToken *MyGameTokenFilterer) ParseOwnershipTransferred(log types.Log) (*MyGameTokenOwnershipTransferred, error) {
	event := new(MyGameTokenOwnershipTransferred)
	if err := _MyGameToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGameTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyGameToken contract.
type MyGameTokenTransferIterator struct {
	Event *MyGameTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyGameTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGameTokenTransfer)
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
		it.Event = new(MyGameTokenTransfer)
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
func (it *MyGameTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGameTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGameTokenTransfer represents a Transfer event raised by the MyGameToken contract.
type MyGameTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyGameToken *MyGameTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyGameTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyGameToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyGameTokenTransferIterator{contract: _MyGameToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyGameToken *MyGameTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyGameTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyGameToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGameTokenTransfer)
				if err := _MyGameToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MyGameToken *MyGameTokenFilterer) ParseTransfer(log types.Log) (*MyGameTokenTransfer, error) {
	event := new(MyGameTokenTransfer)
	if err := _MyGameToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
