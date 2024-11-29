package models

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

type UserProfile struct {
	ID         int64
	UserId     int64
	Age        int
	CurrentJob *string
	Address    *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
}

type UserAndProfile struct {
	User        User
	UserProfile UserProfile
}
