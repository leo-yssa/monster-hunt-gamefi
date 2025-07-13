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

// GovernorContractMetaData contains all meta data concerning the GovernorContract contract.
var GovernorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_votingPowerSource\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorUnableToCancel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"getProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newVotingDelay\",\"type\":\"uint48\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newVotingPeriod\",\"type\":\"uint32\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162006ad038038062006ad0833981810160405281019062000038919062000521565b806001600560006040518060400160405280601081526020017f476f7665726e6f72436f6e7472616374000000000000000000000000000000008152508062000086620001bc60201b60201c565b6200009c600083620001f960201b90919060201c565b6101208181525050620000ba600182620001f960201b90919060201c565b6101408181525050818051906020012060e08181525050808051906020012061010081815250504660a08181525050620000f96200025160201b60201c565b608081815250503073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505050508060039081620001479190620007cd565b50506200015a83620002ae60201b60201c565b6200016b826200032560201b60201c565b6200017c81620003e560201b60201c565b5050508073ffffffffffffffffffffffffffffffffffffffff166101608173ffffffffffffffffffffffffffffffffffffffff1681525050505062000c1e565b60606040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250905090565b60006020835110156200021f5762000217836200042c60201b60201c565b90506200024b565b8262000231836200049960201b60201c565b6000019081620002429190620007cd565b5060ff60001b90505b92915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60e05161010051463060405160200162000293959493929190620008f1565b60405160208183030381529060405280519060200120905090565b7fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93600860009054906101000a900465ffffffffffff1682604051620002f592919062000999565b60405180910390a180600860006101000a81548165ffffffffffff021916908365ffffffffffff16021790555050565b60008163ffffffff1603620003745760006040517ff1cfbf050000000000000000000000000000000000000000000000000000000081526004016200036b919062000a09565b60405180910390fd5b7f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828600860069054906101000a900463ffffffff1682604051620003b992919062000a6f565b60405180910390a180600860066101000a81548163ffffffff021916908363ffffffff16021790555050565b7fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461600754826040516200041a92919062000a9c565b60405180910390a18060078190555050565b600080829050601f815111156200047c57826040517f305a27a900000000000000000000000000000000000000000000000000000000815260040162000473919062000b58565b60405180910390fd5b8051816200048a9062000bae565b60001c1760001b915050919050565b6000819050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620004d582620004a8565b9050919050565b6000620004e982620004c8565b9050919050565b620004fb81620004dc565b81146200050757600080fd5b50565b6000815190506200051b81620004f0565b92915050565b6000602082840312156200053a5762000539620004a3565b5b60006200054a848285016200050a565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620005d557607f821691505b602082108103620005eb57620005ea6200058d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620006557fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000616565b62000661868362000616565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620006ae620006a8620006a28462000679565b62000683565b62000679565b9050919050565b6000819050919050565b620006ca836200068d565b620006e2620006d982620006b5565b84845462000623565b825550505050565b600090565b620006f9620006ea565b62000706818484620006bf565b505050565b5b818110156200072e5762000722600082620006ef565b6001810190506200070c565b5050565b601f8211156200077d576200074781620005f1565b620007528462000606565b8101602085101562000762578190505b6200077a620007718562000606565b8301826200070b565b50505b505050565b600082821c905092915050565b6000620007a26000198460080262000782565b1980831691505092915050565b6000620007bd83836200078f565b9150826002028217905092915050565b620007d88262000553565b67ffffffffffffffff811115620007f457620007f36200055e565b5b620008008254620005bc565b6200080d82828562000732565b600060209050601f83116001811462000845576000841562000830578287015190505b6200083c8582620007af565b865550620008ac565b601f1984166200085586620005f1565b60005b828110156200087f5784890151825560018201915060208501945060208101905062000858565b868310156200089f57848901516200089b601f8916826200078f565b8355505b6001600288020188555050505b505050505050565b6000819050919050565b620008c981620008b4565b82525050565b620008da8162000679565b82525050565b620008eb81620004c8565b82525050565b600060a082019050620009086000830188620008be565b620009176020830187620008be565b620009266040830186620008be565b620009356060830185620008cf565b620009446080830184620008e0565b9695505050505050565b600065ffffffffffff82169050919050565b6000620009816200097b62000975846200094e565b62000683565b62000679565b9050919050565b620009938162000960565b82525050565b6000604082019050620009b0600083018562000988565b620009bf602083018462000988565b9392505050565b6000819050919050565b6000620009f1620009eb620009e584620009c6565b62000683565b62000679565b9050919050565b62000a0381620009d0565b82525050565b600060208201905062000a206000830184620009f8565b92915050565b600063ffffffff82169050919050565b600062000a5762000a5162000a4b8462000a26565b62000683565b62000679565b9050919050565b62000a698162000a36565b82525050565b600060408201905062000a86600083018562000a5e565b62000a95602083018462000a5e565b9392505050565b600060408201905062000ab36000830185620008cf565b62000ac26020830184620008cf565b9392505050565b600082825260208201905092915050565b60005b8381101562000afa57808201518184015260208101905062000add565b60008484015250505050565b6000601f19601f8301169050919050565b600062000b248262000553565b62000b30818562000ac9565b935062000b4281856020860162000ada565b62000b4d8162000b06565b840191505092915050565b6000602082019050818103600083015262000b74818462000b17565b905092915050565b600081519050919050565b6000819050602082019050919050565b600062000ba58251620008b4565b80915050919050565b600062000bbb8262000b7c565b8262000bc78462000b87565b905062000bd48162000b97565b9250602082101562000c175762000c127fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8360200360080262000616565b831692505b5050919050565b60805160a05160c05160e05161010051610120516101405161016051615e4c62000c846000396000611e6d01526000612c0b01526000612bd00152600061370b015260006136ea01526000612e3601526000612e8c01526000612eb50152615e4c6000f3fe6080604052600436106102605760003560e01c80637d5e81e211610144578063c01f9e37116100b6578063e540d01d1161007a578063e540d01d14610aa9578063eb9019d414610ad2578063ece40cc114610b0f578063f23a6e6114610b38578063f8ce560a14610b75578063fc0c546a14610bb2576102d3565b8063c01f9e37146109bd578063c28bc2fa146109fa578063c59057e414610a16578063dd4e2ba514610a53578063deaaa7cc14610a7e576102d3565b80639a802a6d116101085780639a802a6d14610861578063a8f8a6681461089e578063a9a95294146108db578063ab58fb8e14610918578063b58131b014610955578063bc197c8114610980576102d3565b80637d5e81e21461074e5780637ecebe001461078b57806384b0196e146107c85780638ff262e3146107f957806391ddadf414610836576102d3565b80633e4f49e6116101dd57806354fd4d50116101a157806354fd4d501461060657806356781388146106315780635b8d0e0d1461066e5780635f398a14146106ab57806379051887146106e85780637b3c71d314610711576102d3565b80633e4f49e6146104e55780634385963214610522578063452115d61461055f5780634bf5d7e91461059c578063544ffc9c146105c7576102d3565b8063160cbed711610224578063160cbed7146103e55780632656227d146104225780632d63f693146104525780632fe3e2611461048f5780633932abb1146104ba576102d3565b806301ffc9a7146102d857806302a251a31461031557806306fdde0314610340578063143489d01461036b578063150b7a02146103a8576102d3565b366102d3573073ffffffffffffffffffffffffffffffffffffffff16610284610bdd565b73ffffffffffffffffffffffffffffffffffffffff16146102d1576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156102e457600080fd5b506102ff60048036038101906102fa9190613b65565b610be5565b60405161030c9190613bad565b60405180910390f35b34801561032157600080fd5b5061032a610d38565b6040516103379190613be1565b60405180910390f35b34801561034c57600080fd5b50610355610d47565b6040516103629190613c8c565b60405180910390f35b34801561037757600080fd5b50610392600480360381019061038d9190613cda565b610dd9565b60405161039f9190613d48565b60405180910390f35b3480156103b457600080fd5b506103cf60048036038101906103ca9190613ec4565b610e19565b6040516103dc9190613f56565b60405180910390f35b3480156103f157600080fd5b5061040c60048036038101906104079190614213565b610e99565b6040516104199190613be1565b60405180910390f35b61043c60048036038101906104379190614213565b610f95565b6040516104499190613be1565b60405180910390f35b34801561045e57600080fd5b5061047960048036038101906104749190613cda565b61117a565b6040516104869190613be1565b60405180910390f35b34801561049b57600080fd5b506104a46111b4565b6040516104b191906142dd565b60405180910390f35b3480156104c657600080fd5b506104cf6111d8565b6040516104dc9190613be1565b60405180910390f35b3480156104f157600080fd5b5061050c60048036038101906105079190613cda565b6111e7565b604051610519919061436f565b60405180910390f35b34801561052e57600080fd5b506105496004803603810190610544919061438a565b61134b565b6040516105569190613bad565b60405180910390f35b34801561056b57600080fd5b5061058660048036038101906105819190614213565b6113b6565b6040516105939190613be1565b60405180910390f35b3480156105a857600080fd5b506105b1611438565b6040516105be9190613c8c565b60405180910390f35b3480156105d357600080fd5b506105ee60048036038101906105e99190613cda565b6114f7565b6040516105fd939291906143ca565b60405180910390f35b34801561061257600080fd5b5061061b61152f565b6040516106289190613c8c565b60405180910390f35b34801561063d57600080fd5b506106586004803603810190610653919061443a565b61156c565b6040516106659190613be1565b60405180910390f35b34801561067a57600080fd5b50610695600480360381019061069091906144d5565b61159d565b6040516106a29190613be1565b60405180910390f35b3480156106b757600080fd5b506106d260048036038101906106cd91906145bc565b6116ce565b6040516106df9190613be1565b60405180910390f35b3480156106f457600080fd5b5061070f600480360381019061070a919061469e565b611738565b005b34801561071d57600080fd5b50610738600480360381019061073391906146cb565b61174c565b6040516107459190613be1565b60405180910390f35b34801561075a57600080fd5b50610775600480360381019061077091906147e0565b6117b4565b6040516107829190613be1565b60405180910390f35b34801561079757600080fd5b506107b260048036038101906107ad91906148b7565b6118ae565b6040516107bf9190613be1565b60405180910390f35b3480156107d457600080fd5b506107dd6118f7565b6040516107f097969594939291906149dd565b60405180910390f35b34801561080557600080fd5b50610820600480360381019061081b9190614a61565b6119a1565b60405161082d9190613be1565b60405180910390f35b34801561084257600080fd5b5061084b611a77565b6040516108589190614af3565b60405180910390f35b34801561086d57600080fd5b5061088860048036038101906108839190614b0e565b611b03565b6040516108959190613be1565b60405180910390f35b3480156108aa57600080fd5b506108c560048036038101906108c09190614213565b611b19565b6040516108d29190613be1565b60405180910390f35b3480156108e757600080fd5b5061090260048036038101906108fd9190613cda565b611b31565b60405161090f9190613bad565b60405180910390f35b34801561092457600080fd5b5061093f600480360381019061093a9190613cda565b611b38565b60405161094c9190613be1565b60405180910390f35b34801561096157600080fd5b5061096a611b72565b6040516109779190613be1565b60405180910390f35b34801561098c57600080fd5b506109a760048036038101906109a29190614b7d565b611b81565b6040516109b49190613f56565b60405180910390f35b3480156109c957600080fd5b506109e460048036038101906109df9190613cda565b611c02565b6040516109f19190613be1565b60405180910390f35b610a146004803603810190610a0f9190614ca2565b611c73565b005b348015610a2257600080fd5b50610a3d6004803603810190610a389190614213565b611d00565b604051610a4a9190613be1565b60405180910390f35b348015610a5f57600080fd5b50610a68611d3c565b604051610a759190613c8c565b60405180910390f35b348015610a8a57600080fd5b50610a93611d79565b604051610aa091906142dd565b60405180910390f35b348015610ab557600080fd5b50610ad06004803603810190610acb9190614d52565b611d9d565b005b348015610ade57600080fd5b50610af96004803603810190610af49190614d7f565b611db1565b604051610b069190613be1565b60405180910390f35b348015610b1b57600080fd5b50610b366004803603810190610b319190613cda565b611dcd565b005b348015610b4457600080fd5b50610b5f6004803603810190610b5a9190614dbf565b611de1565b604051610b6c9190613f56565b60405180910390f35b348015610b8157600080fd5b50610b9c6004803603810190610b979190613cda565b611e62565b604051610ba99190613be1565b60405180910390f35b348015610bbe57600080fd5b50610bc7611e69565b604051610bd49190614eb5565b60405180910390f35b600030905090565b60007fcdbdfcee000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610cb9575063a8f8a66860e01b7fcdbdfcee00000000000000000000000000000000000000000000000000000000187bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610d2157507f4e2312e0000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610d315750610d3082611e91565b5b9050919050565b6000610d42611efb565b905090565b606060038054610d5690614eff565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8290614eff565b8015610dcf5780601f10610da457610100808354040283529160200191610dcf565b820191906000526020600020905b815481529060010190602001808311610db257829003601f168201915b5050505050905090565b60006004600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60003073ffffffffffffffffffffffffffffffffffffffff16610e3a610bdd565b73ffffffffffffffffffffffffffffffffffffffff1614610e87576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63150b7a0260e01b9050949350505050565b600080610ea886868686611b19565b9050610ebd81610eb86004611f1b565b611f41565b506000610ecd8288888888611fae565b905060008165ffffffffffff1614610f5657806004600084815260200190815260200160002060010160006101000a81548165ffffffffffff021916908365ffffffffffff1602179055507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda28928282604051610f49929190614f61565b60405180910390a1610f88565b6040517f90884a4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8192505050949350505050565b600080610fa486868686611b19565b9050610fc481610fb46005611f1b565b610fbe6004611f1b565b17611f41565b50600160046000838152602001908152602001600020600001601e6101000a81548160ff0219169083151502179055503073ffffffffffffffffffffffffffffffffffffffff16611013610bdd565b73ffffffffffffffffffffffffffffffffffffffff16146110d05760005b86518110156110ce573073ffffffffffffffffffffffffffffffffffffffff1687828151811061106457611063614f8a565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16036110bd576110bc85828151811061109d5761109c614f8a565b5b6020026020010151805190602001206005611fb990919063ffffffff16565b5b806110c790614fe8565b9050611031565b505b6110dd81878787876120b8565b3073ffffffffffffffffffffffffffffffffffffffff166110fc610bdd565b73ffffffffffffffffffffffffffffffffffffffff1614158015611127575061112560056121a8565b155b15611137576111366005612218565b5b7f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f816040516111669190613be1565b60405180910390a180915050949350505050565b60006004600083815260200190815260200160002060000160149054906101000a900465ffffffffffff1665ffffffffffff169050919050565b7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b60006111e2612293565b905090565b600080600460008481526020019081526020016000209050600081600001601e9054906101000a900460ff169050600082600001601f9054906101000a900460ff169050811561123d5760079350505050611346565b801561124f5760029350505050611346565b600061125a8661117a565b9050600081036112a157856040517f6ad060750000000000000000000000000000000000000000000000000000000081526004016112989190613be1565b60405180910390fd5b60006112ab611a77565b65ffffffffffff1690508082106112ca57600095505050505050611346565b60006112d588611c02565b90508181106112ed5760019650505050505050611346565b6112f6886122b7565b15806113085750611306886122fe565b155b1561131c5760039650505050505050611346565b600061132789611b38565b0361133b5760049650505050505050611346565b600596505050505050505b919050565b60006009600084815260200190815260200160002060030160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000806113c586868686611b19565b905060006113d1612329565b90506113dd8282612331565b6114205781816040517f8fe5d8a9000000000000000000000000000000000000000000000000000000008152600401611417929190615030565b60405180910390fd5b61142c878787876123a9565b92505050949350505050565b6060611442611e69565b73ffffffffffffffffffffffffffffffffffffffff16634bf5d7e96040518163ffffffff1660e01b8152600401600060405180830381865afa9250505080156114ae57506040513d6000823e3d601f19601f820116820180604052508101906114ab91906150c9565b60015b6114ef576040518060400160405280601d81526020017f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c7400000081525090506114f4565b809150505b90565b600080600080600960008681526020019081526020016000209050806000015481600101548260020154935093509350509193909250565b60606040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250905090565b600080611577612329565b905061159484828560405180602001604052806000815250612490565b91505092915050565b60008061162b876116257f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c6115d48e6124b0565b8d8d6040516115e4929190615142565b60405180910390208c8051906020012060405160200161160a979695949392919061516a565b60405160208183030381529060405280519060200120612507565b85612521565b90508061166f57866040517f94ab6c070000000000000000000000000000000000000000000000000000000081526004016116669190613d48565b60405180910390fd5b6116c089888a89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050886125d0565b915050979650505050505050565b6000806116d9612329565b905061172c87828888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050876125d0565b91505095945050505050565b6117406126de565b611749816127d4565b50565b600080611757612329565b90506117a986828787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050612490565b915050949350505050565b6000806117bf612329565b90506117cb8184612849565b61180c57806040517fd9b395570000000000000000000000000000000000000000000000000000000081526004016118039190613d48565b60405180910390fd5b6000611816611b72565b90506000811115611895576000611848836001611831611a77565b61183b91906151d9565b65ffffffffffff16611db1565b905081811015611893578281836040517fc242ee1600000000000000000000000000000000000000000000000000000000815260040161188a93929190615213565b60405180910390fd5b505b6118a28787878786612932565b92505050949350505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006060806000806000606061190b612bc7565b611913612c02565b46306000801b600067ffffffffffffffff81111561193457611933613d99565b5b6040519080825280602002602001820160405280156119625781602001602082028036833780820191505090505b507f0f00000000000000000000000000000000000000000000000000000000000000959493929190965096509650965096509650965090919293949596565b600080611a0d84611a077ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78989896119d88b6124b0565b6040516020016119ec95949392919061524a565b60405160208183030381529060405280519060200120612507565b85612521565b905080611a5157836040517f94ab6c07000000000000000000000000000000000000000000000000000000008152600401611a489190613d48565b60405180910390fd5b611a6c86858760405180602001604052806000815250612490565b915050949350505050565b6000611a81611e69565b73ffffffffffffffffffffffffffffffffffffffff166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611ae857506040513d601f19601f82011682018060405250810190611ae591906152b2565b60015b611afb57611af4612c3d565b9050611b00565b809150505b90565b6000611b10848484612c4d565b90509392505050565b6000611b2785858585611d00565b9050949350505050565b6000919050565b60006004600083815260200190815260200160002060010160009054906101000a900465ffffffffffff1665ffffffffffff169050919050565b6000611b7c612cdb565b905090565b60003073ffffffffffffffffffffffffffffffffffffffff16611ba2610bdd565b73ffffffffffffffffffffffffffffffffffffffff1614611bef576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63bc197c8160e01b905095945050505050565b600060046000838152602001908152602001600020600001601a9054906101000a900463ffffffff1663ffffffff166004600084815260200190815260200160002060000160149054906101000a900465ffffffffffff16611c6491906152df565b65ffffffffffff169050919050565b611c7b6126de565b6000808573ffffffffffffffffffffffffffffffffffffffff16858585604051611ca6929190615142565b60006040518083038185875af1925050503d8060008114611ce3576040519150601f19603f3d011682016040523d82523d6000602084013e611ce8565b606091505b5091509150611cf78282612ce5565b50505050505050565b600084848484604051602001611d1994939291906154ee565b6040516020818303038152906040528051906020012060001c9050949350505050565b60606040518060400160405280602081526020017f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e815250905090565b7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b611da56126de565b611dae81612d09565b50565b6000611dc58383611dc0612dc4565b612c4d565b905092915050565b611dd56126de565b611dde81612ddb565b50565b60003073ffffffffffffffffffffffffffffffffffffffff16611e02610bdd565b73ffffffffffffffffffffffffffffffffffffffff1614611e4f576040517fe90a651e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63f23a6e6160e01b905095945050505050565b6000919050565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6000600860069054906101000a900463ffffffff1663ffffffff16905090565b6000816007811115611f3057611f2f6142f8565b5b60ff166001901b60001b9050919050565b600080611f4d846111e7565b90506000801b83611f5d83611f1b565b1603611fa4578381846040517f31b75e4d000000000000000000000000000000000000000000000000000000008152600401611f9b93929190615548565b60405180910390fd5b8091505092915050565b600095945050505050565b60008260000160109054906101000a90046fffffffffffffffffffffffffffffffff1690508260000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16600182016fffffffffffffffffffffffffffffffff1603612037576120366041612e20565b5b81836001016000836fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600181018360000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550505050565b60005b84518110156121a0576000808683815181106120da576120d9614f8a565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1686848151811061210b5761210a614f8a565b5b602002602001015186858151811061212657612125614f8a565b5b602002602001015160405161213b91906155b0565b60006040518083038185875af1925050503d8060008114612178576040519150601f19603f3d011682016040523d82523d6000602084013e61217d565b606091505b509150915061218c8282612ce5565b5050508061219990614fe8565b90506120bb565b505050505050565b60008160000160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff168260000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16149050919050565b60008160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060008160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b6000600860009054906101000a900465ffffffffffff1665ffffffffffff16905090565b600080600960008481526020019081526020016000209050806002015481600101546122e391906155c7565b6122f46122ef8561117a565b611e62565b1115915050919050565b6000806009600084815260200190815260200160002090508060000154816001015411915050919050565b600033905090565b6000806007811115612346576123456142f8565b5b61234f846111e7565b6007811115612361576123606142f8565b5b1480156123a1575061237283610dd9565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b905092915050565b6000806123b886868686611b19565b905061241d816123c86007611f1b565b6123d26006611f1b565b6123dc6002611f1b565b6001806007808111156123f2576123f16142f8565b5b6123fc91906155fb565b60026124089190615763565b61241291906157ae565b60001b181818611f41565b50600160046000838152602001908152602001600020600001601f6101000a81548160ff0219169083151502179055507f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c8160405161247c9190613be1565b60405180910390a180915050949350505050565b60006124a6858585856124a1612dc4565b6125d0565b9050949350505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050559050919050565b600061251a612514612e32565b83612ee9565b9050919050565b6000808473ffffffffffffffffffffffffffffffffffffffff163b036125bb5760008061254e8585612f2a565b509150915060006003811115612567576125666142f8565b5b81600381111561257a576125796142f8565b5b1480156125b257508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b925050506125c9565b6125c6848484612f86565b90505b9392505050565b60006125e5866125e06001611f1b565b611f41565b5060006125fb866125f58961117a565b85612c4d565b9050600061260c88888885886130aa565b90506000845103612670578673ffffffffffffffffffffffffffffffffffffffff167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda48988848960405161266394939291906157e2565b60405180910390a26126c7565b8673ffffffffffffffffffffffffffffffffffffffff167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516126be959493929190615878565b60405180910390a25b6126d0886132ad565b809250505095945050505050565b6126e6612329565b73ffffffffffffffffffffffffffffffffffffffff16612704610bdd565b73ffffffffffffffffffffffffffffffffffffffff161461276357612727612329565b6040517f47096e4700000000000000000000000000000000000000000000000000000000815260040161275a9190613d48565b60405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff16612782610bdd565b73ffffffffffffffffffffffffffffffffffffffff16146127d25760006127a76132b0565b6040516127b5929190615142565b604051809103902090505b806127cb60056132bd565b036127c057505b565b7fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93600860009054906101000a900465ffffffffffff16826040516128199291906158d9565b60405180910390a180600860006101000a81548165ffffffffffff021916908365ffffffffffff16021790555050565b60008082519050603481101561286357600191505061292c565b600061287284603484036133f6565b90507f2370726f706f7365723d0000000000000000000000000000000000000000000075ffffffffffffffffffffffffffffffffffffffffffff19168175ffffffffffffffffffffffffffffffffffffffffffff1916146128d85760019250505061292c565b6000806128e986602a860386613407565b9150915081158061292557508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b9450505050505b92915050565b60006129478686868680519060200120611b19565b90508451865114158061295c57508351865114155b80612968575060008651145b156129b1578551845186516040517f447b05d00000000000000000000000000000000000000000000000000000000081526004016129a8939291906143ca565b60405180910390fd5b60006004600083815260200190815260200160002060000160149054906101000a900465ffffffffffff1665ffffffffffff1614612a3557806129f3826111e7565b6000801b6040517f31b75e4d000000000000000000000000000000000000000000000000000000008152600401612a2c93929190615548565b60405180910390fd5b6000612a3f6111d8565b612a47611a77565b65ffffffffffff16612a5991906155c7565b90506000612a65610d38565b90506000600460008581526020019081526020016000209050848160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612aca8361351e565b8160000160146101000a81548165ffffffffffff021916908365ffffffffffff160217905550612af982613578565b81600001601a6101000a81548163ffffffff021916908363ffffffff1602179055507f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e084868b8b8d5167ffffffffffffffff811115612b5b57612b5a613d99565b5b604051908082528060200260200182016040528015612b8e57816020015b6060815260200190600190039081612b795790505b508c89898b612b9d91906155c7565b8e604051612bb399989796959493929190615a0e565b60405180910390a150505095945050505050565b6060612bfd60007f00000000000000000000000000000000000000000000000000000000000000006135d090919063ffffffff16565b905090565b6060612c3860017f00000000000000000000000000000000000000000000000000000000000000006135d090919063ffffffff16565b905090565b6000612c484361351e565b905090565b6000612c57611e69565b73ffffffffffffffffffffffffffffffffffffffff16633a46b1a885856040518363ffffffff1660e01b8152600401612c91929190615abe565b602060405180830381865afa158015612cae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cd29190615afc565b90509392505050565b6000600754905090565b606082612cfa57612cf582613680565b612d02565b819050612d03565b5b92915050565b60008163ffffffff1603612d555760006040517ff1cfbf05000000000000000000000000000000000000000000000000000000008152600401612d4c9190615b64565b60405180910390fd5b7f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828600860069054906101000a900463ffffffff1682604051612d98929190615bb0565b60405180910390a180600860066101000a81548163ffffffff021916908363ffffffff16021790555050565b606060405180602001604052806000815250905090565b7fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc0546160075482604051612e0e929190615bd9565b60405180910390a18060078190555050565b634e487b71600052806020526024601cfd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148015612eae57507f000000000000000000000000000000000000000000000000000000000000000046145b15612edb577f00000000000000000000000000000000000000000000000000000000000000009050612ee6565b612ee36136c5565b90505b90565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60008060006041845103612f6f5760008060006020870151925060408701519150606087015160001a9050612f618882858561375b565b955095509550505050612f7f565b60006002855160001b9250925092505b9250925092565b60008060008573ffffffffffffffffffffffffffffffffffffffff168585604051602401612fb5929190615c02565b604051602081830303815290604052631626ba7e60e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161300791906155b0565b600060405180830381855afa9150503d8060008114613042576040519150601f19603f3d011682016040523d82523d6000602084013e613047565b606091505b509150915081801561305b57506020815110155b801561309f5750631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168180602001905181019061309d9190615c47565b145b925050509392505050565b6000806009600088815260200190815260200160002090508060030160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561315357856040517f71c6af4900000000000000000000000000000000000000000000000000000000815260040161314a9190613d48565b60405180910390fd5b60018160030160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600060028111156131c1576131c06142f8565b5b60ff168560ff16036131ed57838160000160008282546131e191906155c7565b925050819055506132a0565b60016002811115613201576132006142f8565b5b60ff168560ff160361322d578381600101600082825461322191906155c7565b9250508190555061329f565b6002808111156132405761323f6142f8565b5b60ff168560ff160361326c578381600201600082825461326091906155c7565b9250508190555061329e565b6040517f06b337c200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5b8391505095945050505050565b50565b3660008036915091509091565b6000808260000160009054906101000a90046fffffffffffffffffffffffffffffffff1690508260000160109054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16816fffffffffffffffffffffffffffffffff1603613339576133386031612e20565b5b826001016000826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff168152602001908152602001600020549150826001016000826fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600181018360000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050919050565b600081602001830151905092915050565b600080845183118061341857508284115b156134295760008091509150613516565b600060018561343891906155c7565b841180156134b057507f30780000000000000000000000000000000000000000000000000000000000007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191661348d878761384f565b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b9050600060026134c1831515613860565b6134cb9190615c74565b60286134d791906155c7565b90508086866134e691906157ae565b0361350c576000806134f989898961386c565b9150915081819550955050505050613516565b6000809350935050505b935093915050565b600065ffffffffffff8016821115613570576030826040517f6dfcc650000000000000000000000000000000000000000000000000000000008152600401613567929190615cf1565b60405180910390fd5b819050919050565b600063ffffffff80168211156135c8576020826040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526004016135bf929190615d55565b60405180910390fd5b819050919050565b606060ff60001b83146135ed576135e6836139a4565b905061367a565b8180546135f990614eff565b80601f016020809104026020016040519081016040528092919081815260200182805461362590614eff565b80156136725780601f1061364757610100808354040283529160200191613672565b820191906000526020600020905b81548152906001019060200180831161365557829003601f168201915b505050505090505b92915050565b6000815111156136935780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000004630604051602001613740959493929190615d7e565b60405160208183030381529060405280519060200120905090565b60008060007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08460001c111561379b576000600385925092509250613845565b6000600188888888604051600081526020016040526040516137c09493929190615dd1565b6020604051602081039080840390855afa1580156137e2573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361383657600060016000801b93509350935050613845565b8060008060001b935093509350505b9450945094915050565b600081602001830151905092915050565b60008115159050919050565b6000806000859050600060018661388391906155c7565b851180156138fb57507f30780000000000000000000000000000000000000000000000000000000000007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166138d8838861384f565b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b90506000600261390c831515613860565b6139169190615c74565b9050600080828961392791906155c7565b90505b8781101561398f576000613946613941878461384f565b613a18565b9050600f8160ff161115613966576000809750975050505050505061399c565b6010836139739190615c74565b92508060ff1683019250508061398890614fe8565b905061392a565b5060018195509550505050505b935093915050565b606060006139b183613aa9565b90506000602067ffffffffffffffff8111156139d0576139cf613d99565b5b6040519080825280601f01601f191660200182016040528015613a025781602001600182028036833780820191505090505b5090508181528360208201528092505050919050565b6000808260f81c9050602f8160ff16118015613a375750603a8160ff16105b15613a4757603081039050613a9f565b60608160ff16118015613a5d575060678160ff16105b15613a6d57605781039050613a9e565b60408160ff16118015613a83575060478160ff16105b15613a9357603781039050613a9d565b60ff915050613aa4565b5b5b809150505b919050565b60008060ff8360001c169050601f811115613af0576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613b4281613b0d565b8114613b4d57600080fd5b50565b600081359050613b5f81613b39565b92915050565b600060208284031215613b7b57613b7a613b03565b5b6000613b8984828501613b50565b91505092915050565b60008115159050919050565b613ba781613b92565b82525050565b6000602082019050613bc26000830184613b9e565b92915050565b6000819050919050565b613bdb81613bc8565b82525050565b6000602082019050613bf66000830184613bd2565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613c36578082015181840152602081019050613c1b565b60008484015250505050565b6000601f19601f8301169050919050565b6000613c5e82613bfc565b613c688185613c07565b9350613c78818560208601613c18565b613c8181613c42565b840191505092915050565b60006020820190508181036000830152613ca68184613c53565b905092915050565b613cb781613bc8565b8114613cc257600080fd5b50565b600081359050613cd481613cae565b92915050565b600060208284031215613cf057613cef613b03565b5b6000613cfe84828501613cc5565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613d3282613d07565b9050919050565b613d4281613d27565b82525050565b6000602082019050613d5d6000830184613d39565b92915050565b613d6c81613d27565b8114613d7757600080fd5b50565b600081359050613d8981613d63565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b613dd182613c42565b810181811067ffffffffffffffff82111715613df057613def613d99565b5b80604052505050565b6000613e03613af9565b9050613e0f8282613dc8565b919050565b600067ffffffffffffffff821115613e2f57613e2e613d99565b5b613e3882613c42565b9050602081019050919050565b82818337600083830152505050565b6000613e67613e6284613e14565b613df9565b905082815260208101848484011115613e8357613e82613d94565b5b613e8e848285613e45565b509392505050565b600082601f830112613eab57613eaa613d8f565b5b8135613ebb848260208601613e54565b91505092915050565b60008060008060808587031215613ede57613edd613b03565b5b6000613eec87828801613d7a565b9450506020613efd87828801613d7a565b9350506040613f0e87828801613cc5565b925050606085013567ffffffffffffffff811115613f2f57613f2e613b08565b5b613f3b87828801613e96565b91505092959194509250565b613f5081613b0d565b82525050565b6000602082019050613f6b6000830184613f47565b92915050565b600067ffffffffffffffff821115613f8c57613f8b613d99565b5b602082029050602081019050919050565b600080fd5b6000613fb5613fb084613f71565b613df9565b90508083825260208201905060208402830185811115613fd857613fd7613f9d565b5b835b818110156140015780613fed8882613d7a565b845260208401935050602081019050613fda565b5050509392505050565b600082601f8301126140205761401f613d8f565b5b8135614030848260208601613fa2565b91505092915050565b600067ffffffffffffffff82111561405457614053613d99565b5b602082029050602081019050919050565b600061407861407384614039565b613df9565b9050808382526020820190506020840283018581111561409b5761409a613f9d565b5b835b818110156140c457806140b08882613cc5565b84526020840193505060208101905061409d565b5050509392505050565b600082601f8301126140e3576140e2613d8f565b5b81356140f3848260208601614065565b91505092915050565b600067ffffffffffffffff82111561411757614116613d99565b5b602082029050602081019050919050565b600061413b614136846140fc565b613df9565b9050808382526020820190506020840283018581111561415e5761415d613f9d565b5b835b818110156141a557803567ffffffffffffffff81111561418357614182613d8f565b5b8086016141908982613e96565b85526020850194505050602081019050614160565b5050509392505050565b600082601f8301126141c4576141c3613d8f565b5b81356141d4848260208601614128565b91505092915050565b6000819050919050565b6141f0816141dd565b81146141fb57600080fd5b50565b60008135905061420d816141e7565b92915050565b6000806000806080858703121561422d5761422c613b03565b5b600085013567ffffffffffffffff81111561424b5761424a613b08565b5b6142578782880161400b565b945050602085013567ffffffffffffffff81111561427857614277613b08565b5b614284878288016140ce565b935050604085013567ffffffffffffffff8111156142a5576142a4613b08565b5b6142b1878288016141af565b92505060606142c2878288016141fe565b91505092959194509250565b6142d7816141dd565b82525050565b60006020820190506142f260008301846142ce565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60088110614338576143376142f8565b5b50565b600081905061434982614327565b919050565b60006143598261433b565b9050919050565b6143698161434e565b82525050565b60006020820190506143846000830184614360565b92915050565b600080604083850312156143a1576143a0613b03565b5b60006143af85828601613cc5565b92505060206143c085828601613d7a565b9150509250929050565b60006060820190506143df6000830186613bd2565b6143ec6020830185613bd2565b6143f96040830184613bd2565b949350505050565b600060ff82169050919050565b61441781614401565b811461442257600080fd5b50565b6000813590506144348161440e565b92915050565b6000806040838503121561445157614450613b03565b5b600061445f85828601613cc5565b925050602061447085828601614425565b9150509250929050565b600080fd5b60008083601f84011261449557614494613d8f565b5b8235905067ffffffffffffffff8111156144b2576144b161447a565b5b6020830191508360018202830111156144ce576144cd613f9d565b5b9250929050565b600080600080600080600060c0888a0312156144f4576144f3613b03565b5b60006145028a828b01613cc5565b97505060206145138a828b01614425565b96505060406145248a828b01613d7a565b955050606088013567ffffffffffffffff81111561454557614544613b08565b5b6145518a828b0161447f565b9450945050608088013567ffffffffffffffff81111561457457614573613b08565b5b6145808a828b01613e96565b92505060a088013567ffffffffffffffff8111156145a1576145a0613b08565b5b6145ad8a828b01613e96565b91505092959891949750929550565b6000806000806000608086880312156145d8576145d7613b03565b5b60006145e688828901613cc5565b95505060206145f788828901614425565b945050604086013567ffffffffffffffff81111561461857614617613b08565b5b6146248882890161447f565b9350935050606086013567ffffffffffffffff81111561464757614646613b08565b5b61465388828901613e96565b9150509295509295909350565b600065ffffffffffff82169050919050565b61467b81614660565b811461468657600080fd5b50565b60008135905061469881614672565b92915050565b6000602082840312156146b4576146b3613b03565b5b60006146c284828501614689565b91505092915050565b600080600080606085870312156146e5576146e4613b03565b5b60006146f387828801613cc5565b945050602061470487828801614425565b935050604085013567ffffffffffffffff81111561472557614724613b08565b5b6147318782880161447f565b925092505092959194509250565b600067ffffffffffffffff82111561475a57614759613d99565b5b61476382613c42565b9050602081019050919050565b600061478361477e8461473f565b613df9565b90508281526020810184848401111561479f5761479e613d94565b5b6147aa848285613e45565b509392505050565b600082601f8301126147c7576147c6613d8f565b5b81356147d7848260208601614770565b91505092915050565b600080600080608085870312156147fa576147f9613b03565b5b600085013567ffffffffffffffff81111561481857614817613b08565b5b6148248782880161400b565b945050602085013567ffffffffffffffff81111561484557614844613b08565b5b614851878288016140ce565b935050604085013567ffffffffffffffff81111561487257614871613b08565b5b61487e878288016141af565b925050606085013567ffffffffffffffff81111561489f5761489e613b08565b5b6148ab878288016147b2565b91505092959194509250565b6000602082840312156148cd576148cc613b03565b5b60006148db84828501613d7a565b91505092915050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b614919816148e4565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61495481613bc8565b82525050565b6000614966838361494b565b60208301905092915050565b6000602082019050919050565b600061498a8261491f565b614994818561492a565b935061499f8361493b565b8060005b838110156149d05781516149b7888261495a565b97506149c283614972565b9250506001810190506149a3565b5085935050505092915050565b600060e0820190506149f2600083018a614910565b8181036020830152614a048189613c53565b90508181036040830152614a188188613c53565b9050614a276060830187613bd2565b614a346080830186613d39565b614a4160a08301856142ce565b81810360c0830152614a53818461497f565b905098975050505050505050565b60008060008060808587031215614a7b57614a7a613b03565b5b6000614a8987828801613cc5565b9450506020614a9a87828801614425565b9350506040614aab87828801613d7a565b925050606085013567ffffffffffffffff811115614acc57614acb613b08565b5b614ad887828801613e96565b91505092959194509250565b614aed81614660565b82525050565b6000602082019050614b086000830184614ae4565b92915050565b600080600060608486031215614b2757614b26613b03565b5b6000614b3586828701613d7a565b9350506020614b4686828701613cc5565b925050604084013567ffffffffffffffff811115614b6757614b66613b08565b5b614b7386828701613e96565b9150509250925092565b600080600080600060a08688031215614b9957614b98613b03565b5b6000614ba788828901613d7a565b9550506020614bb888828901613d7a565b945050604086013567ffffffffffffffff811115614bd957614bd8613b08565b5b614be5888289016140ce565b935050606086013567ffffffffffffffff811115614c0657614c05613b08565b5b614c12888289016140ce565b925050608086013567ffffffffffffffff811115614c3357614c32613b08565b5b614c3f88828901613e96565b9150509295509295909350565b60008083601f840112614c6257614c61613d8f565b5b8235905067ffffffffffffffff811115614c7f57614c7e61447a565b5b602083019150836001820283011115614c9b57614c9a613f9d565b5b9250929050565b60008060008060608587031215614cbc57614cbb613b03565b5b6000614cca87828801613d7a565b9450506020614cdb87828801613cc5565b935050604085013567ffffffffffffffff811115614cfc57614cfb613b08565b5b614d0887828801614c4c565b925092505092959194509250565b600063ffffffff82169050919050565b614d2f81614d16565b8114614d3a57600080fd5b50565b600081359050614d4c81614d26565b92915050565b600060208284031215614d6857614d67613b03565b5b6000614d7684828501614d3d565b91505092915050565b60008060408385031215614d9657614d95613b03565b5b6000614da485828601613d7a565b9250506020614db585828601613cc5565b9150509250929050565b600080600080600060a08688031215614ddb57614dda613b03565b5b6000614de988828901613d7a565b9550506020614dfa88828901613d7a565b9450506040614e0b88828901613cc5565b9350506060614e1c88828901613cc5565b925050608086013567ffffffffffffffff811115614e3d57614e3c613b08565b5b614e4988828901613e96565b9150509295509295909350565b6000819050919050565b6000614e7b614e76614e7184613d07565b614e56565b613d07565b9050919050565b6000614e8d82614e60565b9050919050565b6000614e9f82614e82565b9050919050565b614eaf81614e94565b82525050565b6000602082019050614eca6000830184614ea6565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680614f1757607f821691505b602082108103614f2a57614f29614ed0565b5b50919050565b6000614f4b614f46614f4184614660565b614e56565b613bc8565b9050919050565b614f5b81614f30565b82525050565b6000604082019050614f766000830185613bd2565b614f836020830184614f52565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000614ff382613bc8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361502557615024614fb9565b5b600182019050919050565b60006040820190506150456000830185613bd2565b6150526020830184613d39565b9392505050565b600061506c6150678461473f565b613df9565b90508281526020810184848401111561508857615087613d94565b5b615093848285613c18565b509392505050565b600082601f8301126150b0576150af613d8f565b5b81516150c0848260208601615059565b91505092915050565b6000602082840312156150df576150de613b03565b5b600082015167ffffffffffffffff8111156150fd576150fc613b08565b5b6151098482850161509b565b91505092915050565b600081905092915050565b60006151298385615112565b9350615136838584613e45565b82840190509392505050565b600061514f82848661511d565b91508190509392505050565b61516481614401565b82525050565b600060e08201905061517f600083018a6142ce565b61518c6020830189613bd2565b615199604083018861515b565b6151a66060830187613d39565b6151b36080830186613bd2565b6151c060a08301856142ce565b6151cd60c08301846142ce565b98975050505050505050565b60006151e482614660565b91506151ef83614660565b9250828203905065ffffffffffff81111561520d5761520c614fb9565b5b92915050565b60006060820190506152286000830186613d39565b6152356020830185613bd2565b6152426040830184613bd2565b949350505050565b600060a08201905061525f60008301886142ce565b61526c6020830187613bd2565b615279604083018661515b565b6152866060830185613d39565b6152936080830184613bd2565b9695505050505050565b6000815190506152ac81614672565b92915050565b6000602082840312156152c8576152c7613b03565b5b60006152d68482850161529d565b91505092915050565b60006152ea82614660565b91506152f583614660565b9250828201905065ffffffffffff81111561531357615312614fb9565b5b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61534e81613d27565b82525050565b60006153608383615345565b60208301905092915050565b6000602082019050919050565b600061538482615319565b61538e8185615324565b935061539983615335565b8060005b838110156153ca5781516153b18882615354565b97506153bc8361536c565b92505060018101905061539d565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b600061542a82615403565b615434818561540e565b9350615444818560208601613c18565b61544d81613c42565b840191505092915050565b6000615464838361541f565b905092915050565b6000602082019050919050565b6000615484826153d7565b61548e81856153e2565b9350836020820285016154a0856153f3565b8060005b858110156154dc57848403895281516154bd8582615458565b94506154c88361546c565b925060208a019950506001810190506154a4565b50829750879550505050505092915050565b600060808201905081810360008301526155088187615379565b9050818103602083015261551c818661497f565b905081810360408301526155308185615479565b905061553f60608301846142ce565b95945050505050565b600060608201905061555d6000830186613bd2565b61556a6020830185614360565b61557760408301846142ce565b949350505050565b600061558a82615403565b6155948185615112565b93506155a4818560208601613c18565b80840191505092915050565b60006155bc828461557f565b915081905092915050565b60006155d282613bc8565b91506155dd83613bc8565b92508282019050808211156155f5576155f4614fb9565b5b92915050565b600061560682614401565b915061561183614401565b9250828201905060ff81111561562a57615629614fb9565b5b92915050565b60008160011c9050919050565b6000808291508390505b60018511156156875780860481111561566357615662614fb9565b5b60018516156156725780820291505b808102905061568085615630565b9450615647565b94509492505050565b6000826156a0576001905061575c565b816156ae576000905061575c565b81600181146156c457600281146156ce576156fd565b600191505061575c565b60ff8411156156e0576156df614fb9565b5b8360020a9150848211156156f7576156f6614fb9565b5b5061575c565b5060208310610133831016604e8410600b84101617156157325782820a90508381111561572d5761572c614fb9565b5b61575c565b61573f848484600161563d565b9250905081840481111561575657615755614fb9565b5b81810290505b9392505050565b600061576e82613bc8565b915061577983614401565b92506157a67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484615690565b905092915050565b60006157b982613bc8565b91506157c483613bc8565b92508282039050818111156157dc576157db614fb9565b5b92915050565b60006080820190506157f76000830187613bd2565b615804602083018661515b565b6158116040830185613bd2565b81810360608301526158238184613c53565b905095945050505050565b600082825260208201905092915050565b600061584a82615403565b615854818561582e565b9350615864818560208601613c18565b61586d81613c42565b840191505092915050565b600060a08201905061588d6000830188613bd2565b61589a602083018761515b565b6158a76040830186613bd2565b81810360608301526158b98185613c53565b905081810360808301526158cd818461583f565b90509695505050505050565b60006040820190506158ee6000830185614f52565b6158fb6020830184614f52565b9392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061594a82613bfc565b615954818561592e565b9350615964818560208601613c18565b61596d81613c42565b840191505092915050565b6000615984838361593f565b905092915050565b6000602082019050919050565b60006159a482615902565b6159ae818561590d565b9350836020820285016159c08561591e565b8060005b858110156159fc57848403895281516159dd8582615978565b94506159e88361598c565b925060208a019950506001810190506159c4565b50829750879550505050505092915050565b600061012082019050615a24600083018c613bd2565b615a31602083018b613d39565b8181036040830152615a43818a615379565b90508181036060830152615a57818961497f565b90508181036080830152615a6b8188615999565b905081810360a0830152615a7f8187615479565b9050615a8e60c0830186613bd2565b615a9b60e0830185613bd2565b818103610100830152615aae8184613c53565b90509a9950505050505050505050565b6000604082019050615ad36000830185613d39565b615ae06020830184613bd2565b9392505050565b600081519050615af681613cae565b92915050565b600060208284031215615b1257615b11613b03565b5b6000615b2084828501615ae7565b91505092915050565b6000819050919050565b6000615b4e615b49615b4484615b29565b614e56565b613bc8565b9050919050565b615b5e81615b33565b82525050565b6000602082019050615b796000830184615b55565b92915050565b6000615b9a615b95615b9084614d16565b614e56565b613bc8565b9050919050565b615baa81615b7f565b82525050565b6000604082019050615bc56000830185615ba1565b615bd26020830184615ba1565b9392505050565b6000604082019050615bee6000830185613bd2565b615bfb6020830184613bd2565b9392505050565b6000604082019050615c1760008301856142ce565b8181036020830152615c29818461583f565b90509392505050565b600081519050615c41816141e7565b92915050565b600060208284031215615c5d57615c5c613b03565b5b6000615c6b84828501615c32565b91505092915050565b6000615c7f82613bc8565b9150615c8a83613bc8565b9250828202615c9881613bc8565b91508282048414831517615caf57615cae614fb9565b5b5092915050565b6000819050919050565b6000615cdb615cd6615cd184615cb6565b614e56565b614401565b9050919050565b615ceb81615cc0565b82525050565b6000604082019050615d066000830185615ce2565b615d136020830184613bd2565b9392505050565b6000819050919050565b6000615d3f615d3a615d3584615d1a565b614e56565b614401565b9050919050565b615d4f81615d24565b82525050565b6000604082019050615d6a6000830185615d46565b615d776020830184613bd2565b9392505050565b600060a082019050615d9360008301886142ce565b615da060208301876142ce565b615dad60408301866142ce565b615dba6060830185613bd2565b615dc76080830184613d39565b9695505050505050565b6000608082019050615de660008301876142ce565b615df3602083018661515b565b615e0060408301856142ce565b615e0d60608301846142ce565b9594505050505056fea264697066735822122054b03186b0ce45a71e74c10e03d3337e4cc10deac0805d0ef67a2e4d06c7396664736f6c63430008140033",
}

// GovernorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernorContractMetaData.ABI instead.
var GovernorContractABI = GovernorContractMetaData.ABI

// GovernorContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernorContractMetaData.Bin instead.
var GovernorContractBin = GovernorContractMetaData.Bin

// DeployGovernorContract deploys a new Ethereum contract, binding an instance of GovernorContract to it.
func DeployGovernorContract(auth *bind.TransactOpts, backend bind.ContractBackend, _votingPowerSource common.Address) (common.Address, *types.Transaction, *GovernorContract, error) {
	parsed, err := GovernorContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernorContractBin), backend, _votingPowerSource)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernorContract{GovernorContractCaller: GovernorContractCaller{contract: contract}, GovernorContractTransactor: GovernorContractTransactor{contract: contract}, GovernorContractFilterer: GovernorContractFilterer{contract: contract}}, nil
}

// GovernorContract is an auto generated Go binding around an Ethereum contract.
type GovernorContract struct {
	GovernorContractCaller     // Read-only binding to the contract
	GovernorContractTransactor // Write-only binding to the contract
	GovernorContractFilterer   // Log filterer for contract events
}

// GovernorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernorContractSession struct {
	Contract     *GovernorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernorContractCallerSession struct {
	Contract *GovernorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GovernorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernorContractTransactorSession struct {
	Contract     *GovernorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GovernorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernorContractRaw struct {
	Contract *GovernorContract // Generic contract binding to access the raw methods on
}

// GovernorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernorContractCallerRaw struct {
	Contract *GovernorContractCaller // Generic read-only contract binding to access the raw methods on
}

// GovernorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernorContractTransactorRaw struct {
	Contract *GovernorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernorContract creates a new instance of GovernorContract, bound to a specific deployed contract.
func NewGovernorContract(address common.Address, backend bind.ContractBackend) (*GovernorContract, error) {
	contract, err := bindGovernorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernorContract{GovernorContractCaller: GovernorContractCaller{contract: contract}, GovernorContractTransactor: GovernorContractTransactor{contract: contract}, GovernorContractFilterer: GovernorContractFilterer{contract: contract}}, nil
}

// NewGovernorContractCaller creates a new read-only instance of GovernorContract, bound to a specific deployed contract.
func NewGovernorContractCaller(address common.Address, caller bind.ContractCaller) (*GovernorContractCaller, error) {
	contract, err := bindGovernorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernorContractCaller{contract: contract}, nil
}

// NewGovernorContractTransactor creates a new write-only instance of GovernorContract, bound to a specific deployed contract.
func NewGovernorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernorContractTransactor, error) {
	contract, err := bindGovernorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernorContractTransactor{contract: contract}, nil
}

