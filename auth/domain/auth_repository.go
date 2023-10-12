package domain

type AuthRepository interface {
	Create(user User) (User, error)
	GetUser(email string) (User, error)
	AddHouse(id string, idHouse uint) (User, error)
	UsersWithHouseId(idHouse uint) (int64, error)
	ExitHouse(id string) (User, error)
}
