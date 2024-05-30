package main

import (
	database "goAPI/Database"
	router "goAPI/Router"
)

func init() {
	database.LoadEnvVariables()
	database.DBConnect()
	database.SyncDB()
}

func main() {
	router.Routes()
}
