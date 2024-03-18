package main

import (
	"log"

	"github.com/jrandyl/adon-webserver/server"
	"github.com/jrandyl/adon-webserver/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}

	server, err := server.New()
	if err != nil {
		log.Fatal("cannot create a new server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
