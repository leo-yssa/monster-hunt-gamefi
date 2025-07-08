package domain

type PlayerRepository interface {
	FindByAddress(address string) (*Player, error)
	Save(player *Player) error
}