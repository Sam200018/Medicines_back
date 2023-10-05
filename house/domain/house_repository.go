package domain

type HouseRepository interface {
	Create(house House) (House, error)
	GetById(id string) (House, error)
	Update(house House) (House, error)
	Delete(id string) error
}
