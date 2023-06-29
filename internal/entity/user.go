package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username string
	Password string
	Alamat string
	Email string
	NomorTelepon int
}