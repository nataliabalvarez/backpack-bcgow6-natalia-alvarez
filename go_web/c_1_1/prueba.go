package main

import (
	"github.com/gin-gonic/gin"
	structs "github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/structs"
)

func main(){
	router := gin.Default()

	//------------------- ejercicio 2
	router.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola, Nat!",
		})
	})

	//------------------- ejercicio 3
	router.GET("/productos", GetAllHandler)

	router.Run()
}

func GetAllHandler(ctx *gin.Context) {

	var prods structs.Products

	prods.ReadJson("../jsons/productos.json")
	
	ctx.JSON(200, prods)
}
