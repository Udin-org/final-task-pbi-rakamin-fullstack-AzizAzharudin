package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null" valid:"required"`
    Email     string    `gorm:"unique;not null" valid:"email,required"`
    Password  string    `gorm:"not null" valid:"required"`
    CreatedAt time.Time
    UpdatedAt time.Time
}


func (u *User) Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{})
}