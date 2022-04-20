package main

import (
	profile "github.com/lvornholt/go-profiles"
	"hello-fresh-menu-planning-service/internal/adaptors/rest"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"log"
)

// @title HelloFresh Australia MenuPlan RESTFul Apis
// @contact.name Vinod Jagwani
// @contact.email vkj.agwani86@gmail.com
func main() {
	postgres.InitPostgresDB()
	router := rest.SetupRouter()
	err := router.Run(":" + profile.GetStringValue("server.port"))
	if err != nil {
		log.Fatalln(err)
	}
}
