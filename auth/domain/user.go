package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"type:varchar(255);not null" json:"name"`
	LastName  string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string `gorm:"type:varchar(255);not null;unique index" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	HouseID   uint   `gorm:"default:0" json:"house_id" `
}
