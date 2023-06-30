package entity

import "gorm.io/gorm"


type Review struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	UserID      uint `gorm:"foreignKey:UserID"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	BookID      uint `gorm:"foreignKey:BookID"`
	Book        Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book"`
	Text        string
}