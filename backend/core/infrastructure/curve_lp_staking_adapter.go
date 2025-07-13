package infrastructure

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
)

type CurveLPStakingAdapter struct {
	client    *ethclient.Client
	contract  *CurveLPStaking
	contractAddress common.Address
}

var _ domain.CurveLPStakingPort = (*CurveLPStakingAdapter)(nil)

func NewCurveLPStakingAdapter(client *ethclient.Client, contractAddr string) (*CurveLPStakingAdapter, error) {
	addr := common.HexToAddress(contractAddr)
	contract, err := NewCurveLPStaking(addr, client)
	if err != nil {
		return nil, err
	}
	return &CurveLPStakingAdapter{
		client:    client,
		contract:  contract,
		contractAddress: addr,
	}, nil
}

// userPrivKeyHex는 실제 서비스에서는 안전하게 관리되어야 함
func (a *CurveLPStakingAdapter) Stake(ctx context.Context, userPrivKeyHex string, amount int64) (string, error) {
	privKey, err := crypto.HexToECDSA(userPrivKeyHex)
	if err != nil {
		return "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(1)) // chainID는 실제 네트워크에 맞게 수정 필요
	if err != nil {
		return "", err
	}
	auth.Context = ctx
	tx, err := a.contract.Stake(auth, big.NewInt(amount))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (a *CurveLPStakingAdapter) Unstake(ctx context.Context, userPrivKeyHex string, amount int64) (string, error) {
	privKey, err := crypto.HexToECDSA(userPrivKeyHex)
	if err != nil {
		return "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(1)) // chainID는 실제 네트워크에 맞게 수정 필요
	if err != nil {
		return "", err
	}
	auth.Context = ctx
	tx, err := a.contract.Unstake(auth, big.NewInt(amount))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (a *CurveLPStakingAdapter) VotingPower(ctx context.Context, user string) (int64, error) {
	userAddr := common.HexToAddress(user)
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	vp, err := a.contract.VotingPower(callOpts, userAddr)
	if err != nil {
		return 0, err
	}
	return vp.Int64(), nil
} 