package main

import (
	"golang-twitter-clone/api"
	"golang-twitter-clone/migrations"
)

func main() {
	//router.InitializeRouter()
	//migrations.Migrate()
	api.InitializeRouter()
	migrations.Migrate()
}












