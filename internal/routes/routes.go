package routes

import (
	"github.com/aberyotaro/sample_api/internal/di"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, h *di.Handlers) {
	v1 := e.Group("/v1")
	v1.GET("/users/:id", h.UserHandler.GetUserByID)
}
