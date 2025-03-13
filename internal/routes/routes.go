package routes

import (
	"github.com/aberyotaro/go-api-sandbox/internal/di"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine, h *di.Handlers) {
	v1 := g.Group("/v1")
	v1.POST("/users/:id", h.UserHandler.GetUserByID)
	v1.POST("/translate", h.TranslateHandler.Translate)
}
