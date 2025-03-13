package translate

import (
	"github.com/aberyotaro/go-api-sandbox/internal/entity"
	"github.com/aberyotaro/go-api-sandbox/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase usecase.Translator
}

func NewHandler(uc usecase.Translator) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) Translate(c *gin.Context) {
	var (
		err error
		res *entity.ChatGPTResponse
	)

	res, err = h.useCase.Translate(c.Param("text"), entity.GetMatchLang(c.Param("toLang")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  res,
		"message": "success",
	})
}
