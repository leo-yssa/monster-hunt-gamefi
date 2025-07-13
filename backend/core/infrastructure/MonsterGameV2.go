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

// MonsterGameV2MetaData contains all meta data concerning the MonsterGameV2 contract.
var MonsterGameV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterHunted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PlayerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"addMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"}],\"name\":\"huntMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsters\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerPlayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractMyGameToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611f17806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063b3446f851161008c578063c4d66de811610066578063c4d66de8146101ed578063e2eb41ff14610209578063f2fde38b1461023b578063f7c618c114610257576100ea565b8063b3446f8514610183578063b906b3001461019f578063c0576b73146101bb576100ea565b80635c975abb116100c85780635c975abb14610133578063715018a6146101515780638456cb591461015b5780638da5cb5b14610165576100ea565b806318f237bc146100ef5780633f4ba83a1461010b57806354fd4d5014610115575b600080fd5b61010960048036038101906101049190611463565b610275565b005b610113610445565b005b61011d610457565b60405161012a9190611551565b60405180910390f35b61013b610494565b604051610148919061158e565b60405180910390f35b6101596104b9565b005b6101636104cd565b005b61016d6104df565b60405161017a91906115ea565b60405180910390f35b61019d60048036038101906101989190611605565b610517565b005b6101b960048036038101906101b4919061164e565b6106f3565b005b6101d560048036038101906101d0919061164e565b6109a0565b6040516101e49392919061168a565b60405180910390f35b610207600480360381019061020291906116f4565b610a62565b005b610223600480360381019061021e91906116f4565b610c41565b60405161023293929190611721565b60405180910390f35b610255600480360381019061025091906116f4565b610d00565b005b61025f610d86565b60405161026c91906117be565b60405180910390f35b61027d610daa565b610285610e31565b6000835111801561029857506020835111155b6102d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ce90611825565b60405180910390fd5b6000821180156102e957506127108211155b610328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031f90611891565b60405180910390fd5b600081118015610342575069d3c21bcecceda10000008111155b610381576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610378906118fd565b60405180910390fd5b6002604051806060016040528085815260200184815260200183815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000190816103dc9190611b1f565b506020820151816001015560408201518160020155505060016002805490506104059190611c20565b7ffa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b8484846040516104389392919061168a565b60405180910390a2505050565b61044d610daa565b610455610e72565b565b60606040518060400160405280600281526020017f5632000000000000000000000000000000000000000000000000000000000000815250905090565b60008061049f610ee4565b90508060000160009054906101000a900460ff1691505090565b6104c1610daa565b6104cb6000610f0c565b565b6104d5610daa565b6104dd610fe3565b565b6000806104ea611055565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b61051f610e31565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff16156105af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a690611ca0565b60405180910390fd5b600081511180156105c257506020815111155b610601576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f890611d0c565b60405180910390fd5b60405180606001604052808281526020016001815260200160011515815250600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816106749190611b1f565b506020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167fe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a826040516106e89190611551565b60405180910390a250565b6106fb610e31565b61070361107d565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff16610792576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078990611d78565b60405180910390fd5b60028054905081106107d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d090611de4565b60405180910390fd5b6000600282815481106107ef576107ee611e04565b5b90600052602060002090600302016040518060600160405290816000820180546108189061194c565b80601f01602080910402602001604051908101604052809291908181526020018280546108449061194c565b80156108915780601f1061086657610100808354040283529160200191610891565b820191906000526020600020905b81548152906001019060200180831161087457829003601f168201915b5050505050815260200160018201548152602001600282015481525050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f193383604001516040518363ffffffff1660e01b815260040161090f929190611e33565b600060405180830381600087803b15801561092957600080fd5b505af115801561093d573d6000803e3d6000fd5b50505050813373ffffffffffffffffffffffffffffffffffffffff167f3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460836040015160405161098c9190611e5c565b60405180910390a35061099d6110d4565b50565b600281815481106109b057600080fd5b90600052602060002090600302016000915090508060000180546109d39061194c565b80601f01602080910402602001604051908101604052809291908181526020018280546109ff9061194c565b8015610a4c5780601f10610a2157610100808354040283529160200191610a4c565b820191906000526020600020905b815481529060010190602001808311610a2f57829003601f168201915b5050505050908060010154908060020154905083565b6000610a6c6110ed565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff16148015610aba5750825b9050600060018367ffffffffffffffff16148015610aef575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610afd575080155b15610b34576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610b845760018560000160086101000a81548160ff0219169083151502179055505b610b8d33611101565b610b95611115565b610b9d61111f565b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610c395760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610c309190611ec6565b60405180910390a15b505050505050565b6001602052806000526040600020600091509050806000018054610c649061194c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c909061194c565b8015610cdd5780601f10610cb257610100808354040283529160200191610cdd565b820191906000526020600020905b815481529060010190602001808311610cc057829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b610d08610daa565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d7a5760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610d7191906115ea565b60405180910390fd5b610d8381610f0c565b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610db2611131565b73ffffffffffffffffffffffffffffffffffffffff16610dd06104df565b73ffffffffffffffffffffffffffffffffffffffff1614610e2f57610df3611131565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610e2691906115ea565b60405180910390fd5b565b610e39610494565b15610e70576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610e7a611139565b6000610e84610ee4565b905060008160000160006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610ecc611131565b604051610ed991906115ea565b60405180910390a150565b60007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b6000610f16611055565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b610feb610e31565b6000610ff5610ee4565b905060018160000160006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861103d611131565b60405161104a91906115ea565b60405180910390a150565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b6000611087611179565b905060028160000154036110c7576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002816000018190555050565b60006110de611179565b90506001816000018190555050565b6000806110f86111a1565b90508091505090565b6111096111cc565b6111128161120c565b50565b61111d6111cc565b565b6111276111cc565b61112f611292565b565b600033905090565b611141610494565b611177576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00905090565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0060001b905090565b6111d46112b3565b61120a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6112146111cc565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036112865760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161127d91906115ea565b60405180910390fd5b61128f81610f0c565b50565b61129a6111cc565b60006112a4611179565b90506001816000018190555050565b60006112bd6110ed565b60000160089054906101000a900460ff16905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61133a826112f1565b810181811067ffffffffffffffff8211171561135957611358611302565b5b80604052505050565b600061136c6112d3565b90506113788282611331565b919050565b600067ffffffffffffffff82111561139857611397611302565b5b6113a1826112f1565b9050602081019050919050565b82818337600083830152505050565b60006113d06113cb8461137d565b611362565b9050828152602081018484840111156113ec576113eb6112ec565b5b6113f78482856113ae565b509392505050565b600082601f830112611414576114136112e7565b5b81356114248482602086016113bd565b91505092915050565b6000819050919050565b6114408161142d565b811461144b57600080fd5b50565b60008135905061145d81611437565b92915050565b60008060006060848603121561147c5761147b6112dd565b5b600084013567ffffffffffffffff81111561149a576114996112e2565b5b6114a6868287016113ff565b93505060206114b78682870161144e565b92505060406114c88682870161144e565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b60005b8381101561150c5780820151818401526020810190506114f1565b60008484015250505050565b6000611523826114d2565b61152d81856114dd565b935061153d8185602086016114ee565b611546816112f1565b840191505092915050565b6000602082019050818103600083015261156b8184611518565b905092915050565b60008115159050919050565b61158881611573565b82525050565b60006020820190506115a3600083018461157f565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115d4826115a9565b9050919050565b6115e4816115c9565b82525050565b60006020820190506115ff60008301846115db565b92915050565b60006020828403121561161b5761161a6112dd565b5b600082013567ffffffffffffffff811115611639576116386112e2565b5b611645848285016113ff565b91505092915050565b600060208284031215611664576116636112dd565b5b60006116728482850161144e565b91505092915050565b6116848161142d565b82525050565b600060608201905081810360008301526116a48186611518565b90506116b3602083018561167b565b6116c0604083018461167b565b949350505050565b6116d1816115c9565b81146116dc57600080fd5b50565b6000813590506116ee816116c8565b92915050565b60006020828403121561170a576117096112dd565b5b6000611718848285016116df565b91505092915050565b6000606082019050818103600083015261173b8186611518565b905061174a602083018561167b565b611757604083018461157f565b949350505050565b6000819050919050565b600061178461177f61177a846115a9565b61175f565b6115a9565b9050919050565b600061179682611769565b9050919050565b60006117a88261178b565b9050919050565b6117b88161179d565b82525050565b60006020820190506117d360008301846117af565b92915050565b7f496e76616c6964206d6f6e73746572206e616d65000000000000000000000000600082015250565b600061180f6014836114dd565b915061181a826117d9565b602082019050919050565b6000602082019050818103600083015261183e81611802565b9050919050565b7f496e76616c696420485000000000000000000000000000000000000000000000600082015250565b600061187b600a836114dd565b915061188682611845565b602082019050919050565b600060208201905081810360008301526118aa8161186e565b9050919050565b7f496e76616c696420726577617264000000000000000000000000000000000000600082015250565b60006118e7600e836114dd565b91506118f2826118b1565b602082019050919050565b60006020820190508181036000830152611916816118da565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061196457607f821691505b6020821081036119775761197661191d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826119a2565b6119e986836119a2565b95508019841693508086168417925050509392505050565b6000611a1c611a17611a128461142d565b61175f565b61142d565b9050919050565b6000819050919050565b611a3683611a01565b611a4a611a4282611a23565b8484546119af565b825550505050565b600090565b611a5f611a52565b611a6a818484611a2d565b505050565b5b81811015611a8e57611a83600082611a57565b600181019050611a70565b5050565b601f821115611ad357611aa48161197d565b611aad84611992565b81016020851015611abc578190505b611ad0611ac885611992565b830182611a6f565b50505b505050565b600082821c905092915050565b6000611af660001984600802611ad8565b1980831691505092915050565b6000611b0f8383611ae5565b9150826002028217905092915050565b611b28826114d2565b67ffffffffffffffff811115611b4157611b40611302565b5b611b4b825461194c565b611b56828285611a92565b600060209050601f831160018114611b895760008415611b77578287015190505b611b818582611b03565b865550611be9565b601f198416611b978661197d565b60005b82811015611bbf57848901518255600182019150602085019450602081019050611b9a565b86831015611bdc5784890151611bd8601f891682611ae5565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c2b8261142d565b9150611c368361142d565b9250828203905081811115611c4e57611c4d611bf1565b5b92915050565b7f416c726561647920726567697374657265640000000000000000000000000000600082015250565b6000611c8a6012836114dd565b9150611c9582611c54565b602082019050919050565b60006020820190508181036000830152611cb981611c7d565b9050919050565b7f496e76616c6964206e616d650000000000000000000000000000000000000000600082015250565b6000611cf6600c836114dd565b9150611d0182611cc0565b602082019050919050565b60006020820190508181036000830152611d2581611ce9565b9050919050565b7f506c61796572206e6f7420726567697374657265640000000000000000000000600082015250565b6000611d626015836114dd565b9150611d6d82611d2c565b602082019050919050565b60006020820190508181036000830152611d9181611d55565b9050919050565b7f496e76616c6964206d6f6e737465720000000000000000000000000000000000600082015250565b6000611dce600f836114dd565b9150611dd982611d98565b602082019050919050565b60006020820190508181036000830152611dfd81611dc1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604082019050611e4860008301856115db565b611e55602083018461167b565b9392505050565b6000602082019050611e71600083018461167b565b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000611eb0611eab611ea684611e77565b61175f565b611e81565b9050919050565b611ec081611e95565b82525050565b6000602082019050611edb6000830184611eb7565b9291505056fea2646970667358221220ffb1498b04bde97931e76d476801deafbf36a62724cdd7cb73d4021d2cd527a864736f6c63430008140033",
}

