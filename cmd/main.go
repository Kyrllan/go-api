package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada Usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("product", ProductController.CreateProduct)

	server.Run(":8000")

}
