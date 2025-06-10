package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Data      UserData
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserData struct {
	Name     string
	Surname  string
	Age      uint64
	Email    string
	Role     UserRole
	Password string
}

type UserRole int32
