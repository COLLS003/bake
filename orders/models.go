package orders

import (
	"fmt"
	"time"

	"colls.labs.bake/database"
	"colls.labs.bake/items"
	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	ID            uint             `gorm:"primary_key"`
	OrderID       string           `gorm:"size:2048"`
	ClientID      string           `gorm:"size:2048"`
	Price         int              `gorm:"size:2048"`
	SelectedItems []OrderItemModel `gorm:"foreignKey:OrderModelID"`
}

// OrderItemModel in models.go
type OrderItemModel struct {
	gorm.Model
	OrderModelID uint
	ItemID       uint            // Foreign key referencing ItemModel
	Item         items.ItemModel `gorm:"foreignKey:ItemID"`
	Quantity     int
	Price        int
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&OrderModel{}, &OrderItemModel{})
}

// FindSingleOrder finds a single Order based on the provided condition
// func SaveSingleOrder(data interface{}) error {
// 	db := database.GetConnection()
// 	orderModel := data.(*OrderModel)

// 	// Generate a unique 6-digit order ID based on user ID and timestamp
// 	orderID := generateOrderID(orderModel.ClientID)

// 	orderModel.OrderID = orderID

// 	// Save the order
// 	err := db.Save(orderModel).Error
// 	if err != nil {
// 		return err
// 	}

// 	if err := db.Save(&orderModel.SelectedItems).Error; err != nil {
// 		return err
// 	}
// 	// }

//		return nil
//	}
//
// SaveSingleOrder saves a single Order to the database
func SaveSingleOrder(data interface{}) error {
	db := database.GetConnection()
	orderModel := data.(*OrderModel)

	// Generate a unique 6-digit order ID based on user ID and timestamp
	orderID := generateOrderID(orderModel.ClientID)

	orderModel.OrderID = orderID

	// Save the order
	err := db.Save(orderModel).Error
	if err != nil {
		return err
	}

	// Save selected items
	for i := range orderModel.SelectedItems {
		orderModel.SelectedItems[i].OrderModelID = orderModel.ID
		err := db.Save(&orderModel.SelectedItems[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func generateOrderID(clientID string) string {
	// Use client ID and timestamp to generate a unique order ID
	// For simplicity, you can concatenate the client ID and the current timestamp
	// and take the last 6 digits as the order ID
	timestamp := time.Now().UnixNano()
	orderID := fmt.Sprintf("%s%d", clientID, timestamp)[len(clientID):][:10]
	return orderID
}

// SaveSingleOrder saves a single Order to the database
// func SaveSingleOrder(data interface{}) error {
// 	db := database.GetConnection()
// 	err := db.Save(data).Error
// 	return err
// }

// UpdateSingleOrder updates a Order with new data
func UpdateSingleOrder(model *OrderModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleOrder deletes a Order from the database
func DeleteSingleOrder(model *OrderModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllOrders gets all Orders from the database
// GetAllOrders gets all Orders from the database
func GetAllOrders() ([]OrderModel, error) {
	db := database.GetConnection()
	var models []OrderModel
	err := db.Preload("SelectedItems").Find(&models).Error
	fmt.Println("all orders", models)
	return models, err
}

// Add this function to your existing orders package
func GetOrderItemsByOrderID(orderID uint) ([]OrderItemModel, error) {
	db := database.GetConnection()
	var orderItems []OrderItemModel
	err := db.Where("order_model_id = ?", orderID).Find(&orderItems).Error
	return orderItems, err
}

// fix codwa
func GetOrderByID(id uint) (OrderModel, error) {
	db := database.GetConnection()
	var Order OrderModel
	err := db.First(&Order, id).Error
	return Order, err
}

// get Order by  email
func GetOrderByEmail(email string) (OrderModel, error) {
	db := database.GetConnection()
	var Order OrderModel
	err := db.Where("Email = ?", email).First(&Order).Error
	return Order, err
}

func GetOrderByInitiatorID(id uint) ([]OrderModel, error) {
	db := database.GetConnection()
	var Orders []OrderModel
	err := db.Where("Initiator = ?", id).Find(&Orders).Error
	return Orders, err
}
