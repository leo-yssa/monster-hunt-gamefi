package infrastructure

import (
	"errors"

	"github.com/leo-yssa/monster-hunt-gamefi/backend/domain"
)

type InMemoryPlayerRepository struct {
	players map[string]*domain.Player
}

func NewInMemoryPlayerRepository() *InMemoryPlayerRepository {
	return &InMemoryPlayerRepository{players: make(map[string]*domain.Player)}
}

func (r *InMemoryPlayerRepository) FindByAddress(address string) (*domain.Player, error) {
	player, ok := r.players[address]
	if !ok {
		return nil, errors.New("player not found")
	}
	return player, nil
}

func (r *InMemoryPlayerRepository) Save(player *domain.Player) error {
	r.players[player.Address] = player
	return nil
} 