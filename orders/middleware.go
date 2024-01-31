package orders

import (
	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

// A helper to write Order_id and Order_model to the context
func UpdateContextOrderModel(c *gin.Context, my_Order_id uint) {
	var myOrderModel OrderModel
	if my_Order_id != 0 {
		db := database.GetConnection()
		db.First(&myOrderModel, my_Order_id)
	}
	c.Set("my_Order_id", my_Order_id)
	c.Set("my_Order_model", myOrderModel)
}
