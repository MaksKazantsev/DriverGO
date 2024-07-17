package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"gorm.io/gorm"
)

var _ repositories.Repository = &Postgres{}

type Postgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}
