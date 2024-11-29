package servers

import (
	"job-application/apperrors"

	"github.com/gin-gonic/gin"
)

func InvalidRoute(c *gin.Context) {
	c.Error(apperrors.ErrUrlNotFound)
	c.Abort()
}
