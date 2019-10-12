package models

import (
	"time"
)

type User struct {
	Uid       string `gorm:"AUTO_INCREMENT"`
	Username  string
	IsAdmin   bool
	Password  string
	Email     string
	Attempt   uint32
	Accept    uint32
	LoginAt   time.Time
	CreatedAt time.Time
}
