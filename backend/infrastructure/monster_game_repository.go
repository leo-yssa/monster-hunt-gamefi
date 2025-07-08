package infrastructure

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
)

type MonsterGameRepository struct {
	client          *ethclient.Client
	contractAddress common.Address
	contract        *Contract
	privateKey      *ecdsa.PrivateKey
	publicAddress   common.Address
}

func NewMonsterGameRepository(rpcURL, contractAddr, privKeyHex string) (*MonsterGameRepository, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(contractAddr)
	instance, err := NewContract(address, client)
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, err
	}
	publicAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	return &MonsterGameRepository{
		client:          client,
		contractAddress: address,
		contract:        instance,
		privateKey:      privKey,
		publicAddress:   publicAddr,
	}, nil
}

func (r *MonsterGameRepository) getAuth() (*bind.TransactOpts, error) {
	chainID, err := r.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(r.privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (r *MonsterGameRepository) withAuthTx(txFunc func(auth *bind.TransactOpts) (*types.Transaction, error)) (string, error) {
	auth, err := r.getAuth()
	if err != nil {
		return "", err
	}
	tx, err := txFunc(auth)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (r *MonsterGameRepository) RegisterPlayer(name string) (string, error) {
	return r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.contract.RegisterPlayer(auth, name)
	})
}

func (r *MonsterGameRepository) HuntMonster(monsterID int64) (string, error) {
	return r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.contract.HuntMonster(auth, big.NewInt(monsterID))
	})
}

func (r *MonsterGameRepository) AddMonster(name string, hp, reward int) error {
	_, err := r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		return r.contract.AddMonster(auth, name, big.NewInt(int64(hp)), big.NewInt(int64(reward)))
	})
	return err
} 