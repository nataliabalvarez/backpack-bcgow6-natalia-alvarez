package main

import(
	"github.com/gin-gonic/gin"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/cmd/server/handler"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/internal/products"
)
func main(){
	repo := products.NewRepository()
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/", handler.Store())
	pr.GET("/", handler.GetAll())

	router.Run()

}