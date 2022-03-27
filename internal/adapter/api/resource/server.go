package resource

import (
	ports "gas-inventory-service/internal/ports/resource"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	InventoryService ports.InventoryService
	Route            *gin.Engine
}

func NewHTTPHandler(inventoryService ports.InventoryService, router *gin.Engine) *HTTPHandler {
	handler := &HTTPHandler{
		InventoryService: inventoryService,
		Route:            router,
	}
	return handler
}
