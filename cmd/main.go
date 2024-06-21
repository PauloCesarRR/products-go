package main

import (
	"github.com/gin-gonic/gin"
	"products/controller"
	"products/db"
	"products/repository"
	"products/usecase"
)

func main() {
	server := gin.Default()

	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepo := repository.NewProductRepository(dbConn)
	ProductUseCase := usecase.NewProductUsecase(ProductRepo)
	ProductControler := controller.NewProductController(ProductUseCase)

	server.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductControler.GetProductList)
	server.POST("/product", ProductControler.CreateProduct)
	server.GET("/product/{id}", ProductControler.GetProduct)

	server.Run(":8080")
}
