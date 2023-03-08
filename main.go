package main

import (
	"github.com/lucasdk3/maui-oqcomer-api/database"
	"github.com/lucasdk3/maui-oqcomer-api/docs"
	"github.com/lucasdk3/maui-oqcomer-api/server"
)

// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
