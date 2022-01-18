package handler

import (
	"github.com/gin-gonic/gin"
	"w-r-api/internal/controllers/api"
)

func GetAllBooks(service api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, service.GetAll())
	}
}

func GetBook(service api.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")

		book := service.GetByID(idStr)

		if book.Id == "" {
			c.JSON(404, gin.H{
				"error":   true,
				"message": "invalid id",
			})
		} else {
			c.JSON(200, book)
		}
	}
}
