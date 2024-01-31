package items

// package Items

import (
	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

type ItemModelValidator struct {
	Item struct {
		Name  string `form:"name" json:"name" binding:"required" gorm:"size:2048"`
		Price string `form:"price" json:"price" binding:"required" gorm:"size:2048"`
	} `json:"Item"`
	ItemModel ItemModel `json:"-"`
}

// Bind binds the request data to the ItemModelValidator.
func (v *ItemModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	v.ItemModel.Name = v.Item.Name
	v.ItemModel.Price = v.Item.Price

	return nil
}

// NewItemModelValidator creates a new ItemModelValidator instance.
func NewItemModelValidator() ItemModelValidator {
	return ItemModelValidator{}
}

// NewItemModelValidatorFillWith creates a new ItemModelValidator instance and fills it with Item model data.
func NewItemModelValidatorFillWith(ItemModel ItemModel) ItemModelValidator {
	return ItemModelValidator{
		Item: struct {
			Name  string `form:"name" json:"name" binding:"required" gorm:"size:2048"`
			Price string `form:"price" json:"price" binding:"required" gorm:"size:2048"`
		}{
			Name:  ItemModel.Name,
			Price: ItemModel.Price,
		},
	}
}

// LoginValidator represents a validator for Item login.
func sendDate() bool {
	return true
}

// You can put the default value of a Validator here
