package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/internal/products"
)

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product {
		service: s,
	} 
}

type request struct {
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		token := ctx.GetHeader("token")
		if token != "autorizedtoken123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		prods, err := c.service.GetAll()
		if err!= nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, prods)
	}
}

func (c *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//verificar token
		token := ctx.GetHeader("token")
		if token != "autorizedtoken123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		prod, err := c.service.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		token := ctx.GetHeader("token")
		if token != "autorizedtoken123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		
		
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		prod, err := c.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err!= nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, prod)

	}
}
