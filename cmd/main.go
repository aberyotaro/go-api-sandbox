package main

import (
	"log"

	"github.com/aberyotaro/go-api-sandbox/internal/db"
	"github.com/aberyotaro/go-api-sandbox/internal/di"
	"github.com/aberyotaro/go-api-sandbox/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB初期化
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize DB: %v", err)
	}
	defer dbConn.Close()

	g := gin.Default()
	routes.SetupRoutes(g, di.NewHandlers(dbConn))
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
