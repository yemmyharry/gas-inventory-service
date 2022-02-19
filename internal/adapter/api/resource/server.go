package resource

import (
	"github.com/gin-gonic/gin"
	ports "github.com/yemmyharry/gas-inventory-service/internal/ports/resource"
)

type Server struct {
	Inventory ports.Inventory
	Route     *gin.Engine
}
