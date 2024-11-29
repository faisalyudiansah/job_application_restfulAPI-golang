package middlewares

import (
	"errors"
	"net/http"

	"job-application/apperrors"
	"job-application/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	if len(c.Errors) > 0 {
		var ve validator.ValidationErrors
		if errors.As(c.Errors[0].Err, &ve) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": helpers.FormatterErrorInput(ve)})
			return
		}

		errorMappings := map[error]int{
			apperrors.ErrISE:                       http.StatusInternalServerError,
			apperrors.ErrInvalidAccessToken:        http.StatusUnauthorized,
			apperrors.ErrUnauthorization:           http.StatusUnauthorized,
			apperrors.ErrUrlNotFound:               http.StatusNotFound,
			apperrors.ErrRequestBodyInvalid:        http.StatusBadRequest,
			apperrors.ErrUserEmailAlreadyExists:    http.StatusBadRequest,
			apperrors.ErrUserFailedRegister:        http.StatusBadRequest,
			apperrors.ErrUserInvalidEmailPassword:  http.StatusBadRequest,
			apperrors.ErrJobIdNotExists:            http.StatusBadRequest,
			apperrors.ErrFailedApplyJob:            http.StatusBadRequest,
			apperrors.ErrJobIdNotValid:             http.StatusBadRequest,
			apperrors.ErrUserAlreadyApplyToThisJob: http.StatusBadRequest,
			apperrors.ErrQuotaInvalid:              http.StatusBadRequest,
			apperrors.ErrFailedCreateJob:           http.StatusBadRequest,
			apperrors.ErrFailedPatchIsOpen:         http.StatusBadRequest,
		}
		// asd := errorMappings[c.Errors[0].Err]

		for err, statusCode := range errorMappings {
			if errors.Is(c.Errors[0].Err, err) {
				helpers.PrintError(c, statusCode, err.Error())
				return
			}
		}

		helpers.PrintError(c, http.StatusInternalServerError, apperrors.ErrISE.Error())
	}
}
