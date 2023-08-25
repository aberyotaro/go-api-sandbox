package main

import (
	"github.com/aberyotaro/sample_api/internal/di"
	"github.com/aberyotaro/sample_api/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.SetupRoutes(e, di.NewHandlers())
	e.Logger.Fatal(e.Start(":8080"))
}
