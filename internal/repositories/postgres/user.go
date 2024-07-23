package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
)

func (p *Postgres) AboutMe(ctx context.Context, userID string) (entity.UserInfo, error) {
	var info entity.UserInfo

	var user entity.User
	err := p.db.Model(&entity.User{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return entity.UserInfo{}, errors.ErrorDBWrapper(err)
	}
	info.Notifications = user.Notifications

	var profile entity.UserProfile
	err = p.db.Model(&entity.UserProfile{}).Where("id = ?", userID).First(&profile).Error
	if err != nil {
		return entity.UserInfo{}, errors.ErrorDBWrapper(err)
	}

	info.Email = profile.Email
	info.RentHours = profile.RentHours
	info.Bio = profile.Bio
	info.Joined = profile.Joined
	info.Sex = profile.Sex
	info.Username = profile.Username

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "repo layers passed")
	return info, nil
}

func (p *Postgres) GetProfile(ctx context.Context, userID string) (entity.UserProfile, error) {
	var profile entity.UserProfile

	err := p.db.Model(&entity.UserProfile{}).Where("id = ?", userID).First(&profile).Error
	if err != nil {
		return entity.UserProfile{}, errors.ErrorDBWrapper(err)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "repo layers passed")
	return profile, nil
}

func (p *Postgres) GetNotifications(ctx context.Context, userID string) ([]entity.Notification, error) {
	var notifications []entity.Notification

	res := p.db.Model(&entity.User{}).Where("id = ?", userID).Find(&[]entity.User{})
	if res.RowsAffected == 0 {
		return nil, errors.NewError(errors.ERR_NOT_FOUND, "entity not found")
	}

	err := p.db.Model(&entity.Notification{}).Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	err = p.db.Model(&entity.Notification{}).Where("user_id = ?", userID).Delete(&entity.Notification{}).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	err = p.db.Model(&entity.User{}).Where("id = ?", userID).Update("notifications", 0).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "repo layers passed")
	return notifications, nil
}