// MonsterGameV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MonsterGameV2MetaData.ABI instead.
var MonsterGameV2ABI = MonsterGameV2MetaData.ABI

// MonsterGameV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MonsterGameV2MetaData.Bin instead.
var MonsterGameV2Bin = MonsterGameV2MetaData.Bin

// DeployMonsterGameV2 deploys a new Ethereum contract, binding an instance of MonsterGameV2 to it.
func DeployMonsterGameV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterGameV2, error) {
	parsed, err := MonsterGameV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MonsterGameV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterGameV2{MonsterGameV2Caller: MonsterGameV2Caller{contract: contract}, MonsterGameV2Transactor: MonsterGameV2Transactor{contract: contract}, MonsterGameV2Filterer: MonsterGameV2Filterer{contract: contract}}, nil
}

// MonsterGameV2 is an auto generated Go binding around an Ethereum contract.
type MonsterGameV2 struct {
	MonsterGameV2Caller     // Read-only binding to the contract
	MonsterGameV2Transactor // Write-only binding to the contract
	MonsterGameV2Filterer   // Log filterer for contract events
}

// MonsterGameV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterGameV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterGameV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterGameV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterGameV2Session struct {
	Contract     *MonsterGameV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonsterGameV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterGameV2CallerSession struct {
	Contract *MonsterGameV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MonsterGameV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterGameV2TransactorSession struct {
	Contract     *MonsterGameV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MonsterGameV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterGameV2Raw struct {
	Contract *MonsterGameV2 // Generic contract binding to access the raw methods on
}

// MonsterGameV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterGameV2CallerRaw struct {
	Contract *MonsterGameV2Caller // Generic read-only contract binding to access the raw methods on
}

// MonsterGameV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterGameV2TransactorRaw struct {
	Contract *MonsterGameV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterGameV2 creates a new instance of MonsterGameV2, bound to a specific deployed contract.
func NewMonsterGameV2(address common.Address, backend bind.ContractBackend) (*MonsterGameV2, error) {
	contract, err := bindMonsterGameV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2{MonsterGameV2Caller: MonsterGameV2Caller{contract: contract}, MonsterGameV2Transactor: MonsterGameV2Transactor{contract: contract}, MonsterGameV2Filterer: MonsterGameV2Filterer{contract: contract}}, nil
}

// NewMonsterGameV2Caller creates a new read-only instance of MonsterGameV2, bound to a specific deployed contract.
func NewMonsterGameV2Caller(address common.Address, caller bind.ContractCaller) (*MonsterGameV2Caller, error) {
	contract, err := bindMonsterGameV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2Caller{contract: contract}, nil
}

// NewMonsterGameV2Transactor creates a new write-only instance of MonsterGameV2, bound to a specific deployed contract.
func NewMonsterGameV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MonsterGameV2Transactor, error) {
	contract, err := bindMonsterGameV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2Transactor{contract: contract}, nil
}

