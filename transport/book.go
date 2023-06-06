package transport

import (
	"github.com/gin-gonic/gin"
	"go-gin-auth/domain"
	"net/http"
)

func (h *Handler) GetBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &domain.Response{Status: "ok", Message: "Hi there!"})
}
