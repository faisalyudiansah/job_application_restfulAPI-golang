package controllers

import (
	"io"
	"net/http"

	"job-application/apperrors"
	"job-application/constants"
	"job-application/dtos"
	"job-application/helpers"
	"job-application/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(us *services.UserServiceImplementation) *UserController {
	return &UserController{
		UserService: us,
	}
}

func (uc *UserController) PostRegisterUserController(c *gin.Context) {
	var reqBody dtos.RequestRegisterUser
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	result, err := uc.UserService.PostRegisterUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterSuccessRegister(result, constants.SuccessRegister))
}

func (uc *UserController) PostLoginUserController(c *gin.Context) {
	var reqBody dtos.RequestLoginUser
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	result, err := uc.UserService.PostLoginUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessLogin(result, constants.SuccessLogin))
}
