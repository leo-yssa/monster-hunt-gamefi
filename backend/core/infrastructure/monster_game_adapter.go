package infrastructure

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/config"
	"github.com/leo-yssa/monster-hunt-gamefi/backend/core/domain"
)

// MonsterGameAdapter는 domain.MonsterGamePort를 구현하는 블록체인 연동 어댑터입니다.
type MonsterGameAdapter struct {
	client          *ethclient.Client
	contractAddress common.Address
	contract        ContractInterface
	privateKey      *ecdsa.PrivateKey
	publicAddress   common.Address
}

var _ domain.MonsterGamePort = (*MonsterGameAdapter)(nil)

func NewMonsterGameAdapter(rpcURL, contractAddr, privKeyHex string) (*MonsterGameAdapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}
	_, err = client.NetworkID(ctx)
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
	return &MonsterGameAdapter{
		client:          client,
		contractAddress: address,
		contract:        &contractWrapper{inner: instance},
		privateKey:      privKey,
		publicAddress:   publicAddr,
	}, nil
}

// 환경변수 기반 MonsterGameAdapter 초기화
func InitMonsterGameAdapterFromEnv() (*MonsterGameAdapter, error) {
	rpcURL := config.LoadEnvOrDefault("MONSTER_GAME_RPC", "http://localhost:8545")
	contractAddr := os.Getenv("MONSTER_GAME_CONTRACT")
	privKeyHex := os.Getenv("MONSTER_GAME_PRIVKEY")
	if contractAddr == "" || privKeyHex == "" {
		log.Fatal("컨트랙트 주소와 프라이빗키 환경변수(MONSTER_GAME_CONTRACT, MONSTER_GAME_PRIVKEY)가 필요합니다.")
	}
	return NewMonsterGameAdapter(rpcURL, contractAddr, privKeyHex)
}

func (r *MonsterGameAdapter) getAuth() (*bind.TransactOpts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	chainID, err := r.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(r.privateKey, chainID)
	if err != nil {
		return nil, err
	}
	
	gasPrice, err := r.client.SuggestGasPrice(ctx)
	if err == nil {
	    auth.GasPrice = gasPrice
	}
	
	return auth, nil
}

func (r *MonsterGameAdapter) withAuthTx(txFunc func(auth *bind.TransactOpts) (*types.Transaction, error)) (string, error) {
	auth, err := r.getAuth()
	if err != nil {
		return "", err
	}
	
	// 컨텍스트 타임아웃 설정
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	// 트랜잭션 옵션에 컨텍스트 추가
	auth.Context = ctx
	
	tx, err := txFunc(auth)
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (r *MonsterGameAdapter) RegisterPlayer(ctx context.Context, name string) (string, error) {
	return r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		res, err := r.contract.RegisterPlayer(auth, name)
		if err != nil {
			return nil, err
		}
		return res.(*types.Transaction), nil
	})
}

func (r *MonsterGameAdapter) HuntMonster(ctx context.Context, monsterID int64) (string, error) {
	return r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		res, err := r.contract.HuntMonster(auth, big.NewInt(monsterID))
		if err != nil {
			return nil, err
		}
		return res.(*types.Transaction), nil
	})
}

func (r *MonsterGameAdapter) AddMonster(ctx context.Context, name string, hp, reward int) (string, error) {
	return r.withAuthTx(func(auth *bind.TransactOpts) (*types.Transaction, error) {
		res, err := r.contract.AddMonster(auth, name, big.NewInt(int64(hp)), big.NewInt(int64(reward)))
		if err != nil {
			return nil, err
		}
		return res.(*types.Transaction), nil
	})
} 

func (r *MonsterGameAdapter) Client() *ethclient.Client {
	return r.client
} 

type ContractInterface interface {
	RegisterPlayer(auth interface{}, name string) (interface{}, error)
	AddMonster(auth interface{}, name string, hp, reward interface{}) (interface{}, error)
	HuntMonster(auth interface{}, monsterID interface{}) (interface{}, error)
} 

type contractWrapper struct {
	inner *Contract
}

func (w *contractWrapper) RegisterPlayer(auth interface{}, name string) (interface{}, error) {
	return w.inner.RegisterPlayer(auth.(*bind.TransactOpts), name)
}
func (w *contractWrapper) AddMonster(auth interface{}, name string, hp, reward interface{}) (interface{}, error) {
	return w.inner.AddMonster(auth.(*bind.TransactOpts), name, hp.(*big.Int), reward.(*big.Int))
}
func (w *contractWrapper) HuntMonster(auth interface{}, monsterID interface{}) (interface{}, error) {
	return w.inner.HuntMonster(auth.(*bind.TransactOpts), monsterID.(*big.Int))
} 