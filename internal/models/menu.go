package models

import "time"

// Menu represents a menu item in the system
type Menu struct {
	MenuID    uint64    `gorm:"primaryKey;autoIncrement" json:"menu_id"`
	NamaMenu  string    `gorm:"size:255;not null" json:"nama_menu"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	Harga     float64   `gorm:"type:decimal(10,2);not null" json:"harga"`
	Kategori  string    `gorm:"type:enum('cemilan','makanan','minuman');not null" json:"kategori"`
	Gambar    string    `gorm:"size:255" json:"gambar"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
