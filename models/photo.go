package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    URL       string    `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (p *Photo) Migrate(db *gorm.DB) {
    db.AutoMigrate(&Photo{})
}
