package transport

import (
	"github.com/gin-gonic/gin"
	"go-gin-auth/domain"
	"go-gin-auth/pkg/logger"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	var input domain.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(111, err.Error())
		return
	}

	accessToken, refreshToken, err := h.usersService.SignIn(c.Request.Context(), input)
	if err != nil {
		c.JSON(222, err.Error())
		return
	}

	refreshTokenTTL := h.usersService.GetRefreshTokenTTL().Seconds()

	c.SetCookie("refresh-token", refreshToken, int(refreshTokenTTL), "/", "localhost", false, true)
	c.JSON(http.StatusOK, domain.Response{Status: "ok", Token: accessToken})
}

func (h *Handler) refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	
	if err != nil {
		logger.LogError("refresh", err)
		c.JSON(http.StatusBadRequest, domain.Response{Error: domain.ErrRefreshTokenParse.Error()})
		return
	}

	accessToken, refreshToken, err := h.usersService.RefreshTokens(c.Request.Context(), cookie)
	if err != nil {
		logger.LogError("refresh", err)
		c.JSON(http.StatusInternalServerError, domain.Response{Error: domain.ErrRefreshToken.Error()})
		return
	}
	refreshTokenTTL := h.usersService.GetRefreshTokenTTL().Seconds()
	c.SetCookie("refresh-token", refreshToken, int(refreshTokenTTL), "/", "localhost", false, true)
	c.JSON(http.StatusOK, domain.Response{Status: "ok", Token: accessToken})
}
