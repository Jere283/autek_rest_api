package models

type State struct {
	ID        uint    `gorm:"primaryKey;column:id_state"`
	Name      string  `gorm:"unique;not null"`
	CountryID uint    `gorm:"not null;column:country_id"`
	Country   Country `gorm:"foreignKey:CountryID"`
}
