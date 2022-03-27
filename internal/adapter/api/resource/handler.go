package resource

import (
	"gas-inventory-service/internal/core/domain/resource"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

func (s *HTTPHandler) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := uuid.New().String()
		i.Reference = reference
		i.CreatedAt = time.Now()
		err := s.InventoryService.CreateItem(i)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": i.Reference,
		})
	}
}

func (s *HTTPHandler) UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.UpdateItem(reference, i)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": reference,
		})
	}
}

func (s *HTTPHandler) DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.DeleteItem(reference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": reference,
			"message":   "item deleted",
		})
	}
}

func (s *HTTPHandler) DeleteAllUserItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReference := c.Param(strings.TrimSpace("user_reference"))
		err := s.InventoryService.DeleteAllUserItems(userReference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "all items deleted",
		})
	}
}

func (s *HTTPHandler) DeleteItemSoft() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.SoftDeleteItem(reference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": reference,
			"message":   "item soft deleted",
		})
	}
}

func (s *HTTPHandler) RestoreItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.RestoreItem(reference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": reference,
			"message":   "item restored",
		})
	}
}

func (s *HTTPHandler) SoftDeleteAllUserItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReference := c.Param(strings.TrimSpace("user-reference"))
		err := s.InventoryService.SoftDeleteAllUserItems(userReference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "all items soft deleted",
		})
	}
}

func (s *HTTPHandler) CheckItemAvailability() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		available := c.Param(strings.TrimSpace("available"))
		err := s.InventoryService.CheckItemAvailability(reference, available)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"reference": reference,
		})
	}
}

func (s *HTTPHandler) AddDocument() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.AddDocuments(reference, i.DocumentInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"reference": reference,
		})
	}
}

func (s *HTTPHandler) ValidateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.InventoryService.ValidateItem(reference, i.ValidationInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"reference": reference,
		})
	}
}

func (s *HTTPHandler) GetItemDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		i, err := s.InventoryService.GetItemDetail(reference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *HTTPHandler) GetItemList() gin.HandlerFunc {
	return func(c *gin.Context) {
		search_text := c.Param(strings.TrimSpace("search-text"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.InventoryService.GetItemsByMultipleSearch(search_text, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *HTTPHandler) GetItemListByOrgRef() gin.HandlerFunc {
	return func(c *gin.Context) {
		search_text := c.Param(strings.TrimSpace("search-text"))
		orgRef := c.Param(strings.TrimSpace("organization-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.InventoryService.GetItemsByOrganisationSearch(orgRef, search_text, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *HTTPHandler) GetItemListByCatRefAndOrgRef() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchText := c.Param(strings.TrimSpace("search-text"))
		orgRef := c.Param(strings.TrimSpace("organization-reference"))
		catRef := c.Param(strings.TrimSpace("category-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.InventoryService.GetItemsByCategoryAndOrganisationSearch(catRef, orgRef, searchText, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *HTTPHandler) GetItemListByCityReference() gin.HandlerFunc {
	return func(c *gin.Context) {
		search_text := c.Param(strings.TrimSpace("search-text"))
		cityRef := c.Param(strings.TrimSpace("state-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.InventoryService.GetItemsByStateSearch(cityRef, search_text, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *HTTPHandler) GetItemListByCatRefAndCityRef() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchText := c.Param(strings.TrimSpace("search-text"))
		cityRef := c.Param(strings.TrimSpace("state-reference"))
		catRef := c.Param(strings.TrimSpace("category-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.InventoryService.GetItemsByCategoryAndStateSearch(catRef, cityRef, searchText, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}
