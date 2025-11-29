package main

import (
	_ "github.com/lib/pq"
	"log"
	"microblog-api/config"
	_ "microblog-api/docs"
	"microblog-api/server"
	"os"
)

// @title Microblog-api
// @version 1.0
// @description This is simple microblog api LOL

//-- @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
