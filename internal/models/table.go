package models

import "time"

// Table represents a table in the restaurant
type Table struct {
	TableID   uint64    `gorm:"primaryKey;autoIncrement" json:"table_id"`
	NomorMeja int       `gorm:"not null;unique" json:"nomor_meja"`
	Kapasitas int       `gorm:"not null" json:"kapasitas"`
	Status    string    `gorm:"type:enum('dipesan','tersedia','terisi');not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
