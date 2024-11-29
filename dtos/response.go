package dtos

import (
	"time"

	"job-application/models"
)

type ResponseMessageOnly struct {
	Message string `json:"message"`
}

type ResponseApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type ResponseUserProfile struct {
	ID         int64      `json:"id"`
	UserId     int64      `json:"user_id"`
	Age        int        `json:"age"`
	CurrentJob *string    `json:"current_job"`
	Address    *string    `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeleteAt   *time.Time `json:"deleted_at"`
}

type ResponseRegisterUserAndUserProfile struct {
	ID        int64               `json:"id"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Role      string              `json:"role"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeleteAt  *time.Time          `json:"deleted_at"`
	Detail    ResponseUserProfile `json:"detail_user"`
}

type ResponseAccessToken struct {
	AccessToken string `json:"access_token"`
}

type ResponseRegisterUser struct {
	Message string                                        `json:"message"`
	Result  map[string]ResponseRegisterUserAndUserProfile `json:"result"`
}

type ResponseListJob struct {
	Message   string                 `json:"message"`
	TotalData int64                  `json:"total_data"`
	Result    []ResponseJobApplicant `json:"result"`
}

type ResponseAJob struct {
	Message   string      `json:"message"`
	TotalData int64       `json:"total_data"`
	Result    ResponseJob `json:"result"`
}

type ResponseLoginUser struct {
	Message string              `json:"message"`
	Result  ResponseAccessToken `json:"result"`
}

type ResponseJobApplicant struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title"`
	Company   string         `json:"company"`
	IsOpen    bool           `json:"is_open"`
	Quota     int64          `json:"quota"`
	ExpDate   time.Time      `json:"exp_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  *time.Time     `json:"deleted_at"`
	Applicant *[]models.User `json:"applicant,omitempty"`
}

type ResponseJob struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Company   string     `json:"company"`
	IsOpen    bool       `json:"is_open"`
	Quota     int64      `json:"quota"`
	ExpDate   time.Time  `json:"exp_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at"`
}

type ResponseApply struct {
	ID          int64               `json:"id"`
	UserId      int64               `json:"user_id"`
	JobId       int64               `json:"job_id"`
	Status      string              `json:"is_open"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeleteAt    *time.Time          `json:"deleted_at"`
	UserProfile ResponseUserProfile `json:"user_profile"`
}

type ResponseApplyResult struct {
	Message string        `json:"message"`
	Result  ResponseApply `json:"result"`
}
