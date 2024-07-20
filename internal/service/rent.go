package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
)

type Rent interface {
	StartRent(ctx context.Context, token string, carID string) error
	FinishRent(ctx context.Context, token string, rentID string) (entity.Bill, error)
	GetRents(ctx context.Context, token string) ([]entity.Rent, error)
}

type rent struct {
	repo repositories.Rent
}

func NewRent(repo repositories.Rent) Rent {
	return &rent{repo: repo}
}

func (r *rent) StartRent(ctx context.Context, token string, carID string) error {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Info("service layers successfully passed", nil)
	if err = r.repo.StartRent(ctx, id, carID); err != nil {
		return errors.ErrorRepoWrapper(err)
	}
	return nil
}

func (r *rent) FinishRent(ctx context.Context, token string, rentID string) (entity.Bill, error) {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return entity.Bill{}, fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Info("service layers successfully passed", nil)
	bill, err := r.repo.FinishRent(ctx, id, rentID)
	if err != nil {
		return entity.Bill{}, errors.ErrorRepoWrapper(err)
	}

	return bill, nil
}

func (r *rent) GetRents(ctx context.Context, token string) ([]entity.Rent, error) {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to parase token: %w", err)
	}
	id := cl["id"].(string)

	utils.ExtractLogger(ctx).Info("service layers successfully passed", nil)
	rents, err := r.repo.GetRentHistory(ctx, id)
	if err != nil {
		return nil, errors.ErrorRepoWrapper(err)
	}
	return rents, nil
}
