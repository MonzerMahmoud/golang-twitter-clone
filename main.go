package main

import (
	"golang-twitter-clone/api"
	"golang-twitter-clone/database"
	"golang-twitter-clone/migrations"
	log "github.com/sirupsen/logrus"
)

func main() {
	
	database.InitDatabase()
	migrations.Migrate()
	api.InitializeRouter()
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Server started")

}












