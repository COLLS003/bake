package items

import (
	"colls.labs.bake/database"
	"gorm.io/gorm"
)

type ItemModel struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"size:2048"`
	Price int    `gorm:"size:2048"`
}

// Migrate the schema to the database if needed
func AutoMigrate() {
	db := database.GetConnection()
	db.AutoMigrate(&ItemModel{})
}

// FindSingleItem finds a single Item based on the provided condition
func FindSingleItem(condition interface{}) (ItemModel, error) {
	db := database.GetConnection()
	var model ItemModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// SaveSingleItem saves a single Item to the database
func SaveSingleItem(data interface{}) error {
	db := database.GetConnection()
	err := db.Save(data).Error
	return err
}

// UpdateSingleItem updates a Item with new data
func UpdateSingleItem(model *ItemModel, data interface{}) error {
	db := database.GetConnection()
	err := db.Model(model).Updates(data).Error
	return err
}

// DeleteSingleItem deletes a Item from the database
func DeleteSingleItem(model *ItemModel) error {
	db := database.GetConnection()
	err := db.Delete(model).Error
	return err
}

// GetAllItems gets all Items from the database
func GetAllItems() ([]ItemModel, error) {
	db := database.GetConnection()
	var models []ItemModel
	err := db.Find(&models).Error
	return models, err
}

// fix codwa
func GetItemByID(id uint) (ItemModel, error) {
	db := database.GetConnection()
	var Item ItemModel
	err := db.First(&Item, id).Error
	return Item, err
}

// get Item by  email
