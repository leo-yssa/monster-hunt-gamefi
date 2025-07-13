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

// MonsterGameMetaData contains all meta data concerning the MonsterGame contract.
var MonsterGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"MonsterHunted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"PlayerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"addMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"monsterId\",\"type\":\"uint256\"}],\"name\":\"huntMonster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monsters\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"registerPlayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractMyGameToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ea1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063b3446f851161008c578063c4d66de811610066578063c4d66de8146101b4578063e2eb41ff146101d0578063f2fde38b14610202578063f7c618c11461021e576100cf565b8063b3446f851461014a578063b906b30014610166578063c0576b7314610182576100cf565b806318f237bc146100d45780633f4ba83a146100f05780635c975abb146100fa578063715018a6146101185780638456cb59146101225780638da5cb5b1461012c575b600080fd5b6100ee60048036038101906100e991906113ed565b61023c565b005b6100f861040c565b005b61010261041e565b60405161010f9190611477565b60405180910390f35b610120610443565b005b61012a610457565b005b610134610469565b60405161014191906114d3565b60405180910390f35b610164600480360381019061015f91906114ee565b6104a1565b005b610180600480360381019061017b9190611537565b61067d565b005b61019c60048036038101906101979190611537565b61092a565b6040516101ab939291906115f2565b60405180910390f35b6101ce60048036038101906101c9919061165c565b6109ec565b005b6101ea60048036038101906101e5919061165c565b610bcb565b6040516101f993929190611689565b60405180910390f35b61021c6004803603810190610217919061165c565b610c8a565b005b610226610d10565b6040516102339190611726565b60405180910390f35b610244610d34565b61024c610dbb565b6000835111801561025f57506020835111155b61029e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102959061178d565b60405180910390fd5b6000821180156102b057506127108211155b6102ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e6906117f9565b60405180910390fd5b600081118015610309575069d3c21bcecceda10000008111155b610348576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033f90611865565b60405180910390fd5b6002604051806060016040528085815260200184815260200183815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000190816103a39190611a87565b506020820151816001015560408201518160020155505060016002805490506103cc9190611b88565b7ffa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b8484846040516103ff939291906115f2565b60405180910390a2505050565b610414610d34565b61041c610dfc565b565b600080610429610e6e565b90508060000160009054906101000a900460ff1691505090565b61044b610d34565b6104556000610e96565b565b61045f610d34565b610467610f6d565b565b600080610474610fdf565b90508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6104a9610dbb565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1615610539576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053090611c08565b60405180910390fd5b6000815111801561054c57506020815111155b61058b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058290611c74565b60405180606001604052808281526020016001815260200160011515815250600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816105fe9190611a87565b506020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167fe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a826040516106729190611c94565b60405180910390a2505050565b610685610dbb565b61068d611007565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff1661071c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071390611d02565b60405180910390fd5b6002805490508110610763576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075a90611d6e565b60405180910390fd5b60006002828154811061077957610778611d8e565b5b90600052602060002090600302016040518060600160405290816000820180546107a2906118b4565b80601f01602080910402602001604051908101604052809291908181526020018280546107ce906118b4565b801561081b5780601f106107f05761010080835404028352916020019161081b565b820191906000526020600020905b8154815290600101906020018083116107fe57829003601f168201915b5050505050815260200160018201548152602001600282015481525050905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f193383604001516040518363ffffffff1660e01b8152600401610899929190611dbd565b600060405180830381600087803b1580156108b357600080fd5b505af11580156108c7573d6000803e3d6000fd5b50505050813373ffffffffffffffffffffffffffffffffffffffff167f3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c26546083604001516040516109169190611de6565b60405180910390a35061092761105e565b50565b6002818154811061093a57600080fd5b906000526020600020906003020160009150905080600001805461095d906118b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610989906118b4565b80156109d65780601f106109ab576101008083540402835291602001916109d6565b820191906000526020600020905b8154815290600101906020018083116109b957829003601f168201915b5050505050908060010154908060020154905083565b60006109f6611077565b905060008160000160089054906101000a900460ff1615905060008260000160009054906101000a900467ffffffffffffffff1690506000808267ffffffffffffffff16148015610a445750825b9050600060018367ffffffffffffffff16148015610a79575060003073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610a87575080155b15610abe576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018560000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610b0e5760018560000160086101000a81548160ff0219169083151502179055505b610b173361108b565b610b1f61109f565b610b276110a9565b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610bc35760008560000160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610bba9190611e50565b60405180910390a15b505050505050565b6001602052806000526040600020600091509050806000018054610bee906118b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1a906118b4565b8015610c675780601f10610c3c57610100808354040283529160200191610c67565b820191906000526020600020905b815481529060010190602001808311610c4a57829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b610c92610d34565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d045760006040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610cfb91906114d3565b60405180910390fd5b610d0d81610e96565b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610d3c6110bb565b73ffffffffffffffffffffffffffffffffffffffff16610d5a610469565b73ffffffffffffffffffffffffffffffffffffffff1614610db957610d7d6110bb565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610db091906114d3565b60405180910390fd5b565b610dc361041e565b15610dfa576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610e046110c3565b6000610e0e610e6e565b905060008160000160006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610e566110bb565b604051610e6391906114d3565b60405180910390a150565b60007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300905090565b6000610ea0610fdf565b905060008160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050828260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b610f75610dbb565b6000610f7f610e6e565b905060018160000160006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610fc76110bb565b604051610fd491906114d3565b60405180910390a150565b60007f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b6000611011611103565b90506002816000015403611051576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002816000018190555050565b6000611068611103565b90506001816000018190555050565b60008061108261112b565b90508091505090565b611093611156565b61109c81611196565b50565b6110a7611156565b565b6110b1611156565b6110b961121c565b565b600033905090565b6110cb61041e565b611101576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00905090565b60007ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0060001b905090565b61115e61123d565b611194576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61119e611156565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036112105760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161120791906114d3565b60405180910390fd5b61121981610e96565b50565b611224611156565b600061122e611103565b90506001816000018190555050565b6000611247611077565b60000160089054906101000a900460ff16905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6112c48261127b565b810181811067ffffffffffffffff821117156112e3576112e261128c565b5b80604052505050565b60006112f661125d565b905061130282826112bb565b919050565b600067ffffffffffffffff8211156113225761132161128c565b5b61132b8261127b565b9050602081019050919050565b82818337600083830152505050565b600061135a61135584611307565b6112ec565b90508281526020810184848401111561137657611375611276565b5b611381848285611338565b509392505050565b600082601f83011261139e5761139d611271565b5b81356113ae848260208601611347565b91505092915050565b6000819050919050565b6113ca816113b7565b81146113d557600080fd5b50565b6000813590506113e7816113c1565b92915050565b60008060006060848603121561140657611405611267565b5b600084013567ffffffffffffffff8111156114245761142361126c565b5b61143086828701611389565b9350506020611441868287016113d8565b9250506040611452868287016113d8565b9150509250925092565b60008115159050919050565b6114718161145c565b82525050565b600060208201905061148c6000830184611468565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114bd82611492565b9050919050565b6114cd816114b2565b82525050565b60006020820190506114e860008301846114c4565b92915050565b60006020828403121561150457611503611267565b5b600061155b848285016113d8565b91505092915050565b6000602082019050818103600083015261157a81866115aa565b905061158960208301856115e3565b6115966040830184611468565b949350505050565b611639816114b2565b811461164457600080fd5b50565b60008135905061165681611630565b92915050565b60006020828403121561167257611671611267565b5b600061168084828501611647565b91505092915050565b600060608201905081810360008301526116a381866115aa565b90506116b260208301856115e3565b6116bf6040830184611468565b949350505050565b6000819050919050565b60006116ec6116e76116e284611492565b6116c7565b611492565b9050919050565b60006116fe826116d1565b9050919050565b6000611710826116f3565b9050919050565b61172081611705565b82525050565b600060208201905061173b6000830184611717565b92915050565b7f496e76616c6964206d6f6e73746572206e616d65000000000000000000000000600082015250565b600061177760148361156f565b915061178282611741565b602082019050919050565b600060208201905081810360008301526117a68161176a565b9050919050565b7f496e76616c696420485000000000000000000000000000000000000000000000600082015250565b60006117e3600a8361156f565b91506117ee826117ad565b602082019050919050565b60006020820190508181036000830152611812816117d6565b9050919050565b7f496e76616c696420726577617264000000000000000000000000000000000000600082015250565b600061184f600e8361156f565b915061185a82611819565b602082019050919050565b6000602082019050818103600083015261187e81611842565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806118cc57607f821691505b6020821081036118df576118de611885565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119477fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261190a565b611951868361190a565b95508019841693508086168417925050509392505050565b600061198461197f61197a846113b7565b6116c7565b6113b7565b9050919050565b6000819050919050565b61199e83611969565b6119b26119aa8261198b565b848454611917565b825550505050565b600090565b6119c76119ba565b6119d2818484611995565b505050565b5b818110156119f6576119eb6000826119bf565b6001810190506119d8565b5050565b601f821115611a3b57611a0c816118e5565b611a15846118fa565b81016020851015611a24578190505b611a38611a30856118fa56b8301826119d7565b50505b505050565b600082821c905092915050565b6000611a5e60001984600802611a40565b1980831691505092915050565b6000611a778383611a4d565b9150826002028217905092915050565b611a9082611564565b67ffffffffffffffff811115611aa957611aa861128c565b5b611ab382546118b4565b611abe8282856119fa565b600060209050601f831160018114611af15760008415611adf578287015190505b611ae98582611a6b565b865550611b51565b601f198416611aff866118e5565b60005b82811015611b2757848901518255600182019150602085019450602081019050611b02565b86831015611b445784890151611b40601f891682611a4d565b8355505b60016002880201885550505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b93826113b7565b9150611b9e836113b7565b9250828203905081811115611bb657611bb5611b59565b5b92915050565b7f416c726561647920726567697374657265640000000000000000000000000000600082015250565b6000611bf260128361156f565b9150611bfd82611bbc565b602082019050919050565b60006020820190508181036000830152611c2181611be5565b9050919050565b7f496e76616c6964206e616d650000000000000000000000000000000000000000600082015250565b6000611c5e600c8361156f565b9150611c6982611c28565b602082019050919050565b60006020820190508181036000830152611c8d81611c51565b9050919050565b60006020820190508181036000830152611cae81846115aa565b905092915050565b7f506c61796572206e6f7420726567697374657265640000000000000000000000600082015250565b6000611cec60158361156f565b9150611cf782611cb6565b602082019050919050565b60006020820190508181036000830152611d1b81611cdf565b9050919050565b7f496e76616c6964206d6f6e737465720000000000000000000000000000000000600082015250565b6000611d58600f8361156f565b9150611d6382611d22565b602082019050919050565b60006020820190508181036000830152611d8781611d4b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604082019050611dd260008301856114c4565b611ddf60208301846115e3565b9392505050565b6000602082019050611dfb60008301846115e3565b92915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000611e3a611e35611e3084611e01565b6116c7565b611e0b565b9050919050565b611e4a81611e1f565b82525050565b6000602082019050611e656000830184611e41565b9291505056fea264697066735822122028397716324f4124547133e094da1728c2b03f80eb51e0305a8accada3c996c564736f6c63430008140033",
}

