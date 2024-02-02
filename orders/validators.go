// OrderModelValidator.go
package orders

import (
	"colls.labs.bake/database"
	"colls.labs.bake/items"
	"github.com/gin-gonic/gin"
)

type OrderModelValidator struct {
	Order struct {
		Client struct {
			ClientID string `json:"clientID" binding:"required" gorm:"size:2048"`
		} `json:"client" binding:"required"`
		SelectedItems []struct {
			ItemID   int    `json:"itemID" binding:"required" gorm:"size:2048"`
			Quantity int    `json:"quantity" binding:"required"`
			Price    int    `json:"price" binding:"required"`
			Name     string `json:"name" binding:"required"`
		} `json:"selectedItems" binding:"required"`
	} `json:"order"`
	OrderModel OrderModel `json:"-"`
}

func (v *OrderModelValidator) Bind(c *gin.Context) error {
	err := database.Bind(c, v)
	if err != nil {
		return err
	}

	// Extracting values from the new structure
	v.OrderModel.ClientID = v.Order.Client.ClientID
	v.OrderModel.Price = v.calculateTotalPrice()

	for _, item := range v.Order.SelectedItems {
		orderItem := OrderItemModel{
			Item: items.ItemModel{
				// Assuming you want to map 'Name' and 'Price' from ItemModel
				ID:    uint(item.ItemID),
				Name:  item.Name,
				Price: item.Price,
			},
			Quantity: item.Quantity,
			Price:    item.Price,
		}
		v.OrderModel.SelectedItems = append(v.OrderModel.SelectedItems, orderItem)
	}

	// Other assignments...

	return nil
}

func (v *OrderModelValidator) calculateTotalPrice() int {
	totalPrice := 0
	for _, item := range v.Order.SelectedItems {
		totalPrice += item.Price
	}
	return totalPrice
}

func NewOrderModelValidator() OrderModelValidator {
	return OrderModelValidator{}
}

func NewOrderModelValidatorFillWith(orderModel OrderModel) OrderModelValidator {
	return OrderModelValidator{
		OrderModel: orderModel,
	}
}