// NewGovernorContractFilterer creates a new log filterer instance of GovernorContract, bound to a specific deployed contract.
func NewGovernorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernorContractFilterer, error) {
	contract, err := bindGovernorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernorContractFilterer{contract: contract}, nil
}

// bindGovernorContract binds a generic wrapper to an already deployed contract.
func bindGovernorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernorContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernorContract *GovernorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernorContract.Contract.GovernorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernorContract *GovernorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernorContract.Contract.GovernorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernorContract *GovernorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernorContract.Contract.GovernorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernorContract *GovernorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernorContract *GovernorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernorContract *GovernorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernorContract.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _GovernorContract.Contract.BALLOTTYPEHASH(&_GovernorContract.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _GovernorContract.Contract.BALLOTTYPEHASH(&_GovernorContract.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernorContract *GovernorContractCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernorContract *GovernorContractSession) CLOCKMODE() (string, error) {
	return _GovernorContract.Contract.CLOCKMODE(&_GovernorContract.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GovernorContract *GovernorContractCallerSession) CLOCKMODE() (string, error) {
	return _GovernorContract.Contract.CLOCKMODE(&_GovernorContract.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GovernorContract *GovernorContractCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GovernorContract *GovernorContractSession) COUNTINGMODE() (string, error) {
	return _GovernorContract.Contract.COUNTINGMODE(&_GovernorContract.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GovernorContract *GovernorContractCallerSession) COUNTINGMODE() (string, error) {
	return _GovernorContract.Contract.COUNTINGMODE(&_GovernorContract.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _GovernorContract.Contract.EXTENDEDBALLOTTYPEHASH(&_GovernorContract.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GovernorContract *GovernorContractCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _GovernorContract.Contract.EXTENDEDBALLOTTYPEHASH(&_GovernorContract.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernorContract *GovernorContractCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernorContract *GovernorContractSession) Clock() (*big.Int, error) {
	return _GovernorContract.Contract.Clock(&_GovernorContract.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GovernorContract *GovernorContractCallerSession) Clock() (*big.Int, error) {
	return _GovernorContract.Contract.Clock(&_GovernorContract.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernorContract *GovernorContractCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernorContract *GovernorContractSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GovernorContract.Contract.Eip712Domain(&_GovernorContract.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GovernorContract *GovernorContractCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GovernorContract.Contract.Eip712Domain(&_GovernorContract.CallOpts)
}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) GetProposalId(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "getProposalId", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_GovernorContract *GovernorContractSession) GetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GovernorContract.Contract.GetProposalId(&_GovernorContract.CallOpts, targets, values, calldatas, descriptionHash)
}

// GetProposalId is a free data retrieval call binding the contract method 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) GetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GovernorContract.Contract.GetProposalId(&_GovernorContract.CallOpts, targets, values, calldatas, descriptionHash)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernorContract *GovernorContractSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.GetVotes(&_GovernorContract.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.GetVotes(&_GovernorContract.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GovernorContract *GovernorContractSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _GovernorContract.Contract.GetVotesWithParams(&_GovernorContract.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _GovernorContract.Contract.GetVotesWithParams(&_GovernorContract.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GovernorContract *GovernorContractCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GovernorContract *GovernorContractSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _GovernorContract.Contract.HasVoted(&_GovernorContract.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GovernorContract *GovernorContractCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _GovernorContract.Contract.HasVoted(&_GovernorContract.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GovernorContract *GovernorContractCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GovernorContract *GovernorContractSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GovernorContract.Contract.HashProposal(&_GovernorContract.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GovernorContract.Contract.HashProposal(&_GovernorContract.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernorContract *GovernorContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernorContract *GovernorContractSession) Name() (string, error) {
	return _GovernorContract.Contract.Name(&_GovernorContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GovernorContract *GovernorContractCallerSession) Name() (string, error) {
	return _GovernorContract.Contract.Name(&_GovernorContract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernorContract *GovernorContractSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GovernorContract.Contract.Nonces(&_GovernorContract.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GovernorContract.Contract.Nonces(&_GovernorContract.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalDeadline(&_GovernorContract.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalDeadline(&_GovernorContract.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalEta(&_GovernorContract.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalEta(&_GovernorContract.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_GovernorContract *GovernorContractCaller) ProposalNeedsQueuing(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalNeedsQueuing", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_GovernorContract *GovernorContractSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _GovernorContract.Contract.ProposalNeedsQueuing(&_GovernorContract.CallOpts, arg0)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 ) view returns(bool)
func (_GovernorContract *GovernorContractCallerSession) ProposalNeedsQueuing(arg0 *big.Int) (bool, error) {
	return _GovernorContract.Contract.ProposalNeedsQueuing(&_GovernorContract.CallOpts, arg0)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GovernorContract *GovernorContractCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GovernorContract *GovernorContractSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _GovernorContract.Contract.ProposalProposer(&_GovernorContract.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GovernorContract *GovernorContractCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _GovernorContract.Contract.ProposalProposer(&_GovernorContract.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalSnapshot(&_GovernorContract.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.ProposalSnapshot(&_GovernorContract.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GovernorContract *GovernorContractCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GovernorContract *GovernorContractSession) ProposalThreshold() (*big.Int, error) {
	return _GovernorContract.Contract.ProposalThreshold(&_GovernorContract.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) ProposalThreshold() (*big.Int, error) {
	return _GovernorContract.Contract.ProposalThreshold(&_GovernorContract.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GovernorContract *GovernorContractCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GovernorContract *GovernorContractSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _GovernorContract.Contract.ProposalVotes(&_GovernorContract.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GovernorContract *GovernorContractCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _GovernorContract.Contract.ProposalVotes(&_GovernorContract.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) pure returns(uint256)
func (_GovernorContract *GovernorContractCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) pure returns(uint256)
func (_GovernorContract *GovernorContractSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.Quorum(&_GovernorContract.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) pure returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _GovernorContract.Contract.Quorum(&_GovernorContract.CallOpts, blockNumber)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GovernorContract *GovernorContractCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GovernorContract *GovernorContractSession) State(proposalId *big.Int) (uint8, error) {
	return _GovernorContract.Contract.State(&_GovernorContract.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GovernorContract *GovernorContractCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _GovernorContract.Contract.State(&_GovernorContract.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernorContract *GovernorContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernorContract *GovernorContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernorContract.Contract.SupportsInterface(&_GovernorContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernorContract *GovernorContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernorContract.Contract.SupportsInterface(&_GovernorContract.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GovernorContract *GovernorContractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GovernorContract *GovernorContractSession) Token() (common.Address, error) {
	return _GovernorContract.Contract.Token(&_GovernorContract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GovernorContract *GovernorContractCallerSession) Token() (common.Address, error) {
	return _GovernorContract.Contract.Token(&_GovernorContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GovernorContract *GovernorContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GovernorContract *GovernorContractSession) Version() (string, error) {
	return _GovernorContract.Contract.Version(&_GovernorContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GovernorContract *GovernorContractCallerSession) Version() (string, error) {
	return _GovernorContract.Contract.Version(&_GovernorContract.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GovernorContract *GovernorContractCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GovernorContract *GovernorContractSession) VotingDelay() (*big.Int, error) {
	return _GovernorContract.Contract.VotingDelay(&_GovernorContract.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) VotingDelay() (*big.Int, error) {
	return _GovernorContract.Contract.VotingDelay(&_GovernorContract.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GovernorContract *GovernorContractCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernorContract.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GovernorContract *GovernorContractSession) VotingPeriod() (*big.Int, error) {
	return _GovernorContract.Contract.VotingPeriod(&_GovernorContract.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GovernorContract *GovernorContractCallerSession) VotingPeriod() (*big.Int, error) {
	return _GovernorContract.Contract.VotingPeriod(&_GovernorContract.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Cancel(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Cancel(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GovernorContract *GovernorContractSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVote(&_GovernorContract.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVote(&_GovernorContract.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteBySig(&_GovernorContract.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteBySig(&_GovernorContract.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GovernorContract *GovernorContractSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReason(&_GovernorContract.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReason(&_GovernorContract.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GovernorContract *GovernorContractSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReasonAndParams(&_GovernorContract.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReasonAndParams(&_GovernorContract.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReasonAndParamsBySig(&_GovernorContract.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.CastVoteWithReasonAndParamsBySig(&_GovernorContract.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GovernorContract *GovernorContractTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GovernorContract *GovernorContractSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Execute(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Execute(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC1155BatchReceived(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC1155BatchReceived(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC1155Received(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC1155Received(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC721Received(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernorContract *GovernorContractTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.OnERC721Received(&_GovernorContract.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GovernorContract *GovernorContractSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GovernorContract.Contract.Propose(&_GovernorContract.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GovernorContract.Contract.Propose(&_GovernorContract.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Queue(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GovernorContract *GovernorContractTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Queue(&_GovernorContract.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GovernorContract *GovernorContractTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GovernorContract *GovernorContractSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Relay(&_GovernorContract.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GovernorContract *GovernorContractTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernorContract.Contract.Relay(&_GovernorContract.TransactOpts, target, value, data)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GovernorContract *GovernorContractTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GovernorContract *GovernorContractSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetProposalThreshold(&_GovernorContract.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GovernorContract *GovernorContractTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetProposalThreshold(&_GovernorContract.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GovernorContract *GovernorContractTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GovernorContract *GovernorContractSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetVotingDelay(&_GovernorContract.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GovernorContract *GovernorContractTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetVotingDelay(&_GovernorContract.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GovernorContract *GovernorContractTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod uint32) (*types.Transaction, error) {
	return _GovernorContract.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GovernorContract *GovernorContractSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetVotingPeriod(&_GovernorContract.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GovernorContract *GovernorContractTransactorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _GovernorContract.Contract.SetVotingPeriod(&_GovernorContract.TransactOpts, newVotingPeriod)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernorContract *GovernorContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernorContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernorContract *GovernorContractSession) Receive() (*types.Transaction, error) {
	return _GovernorContract.Contract.Receive(&_GovernorContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernorContract *GovernorContractTransactorSession) Receive() (*types.Transaction, error) {
	return _GovernorContract.Contract.Receive(&_GovernorContract.TransactOpts)
}

// GovernorContractEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GovernorContract contract.
type GovernorContractEIP712DomainChangedIterator struct {
	Event *GovernorContractEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *GovernorContractEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractEIP712DomainChanged)
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
		it.Event = new(GovernorContractEIP712DomainChanged)
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
func (it *GovernorContractEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractEIP712DomainChanged represents a EIP712DomainChanged event raised by the GovernorContract contract.
type GovernorContractEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernorContract *GovernorContractFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GovernorContractEIP712DomainChangedIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GovernorContractEIP712DomainChangedIterator{contract: _GovernorContract.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernorContract *GovernorContractFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GovernorContractEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractEIP712DomainChanged)
				if err := _GovernorContract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GovernorContract *GovernorContractFilterer) ParseEIP712DomainChanged(log types.Log) (*GovernorContractEIP712DomainChanged, error) {
	event := new(GovernorContractEIP712DomainChanged)
	if err := _GovernorContract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the GovernorContract contract.
type GovernorContractProposalCanceledIterator struct {
	Event *GovernorContractProposalCanceled // Event containing the contract specifics and raw log

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
func (it *GovernorContractProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractProposalCanceled)
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
		it.Event = new(GovernorContractProposalCanceled)
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
func (it *GovernorContractProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractProposalCanceled represents a ProposalCanceled event raised by the GovernorContract contract.
type GovernorContractProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*GovernorContractProposalCanceledIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &GovernorContractProposalCanceledIterator{contract: _GovernorContract.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GovernorContractProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractProposalCanceled)
				if err := _GovernorContract.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) ParseProposalCanceled(log types.Log) (*GovernorContractProposalCanceled, error) {
	event := new(GovernorContractProposalCanceled)
	if err := _GovernorContract.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the GovernorContract contract.
type GovernorContractProposalCreatedIterator struct {
	Event *GovernorContractProposalCreated // Event containing the contract specifics and raw log

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
func (it *GovernorContractProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractProposalCreated)
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
		it.Event = new(GovernorContractProposalCreated)
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
func (it *GovernorContractProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractProposalCreated represents a ProposalCreated event raised by the GovernorContract contract.
type GovernorContractProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GovernorContract *GovernorContractFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GovernorContractProposalCreatedIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GovernorContractProposalCreatedIterator{contract: _GovernorContract.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GovernorContract *GovernorContractFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernorContractProposalCreated) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractProposalCreated)
				if err := _GovernorContract.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GovernorContract *GovernorContractFilterer) ParseProposalCreated(log types.Log) (*GovernorContractProposalCreated, error) {
	event := new(GovernorContractProposalCreated)
	if err := _GovernorContract.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GovernorContract contract.
type GovernorContractProposalExecutedIterator struct {
	Event *GovernorContractProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovernorContractProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractProposalExecuted)
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
		it.Event = new(GovernorContractProposalExecuted)
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
func (it *GovernorContractProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractProposalExecuted represents a ProposalExecuted event raised by the GovernorContract contract.
type GovernorContractProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GovernorContractProposalExecutedIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GovernorContractProposalExecutedIterator{contract: _GovernorContract.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernorContractProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractProposalExecuted)
				if err := _GovernorContract.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GovernorContract *GovernorContractFilterer) ParseProposalExecuted(log types.Log) (*GovernorContractProposalExecuted, error) {
	event := new(GovernorContractProposalExecuted)
	if err := _GovernorContract.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the GovernorContract contract.
type GovernorContractProposalQueuedIterator struct {
	Event *GovernorContractProposalQueued // Event containing the contract specifics and raw log

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
func (it *GovernorContractProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractProposalQueued)
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
		it.Event = new(GovernorContractProposalQueued)
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
func (it *GovernorContractProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractProposalQueued represents a ProposalQueued event raised by the GovernorContract contract.
type GovernorContractProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GovernorContract *GovernorContractFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*GovernorContractProposalQueuedIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &GovernorContractProposalQueuedIterator{contract: _GovernorContract.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GovernorContract *GovernorContractFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernorContractProposalQueued) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractProposalQueued)
				if err := _GovernorContract.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GovernorContract *GovernorContractFilterer) ParseProposalQueued(log types.Log) (*GovernorContractProposalQueued, error) {
	event := new(GovernorContractProposalQueued)
	if err := _GovernorContract.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the GovernorContract contract.
type GovernorContractProposalThresholdSetIterator struct {
	Event *GovernorContractProposalThresholdSet // Event containing the contract specifics and raw log

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
func (it *GovernorContractProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractProposalThresholdSet)
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
		it.Event = new(GovernorContractProposalThresholdSet)
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
func (it *GovernorContractProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractProposalThresholdSet represents a ProposalThresholdSet event raised by the GovernorContract contract.
type GovernorContractProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GovernorContract *GovernorContractFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*GovernorContractProposalThresholdSetIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &GovernorContractProposalThresholdSetIterator{contract: _GovernorContract.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GovernorContract *GovernorContractFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *GovernorContractProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractProposalThresholdSet)
				if err := _GovernorContract.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
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

// ParseProposalThresholdSet is a log parse operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GovernorContract *GovernorContractFilterer) ParseProposalThresholdSet(log types.Log) (*GovernorContractProposalThresholdSet, error) {
	event := new(GovernorContractProposalThresholdSet)
	if err := _GovernorContract.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the GovernorContract contract.
type GovernorContractVoteCastIterator struct {
	Event *GovernorContractVoteCast // Event containing the contract specifics and raw log

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
func (it *GovernorContractVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractVoteCast)
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
		it.Event = new(GovernorContractVoteCast)
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
func (it *GovernorContractVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractVoteCast represents a VoteCast event raised by the GovernorContract contract.
type GovernorContractVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GovernorContract *GovernorContractFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*GovernorContractVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &GovernorContractVoteCastIterator{contract: _GovernorContract.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GovernorContract *GovernorContractFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *GovernorContractVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractVoteCast)
				if err := _GovernorContract.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GovernorContract *GovernorContractFilterer) ParseVoteCast(log types.Log) (*GovernorContractVoteCast, error) {
	event := new(GovernorContractVoteCast)
	if err := _GovernorContract.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the GovernorContract contract.
type GovernorContractVoteCastWithParamsIterator struct {
	Event *GovernorContractVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *GovernorContractVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractVoteCastWithParams)
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
		it.Event = new(GovernorContractVoteCastWithParams)
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
func (it *GovernorContractVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractVoteCastWithParams represents a VoteCastWithParams event raised by the GovernorContract contract.
type GovernorContractVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GovernorContract *GovernorContractFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*GovernorContractVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &GovernorContractVoteCastWithParamsIterator{contract: _GovernorContract.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GovernorContract *GovernorContractFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *GovernorContractVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractVoteCastWithParams)
				if err := _GovernorContract.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GovernorContract *GovernorContractFilterer) ParseVoteCastWithParams(log types.Log) (*GovernorContractVoteCastWithParams, error) {
	event := new(GovernorContractVoteCastWithParams)
	if err := _GovernorContract.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the GovernorContract contract.
type GovernorContractVotingDelaySetIterator struct {
	Event *GovernorContractVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *GovernorContractVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractVotingDelaySet)
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
		it.Event = new(GovernorContractVotingDelaySet)
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
func (it *GovernorContractVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractVotingDelaySet represents a VotingDelaySet event raised by the GovernorContract contract.
type GovernorContractVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GovernorContract *GovernorContractFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*GovernorContractVotingDelaySetIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &GovernorContractVotingDelaySetIterator{contract: _GovernorContract.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GovernorContract *GovernorContractFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *GovernorContractVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractVotingDelaySet)
				if err := _GovernorContract.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GovernorContract *GovernorContractFilterer) ParseVotingDelaySet(log types.Log) (*GovernorContractVotingDelaySet, error) {
	event := new(GovernorContractVotingDelaySet)
	if err := _GovernorContract.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernorContractVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the GovernorContract contract.
type GovernorContractVotingPeriodSetIterator struct {
	Event *GovernorContractVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *GovernorContractVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernorContractVotingPeriodSet)
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
		it.Event = new(GovernorContractVotingPeriodSet)
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
func (it *GovernorContractVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernorContractVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernorContractVotingPeriodSet represents a VotingPeriodSet event raised by the GovernorContract contract.
type GovernorContractVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GovernorContract *GovernorContractFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*GovernorContractVotingPeriodSetIterator, error) {

	logs, sub, err := _GovernorContract.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &GovernorContractVotingPeriodSetIterator{contract: _GovernorContract.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GovernorContract *GovernorContractFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *GovernorContractVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _GovernorContract.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernorContractVotingPeriodSet)
				if err := _GovernorContract.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GovernorContract *GovernorContractFilterer) ParseVotingPeriodSet(log types.Log) (*GovernorContractVotingPeriodSet, error) {
	event := new(GovernorContractVotingPeriodSet)
	if err := _GovernorContract.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
