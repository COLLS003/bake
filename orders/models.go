package orders

import (
	"colls.labs.bake/database"
	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	ID              uint   `gorm:"primary_key"`
	ClientID        string `gorm:"size:2048"`
	ItemID          string `gorm:"size:2048"`
	Quantity        string `gorm:"size:2048"`
	DueDate         string `gorm:"size:2048"`
	DateInitialized string `gorm:"size:2048"`
	Price           string `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&OrderModel{})
}

// FindSingleOrder finds a single Order based on the provided condition
func FindSingleOrder(condition interface{}) (OrderModel, error) {
	db := database.GetConnection()
	var model OrderModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleOrder saves a single Order to the database
func SaveSingleOrder(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

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
func GetAllOrders() ([]OrderModel, error) {
	db := database.GetConnection()
	var models []OrderModel
	err := db.Find(&models).Error
	return models, err
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
