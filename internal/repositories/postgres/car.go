package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
)

func (p *Postgres) AddCar(ctx context.Context, car entity.Car) error {
	err := p.db.Model(&entity.Car{}).Create(&car).Error
	if err != nil {
		return errors.ErrorDBWrapper(err)
	}
	return nil
}

func (p *Postgres) RemoveCar(ctx context.Context, carID string) error {
	var rentCount int64

	err := p.db.Model(&entity.Rent{}).Where("car_id = ?", carID).Count(&rentCount).Error
	if err != nil {
		return errors.ErrorDBWrapper(err)
	}

	if rentCount > 0 {
		return errors.NewError(errors.ERR_BAD_REQUEST, "car is rented")
	}

	res := p.db.Model(&entity.Car{}).Where("id = ?", carID).Delete(&entity.Car{})
	if res.RowsAffected == 0 {
		return errors.NewError(errors.ERR_NOT_FOUND, "car does not exist")
	}
	if err != nil {
		return errors.ErrorDBWrapper(res.Error)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "repo layers successfully passed")
	return nil
}

func (p *Postgres) EditCar(ctx context.Context, data models.CarReq, carID string) error {
	var rentCount int64

	err := p.db.Model(&entity.Rent{}).Where("car_id = ?", carID).Count(&rentCount).Error
	if err != nil {
		return errors.ErrorDBWrapper(err)
	}

	if rentCount > 0 {
		return errors.NewError(errors.ERR_BAD_REQUEST, "car is rented")
	}

	res := p.db.Where("id = ?", carID).Updates(&entity.Car{Class: data.Class, Brand: data.Brand})
	if res.RowsAffected == 0 {
		return errors.NewError(errors.ERR_NOT_FOUND, "car does not exist")
	}
	if res.Error != nil {
		return errors.ErrorDBWrapper(res.Error)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "repo layers successfully passed")
	return nil
}
