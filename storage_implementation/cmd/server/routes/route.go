package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/cmd/server/handler"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/internal/movie"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

// llamado por el main
// no defino el grupo aca
func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r: r, 
		db: db,
	}
}

// llamado por el main
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildMovieRoutes()
}

// defino el grupo
func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildMovieRoutes() {

	repo := movie.NewRepository(r.db)
	service := movie.NewService(repo)
	handler := handler.NewMovie(service)

	r.rg.GET("/movies", handler.GetAll())
	r.rg.GET("/movies/:id", handler.Get())
	r.rg.POST("/movies", handler.Create())
	r.rg.DELETE("/movies/:id", handler.Delete())
	r.rg.PATCH("/movies/:id", handler.Update())
}