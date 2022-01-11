package handler

import (
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
)

func GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, db.SelectAll())
	}
}

func GetBook() gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")

		book := db.Select(idStr)

		c.JSON(200, book)
	}
}
