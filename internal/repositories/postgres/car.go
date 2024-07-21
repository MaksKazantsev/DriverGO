package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
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

	err = p.db.Where("id = ?", carID).Delete(&entity.Car{}).Error
	if err != nil {
		return errors.NewError(errors.ERR_INTERNAL, err.Error())
	}
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

	err = p.db.Where("id = ?", carID).Updates(map[string]interface{}{"class": data.Class, "brand": data.Brand}).Error
	if err != nil {
		return errors.NewError(errors.ERR_INTERNAL, err.Error())
	}
	return nil
}
