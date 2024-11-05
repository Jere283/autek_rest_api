// models/appointment_status.go
package models

type AppointmentStatus struct {
	ID   uint   `gorm:"primaryKey;column:id_appointment_status"`
	Name string `gorm:"unique;not null"`
}
