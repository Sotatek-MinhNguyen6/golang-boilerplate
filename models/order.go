package models

import (
	"gorm.io/gorm"
)

// OrderStatus defines the possible statuses of an order.
type OrderStatus string

const (
	Created   OrderStatus = "created"
	Confirmed OrderStatus = "confirmed"
	Delivered OrderStatus = "delivered"
	Cancelled OrderStatus = "cancelled"
)

// Order represents the structure of an order in the database.
type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`           // Foreign key for User
	User       User        `gorm:"foreignKey:UserID"` // Relation to User
	Status     OrderStatus `json:"status"`            // Order status
	TotalPrice uint        `json:"total_price"`
	Items      []OrderItem `json:"items"` // Relation to OrderItems
}

// TableName overrides the table name used by Order to prevent gorm from inferring it incorrectly.
func (e *Order) TableName() string {
	return "orders"
}
