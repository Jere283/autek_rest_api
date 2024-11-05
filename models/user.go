package models

import "time"

type User struct {
	ID           string    `gorm:"primaryKey;column:id_user"`
	FirstName    string    `gorm:"size:60;not null"`
	LastName     string    `gorm:"size:60;not null"`
	Email        string    `gorm:"size:120;unique;not null"`
	Password     string    `gorm:"size:200;not null"`
	DateOfBirth  time.Time `gorm:"not null"`
	RegisterDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	IsActive     bool      `gorm:"default:true"`
	RoleID       uint      `gorm:"column:role"`
	Role         Role      `gorm:"foreignKey:RoleID"`
}
