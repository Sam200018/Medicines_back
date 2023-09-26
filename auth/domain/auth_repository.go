package domain

type AuthRepository interface {
	Create(user User) (User, error)
	GetUser(email string) (User, error)
}
