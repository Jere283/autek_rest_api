// models/car_model.go
package models

type CarModel struct {
	ID   uint   `gorm:"primaryKey;column:id_model"`
	Name string `gorm:"unique;not null"`
}
