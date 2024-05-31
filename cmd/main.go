package main

import (
	"github.com/aberyotaro/sample_api/internal/di"
	"github.com/aberyotaro/sample_api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	routes.SetupRoutes(g, di.NewHandlers())
	err := g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
