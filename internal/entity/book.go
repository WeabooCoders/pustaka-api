package entity

import "gorm.io/gorm"


type Book struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Judul string
	Pengarang string
	Penerbit string
	Status   bool
	TahunTerbit string
}