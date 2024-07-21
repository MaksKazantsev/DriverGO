package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.Repository {
	return &Postgres{
		db: db,
	}
}
