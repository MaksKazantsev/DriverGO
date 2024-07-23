package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
)

type User interface {
	GetNotifications(ctx context.Context, token string) ([]entity.Notification, error)
	GetProfile(ctx context.Context, userID string) (entity.UserProfile, error)
	AboutMe(ctx context.Context, token string) (entity.UserInfo, error)
}

type user struct {
	repo repositories.User
}

func NewUser(repo repositories.User) User {
	return &user{repo: repo}
}

func (u *user) GetNotifications(ctx context.Context, token string) ([]entity.Notification, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	id := claims["id"].(string)

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers passed")
	notifications, err := u.repo.GetNotifications(ctx, id)
	if err != nil {
		return nil, errors.ErrorRepoWrapper(err)
	}
	return notifications, nil
}

func (u *user) GetProfile(ctx context.Context, userID string) (entity.UserProfile, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers passed")
	profile, err := u.repo.GetProfile(ctx, userID)
	if err != nil {
		return entity.UserProfile{}, errors.ErrorRepoWrapper(err)
	}
	return profile, nil
}

func (u *user) AboutMe(ctx context.Context, token string) (entity.UserInfo, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return entity.UserInfo{}, fmt.Errorf("failed to parse token: %w", err)
	}
	id := claims["id"].(string)

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers passed")
	info, err := u.repo.AboutMe(ctx, id)
	if err != nil {
		return entity.UserInfo{}, errors.ErrorRepoWrapper(err)
	}
	return info, nil
}
