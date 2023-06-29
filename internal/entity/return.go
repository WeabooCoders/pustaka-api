package entity

import (
	"time"

	"github.com/google/uuid"
)

type Return struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uint `gorm:"foreignKey:UserID"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	BookID      uint `gorm:"foreignKey:BookID"`
	Book        Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book"`
	TanggalPengembalian time.Time 
}