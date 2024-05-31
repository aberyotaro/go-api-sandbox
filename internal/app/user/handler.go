package user

import (
	"net/http"
	"strconv"

	"github.com/aberyotaro/sample_api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.User
}

func NewHandler(uc *usecase.User) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "success",
	})
}
