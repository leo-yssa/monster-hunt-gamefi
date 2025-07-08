package domain

type MonsterRepository interface {
	FindByID(id int) (*Monster, error)
	Save(monster *Monster) error
	List() ([]*Monster, error)
}