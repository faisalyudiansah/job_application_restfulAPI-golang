package dtos

import "time"

type RequestValidationMiddleware struct {
	Title    string `json:"title" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      string `json:"age" binding:"required"`
}

type RequestRegisterUser struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required"`
	Password   string  `json:"password" binding:"required"`
	Age        int     `json:"age" binding:"required"`
	CurrentJob *string `json:"current_job"`
	Address    *string `json:"address"`
}

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestCreateJob struct {
	Title   string    `json:"title" binding:"required"`
	Company string    `json:"company" binding:"required"`
	IsOpen  bool      `json:"is_open" binding:"required"`
	Quota   int64     `json:"quota" binding:"required"`
	ExpDate time.Time `json:"exp_date" binding:"required"`
}
