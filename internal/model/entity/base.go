package entity

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt time.Time `gorm:"type:DATETIME;column:create_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME;column:update_at"`
	Deleted   gorm.DeletedAt
}
