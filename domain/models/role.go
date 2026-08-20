package models

import "time"

// Untuk migrate Table di DB
type Role struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Code      string `gorm:"type:varchar(15);not null"`
	Name      string `gorm:"type:varchar(25);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
