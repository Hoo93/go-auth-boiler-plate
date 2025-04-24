package models

type User struct {
	ID       int32  `gorm:"primaryKey"`
	UserName string `gorm:"column:userName"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

// TableName ::
func (User) TableName() string { return "User" }
