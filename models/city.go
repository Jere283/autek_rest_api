// models/city.go
package models

type City struct {
	ID      uint   `gorm:"primaryKey;column:id_city"`
	Name    string `gorm:"unique;not null"`
	StateID uint   `gorm:"not null;column:state_id"`
	State   State  `gorm:"foreignKey:StateID"`
}
