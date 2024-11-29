package models

import (
	"time"
)

type Job struct {
	ID        int64
	Title     string
	Company   string
	IsOpen    bool
	Quota     int64
	ExpDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

type JobWithListApplicant struct {
	ID         int64
	Title      string
	Company    string
	IsOpen     bool
	Quota      int64
	ExpDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
	Applicants []User
}
