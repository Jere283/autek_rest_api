// models/address.go
package models

type Address struct {
	ID      uint   `gorm:"primaryKey;column:id_address"`
	CityID  uint   `gorm:"not null;column:city_id"`
	City    City   `gorm:"foreignKey:CityID"`
	Address string `gorm:"not null"`
}