// MonsterGameABI is the input ABI used to generate the binding from.
// Deprecated: Use MonsterGameMetaData.ABI instead.
var MonsterGameABI = MonsterGameMetaData.ABI

// MonsterGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MonsterGameMetaData.Bin instead.
var MonsterGameBin = MonsterGameMetaData.Bin

// DeployMonsterGame deploys a new Ethereum contract, binding an instance of MonsterGame to it.
func DeployMonsterGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MonsterGame, error) {
	parsed, err := MonsterGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MonsterGameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MonsterGame{MonsterGameCaller: MonsterGameCaller{contract: contract}, MonsterGameTransactor: MonsterGameTransactor{contract: contract}, MonsterGameFilterer: MonsterGameFilterer{contract: contract}}, nil
}

// MonsterGame is an auto generated Go binding around an Ethereum contract.
type MonsterGame struct {
	MonsterGameCaller     // Read-only binding to the contract
	MonsterGameTransactor // Write-only binding to the contract
	MonsterGameFilterer   // Log filterer for contract events
}

// MonsterGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type MonsterGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MonsterGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MonsterGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MonsterGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MonsterGameSession struct {
	Contract     *MonsterGame      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MonsterGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MonsterGameCallerSession struct {
	Contract *MonsterGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MonsterGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MonsterGameTransactorSession struct {
	Contract     *MonsterGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MonsterGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type MonsterGameRaw struct {
	Contract *MonsterGame // Generic contract binding to access the raw methods on
}

// MonsterGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MonsterGameCallerRaw struct {
	Contract *MonsterGameCaller // Generic read-only contract binding to access the raw methods on
}

// MonsterGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MonsterGameTransactorRaw struct {
	Contract *MonsterGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMonsterGame creates a new instance of MonsterGame, bound to a specific deployed contract.
func NewMonsterGame(address common.Address, backend bind.ContractBackend) (*MonsterGame, error) {
	contract, err := bindMonsterGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MonsterGame{MonsterGameCaller: MonsterGameCaller{contract: contract}, MonsterGameTransactor: MonsterGameTransactor{contract: contract}, MonsterGameFilterer: MonsterGameFilterer{contract: contract}}, nil
}

// NewMonsterGameCaller creates a new read-only instance of MonsterGame, bound to a specific deployed contract.
func NewMonsterGameCaller(address common.Address, caller bind.ContractCaller) (*MonsterGameCaller, error) {
	contract, err := bindMonsterGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterGameCaller{contract: contract}, nil
}

// NewMonsterGameTransactor creates a new write-only instance of MonsterGame, bound to a specific deployed contract.
func NewMonsterGameTransactor(address common.Address, transactor bind.ContractTransactor) (*MonsterGameTransactor, error) {
	contract, err := bindMonsterGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MonsterGameTransactor{contract: contract}, nil
}

// NewMonsterGameFilterer creates a new log filterer instance of MonsterGame, bound to a specific deployed contract.
func NewMonsterGameFilterer(address common.Address, filterer bind.ContractFilterer) (*MonsterGameFilterer, error) {
	contract, err := bindMonsterGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MonsterGameFilterer{contract: contract}, nil
}

// bindMonsterGame binds a generic wrapper to an already deployed contract.
func bindMonsterGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MonsterGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterGame *MonsterGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MonsterGame.Contract.MonsterGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterGame *MonsterGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGame.Contract.MonsterGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterGame *MonsterGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterGame.Contract.MonsterGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MonsterGame *MonsterGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MonsterGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MonsterGame *MonsterGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MonsterGame *MonsterGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MonsterGame.Contract.contract.Transact(opts, method, params...)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_MonsterGame *MonsterGameCaller) Monsters(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	var out []interface{}
	err := _MonsterGame.contract.Call(opts, &out, "monsters", arg0)

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
func (_MonsterGame *MonsterGameSession) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _MonsterGame.Contract.Monsters(&_MonsterGame.CallOpts, arg0)
}

