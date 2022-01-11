package handler

import (
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
)

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		if db.Delete(idStr) {
			c.JSON(200, gin.H{
				"error":   false,
				"message": "",
			})
		} else {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid id",
			})
		}

	}
}
