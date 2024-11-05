// models/appointment_service.go
package models

type AppointmentService struct {
	ID            uint            `gorm:"primaryKey;column:id_appointment_service"`
	AppointmentID uint            `gorm:"not null;column:appointment_id"`
	Appointment   Appointment     `gorm:"foreignKey:AppointmentID"`
	ServiceID     uint            `gorm:"not null;column:service_id"`
	Service       WorkshopService `gorm:"foreignKey:ServiceID"`
}
