package main

import (
	database "goAPI/database"
	"goAPI/migrate"
	router "goAPI/router"
)

func init() {
	database.LoadEnvVariables()
	database.DBConnect()
	migrate.SyncDB()
}

func main() {
	router.Routes()
}
