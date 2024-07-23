package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"gorm.io/gorm"
)

func (p *Postgres) GetFBToken(ctx context.Context, userID string) (string, error) {
	var user entity.User
	err := p.db.Model(&entity.User{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", errors.ErrorDBWrapper(err)
	}
	return user.FBToken, nil
}

func (p *Postgres) SaveNotification(ctx context.Context, noti entity.Notification) error {
	if err := p.db.Create(&noti).Error; err != nil {
		return errors.ErrorDBWrapper(err)
	}

	if err := p.db.Model(&entity.User{}).Where("id = ?", noti.UserID).Update("notifications", gorm.Expr("notifications + 1")).Error; err != nil {
		return errors.ErrorDBWrapper(err)
	}

	return nil
}
