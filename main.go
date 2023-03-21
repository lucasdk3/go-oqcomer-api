package main

import (
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/docs"
	"github.com/lucasdk3/go-oqcomer-api/server"
	"github.com/lucasdk3/go-oqcomer-api/server/storage"
)

// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	database.StartDB()
	storage.StartStorage()
	s := server.NewServer()

	s.Run()
	// log.SetFlags(log.LstdFlags)
}
