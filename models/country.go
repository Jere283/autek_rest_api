package models

type Country struct {
	ID   uint   `gorm:"primaryKey;column:id_country"`
	Name string `gorm:"unique;not null"`
}
