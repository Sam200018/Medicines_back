package infrastructure

import (
	"github.com/Samuel200018/pills_backend/house/domain"
	"gorm.io/gorm"
)

type DatabaseHouseRepository struct {
	db *gorm.DB
}

func NewDatabaseHouseRepository(db *gorm.DB) domain.HouseRepository {
	return &DatabaseHouseRepository{db}
}

func (d *DatabaseHouseRepository) Create(house domain.House) (domain.House, error) {
	err := d.db.Create(&house).Error

	return house, err
}

func (d *DatabaseHouseRepository) GetById(id string) (domain.House, error) {
	var house domain.House

	err := d.db.First(&house, id).Error

	return house, err
}

func (d *DatabaseHouseRepository) Update(house domain.House) (domain.House, error) {
	err := d.db.Model(&domain.House{}).Where("id = ?", house.ID).Updates(&house).Error
	return house, err
}

func (d *DatabaseHouseRepository) Delete(id string) error {
	var house domain.House
	d.db.First(&house, id)

	return d.db.Delete(&house).Error
}
