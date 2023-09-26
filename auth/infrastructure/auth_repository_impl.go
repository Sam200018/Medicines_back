package infrastructure

import (
	"errors"

	"github.com/Samuel200018/pills_backend/auth/domain"
	"gorm.io/gorm"
)

type DatabaseAuthRepository struct {
	db *gorm.DB
}

func NewDatabaseAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &DatabaseAuthRepository{db}
}

// Create implements domain.AuthRepository.
func (d *DatabaseAuthRepository) Create(user domain.User) (domain.User, error) {
	var auxUser domain.User

	d.db.Where("email = ?", user.Email).First(&auxUser)

	if auxUser.ID != 0 {
		return domain.User{}, errors.New("Email already exists")
	}
	d.db.Create(&user)
	return user, nil
}

// GetUser implements domain.AuthRepository.
func (d *DatabaseAuthRepository) GetUser(email string) (domain.User, error) {
	panic("unimplemented")
}
