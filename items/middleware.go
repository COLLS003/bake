package items

import (
	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

// A helper to write Item_id and Item_model to the context
func UpdateContextItemModel(c *gin.Context, my_Item_id uint) {
	var myItemModel ItemModel
	if my_Item_id != 0 {
		db := database.GetConnection()
		db.First(&myItemModel, my_Item_id)
	}
	c.Set("my_Item_id", my_Item_id)
	c.Set("my_Item_model", myItemModel)
}
