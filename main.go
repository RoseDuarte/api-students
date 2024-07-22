package main

import (
	"api-student/api"
	
	"github.com/rs/zerolog/log"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server")
	}
}
