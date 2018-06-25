package model

import "time"

type User struct {
	ID      int `gorm:"primary_key:id"`
	Name    string
	Phone   string
	Address string
	CreateTime time.Time
	UpdateTime time.Time
}
