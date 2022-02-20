package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yemmyharry/gas-inventory-service/internal/core/domain/resource"
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
			"message":   "item availability is " + available,
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
