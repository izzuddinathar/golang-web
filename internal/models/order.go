package models

import "time"

// Order represents an order
type Order struct {
	OrderID       uint64    `gorm:"primaryKey;autoIncrement" json:"order_id"`
	NomorMeja     int       `gorm:"not null" json:"nomor_meja"`
	MenuID        uint64    `gorm:"not null" json:"menu_id"`
	Jumlah        int       `gorm:"not null" json:"jumlah"`
	StatusPesanan string    `gorm:"type:enum('dipesan','diproses','disajikan');default:'dipesan'" json:"status_pesanan"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Association with Menu
	Menu Menu `gorm:"foreignKey:MenuID;references:MenuID" json:"menu"`
}
