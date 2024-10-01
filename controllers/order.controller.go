package controllers

import (
	"gin-example/infra/database"
	"gin-example/models"
	"gin-example/validation"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func (ctrl *OrderController) CreateOrder(ctx *gin.Context) {
	body := validation.ValidateCreateOrderRequest(ctx)
	if body == nil {
		return
	}

	// Extract UserID from JWT claims stored in context
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var order models.Order
	order.UserID = uint(userId.(float64)) // Ensure the type assertion is correct based on how userID is stored
	order.Status = models.Created         // Set default status to "created"
	order.Items = make([]models.OrderItem, len(body.Items))

	// Fetch product prices and populate order items
	for i, item := range body.Items {
		var product models.Product
		if err := database.DB.First(&product, item.ProductId).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product not found", "productId": item.ProductId})
			return
		}
		order.Items[i] = models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Amount,
			Price:     product.Price, // Set price from the product
		}
	}

	// Save the order to the database
	if err := database.DB.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (ctrl *OrderController) CheckoutOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, orderId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Simulate payment
	if ctrl.SimulatePayment() {
		order.Status = models.Confirmed
	} else {
		order.Status = models.Cancelled
	}

	// Update the order status in the database
	if err := database.DB.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	// Respond to the client
	ctx.JSON(http.StatusOK, gin.H{"message": "Checkout processed", "orderStatus": order.Status})

	// If payment is successful, schedule status change to "delivered" after 10 seconds
	if order.Status == models.Confirmed {
		go func(orderId uint) {
			time.Sleep(10 * time.Second)
			var order models.Order
			if err := database.DB.First(&order, orderId).Error; err != nil {
				log.Printf("Failed to find order: %v", err)
				return
			}
			order.Status = models.Delivered
			if err := database.DB.Save(&order).Error; err != nil {
				log.Printf("Failed to update order status to delivered: %v", err)
			}
		}(order.ID)
	}
}

func (ctrl *OrderController) SimulatePayment() bool {
	return rand.Intn(10) < 8
}
