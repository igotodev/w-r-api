package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"w-r-api/config"
	"w-r-api/internal/controllers/api/handler"
	"w-r-api/internal/controllers/db"
	"w-r-api/internal/domain/book/service"
	"w-r-api/pkg/postgres"
)

func main() {
	var address string

	if len(os.Args) < 2 {
		address = "0.0.0.0:8080"
	} else {
		address = os.Args[1]
	}

	psqlCfg := config.InitConfig()
	myDB, err := postgres.OpenDB(psqlCfg.User, psqlCfg.Password, psqlCfg.Host, psqlCfg.Port, psqlCfg.Name)
	if err != nil {
		log.Println(err)
	}
	defer myDB.Close()

	dbStorage := db.NewStorage(myDB)
	bookService := service.NewService(dbStorage)

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gr := r.Group("/api/v1/book")
	{
		gr.GET("/", handler.GetAllBooks(bookService))
		gr.GET("/:id", handler.GetBook(bookService))
		gr.POST("/", handler.PostBook(bookService))
		gr.PUT("/:id", handler.UpdateBook(bookService))
		gr.DELETE("/:id", handler.DeleteBook(bookService))
	}

	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}

}
