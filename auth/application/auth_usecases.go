package application

import "github.com/Samuel200018/pills_backend/auth/domain"

type AuthUseCase struct {
	authRepo domain.AuthRepository
}

func NewUseCasesAuth(authRepo domain.AuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo}
}

func (auc *AuthUseCase) CreateUser(user domain.User) (domain.User, error) {
	return auc.authRepo.Create(user)
}

func (auc *AuthUseCase) GetUser(email string) (domain.User, error) {
	return auc.authRepo.GetUser(email)
}

func (auc *AuthUseCase) AddHouse(id string, idHouse uint) (domain.User, error) {
	return auc.authRepo.AddHouse(id, idHouse)
}
