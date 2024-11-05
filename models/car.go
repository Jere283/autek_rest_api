// models/car.go
package models

type Car struct {
	ID           uint     `gorm:"primaryKey;column:id_car"`
	BrandID      uint     `gorm:"not null;column:brand_id"`
	Brand        CarBrand `gorm:"foreignKey:BrandID"`
	ModelID      uint     `gorm:"not null;column:model_id"`
	Model        CarModel `gorm:"foreignKey:ModelID"`
	ColorID      uint     `gorm:"not null;column:color_id"`
	Color        Color    `gorm:"foreignKey:ColorID"`
	LicensePlate string   `gorm:"size:15;unique;not null"`
	Year         string   `gorm:"size:4;not null"`
}
