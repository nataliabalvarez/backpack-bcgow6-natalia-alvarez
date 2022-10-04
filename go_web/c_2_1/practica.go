package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	structs "github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/structs"
)

var products []structs.Product

func main() {
	router := gin.Default()

	group := router.Group("/products")

	group.POST("/", Save())

	router.Run()
}

func Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request structs.Product

		token := ctx.GetHeader("token")
		if token != "autorizedtoken123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		request.Id = len(products) + 1
		products = append(products, request)
		ctx.JSON(http.StatusOK, request)

	}
}
