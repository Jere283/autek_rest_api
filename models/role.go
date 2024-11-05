package models

type Role struct {
	ID   uint   `gorm:"primaryKey;column:id_role"`
	Name string `gorm:"unique;not null"`
}
