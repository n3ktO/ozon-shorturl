package main

import (
	"shorturl/src/database"
	"shorturl/src/server"
)

func main() {

	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	defer dbInstance.DB.Close()

	serverInstance := &server.Server{DBInstance: dbInstance}
	serverInstance.Run()
}
