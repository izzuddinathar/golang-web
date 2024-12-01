package models

import "time"

type User struct {
	UserID    uint64    `gorm:"primaryKey;autoIncrement"`
	Nama      string    `gorm:"size:255"`
	NoTelp    string    `gorm:"size:255"`
	Email     string    `gorm:"size:255;unique;not null"`
	Foto      string    `gorm:"size:255"`
	Username  string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"type:enum('admin','owner','waiters','cook','cleaner');not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
