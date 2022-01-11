package handler

import (
	"w-r-api/platform"
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book platform.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": "invalid json",
			})

			return
		}

		//user cannot set id it is generated automatically
		book.Id = uuid.New().String()

		db.Insert(book)

		c.JSON(201, gin.H{
			"error":   false,
			"message": "",
		})
	}
}
