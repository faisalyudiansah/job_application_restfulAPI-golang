package helpers

import (
	"reflect"

	"job-application/dtos"
	"job-application/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PrintError(c *gin.Context, statusCode int, msg string) {
	res := dtos.ResponseMessageOnly{
		Message: msg,
	}
	c.AbortWithStatusJSON(statusCode, res)
}

func PrintResponse(c *gin.Context, statusCode int, res interface{}) {
	c.JSON(statusCode, res)
}

func FormatterErrorInput(ve validator.ValidationErrors) []dtos.ResponseApiError {
	result := make([]dtos.ResponseApiError, len(ve))
	for i, fe := range ve {
		result[i] = dtos.ResponseApiError{
			Field: jsonFieldName(fe.Field()),
			Msg:   msgForTag(fe.Tag()),
		}
	}
	return result
}

func jsonFieldName(fieldName string) string {
	t := reflect.TypeOf(dtos.RequestValidationMiddleware{})
	field, found := t.FieldByName(fieldName)
	if !found {
		return ""
	}
	jsonTag := field.Tag.Get("json")
	return jsonTag
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}

func FormatterSuccessRegister(data *dtos.ResponseRegisterUserAndUserProfile, msg string) dtos.ResponseRegisterUser {
	mapBook := map[string]dtos.ResponseRegisterUserAndUserProfile{}
	mapBook["data"] = *data
	res := dtos.ResponseRegisterUser{
		Message: msg,
		Result:  mapBook,
	}
	return res
}

func FormatterSuccessLogin(data *dtos.ResponseAccessToken, msg string) dtos.ResponseLoginUser {
	res := dtos.ResponseLoginUser{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterListJob(data []dtos.ResponseJobApplicant, msg string) dtos.ResponseListJob {
	res := dtos.ResponseListJob{
		Message:   msg,
		TotalData: int64(len(data)),
		Result:    data,
	}
	return res
}

func FormatterAJob(data dtos.ResponseJob, msg string) dtos.ResponseAJob {
	res := dtos.ResponseAJob{
		Message:   msg,
		TotalData: int64(1),
		Result:    data,
	}
	return res
}

func FormatterApplyJob(data *dtos.ResponseApply, msg string) dtos.ResponseApplyResult {
	res := dtos.ResponseApplyResult{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterToJobApplicantWithDataUser(ja *models.JobApplicant, up *models.UserProfile) *dtos.ResponseApply {
	return &dtos.ResponseApply{
		ID:          ja.ID,
		UserId:      ja.UserId,
		JobId:       ja.JobId,
		Status:      ja.Status,
		CreatedAt:   ja.CreatedAt,
		UpdatedAt:   ja.UpdatedAt,
		DeleteAt:    ja.DeleteAt,
		UserProfile: dtos.ResponseUserProfile(*up),
	}
}
