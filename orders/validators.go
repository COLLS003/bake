package orders

import (
	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

type OrderModelValidator struct {
	Order struct {
		ClientID        string `json:"client_id" binding:"required" gorm:"size:2048"`
		ItemID          string `json:"item_id" binding:"required" gorm:"size:2048"`
		Quantity        string `json:"quantity" binding:"required" gorm:"size:2048"`
		DueDate         string `json:"due_date" binding:"required" gorm:"size:2048"`
		DateInitialized string `json:"date_initialized" binding:"required" gorm:"size:2048"`
		Price           string `json:"price" binding:"required" gorm:"size:2048"`
	} `json:"order"`
	OrderModel OrderModel `json:"-"`
}

func (v *OrderModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.OrderModel.ClientID = v.Order.ClientID
	v.OrderModel.ItemID = v.Order.ItemID
	v.OrderModel.Quantity = v.Order.Quantity
	v.OrderModel.DueDate = v.Order.DueDate
	v.OrderModel.DateInitialized = v.Order.DateInitialized
	v.OrderModel.Price = v.Order.Price

	return nil
}

func NewOrderModelValidator() OrderModelValidator {
	return OrderModelValidator{}
}

func NewOrderModelValidatorFillWith(orderModel OrderModel) OrderModelValidator {
	return OrderModelValidator{
		Order: struct {
			ClientID        string `json:"client_id" binding:"required" gorm:"size:2048"`
			ItemID          string `json:"item_id" binding:"required" gorm:"size:2048"`
			Quantity        string `json:"quantity" binding:"required" gorm:"size:2048"`
			DueDate         string `json:"due_date" binding:"required" gorm:"size:2048"`
			DateInitialized string `json:"date_initialized" binding:"required" gorm:"size:2048"`
			Price           string `json:"price" binding:"required" gorm:"size:2048"`
		}{
			ClientID:        orderModel.ClientID,
			ItemID:          orderModel.ItemID,
			Quantity:        orderModel.Quantity,
			DueDate:         orderModel.DueDate,
			DateInitialized: orderModel.DateInitialized,
			Price:           orderModel.Price,
		},
	}
}
