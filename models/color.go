// models/color.go
package models

type Color struct {
	ID   uint   `gorm:"primaryKey;column:id_color"`
	Name string `gorm:"unique;not null"`
}
