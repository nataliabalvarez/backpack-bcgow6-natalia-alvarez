package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/internal/products"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/pkg/web"
)

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

type request struct {
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		prods, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prods, ""))
	}
}

func (c *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//verificar token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		prod, err := c.service.Get(int(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))

	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
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

		if req.Name == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Nombre es requerido"})
			return
		}
		if req.Color == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Color es requerido"})
			return
		}
		if req.Price == 0 {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Precio es requerido"})
			return
		}
		if req.Stock == 0 {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Stock es requerido"})
			return
		}
		if req.Code == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Codigo es requerido"})
			return
		}
		if !req.Published {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El campo Publicado es requerido"})
			return
		}
		if req.CreationDate == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La Fecha de Creacion es requerido"})
			return
		}

		prod, err := c.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
	}
}

func (c *Product) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err = strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		if req.Name == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Nombre es requerido"})
			return
		}
		if req.Color == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Color es requerido"})
			return
		}
		if req.Price == 0 {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Precio es requerido"})
			return
		}
		if req.Stock == 0 {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Stock es requerido"})
			return
		}
		if req.Code == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El Codigo es requerido"})
			return
		}
		if !req.Published {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El campo Publicado es requerido"})
			return
		}
		if req.CreationDate == "" {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "La Fecha de Creacion es requerido"})
			return
		}

		product, err := c.service.Put(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (c *Product) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := c.service.Patch(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//verificar token
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
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

		if err := c.service.Delete(int(id)); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusOK, nil, ""))
	}
}
