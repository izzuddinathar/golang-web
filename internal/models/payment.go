package models

import "time"

// Payment represents a payment
type Payment struct {
	PaymentID        uint64    `gorm:"primaryKey;autoIncrement" json:"payment_id"`
	NomorMeja        int       `gorm:"not null" json:"nomor_meja"`
	MenuID           uint64    `gorm:"not null" json:"menu_id"`
	Jumlah           int       `gorm:"not null" json:"jumlah"`
	MetodePembayaran string    `gorm:"type:enum('tunai','kartu kredit','kartu debit','qris');not null" json:"metode_pembayaran"`
	Status           string    `gorm:"type:enum('belum dibayar','lunas');default:'belum dibayar'" json:"status"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Association with Menu
	Menu Menu `gorm:"foreignKey:MenuID;references:MenuID" json:"menu"`
}
