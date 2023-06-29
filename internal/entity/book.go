package entity

import "github.com/google/uuid"

type Book struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Judul string
	Pengarang string
	Penerbit string
	Status   bool
	TahunTerbit string
}