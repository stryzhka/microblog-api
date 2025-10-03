package main

import (
	_ "github.com/lib/pq"
	"log"
	"microblog-api/config"
	"microblog-api/server"
	"os"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	app := server.NewApp()
	if err = app.Run(os.Getenv("port")); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
