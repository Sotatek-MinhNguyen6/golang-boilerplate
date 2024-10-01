package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"unique" json:"username"`
	Password string  `json:"-"`
	Role     string  `json:"role"`   // Can be "user" or "admin"
	Orders   []Order `json:"orders"` // Relation to Orders
}
