package resource

import (
	"gas-inventory-service/internal/core/helper"
	"gas-inventory-service/internal/core/shared"
	"github.com/gin-gonic/gin"
)

func (s *Server) Routes(router *gin.Engine) {
	apirouter := router.Group("api")
	apirouter.POST("/inventory/items", s.CreateItem())
	apirouter.PUT("/inventory/items/:reference", s.UpdateItem())
	apirouter.PUT("/inventory/items/:reference/availability/:available", s.CheckItemAvailability())
	apirouter.DELETE("/inventory/items/:reference", s.DeleteItem())
	apirouter.DELETE("/inventory/items/user/:user_reference", s.DeleteAllUserItems())
	apirouter.PUT("/inventory/items/delete/:reference", s.DeleteItemSoft())
	apirouter.PUT("/inventory/items/restore/:reference", s.RestoreItem())
	apirouter.PUT("/inventory/items/user/delete/:user-reference", s.SoftDeleteAllUserItems())
	apirouter.PUT("/inventory/items/documents/:reference", s.AddDocument())
	apirouter.PUT("/inventory/items/validate-item/:reference", s.ValidateItem())
	apirouter.GET("/inventory/items/:reference", s.GetItemDetail())
	apirouter.GET("/inventory/items/list/:search-text/page/:page", s.GetItemList())
	apirouter.GET("/inventory/items/organisation/:organization-reference/:search-text/list/page/:page", s.GetItemListByOrgRef())
	apirouter.GET("/inventory/items/category/organisation/:category-reference/:organization-reference/:search-text/list/page/:page", s.GetItemListByCatRefAndOrgRef())
	apirouter.GET("/inventory/items/state/:state-reference/:search-text/list/page/:page", s.GetItemListByCityReference())
	apirouter.GET("/inventory/items/category/state/:category-reference/:state-reference/:search-text/list/page/:page", s.GetItemListByCatRefAndCityRef())
	router.NoRoute(func(c *gin.Context) { c.JSON(404, helper.PrintErrorMessage("404", shared.NoResourceFound)) })
}
