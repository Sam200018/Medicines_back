package domain

import (
	"gorm.io/gorm"
	"time"
)

type Medicine struct {
	gorm.Model

	Name            string    `gorm:"type:varchar(255);not null;unique index" json:"name"`
	Dose            float64   `gorm:"not null" json:"dose"`
	AmountAvailable float64   `gorm:"not null" json:"amountAvailable"`
	DueDate         time.Time `gorm:"not null" json:"dueDate"`
	ActiveCompounds string    `gorm:"type:varchar(455);not null" json:"activeCompounds"`
	HouseID         uint      `gorm:"default:0" json:"houseId" `
}
