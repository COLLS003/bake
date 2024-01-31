package items

import (
	"fmt"
	"net/http"
	"strconv"

	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreateItem)
	router.GET("/read/:id", ReadSingleItem)
	router.GET("/read/Initiator/:Initiator", ReadItemUsingInitiatorID)
	router.PUT("/update/:id", UpdateItem)
	router.DELETE("/delete/:id", DeleteItem)
	router.GET("/list", ItemsList)
}

func CreateItem(c *gin.Context) {
	modelValidator := NewItemModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleItem(&modelValidator.ItemModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Item_model", modelValidator.ItemModel)
	serializer := NewItemSerializer(c, modelValidator.ItemModel)
	c.JSON(http.StatusCreated, gin.H{"Item": serializer.Response()})

	fmt.Println("Item saved ...")
}

func ReadSingleItem(c *gin.Context) {
	ItemID := c.Param("id")
	ItemIDUint, err := strconv.ParseUint(ItemID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}

	ItemModel, err := GetItemByID(uint(ItemIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Item", err))
		return
	}

	serializer := NewItemSerializer(c, ItemModel)
	c.JSON(http.StatusOK, gin.H{"Item": serializer.Response()})
}

func UpdateItem(c *gin.Context) {
	ItemID := c.Param("id")
	ItemIDUint, err := strconv.ParseUint(ItemID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}
	ItemModel, err := GetItemByID(uint(ItemIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Item", err))
		return
	}

	// Bind and update ItemModel with new data
	modelValidator := NewItemModelValidatorFillWith(ItemModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleItem function with the ItemModel and updated data
	if err := UpdateSingleItem(&ItemModel, modelValidator.ItemModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewItemSerializer(c, ItemModel)
	c.JSON(http.StatusOK, gin.H{"Item": serializer.Response()})
}

// update sent Items
func DeleteItem(c *gin.Context) {
	ItemID := c.Param("id")
	ItemIDUint, err := strconv.ParseUint(ItemID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}
	ItemModel, err := GetItemByID(uint(ItemIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Item", err))
		return
	}

	if err := DeleteSingleItem(&ItemModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"Item": "Item deleted successfully"})
}

func ItemsList(c *gin.Context) {
	ItemsModels, err := GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Items", err))
		return
	}
	serializer := NewItemsSerializer(c, ItemsModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Items": response})
}

func ReadItemUsingInitiatorID(c *gin.Context) {
	ItemID := c.Param("Initiator")
	ItemIDUint, err := strconv.ParseUint(ItemID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Initiator  ID"})
		return
	}

	ItemModels, err := GetItemByInitiatorID(uint(ItemIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Item", err))
		return
	}

	if len(ItemModels) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Create a slice to store serialized Items
	var serializedItems []interface{}

	// Loop through the array of ItemModels and create serializers for each
	for _, org := range ItemModels {
		serializer := NewItemSerializer(c, org)
		serializedItems = append(serializedItems, serializer.Response())
	}

	c.JSON(http.StatusOK, gin.H{"Items": serializedItems})
}

// aux functions

// get an Initiator scheduled Items

// update the Item status

// read ready Items
