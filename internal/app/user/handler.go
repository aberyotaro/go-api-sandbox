package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aberyotaro/sample_api/internal/entity"
	"github.com/aberyotaro/sample_api/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase *usecase.User
}

func NewHandler(uc *usecase.User) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("message: %s", err.Error())})
	}

	return c.JSON(
		http.StatusOK,
		&Response{
			StatusCode: http.StatusOK,
			User:       user,
		},
	)
}

type Response struct {
	StatusCode int          `json:"status_code"`
	User       *entity.User `json:"user"`
}
