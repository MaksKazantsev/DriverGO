package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/db"
	"gorm.io/gorm"
)

var _ db.Repository = &Postgres{}

type Postgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}
