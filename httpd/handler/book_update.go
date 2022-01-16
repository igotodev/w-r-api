package handler

import (
	"w-r-api/platform/bookvalidator"
	"w-r-api/platform/db"
	"w-r-api/platform/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		var book entity.Book

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

		if err := bookvalidator.IsValid(book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": err.Error(),
			})

			return
		}

		db.Update(book, idStr)

		c.JSON(200, gin.H{
			"error":   false,
			"message": "",
		})
	}
}
