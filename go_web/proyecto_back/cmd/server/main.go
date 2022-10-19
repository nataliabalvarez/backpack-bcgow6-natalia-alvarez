package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/cmd/server/handler"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/internal/products"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/pkg/store"
)

func main() {
	_ = godotenv.Load()
	db := store.NewStore(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	handler := handler.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/", handler.Store())
	pr.GET("/", handler.GetAll())
	pr.GET("/:id", handler.Get())
	// pr.PUT("/:id", handler.Put()) falta implementar con json
	// pr.PATCH("/:id", handler.Patch()) falta implementar con json
	// pr.DELETE("/:id", handler.Delete()) falta implementar con json

	router.Run()
}
