package dto

type CreateOrderRequest struct {
	Items []CreateOrderItemRequest
}

type CreateOrderItemRequest struct {
	ProductId uint `json:"productId" validate:"required"`
	Amount    uint `json:"amount" validate:"required,gt=0"`
}
