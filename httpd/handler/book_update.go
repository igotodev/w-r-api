package handler

import (
	"w-r-api/platform"
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		var book platform.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid json",
			})

			return
		}

		//user cannot update id
		//new id is generated automatically
		book.Id = uuid.New().String()

		db.Update(book, idStr)

		c.JSON(200, gin.H{
			"error":   false,
			"message": "",
		})
	}
}
