package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yemmyharry/gas-inventory-service/internal/core/domain/resource"
	"strconv"
	"strings"
	"time"
)

func (s *Server) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := uuid.New().String()
		i.Reference = reference
		i.CreatedAt = time.Now()
		err := s.Inventory.CreateItem(i)
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

func (s *Server) UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.UpdateItem(reference, i)
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

func (s *Server) DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.DeleteItem(reference)
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

func (s *Server) DeleteAllUserItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReference := c.Param(strings.TrimSpace("user_reference"))
		err := s.Inventory.DeleteAllUserItems(userReference)
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

func (s *Server) DeleteItemSoft() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.SoftDeleteItem(reference)
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

func (s *Server) RestoreItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.RestoreItem(reference)
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

func (s *Server) SoftDeleteAllUserItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReference := c.Param(strings.TrimSpace("user-reference"))
		err := s.Inventory.SoftDeleteAllUserItems(userReference)
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

func (s *Server) CheckItemAvailability() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		available := c.Param(strings.TrimSpace("available"))
		err := s.Inventory.CheckItemAvailability(reference, available)
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

func (s *Server) AddDocument() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.AddDocuments(reference, i.DocumentInfo)
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

func (s *Server) ValidateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := &resource.Inventory{}
		c.ShouldBindJSON(i)
		reference := c.Param(strings.TrimSpace("reference"))
		err := s.Inventory.ValidateItem(reference, i.ValidationInfo)
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

func (s *Server) GetItemDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		reference := c.Param(strings.TrimSpace("reference"))
		i, err := s.Inventory.GetItemDetail(reference)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *Server) GetItemList() gin.HandlerFunc {
	return func(c *gin.Context) {
		search_text := c.Param(strings.TrimSpace("search-text"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.Inventory.GetItemsByMultipleSearch(search_text, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *Server) GetItemListByOrgRef() gin.HandlerFunc {
	return func(c *gin.Context) {
		search_text := c.Param(strings.TrimSpace("search-text"))
		orgRef := c.Param(strings.TrimSpace("organization-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.Inventory.GetItemsByOrganisationSearch(orgRef, search_text, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}

func (s *Server) GetItemListByCatRefAndOrgRef() gin.HandlerFunc {
	return func(c *gin.Context) {
		searchText := c.Param(strings.TrimSpace("search-text"))
		orgRef := c.Param(strings.TrimSpace("organization-reference"))
		catRef := c.Param(strings.TrimSpace("category-reference"))
		page, _ := strconv.Atoi(c.Param(strings.TrimSpace("page")))
		i, err := s.Inventory.GetItemsByCategoryAndOrganisationSearch(catRef, orgRef, searchText, page)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, i)
	}
}
