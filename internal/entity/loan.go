package entity

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	UserID      uint `gorm:"foreignKey:UserID"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	BookID      uint `gorm:"foreignKey:BookID"`
	Book        Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book"`
	TanggalPeminjaman time.Time
	TanggalPengembalian time.Time 
}
