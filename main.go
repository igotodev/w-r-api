package main

import (
	"log"
	"os"
	"w-r-api/httpd/handler"
	"w-r-api/platform/db"

	"github.com/gin-gonic/gin"
)

func main() {
	var address string

	if len(os.Args) < 2 {
		address = "0.0.0.0:8080"
	} else {
		address = os.Args[1]
	}

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

	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}

}
