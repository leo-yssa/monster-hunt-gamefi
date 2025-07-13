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

type CurveLPTokenAdapter struct {
	client    *ethclient.Client
	contract  *CurveLPToken
	contractAddress common.Address
}

var _ domain.CurveLPTokenPort = (*CurveLPTokenAdapter)(nil)

func NewCurveLPTokenAdapter(client *ethclient.Client, contractAddr string) (*CurveLPTokenAdapter, error) {
	addr := common.HexToAddress(contractAddr)
	contract, err := NewCurveLPToken(addr, client)
	if err != nil {
		return nil, err
	}
	return &CurveLPTokenAdapter{
		client:    client,
		contract:  contract,
		contractAddress: addr,
	}, nil
}

// ownerPrivKeyHex는 실제 서비스에서는 안전하게 관리되어야 함
func (a *CurveLPTokenAdapter) Approve(ctx context.Context, ownerPrivKeyHex, spender string, amount int64) (string, error) {
	privKey, err := crypto.HexToECDSA(ownerPrivKeyHex)
	if err != nil {
		return "", err
	}
	spenderAddr := common.HexToAddress(spender)
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(1)) // chainID는 실제 네트워크에 맞게 수정 필요
	if err != nil {
		return "", err
	}
	auth.Context = ctx
	tx, err := a.contract.Approve(auth, spenderAddr, big.NewInt(amount))
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (a *CurveLPTokenAdapter) BalanceOf(ctx context.Context, owner string) (int64, error) {
	ownerAddr := common.HexToAddress(owner)
	callOpts := &bind.CallOpts{
		Context: ctx,
	}
	bal, err := a.contract.BalanceOf(callOpts, ownerAddr)
	if err != nil {
		return 0, err
	}
	return bal.Int64(), nil
} 