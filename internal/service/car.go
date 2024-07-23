package service

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/google/uuid"
)

type CarManagement interface {
	AddCar(ctx context.Context, data models.CarReq) error
	RemoveCar(ctx context.Context, carID string) error
	EditCar(ctx context.Context, data models.CarReq, carID string) error
}

type carManagement struct {
	repo repositories.CarManagement
}

func NewCarManagement(repo repositories.CarManagement) CarManagement {
	return &carManagement{
		repo: repo,
	}
}

func (c *carManagement) AddCar(ctx context.Context, data models.CarReq) error {
	carID := uuid.New().String()

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	if err := c.repo.AddCar(ctx, entity.Car{
		ID:    carID,
		Brand: data.Brand,
		Class: data.Class,
	}); err != nil {
		return errors.ErrorRepoWrapper(err)
	}
	return nil
}

func (c *carManagement) RemoveCar(ctx context.Context, carID string) error {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	if err := c.repo.RemoveCar(ctx, carID); err != nil {
		return errors.ErrorRepoWrapper(err)
	}
	return nil
}

func (c *carManagement) EditCar(ctx context.Context, data models.CarReq, carID string) error {
	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	if err := c.repo.EditCar(ctx, data, carID); err != nil {
		return errors.ErrorRepoWrapper(err)
	}
	return nil
}
