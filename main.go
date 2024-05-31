package main

import (
	database "goAPI/database"
	router "goAPI/router"
)

func init() {
	database.LoadEnvVariables()
	database.DBConnect()
	database.SyncDB()
}

func main() {
	router.Routes()
}
