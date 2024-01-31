package orders

import (
	"github.com/gin-gonic/gin"
)

// OrderSerializer represents a serializer for a single OrderModel.
type OrderSerializer struct {
	C         *gin.Context
	OrderModel OrderModel
}

// OrdersSerializer represents a serializer for a slice of OrderModel.
type OrdersSerializer struct {
	C     *gin.Context
	Orders []OrderModel
}

// OrderResponse represents the response structure for a single OrderModel.
type OrderResponse struct {
	OrderModel
}

// NewOrderSerializer creates a new OrderSerializer instance for a single OrderModel.
func NewOrderSerializer(c *gin.Context, Order OrderModel) OrderSerializer {
	return OrderSerializer{
		C:         c,
		OrderModel: Order,
	}
}

// NewOrdersSerializer creates a new OrdersSerializer instance for a slice of OrderModel.
func NewOrdersSerializer(c *gin.Context, Orders []OrderModel) OrdersSerializer {
	return OrdersSerializer{
		C:     c,
		Orders: Orders,
	}
}

// Response returns a OrderResponse for a single OrderModel.
func (s *OrderSerializer) Response() OrderResponse {
	return OrderResponse{
		OrderModel: s.OrderModel,
	}
}

// Response returns a slice of OrderResponse for a slice of OrderModel.
func (s *OrdersSerializer) Response() []OrderResponse {
	response := make([]OrderResponse, len(s.Orders))
	for i, OrderModel := range s.Orders {
		serializer := NewOrderSerializer(s.C, OrderModel)
		response[i] = serializer.Response()
	}
	return response
}
