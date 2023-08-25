package main

import (
	"github.com/aberyotaro/sample_api/internal/di"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v1 := e.Group("/v1")

	h := di.NewHandlers()

	v1.GET("/users/:id", h.UserHandler.GetUserByID)

	e.Start(":8080")
}
