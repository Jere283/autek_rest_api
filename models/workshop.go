package models

type Workshop struct {
	ID          uint    `gorm:"primaryKey;column:id_workshop"`
	Name        string  `gorm:"size:60;not null"`
	PhoneNumber string  `gorm:"size:10;not null"`
	Email       string  `gorm:"size:120;unique;not null"`
	AddressID   uint    `gorm:"not null;column:address_id"`
	Address     Address `gorm:"foreignKey:AddressID"`
}
