package main

import (
	"example.com/go-api/db"
	"example.com/go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server) // we can pass the "server" instance because gin returns a pointer to that server

	server.Run(":8080") // Start the server on port 8080

}