// NewMonsterGameV2Filterer creates a new log filterer instance of MonsterGameV2, bound to a specific deployed contract.
func NewMonsterGameV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MonsterGameV2Filterer, error) {
	contract, err := bindMonsterGameV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2Filterer{contract: contract}, nil
}

// bindMonsterGameV2 binds a generic wrapper to an already deployed contract.
func bindMonsterGameV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MonsterGameV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterGameV2 *MonsterGameV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MonsterGameV2.Contract.MonsterGameV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterGameV2 *MonsterGameV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.MonsterGameV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterGameV2 *MonsterGameV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.MonsterGameV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterGameV2 *MonsterGameV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MonsterGameV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterGameV2 *MonsterGameV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterGameV2 *MonsterGameV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.contract.Transact(opts, method, params...)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2Caller) Monsters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "monsters", arg0)

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
func (_MonsterGameV2 *MonsterGameV2Session) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _MonsterGameV2.Contract.Monsters(&_MonsterGameV2.CallOpts, arg0)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2CallerSession) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _MonsterGameV2.Contract.Monsters(&_MonsterGameV2.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGameV2 *MonsterGameV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGameV2 *MonsterGameV2Session) Owner() (common.Address, error) {
	return _MonsterGameV2.Contract.Owner(&_MonsterGameV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGameV2 *MonsterGameV2CallerSession) Owner() (common.Address, error) {
	return _MonsterGameV2.Contract.Owner(&_MonsterGameV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGameV2 *MonsterGameV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGameV2 *MonsterGameV2Session) Paused() (bool, error) {
	return _MonsterGameV2.Contract.Paused(&_MonsterGameV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGameV2 *MonsterGameV2CallerSession) Paused() (bool, error) {
	return _MonsterGameV2.Contract.Paused(&_MonsterGameV2.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_MonsterGameV2 *MonsterGameV2Caller) Players(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "players", arg0)

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
func (_MonsterGameV2 *MonsterGameV2Session) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _MonsterGameV2.Contract.Players(&_MonsterGameV2.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_MonsterGameV2 *MonsterGameV2CallerSession) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _MonsterGameV2.Contract.Players(&_MonsterGameV2.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGameV2 *MonsterGameV2Caller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGameV2 *MonsterGameV2Session) RewardToken() (common.Address, error) {
	return _MonsterGameV2.Contract.RewardToken(&_MonsterGameV2.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGameV2 *MonsterGameV2CallerSession) RewardToken() (common.Address, error) {
	return _MonsterGameV2.Contract.RewardToken(&_MonsterGameV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MonsterGameV2 *MonsterGameV2Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MonsterGameV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MonsterGameV2 *MonsterGameV2Session) Version() (string, error) {
	return _MonsterGameV2.Contract.Version(&_MonsterGameV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_MonsterGameV2 *MonsterGameV2CallerSession) Version() (string, error) {
	return _MonsterGameV2.Contract.Version(&_MonsterGameV2.CallOpts)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) AddMonster(opts *bind.TransactOpts, name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "addMonster", name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGameV2 *MonsterGameV2Session) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.AddMonster(&_MonsterGameV2.TransactOpts, name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.AddMonster(&_MonsterGameV2.TransactOpts, name, hp, reward)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) HuntMonster(opts *bind.TransactOpts, monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "huntMonster", monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGameV2 *MonsterGameV2Session) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.HuntMonster(&_MonsterGameV2.TransactOpts, monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.HuntMonster(&_MonsterGameV2.TransactOpts, monsterId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) Initialize(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "initialize", tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGameV2 *MonsterGameV2Session) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Initialize(&_MonsterGameV2.TransactOpts, tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Initialize(&_MonsterGameV2.TransactOpts, tokenAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGameV2 *MonsterGameV2Session) Pause() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Pause(&_MonsterGameV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Pause(&_MonsterGameV2.TransactOpts)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) RegisterPlayer(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "registerPlayer", name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGameV2 *MonsterGameV2Session) RegisterPlayer(name string) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.RegisterPlayer(&_MonsterGameV2.TransactOpts, name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) RegisterPlayer(name string) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.RegisterPlayer(&_MonsterGameV2.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGameV2 *MonsterGameV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.RenounceOwnership(&_MonsterGameV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.RenounceOwnership(&_MonsterGameV2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGameV2 *MonsterGameV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.TransferOwnership(&_MonsterGameV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGameV2.Contract.TransferOwnership(&_MonsterGameV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGameV2 *MonsterGameV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGameV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGameV2 *MonsterGameV2Session) Unpause() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Unpause(&_MonsterGameV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGameV2 *MonsterGameV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterGameV2.Contract.Unpause(&_MonsterGameV2.TransactOpts)
}

// MonsterGameV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MonsterGameV2 contract.
type MonsterGameV2InitializedIterator struct {
	Event *MonsterGameV2Initialized // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2Initialized)
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
		it.Event = new(MonsterGameV2Initialized)
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
func (it *MonsterGameV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2Initialized represents a Initialized event raised by the MonsterGameV2 contract.
type MonsterGameV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*MonsterGameV2InitializedIterator, error) {

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2InitializedIterator{contract: _MonsterGameV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MonsterGameV2Initialized) (event.Subscription, error) {

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2Initialized)
				if err := _MonsterGameV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MonsterGameV2 *MonsterGameV2Filterer) ParseInitialized(log types.Log) (*MonsterGameV2Initialized, error) {
	event := new(MonsterGameV2Initialized)
	if err := _MonsterGameV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2MonsterAddedIterator is returned from FilterMonsterAdded and is used to iterate over the raw logs and unpacked data for MonsterAdded events raised by the MonsterGameV2 contract.
type MonsterGameV2MonsterAddedIterator struct {
	Event *MonsterGameV2MonsterAdded // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2MonsterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2MonsterAdded)
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
		it.Event = new(MonsterGameV2MonsterAdded)
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
func (it *MonsterGameV2MonsterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2MonsterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2MonsterAdded represents a MonsterAdded event raised by the MonsterGameV2 contract.
type MonsterGameV2MonsterAdded struct {
	MonsterId *big.Int
	Name      string
	Hp        *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterAdded is a free log retrieval operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterMonsterAdded(opts *bind.FilterOpts, monsterId []*big.Int) (*MonsterGameV2MonsterAddedIterator, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2MonsterAddedIterator{contract: _MonsterGameV2.contract, event: "MonsterAdded", logs: logs, sub: sub}, nil
}

// WatchMonsterAdded is a free log subscription operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchMonsterAdded(opts *bind.WatchOpts, sink chan<- *MonsterGameV2MonsterAdded, monsterId []*big.Int) (event.Subscription, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2MonsterAdded)
				if err := _MonsterGameV2.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParseMonsterAdded(log types.Log) (*MonsterGameV2MonsterAdded, error) {
	event := new(MonsterGameV2MonsterAdded)
	if err := _MonsterGameV2.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2MonsterHuntedIterator is returned from FilterMonsterHunted and is used to iterate over the raw logs and unpacked data for MonsterHunted events raised by the MonsterGameV2 contract.
type MonsterGameV2MonsterHuntedIterator struct {
	Event *MonsterGameV2MonsterHunted // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2MonsterHuntedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2MonsterHunted)
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
		it.Event = new(MonsterGameV2MonsterHunted)
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
func (it *MonsterGameV2MonsterHuntedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2MonsterHuntedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2MonsterHunted represents a MonsterHunted event raised by the MonsterGameV2 contract.
type MonsterGameV2MonsterHunted struct {
	Player    common.Address
	MonsterId *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterHunted is a free log retrieval operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterMonsterHunted(opts *bind.FilterOpts, player []common.Address, monsterId []*big.Int) (*MonsterGameV2MonsterHuntedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2MonsterHuntedIterator{contract: _MonsterGameV2.contract, event: "MonsterHunted", logs: logs, sub: sub}, nil
}

// WatchMonsterHunted is a free log subscription operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchMonsterHunted(opts *bind.WatchOpts, sink chan<- *MonsterGameV2MonsterHunted, player []common.Address, monsterId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2MonsterHunted)
				if err := _MonsterGameV2.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParseMonsterHunted(log types.Log) (*MonsterGameV2MonsterHunted, error) {
	event := new(MonsterGameV2MonsterHunted)
	if err := _MonsterGameV2.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonsterGameV2 contract.
type MonsterGameV2OwnershipTransferredIterator struct {
	Event *MonsterGameV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2OwnershipTransferred)
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
		it.Event = new(MonsterGameV2OwnershipTransferred)
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
func (it *MonsterGameV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2OwnershipTransferred represents a OwnershipTransferred event raised by the MonsterGameV2 contract.
type MonsterGameV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonsterGameV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2OwnershipTransferredIterator{contract: _MonsterGameV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonsterGameV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2OwnershipTransferred)
				if err := _MonsterGameV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParseOwnershipTransferred(log types.Log) (*MonsterGameV2OwnershipTransferred, error) {
	event := new(MonsterGameV2OwnershipTransferred)
	if err := _MonsterGameV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MonsterGameV2 contract.
type MonsterGameV2PausedIterator struct {
	Event *MonsterGameV2Paused // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2Paused)
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
		it.Event = new(MonsterGameV2Paused)
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
func (it *MonsterGameV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2Paused represents a Paused event raised by the MonsterGameV2 contract.
type MonsterGameV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterPaused(opts *bind.FilterOpts) (*MonsterGameV2PausedIterator, error) {

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2PausedIterator{contract: _MonsterGameV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MonsterGameV2Paused) (event.Subscription, error) {

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2Paused)
				if err := _MonsterGameV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParsePaused(log types.Log) (*MonsterGameV2Paused, error) {
	event := new(MonsterGameV2Paused)
	if err := _MonsterGameV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2PlayerRegisteredIterator is returned from FilterPlayerRegistered and is used to iterate over the raw logs and unpacked data for PlayerRegistered events raised by the MonsterGameV2 contract.
type MonsterGameV2PlayerRegisteredIterator struct {
	Event *MonsterGameV2PlayerRegistered // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2PlayerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2PlayerRegistered)
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
		it.Event = new(MonsterGameV2PlayerRegistered)
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
func (it *MonsterGameV2PlayerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2PlayerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2PlayerRegistered represents a PlayerRegistered event raised by the MonsterGameV2 contract.
type MonsterGameV2PlayerRegistered struct {
	Player common.Address
	Name   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlayerRegistered is a free log retrieval operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterPlayerRegistered(opts *bind.FilterOpts, player []common.Address) (*MonsterGameV2PlayerRegisteredIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2PlayerRegisteredIterator{contract: _MonsterGameV2.contract, event: "PlayerRegistered", logs: logs, sub: sub}, nil
}

// WatchPlayerRegistered is a free log subscription operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchPlayerRegistered(opts *bind.WatchOpts, sink chan<- *MonsterGameV2PlayerRegistered, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2PlayerRegistered)
				if err := _MonsterGameV2.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParsePlayerRegistered(log types.Log) (*MonsterGameV2PlayerRegistered, error) {
	event := new(MonsterGameV2PlayerRegistered)
	if err := _MonsterGameV2.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MonsterGameV2 contract.
type MonsterGameV2UnpausedIterator struct {
	Event *MonsterGameV2Unpaused // Event containing the contract specifics and raw log

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
func (it *MonsterGameV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameV2Unpaused)
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
		it.Event = new(MonsterGameV2Unpaused)
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
func (it *MonsterGameV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameV2Unpaused represents a Unpaused event raised by the MonsterGameV2 contract.
type MonsterGameV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MonsterGameV2 *MonsterGameV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*MonsterGameV2UnpausedIterator, error) {

	logs, sub, err := _MonsterGameV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MonsterGameV2UnpausedIterator{contract: _MonsterGameV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MonsterGameV2 *MonsterGameV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MonsterGameV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _MonsterGameV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameV2Unpaused)
				if err := _MonsterGameV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MonsterGameV2 *MonsterGameV2Filterer) ParseUnpaused(log types.Log) (*MonsterGameV2Unpaused, error) {
	event := new(MonsterGameV2Unpaused)
	if err := _MonsterGameV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
