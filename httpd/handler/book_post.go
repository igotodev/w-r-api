package handler

import (
	"w-r-api/platform/bookvalidator"
	"w-r-api/platform/db"
	"w-r-api/platform/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book entity.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid json",
			})

			return
		}

		//user cannot set id it is generated automatically
		book.Id = uuid.New().String()

		if err := bookvalidator.IsValid(book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": err.Error(),
			})

			return
		}

		db.Insert(book)

		c.JSON(201, gin.H{
			"error":   false,
			"message": "",
		})
	}
}
