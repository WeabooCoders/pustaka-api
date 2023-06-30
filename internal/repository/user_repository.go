package repository

import (
	"github.com/AvinFajarF/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface{
	Save(user *entity.User) error
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
        db: db,
    }
}

func (repo *PostgresUserRepository) Save(user *entity.User) error {
	return repo.db.Create(user).Error
}