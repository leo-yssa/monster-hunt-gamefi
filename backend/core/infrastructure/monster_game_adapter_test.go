package infrastructure

import (
	"context"
	"errors"
	"testing"
)

type mockContractInterface interface {
	RegisterPlayer(auth interface{}, name string) (interface{}, error)
	AddMonster(auth interface{}, name string, hp, reward interface{}) (interface{}, error)
	HuntMonster(auth interface{}, monsterID interface{}) (interface{}, error)
}

type mockContract struct {
	registerPlayerCalled bool
	addMonsterCalled     bool
	huntMonsterCalled    bool
}

func (m *mockContract) RegisterPlayer(auth interface{}, name string) (interface{}, error) {
	m.registerPlayerCalled = true
	if name == "fail" {
		return nil, errors.New("fail")
	}
	return "txhash", nil
}
func (m *mockContract) AddMonster(auth interface{}, name string, hp, reward interface{}) (interface{}, error) {
	m.addMonsterCalled = true
	if name == "fail" {
		return nil, errors.New("fail")
	}
	return "txhash", nil
}
func (m *mockContract) HuntMonster(auth interface{}, monsterID interface{}) (interface{}, error) {
	m.huntMonsterCalled = true
	if monsterID == -1 {
		return nil, errors.New("fail")
	}
	return "txhash", nil
}

type testMonsterGameAdapter struct {
	contract ContractInterface
}

func (a *testMonsterGameAdapter) RegisterPlayer(ctx context.Context, name string) (string, error) {
	_, err := a.contract.RegisterPlayer(nil, name)
	return "mock-txhash", err
}
func (a *testMonsterGameAdapter) AddMonster(ctx context.Context, name string, hp, reward int) (string, error) {
	_, err := a.contract.AddMonster(nil, name, hp, reward)
	return "mock-txhash", err
}
func (a *testMonsterGameAdapter) HuntMonster(ctx context.Context, monsterID int64) (string, error) {
	_, err := a.contract.HuntMonster(nil, monsterID)
	return "mock-txhash", err
}

func newTestAdapter() *testMonsterGameAdapter {
	return &testMonsterGameAdapter{
		contract: &mockContract{},
	}
}

func TestMonsterGameAdapter_RegisterPlayer(t *testing.T) {
	a := newTestAdapter()
	_, err := a.RegisterPlayer(context.Background(), "Alice")
	if err != nil {
		t.Errorf("RegisterPlayer failed: %v", err)
	}
}

func TestMonsterGameAdapter_AddMonster(t *testing.T) {
	a := newTestAdapter()
	_, err := a.AddMonster(context.Background(), "Goblin", 10, 100)
	if err != nil {
		t.Errorf("AddMonster failed: %v", err)
	}
}

func TestMonsterGameAdapter_HuntMonster(t *testing.T) {
	a := newTestAdapter()
	_, err := a.HuntMonster(context.Background(), 1)
	if err != nil {
		t.Errorf("HuntMonster failed: %v", err)
	}
} 