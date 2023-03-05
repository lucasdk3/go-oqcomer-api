package main

import (
	"github.com/lucasdk3/maui-oqcomer-api/database"
	_ "github.com/lucasdk3/maui-oqcomer-api/docs"
	"github.com/lucasdk3/maui-oqcomer-api/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
