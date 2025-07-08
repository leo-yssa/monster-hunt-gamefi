package domain

import "testing"

func TestPlayer(t *testing.T) {
	player := &Player{Address: "0x123", Name: "Alice", Level: 1}
	if player.Name != "Alice" {
		t.Errorf("expected name Alice, got %s", player.Name)
	}
} 