package main

import (
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/cmd/server/routes"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/storage_implementation/pkg/db"
)

func main() {

	engine, db := db.ConnectDatabase()

	router := routes.NewRouter(engine, db)
	
	router.MapRoutes()

	engine.Run(":8080")
}
