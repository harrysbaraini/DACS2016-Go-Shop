package models

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Product ...
// Product model
type Product struct {
	BaseModel
	Name        string
	Description string
	Price       float32
}
