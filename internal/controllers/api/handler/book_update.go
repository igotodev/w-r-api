package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"w-r-api/internal/controllers/api"
	"w-r-api/internal/domain/book/model"
	"w-r-api/internal/domain/book/validator"
)

func UpdateBook(service api.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		var book model.Book

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

		if err := validator.IsValid(book); err != nil {
			c.JSON(422, gin.H{
				"error":   true,
				"message": err.Error(),
			})

			return
		}

		service.Update(&book, idStr)

		c.JSON(200, gin.H{
			"error":   false,
			"message": "",
		})
	}
}
