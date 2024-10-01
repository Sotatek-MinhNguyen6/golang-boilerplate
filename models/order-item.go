package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`             // Foreign key for Order
	Order     Order   `gorm:"foreignKey:OrderID"`   // Relation to Order
	ProductID uint    `json:"product_id"`           // Foreign key for Product
	Product   Product `gorm:"foreignKey:ProductID"` // Relation to Product
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"` // Price at the time of order
}

func (e *OrderItem) TableName() string {
	return "order_items"
}
