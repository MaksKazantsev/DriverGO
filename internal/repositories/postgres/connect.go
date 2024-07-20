package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
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
		&entity.Car{},
		&entity.Rent{},
		&entity.RentHistory{},
		&entity.Bill{},
	); err != nil {
		panic("failed to migrate db: " + err.Error())
	}

	initAdmin(db)
	return db
}

func initAdmin(db *gorm.DB) {
	id := "1"
	token, err := utils.NewToken(utils.REFRESH, utils.TokenData{ID: id, Role: "admin"})
	if err != nil {
		panic("failed to make admin")
	}
	db.Model(&entity.User{}).Create(&entity.User{ID: id, RFToken: token, Email: "admin@gmail.com", Password: "adminpassword", Username: "admin"})
}
