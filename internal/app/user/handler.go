package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aberyotaro/sample_api/internal/entity"
	"github.com/aberyotaro/sample_api/internal/usecase"
)

type Handler struct {
	usecase *usecase.User
}

func NewHandler(uc *usecase.User) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
	}

	res := &Response{
		StatusCode: http.StatusOK,
		User:       user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

type Response struct {
	StatusCode int          `json:"status_code"`
	User       *entity.User `json:"user"`
}
