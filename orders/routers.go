package orders

import (
	"fmt"
	"net/http"
	"strconv"

	"colls.labs.bake/database"
	"github.com/gin-gonic/gin"
)

func Create(router *gin.RouterGroup) {
	router.POST("/create", CreateOrder)
	router.GET("/read/:id", ReadSingleOrder)
	router.GET("/read/Initiator/:Initiator", ReadOrderUsingInitiatorID)
	router.PUT("/update/:id", UpdateOrder)
	router.DELETE("/delete/:id", DeleteOrder)
	router.GET("/list", OrdersList)
}

func CreateOrder(c *gin.Context) {
	modelValidator := NewOrderModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	if err := SaveSingleOrder(&modelValidator.OrderModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.Set("my_Order_model", modelValidator.OrderModel)
	serializer := NewOrderSerializer(c, modelValidator.OrderModel)
	c.JSON(http.StatusCreated, gin.H{"Order": serializer.Response()})

	fmt.Println("Order saved ...")
}

func ReadSingleOrder(c *gin.Context) {
	OrderID := c.Param("id")
	OrderIDUint, err := strconv.ParseUint(OrderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}

	OrderModel, err := GetOrderByID(uint(OrderIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Order", err))
		return
	}

	serializer := NewOrderSerializer(c, OrderModel)
	c.JSON(http.StatusOK, gin.H{"Order": serializer.Response()})
}

func UpdateOrder(c *gin.Context) {
	OrderID := c.Param("id")
	OrderIDUint, err := strconv.ParseUint(OrderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}
	OrderModel, err := GetOrderByID(uint(OrderIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Order", err))
		return
	}

	// Bind and update OrderModel with new data
	modelValidator := NewOrderModelValidatorFillWith(OrderModel)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewValidatorError(err))
		return
	}

	// Call UpdateSingleOrder function with the OrderModel and updated data
	if err := UpdateSingleOrder(&OrderModel, modelValidator.OrderModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	serializer := NewOrderSerializer(c, OrderModel)
	c.JSON(http.StatusOK, gin.H{"Order": serializer.Response()})
}

// update sent Orders
func DeleteOrder(c *gin.Context) {
	OrderID := c.Param("id")
	OrderIDUint, err := strconv.ParseUint(OrderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}
	OrderModel, err := GetOrderByID(uint(OrderIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Order", err))
		return
	}

	if err := DeleteSingleOrder(&OrderModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, database.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"Order": "Order deleted successfully"})
}

func OrdersList(c *gin.Context) {
	OrdersModels, err := GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Orders", err))
		return
	}
	serializer := NewOrdersSerializer(c, OrdersModels)
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"Orders": response})
}

func ReadOrderUsingInitiatorID(c *gin.Context) {
	OrderID := c.Param("Initiator")
	OrderIDUint, err := strconv.ParseUint(OrderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Initiator  ID"})
		return
	}

	OrderModels, err := GetOrderByInitiatorID(uint(OrderIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, database.NewError("Order", err))
		return
	}

	if len(OrderModels) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Create a slice to store serialized Orders
	var serializedOrders []interface{}

	// Loop through the array of OrderModels and create serializers for each
	for _, org := range OrderModels {
		serializer := NewOrderSerializer(c, org)
		serializedOrders = append(serializedOrders, serializer.Response())
	}

	c.JSON(http.StatusOK, gin.H{"Orders": serializedOrders})
}

// aux functions

// get an Initiator scheduled Orders

// update the Order status

// read ready Orders
