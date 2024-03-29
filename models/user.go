package models

import "time"

type User struct {
	Id       int64      `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Birthday *time.Time `json:"birthday"`
}
