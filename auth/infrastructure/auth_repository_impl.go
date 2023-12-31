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
	var user domain.User

	err := d.db.Where("email= ?", email).First(&user).Error

	return user, err
}

func (d *DatabaseAuthRepository) AddHouse(id string, idHouse uint) (domain.User, error) {
	var user domain.User

	err := d.db.First(&user, id).Error

	if err != nil {
		return domain.User{}, errors.New("User not found")
	}

	user.HouseID = idHouse

	err = d.db.Updates(&user).Error

	return user, err
}

func (d *DatabaseAuthRepository) UsersWithHouseId(idHouse uint) (int64, error) {
	var countUsers int64

	err := d.db.Model(&domain.User{}).Where("house_id =?", idHouse).Count(&countUsers).Error
	return countUsers, err
}

func (d *DatabaseAuthRepository) ExitHouse(id string) (domain.User, error) {
	var user domain.User
	err := d.db.First(&user, id).Error

	if err != nil {
		return domain.User{}, errors.New("Users not found")
	}

	user.HouseID = 0

	err = d.db.Save(&user).Error

	return user, err
}
