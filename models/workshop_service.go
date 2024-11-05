// models/workshop_service.go
package models

type WorkshopService struct {
	ID         uint     `gorm:"primaryKey;column:id_workshop_service"`
	WorkshopID uint     `gorm:"not null;column:workshop_id"`
	Workshop   Workshop `gorm:"foreignKey:WorkshopID"`
	ServiceID  uint     `gorm:"not null;column:service_id"`
	Service    Service  `gorm:"foreignKey:ServiceID"`
	Price      float64  `gorm:"type:numeric(10,2);not null"`
}
