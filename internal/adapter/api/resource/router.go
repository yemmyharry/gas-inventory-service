package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/yemmyharry/gas-inventory-service/internal/core/helper"
	"github.com/yemmyharry/gas-inventory-service/internal/core/shared"
)

func (s *Server) Routes(router *gin.Engine) {
	apirouter := router.Group("api")
	apirouter.POST("/inventory/items", s.CreateItem())
	apirouter.PUT("/inventory/items/:reference", s.UpdateItem())
	apirouter.PUT("/inventory/items/:reference/availability/:available", s.CheckItemAvailability())
	apirouter.DELETE("/inventory/items/:reference", s.DeleteItem())
	apirouter.DELETE("/inventory/items/user/:user_reference", s.DeleteAllUserItems())
	apirouter.PUT("/inventory/items/documents/:reference", s.AddDocument())
	apirouter.PUT("/inventory/items/validate-item/:reference", s.ValidateItem())
	apirouter.GET("/inventory/items/:reference", s.GetItemDetail())
	router.NoRoute(func(c *gin.Context) { c.JSON(404, helper.PrintErrorMessage("404", shared.NoResourceFound)) })
}
