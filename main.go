package main

import (
	"w-r-api/httpd/handler"
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
)

const address string = "0.0.0.0:8080"

func main() {
	db.OpenDB()
	defer db.CloseDB()

	r := gin.Default()

	r.Use(gin.Logger())

	gr := r.Group("/api/v1/book")
	{
		gr.GET("/", handler.GetAllBooks())
		gr.GET("/:id", handler.GetBook())
		gr.POST("/", handler.PostBook())
		gr.PUT("/:id", handler.UpdateBook())
		gr.DELETE("/:id", handler.DeleteBook())
	}

	r.Run(address)

}
