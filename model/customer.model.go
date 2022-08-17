package model

import (
	"time"
)

type Customer struct {
	Id        int64
	FirstName string    `form:"firstname" json:"firstname" binding:"required"`
	LastName  string    `form:"lastname" json:"lastname" binding:"required"`
	BirthDate time.Time `form:"birthdate" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	Gender    string    `form:"gender" json:"gender" binding:"required"`
	Email     string    `form:"email" json:"email" binding:"required,email"`
	Address   string    `form:"address" json:"address" binding:"required"`
}
