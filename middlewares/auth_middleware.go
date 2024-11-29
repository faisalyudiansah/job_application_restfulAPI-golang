package middlewares

import (
	"context"
	"strings"

	"job-application/apperrors"
	"job-application/helpers"
	"job-application/models"

	"github.com/gin-gonic/gin"
)

func AuthorizationJob(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	if reqToken == "" || len(reqToken) == 0 {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	jwtProvider := helpers.NewJWTProviderHS256()
	result, err := jwtProvider.VerifyToken(splitToken[1])
	if err != nil {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}

	var userId models.ID = "userId"
	ctxId := context.WithValue(c.Request.Context(), userId, result.UserID)
	c.Request = c.Request.WithContext(ctxId)

	var role models.Role = "role_user"
	ctxRole := context.WithValue(c.Request.Context(), role, result.Role)
	c.Request = c.Request.WithContext(ctxRole)

	c.Next()
}
