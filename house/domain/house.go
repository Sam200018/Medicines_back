package domain

import "gorm.io/gorm"

type House struct {
	gorm.Model

	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
