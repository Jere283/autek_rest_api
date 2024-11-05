// models/car_brand.go
package models

type CarBrand struct {
	ID   uint   `gorm:"primaryKey;column:id_brand"`
	Name string `gorm:"unique;not null"`
}
