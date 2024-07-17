package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustConnect(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to open db: " + err.Error())
	}

	if err = db.AutoMigrate(
		&entity.User{},
		&entity.UserProfile{},
	); err != nil {
		panic("failed to migrate db: " + err.Error())
	}

	return db
}
