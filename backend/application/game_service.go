package application

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/domain"
)

type GameService struct {
	PlayerRepo  domain.PlayerRepository
	MonsterRepo domain.MonsterRepository
}

func (s *GameService) RegisterPlayer(address, name string) error {
	player := &domain.Player{Address: address, Name: name, Level: 1}
	return s.PlayerRepo.Save(player)
}

func (s *GameService) AddMonster(name string, hp, reward int) error {
	monsters, _ := s.MonsterRepo.List()
	monster := &domain.Monster{ID: len(monsters), Name: name, HP: hp, Reward: reward}
	return s.MonsterRepo.Save(monster)
}

func (s *GameService) HuntMonster(playerAddr string, monsterID int) (int, error) {
	player, err := s.PlayerRepo.FindByAddress(playerAddr)
	if err != nil {
		return 0, err
	}
	monster, err := s.MonsterRepo.FindByID(monsterID)
	if err != nil {
		return 0, err
	}
	// 실제 컨트랙트 연동 및 보상 지급 로직은 infrastructure에서 구현
	return monster.Reward, nil
}

func (s *GameService) HuntMonster(auth *bind.TransactOpts, monsterID *big.Int) (string, error) {
	input, err := s.contractAbi.Pack("huntMonster", monsterID)
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(
		nonce, s.address, big.NewInt(0), gasLimit, gasPrice, input,
	)
	// ... 서명 및 전송 생략 (bind.NewKeyedTransactor 등 활용)
	// 실제로는 bind.ContractTransactor 사용 권장
	return tx.Hash().Hex(), nil
} 