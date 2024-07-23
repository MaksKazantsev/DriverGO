package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/notifications"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"log"
)

type Rent interface {
	StartRent(ctx context.Context, token string, carID string) error
	FinishRent(ctx context.Context, token string, rentID string) (entity.Bill, error)
	GetRentHistory(ctx context.Context, token string) ([]entity.RentHistory, error)
	GetAvailableCars(ctx context.Context) ([]entity.Car, error)
}

type rent struct {
	repo     repositories.Rent
	notifier notifications.Notifier
}

func NewRent(repo repositories.Rent, notifier notifications.Notifier) Rent {
	return &rent{repo: repo, notifier: notifier}
}

func (r *rent) StartRent(ctx context.Context, token string, carID string) error {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	if err = r.repo.StartRent(ctx, id, carID); err != nil {
		return errors.ErrorRepoWrapper(err)
	}

	if err = r.notifier.Notify("rent_started", id); err != nil {
		log.Println("failed to send notification: %w", err)
	}

	return nil
}

func (r *rent) FinishRent(ctx context.Context, token string, rentID string) (entity.Bill, error) {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return entity.Bill{}, fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")

	bill, err := r.repo.FinishRent(ctx, id, rentID)
	if err != nil {
		return entity.Bill{}, errors.ErrorRepoWrapper(err)
	}

	if err = r.notifier.Notify("rent_finished", id); err != nil {
		log.Println("failed to send notification: %w", err)
	}

	return bill, nil
}

func (r *rent) GetRentHistory(ctx context.Context, token string) ([]entity.RentHistory, error) {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	rents, err := r.repo.GetRentHistory(ctx, id)

	if err != nil {
		return nil, errors.ErrorRepoWrapper(err)
	}
	return rents, nil
}

func (r *rent) GetAvailableCars(ctx context.Context) ([]entity.Car, error) {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")

	cars, err := r.repo.GetAvailableCars(ctx)
	if err != nil {
		return nil, errors.ErrorRepoWrapper(err)
	}

	return cars, nil
}
