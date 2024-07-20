package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"gorm.io/gorm"
	"time"
)

func (p *Postgres) Register(ctx context.Context, data entity.User) (models.AuthResponse, error) {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.User{}).Create(&data).Error
		if err != nil {
			return err
		}
		err = tx.Model(&entity.UserProfile{}).Create(&entity.UserProfile{
			ID:        data.ID,
			Username:  data.Username,
			Email:     data.Email,
			RentHours: 0 * time.Second,
		}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.AuthResponse{}, errors.ErrorDBWrapper(err)
	}

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return models.AuthResponse{RefreshToken: data.RFToken, UUID: data.ID}, nil
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
	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return models.AuthResponse{RefreshToken: data.RToken, UUID: user.ID}, nil
}

func (p *Postgres) Refresh(ctx context.Context, uuid, token string) (models.AuthResponse, error) {
	err := p.db.Model(&entity.User{}).Where("id = ?", uuid).Update("rf_token", token).Error
	if err != nil {
		return models.AuthResponse{}, errors.ErrorDBWrapper(err)
	}

	return models.AuthResponse{
		RefreshToken: token,
		UUID:         uuid,
	}, nil
}

func (p *Postgres) GetPasswordAndID(ctx context.Context, email string) (string, string, error) {
	var user entity.User
	if err := p.db.Model(&entity.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", errors.ErrorRepoWrapper(err)
	}

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return user.ID, user.Password, nil
}
