package handler

import (
	"github.com/gin-gonic/gin"
	"w-r-api/internal/controllers/api"
)

func DeleteBook(service api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		if err := service.Delete(idStr); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid id",
			})
		} else {
			c.JSON(200, gin.H{
				"error":   false,
				"message": "",
			})
		}

	}
}
