package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

var _ repositories.Repository = &Postgres{}

func NewRepository(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}
