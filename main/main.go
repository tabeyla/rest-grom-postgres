package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tabeyla/rest-grom-postgres/controllers"
	"github.com/tabeyla/rest-grom-postgres/models" // new
)

func main() {

	r := gin.Default()
	models.ConnectDatabase() // new

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books/:id", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
