// models/service.go
package models

type Service struct {
	ID          uint   `gorm:"primaryKey;column:id_service"`
	Name        string `gorm:"size:100;not null"`
	Description string
}
