package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

func (e *Product) TableName() string {
	return "products"
}
