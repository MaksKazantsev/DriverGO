package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"gorm.io/gorm"
)

func (p *Postgres) Register(ctx context.Context, data models.RegisterReq) (models.AuthResponse, error) {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.User{}).Create(&data).Error
		if err != nil {
			return err
		}
		err = tx.Model(&entity.UserProfile{}).Create(&data).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.AuthResponse{}, errors.ErrorDBWrapper(err)
	}
	return models.AuthResponse{RefreshToken: data.RToken, UUID: data.ID}, nil
}

func (p *Postgres) Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error) {
	var user entity.User
	err := p.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.User{}).Where("email = ?", data.Email).Update("rf_token", data.RToken).Update("fb_token", data.FBToken).First(&user).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.AuthResponse{}, errors.ErrorDBWrapper(err)
	}
	return models.AuthResponse{RefreshToken: data.RToken, UUID: user.ID}, nil
}
