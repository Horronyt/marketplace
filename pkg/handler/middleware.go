package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		if c.Request.Method != http.MethodGet {
			newErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		}
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		if c.Request.Method != http.MethodGet {
			newErrorResponse(c, http.StatusUnauthorized, "Incorrect authorization header format")
		}
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		if c.Request.Method != http.MethodGet {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		if c.Request.Method != http.MethodGet {
			newErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		}
		return 0, errors.New("User not authenticated")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "Incorrect user id format")
		return 0, errors.New("User not authenticated")
	}

	return idInt, nil
}
