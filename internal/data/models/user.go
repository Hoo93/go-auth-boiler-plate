package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"column:userName"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
