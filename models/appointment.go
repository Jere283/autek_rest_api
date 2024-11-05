// models/appointment.go
package models

import "time"

type Appointment struct {
	ID                  uint     `gorm:"primaryKey;column:id_appointment"`
	UserID              string   `gorm:"column:id_user"`
	User                User     `gorm:"foreignKey:UserID"`
	CarID               uint     `gorm:"column:id_car"`
	Car                 Car      `gorm:"foreignKey:CarID"`
	WorkshopID          uint     `gorm:"column:id_workshops"`
	Workshop            Workshop `gorm:"foreignKey:WorkshopID"`
	Description         string
	Date                time.Time         `gorm:"not null"`
	AppointmentStatus   AppointmentStatus `gorm:"foreignKey:AppointmentStatusID"`
	AppointmentStatusID uint
}
