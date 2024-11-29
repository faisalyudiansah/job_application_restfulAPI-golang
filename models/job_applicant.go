package models

import "time"

type JobApplicant struct {
	ID        int64
	UserId    int64
	JobId     int64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

type JobApplicantWithDataUser struct {
	ID          int64
	UserId      int64
	JobId       int64
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
	UserProfile UserProfile
}
