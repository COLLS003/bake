package items

import (
	"github.com/gin-gonic/gin"
)

// ItemSerializer represents a serializer for a single ItemModel.
type ItemSerializer struct {
	C         *gin.Context
	ItemModel ItemModel
}

// ItemsSerializer represents a serializer for a slice of ItemModel.
type ItemsSerializer struct {
	C     *gin.Context
	Items []ItemModel
}

// ItemResponse represents the response structure for a single ItemModel.
type ItemResponse struct {
	ItemModel
}

// NewItemSerializer creates a new ItemSerializer instance for a single ItemModel.
func NewItemSerializer(c *gin.Context, Item ItemModel) ItemSerializer {
	return ItemSerializer{
		C:         c,
		ItemModel: Item,
	}
}

// NewItemsSerializer creates a new ItemsSerializer instance for a slice of ItemModel.
func NewItemsSerializer(c *gin.Context, Items []ItemModel) ItemsSerializer {
	return ItemsSerializer{
		C:     c,
		Items: Items,
	}
}

// Response returns a ItemResponse for a single ItemModel.
func (s *ItemSerializer) Response() ItemResponse {
	return ItemResponse{
		ItemModel: s.ItemModel,
	}
}

// Response returns a slice of ItemResponse for a slice of ItemModel.
func (s *ItemsSerializer) Response() []ItemResponse {
	response := make([]ItemResponse, len(s.Items))
	for i, ItemModel := range s.Items {
		serializer := NewItemSerializer(s.C, ItemModel)
		response[i] = serializer.Response()
	}
	return response
}
