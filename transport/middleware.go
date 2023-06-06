package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-gin-auth/domain"
	"go-gin-auth/pkg/logger"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	ctxUserID           = "userID"
)

func (h *Handler) getTokenFromRequest(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", domain.ErrEmptyAuthHeader
	}

	return header, nil
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := h.getTokenFromRequest(c)
		if err != nil {
			logger.LogError("authMiddleware", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Response{Error: domain.ErrAccessTokenParse.Error()})
			return
		}

		id, err := h.usersService.ParseToken(token)
		if err != nil {
			logger.LogError("authMiddleware", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.Response{Error: domain.ErrAccessTokenExpired.Error()})
			return
		}

		logrus.Infof("id %d", id)
		c.Set(ctxUserID, id)
		c.Next()
	}
}
