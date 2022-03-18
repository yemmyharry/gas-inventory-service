package resource

import (
	ports "gas-inventory-service/internal/ports/resource"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Inventory ports.Inventory
	Route     *gin.Engine
}
