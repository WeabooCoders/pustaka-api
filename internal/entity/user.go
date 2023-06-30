package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Username string
	Password string
	Alamat string
	Email string
	NomorTelepon int
}