// Monsters is a free data retrieval call binding the contract method 0xc0576b73.
//
// Solidity: function monsters(uint256 ) view returns(string name, uint256 hp, uint256 reward)
func (_MonsterGame *MonsterGameCallerSession) Monsters(arg0 *big.Int) (struct {
	Name   string
	Hp     *big.Int
	Reward *big.Int
}, error) {
	return _MonsterGame.Contract.Monsters(&_MonsterGame.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGame *MonsterGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MonsterGame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGame *MonsterGameSession) Owner() (common.Address, error) {
	return _MonsterGame.Contract.Owner(&_MonsterGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MonsterGame *MonsterGameCallerSession) Owner() (common.Address, error) {
	return _MonsterGame.Contract.Owner(&_MonsterGame.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGame *MonsterGameCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MonsterGame.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGame *MonsterGameSession) Paused() (bool, error) {
	return _MonsterGame.Contract.Paused(&_MonsterGame.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MonsterGame *MonsterGameCallerSession) Paused() (bool, error) {
	return _MonsterGame.Contract.Paused(&_MonsterGame.CallOpts)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_MonsterGame *MonsterGameCaller) Players(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	var out []interface{}
	err := _MonsterGame.contract.Call(opts, &out, "players", arg0)

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
func (_MonsterGame *MonsterGameSession) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _MonsterGame.Contract.Players(&_MonsterGame.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players(address ) view returns(string name, uint256 level, bool registered)
func (_MonsterGame *MonsterGameCallerSession) Players(arg0 common.Address) (struct {
	Name       string
	Level      *big.Int
	Registered bool
}, error) {
	return _MonsterGame.Contract.Players(&_MonsterGame.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGame *MonsterGameCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MonsterGame.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGame *MonsterGameSession) RewardToken() (common.Address, error) {
	return _MonsterGame.Contract.RewardToken(&_MonsterGame.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_MonsterGame *MonsterGameCallerSession) RewardToken() (common.Address, error) {
	return _MonsterGame.Contract.RewardToken(&_MonsterGame.CallOpts)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGame *MonsterGameTransactor) AddMonster(opts *bind.TransactOpts, name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "addMonster", name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGame *MonsterGameSession) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGame.Contract.AddMonster(&_MonsterGame.TransactOpts, name, hp, reward)
}

// AddMonster is a paid mutator transaction binding the contract method 0x18f237bc.
//
// Solidity: function addMonster(string name, uint256 hp, uint256 reward) returns()
func (_MonsterGame *MonsterGameTransactorSession) AddMonster(name string, hp *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _MonsterGame.Contract.AddMonster(&_MonsterGame.TransactOpts, name, hp, reward)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGame *MonsterGameTransactor) HuntMonster(opts *bind.TransactOpts, monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "huntMonster", monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGame *MonsterGameSession) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGame.Contract.HuntMonster(&_MonsterGame.TransactOpts, monsterId)
}

// HuntMonster is a paid mutator transaction binding the contract method 0xb906b300.
//
// Solidity: function huntMonster(uint256 monsterId) returns()
func (_MonsterGame *MonsterGameTransactorSession) HuntMonster(monsterId *big.Int) (*types.Transaction, error) {
	return _MonsterGame.Contract.HuntMonster(&_MonsterGame.TransactOpts, monsterId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGame *MonsterGameTransactor) Initialize(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "initialize", tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGame *MonsterGameSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGame.Contract.Initialize(&_MonsterGame.TransactOpts, tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_MonsterGame *MonsterGameTransactorSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _MonsterGame.Contract.Initialize(&_MonsterGame.TransactOpts, tokenAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGame *MonsterGameTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGame *MonsterGameSession) Pause() (*types.Transaction, error) {
	return _MonsterGame.Contract.Pause(&_MonsterGame.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MonsterGame *MonsterGameTransactorSession) Pause() (*types.Transaction, error) {
	return _MonsterGame.Contract.Pause(&_MonsterGame.TransactOpts)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGame *MonsterGameTransactor) RegisterPlayer(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "registerPlayer", name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGame *MonsterGameSession) RegisterPlayer(name string) (*types.Transaction, error) {
	return _MonsterGame.Contract.RegisterPlayer(&_MonsterGame.TransactOpts, name)
}

// RegisterPlayer is a paid mutator transaction binding the contract method 0xb3446f85.
//
// Solidity: function registerPlayer(string name) returns()
func (_MonsterGame *MonsterGameTransactorSession) RegisterPlayer(name string) (*types.Transaction, error) {
	return _MonsterGame.Contract.RegisterPlayer(&_MonsterGame.TransactOpts, name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGame *MonsterGameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGame *MonsterGameSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonsterGame.Contract.RenounceOwnership(&_MonsterGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MonsterGame *MonsterGameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MonsterGame.Contract.RenounceOwnership(&_MonsterGame.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGame *MonsterGameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGame *MonsterGameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGame.Contract.TransferOwnership(&_MonsterGame.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MonsterGame *MonsterGameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MonsterGame.Contract.TransferOwnership(&_MonsterGame.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGame *MonsterGameTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MonsterGame.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGame *MonsterGameSession) Unpause() (*types.Transaction, error) {
	return _MonsterGame.Contract.Unpause(&_MonsterGame.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MonsterGame *MonsterGameTransactorSession) Unpause() (*types.Transaction, error) {
	return _MonsterGame.Contract.Unpause(&_MonsterGame.TransactOpts)
}

// MonsterGameInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MonsterGame contract.
type MonsterGameInitializedIterator struct {
	Event *MonsterGameInitialized // Event containing the contract specifics and raw log

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
func (it *MonsterGameInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameInitialized)
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
		it.Event = new(MonsterGameInitialized)
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
func (it *MonsterGameInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameInitialized represents a Initialized event raised by the MonsterGame contract.
type MonsterGameInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MonsterGame *MonsterGameFilterer) FilterInitialized(opts *bind.FilterOpts) (*MonsterGameInitializedIterator, error) {

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MonsterGameInitializedIterator{contract: _MonsterGame.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MonsterGame *MonsterGameFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MonsterGameInitialized) (event.Subscription, error) {

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameInitialized)
				if err := _MonsterGame.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParseInitialized(log types.Log) (*MonsterGameInitialized, error) {
	event := new(MonsterGameInitialized)
	if err := _MonsterGame.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameMonsterAddedIterator is returned from FilterMonsterAdded and is used to iterate over the raw logs and unpacked data for MonsterAdded events raised by the MonsterGame contract.
type MonsterGameMonsterAddedIterator struct {
	Event *MonsterGameMonsterAdded // Event containing the contract specifics and raw log

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
func (it *MonsterGameMonsterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameMonsterAdded)
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
		it.Event = new(MonsterGameMonsterAdded)
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
func (it *MonsterGameMonsterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameMonsterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameMonsterAdded represents a MonsterAdded event raised by the MonsterGame contract.
type MonsterGameMonsterAdded struct {
	MonsterId *big.Int
	Name      string
	Hp        *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterAdded is a free log retrieval operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_MonsterGame *MonsterGameFilterer) FilterMonsterAdded(opts *bind.FilterOpts, monsterId []*big.Int) (*MonsterGameMonsterAddedIterator, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameMonsterAddedIterator{contract: _MonsterGame.contract, event: "MonsterAdded", logs: logs, sub: sub}, nil
}

// WatchMonsterAdded is a free log subscription operation binding the contract event 0xfa2e8024f52f684312a1697efe6ac91d76d4c26b6ace005acb24825d30e23b0b.
//
// Solidity: event MonsterAdded(uint256 indexed monsterId, string name, uint256 hp, uint256 reward)
func (_MonsterGame *MonsterGameFilterer) WatchMonsterAdded(opts *bind.WatchOpts, sink chan<- *MonsterGameMonsterAdded, monsterId []*big.Int) (event.Subscription, error) {

	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "MonsterAdded", monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameMonsterAdded)
				if err := _MonsterGame.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParseMonsterAdded(log types.Log) (*MonsterGameMonsterAdded, error) {
	event := new(MonsterGameMonsterAdded)
	if err := _MonsterGame.contract.UnpackLog(event, "MonsterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameMonsterHuntedIterator is returned from FilterMonsterHunted and is used to iterate over the raw logs and unpacked data for MonsterHunted events raised by the MonsterGame contract.
type MonsterGameMonsterHuntedIterator struct {
	Event *MonsterGameMonsterHunted // Event containing the contract specifics and raw log

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
func (it *MonsterGameMonsterHuntedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameMonsterHunted)
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
		it.Event = new(MonsterGameMonsterHunted)
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
func (it *MonsterGameMonsterHuntedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameMonsterHuntedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameMonsterHunted represents a MonsterHunted event raised by the MonsterGame contract.
type MonsterGameMonsterHunted struct {
	Player    common.Address
	MonsterId *big.Int
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMonsterHunted is a free log retrieval operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_MonsterGame *MonsterGameFilterer) FilterMonsterHunted(opts *bind.FilterOpts, player []common.Address, monsterId []*big.Int) (*MonsterGameMonsterHuntedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameMonsterHuntedIterator{contract: _MonsterGame.contract, event: "MonsterHunted", logs: logs, sub: sub}, nil
}

// WatchMonsterHunted is a free log subscription operation binding the contract event 0x3e01e3eb71f0824366d4e3e5c4d4614c9f035a178a50c3f68243d3357c265460.
//
// Solidity: event MonsterHunted(address indexed player, uint256 indexed monsterId, uint256 reward)
func (_MonsterGame *MonsterGameFilterer) WatchMonsterHunted(opts *bind.WatchOpts, sink chan<- *MonsterGameMonsterHunted, player []common.Address, monsterId []*big.Int) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var monsterIdRule []interface{}
	for _, monsterIdItem := range monsterId {
		monsterIdRule = append(monsterIdRule, monsterIdItem)
	}

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "MonsterHunted", playerRule, monsterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameMonsterHunted)
				if err := _MonsterGame.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParseMonsterHunted(log types.Log) (*MonsterGameMonsterHunted, error) {
	event := new(MonsterGameMonsterHunted)
	if err := _MonsterGame.contract.UnpackLog(event, "MonsterHunted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MonsterGame contract.
type MonsterGameOwnershipTransferredIterator struct {
	Event *MonsterGameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MonsterGameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameOwnershipTransferred)
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
		it.Event = new(MonsterGameOwnershipTransferred)
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
func (it *MonsterGameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameOwnershipTransferred represents a OwnershipTransferred event raised by the MonsterGame contract.
type MonsterGameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MonsterGame *MonsterGameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MonsterGameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGameOwnershipTransferredIterator{contract: _MonsterGame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MonsterGame *MonsterGameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MonsterGameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameOwnershipTransferred)
				if err := _MonsterGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParseOwnershipTransferred(log types.Log) (*MonsterGameOwnershipTransferred, error) {
	event := new(MonsterGameOwnershipTransferred)
	if err := _MonsterGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGamePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MonsterGame contract.
type MonsterGamePausedIterator struct {
	Event *MonsterGamePaused // Event containing the contract specifics and raw log

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
func (it *MonsterGamePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGamePaused)
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
		it.Event = new(MonsterGamePaused)
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
func (it *MonsterGamePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGamePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGamePaused represents a Paused event raised by the MonsterGame contract.
type MonsterGamePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MonsterGame *MonsterGameFilterer) FilterPaused(opts *bind.FilterOpts) (*MonsterGamePausedIterator, error) {

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MonsterGamePausedIterator{contract: _MonsterGame.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MonsterGame *MonsterGameFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MonsterGamePaused) (event.Subscription, error) {

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGamePaused)
				if err := _MonsterGame.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParsePaused(log types.Log) (*MonsterGamePaused, error) {
	event := new(MonsterGamePaused)
	if err := _MonsterGame.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGamePlayerRegisteredIterator is returned from FilterPlayerRegistered and is used to iterate over the raw logs and unpacked data for PlayerRegistered events raised by the MonsterGame contract.
type MonsterGamePlayerRegisteredIterator struct {
	Event *MonsterGamePlayerRegistered // Event containing the contract specifics and raw log

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
func (it *MonsterGamePlayerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGamePlayerRegistered)
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
		it.Event = new(MonsterGamePlayerRegistered)
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
func (it *MonsterGamePlayerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGamePlayerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGamePlayerRegistered represents a PlayerRegistered event raised by the MonsterGame contract.
type MonsterGamePlayerRegistered struct {
	Player common.Address
	Name   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlayerRegistered is a free log retrieval operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_MonsterGame *MonsterGameFilterer) FilterPlayerRegistered(opts *bind.FilterOpts, player []common.Address) (*MonsterGamePlayerRegisteredIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return &MonsterGamePlayerRegisteredIterator{contract: _MonsterGame.contract, event: "PlayerRegistered", logs: logs, sub: sub}, nil
}

// WatchPlayerRegistered is a free log subscription operation binding the contract event 0xe04d23b73ae3d3b27074da9719b5a4cd521395aaf86bfe7e374933ecb567731a.
//
// Solidity: event PlayerRegistered(address indexed player, string name)
func (_MonsterGame *MonsterGameFilterer) WatchPlayerRegistered(opts *bind.WatchOpts, sink chan<- *MonsterGamePlayerRegistered, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "PlayerRegistered", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGamePlayerRegistered)
				if err := _MonsterGame.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParsePlayerRegistered(log types.Log) (*MonsterGamePlayerRegistered, error) {
	event := new(MonsterGamePlayerRegistered)
	if err := _MonsterGame.contract.UnpackLog(event, "PlayerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MonsterGameUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MonsterGame contract.
type MonsterGameUnpausedIterator struct {
	Event *MonsterGameUnpaused // Event containing the contract specifics and raw log

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
func (it *MonsterGameUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MonsterGameUnpaused)
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
		it.Event = new(MonsterGameUnpaused)
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
func (it *MonsterGameUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MonsterGameUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MonsterGameUnpaused represents a Unpaused event raised by the MonsterGame contract.
type MonsterGameUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MonsterGame *MonsterGameFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MonsterGameUnpausedIterator, error) {

	logs, sub, err := _MonsterGame.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MonsterGameUnpausedIterator{contract: _MonsterGame.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MonsterGame *MonsterGameFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MonsterGameUnpaused) (event.Subscription, error) {

	logs, sub, err := _MonsterGame.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MonsterGameUnpaused)
				if err := _MonsterGame.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MonsterGame *MonsterGameFilterer) ParseUnpaused(log types.Log) (*MonsterGameUnpaused, error) {
	event := new(MonsterGameUnpaused)
	if err := _MonsterGame.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
