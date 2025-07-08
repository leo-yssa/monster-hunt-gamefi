package infrastructure

import (
	"errors"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/domain"
)

type InMemoryMonsterRepository struct {
	monsters map[int]*domain.Monster
}

func NewInMemoryMonsterRepository() *InMemoryMonsterRepository {
	return &InMemoryMonsterRepository{monsters: make(map[int]*domain.Monster)}
}

func (r *InMemoryMonsterRepository) FindByID(id int) (*domain.Monster, error) {
	monster, ok := r.monsters[id]
	if !ok {
		return nil, errors.New("monster not found")
	}
	return monster, nil
}

func (r *InMemoryMonsterRepository) Save(monster *domain.Monster) error {
	r.monsters[monster.ID] = monster
	return nil
}

func (r *InMemoryMonsterRepository) List() ([]*domain.Monster, error) {
	var list []*domain.Monster
	for _, m := range r.monsters {
		list = append(list, m)
	}
	return list, nil
} 