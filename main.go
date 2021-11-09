package main

import (
	"example/webapi/database"
	"example/webapi/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
