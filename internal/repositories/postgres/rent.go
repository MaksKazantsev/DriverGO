package postgres

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func (p *Postgres) StartRent(ctx context.Context, userID, carID string) error {
	var car entity.Car
	err := p.db.Model(&entity.Car{}).Where("id = ?", carID).First(&car).Error
	if err != nil {
		return errors.ErrorDBWrapperC(err, "car does not exist", "")
	}

	err = p.db.Model(&entity.Rent{}).Create(&entity.Rent{ID: uuid.NewString(), CarID: carID, UserID: userID, CarClass: car.Class, StartTime: time.Now().In(time.Local)}).Error
	if err != nil {
		return errors.ErrorDBWrapperC(err, "", "car already in rent")
	}

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return nil
}

func (p *Postgres) FinishRent(ctx context.Context, userID, rentID string) (entity.Bill, error) {
	var rent entity.Rent
	err := p.db.Model(&entity.Rent{}).Where("id = ?", rentID).
		Where("user_id = ?", userID).
		First(&rent).
		Delete(&rent).
		Error

	finishTime := time.Since(rent.StartTime)

	err = p.db.Model(&entity.RentHistory{}).Create(&entity.RentHistory{ID: rentID, CarID: rent.CarID, UserID: userID, CarClass: rent.CarClass, RentDuration: finishTime}).Error
	if err != nil {
		return entity.Bill{}, errors.ErrorDBWrapper(err)
	}

	price, err := calculateRentPrice(rent.CarClass, rent.StartTime)
	if err != nil {
		return entity.Bill{}, err
	}

	bill := entity.Bill{ID: uuid.NewString(), UserID: userID, Price: price}
	err = p.db.Model(&entity.Bill{}).Create(&bill).Error
	if err != nil {
		return entity.Bill{}, errors.ErrorDBWrapper(err)
	}

	p.db.Model(&entity.UserProfile{}).Where("id = ?", userID).Update("rent_hours", gorm.Expr("rent_hours + ?", finishTime))

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return bill, nil
}

func (p *Postgres) GetRentHistory(ctx context.Context, userID string) ([]entity.RentHistory, error) {
	var rents []entity.RentHistory
	err := p.db.Model(&entity.RentHistory{}).Where("user_id = ?", userID).Find(&rents).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return rents, nil
}

func (p *Postgres) GetAvailableCars(ctx context.Context) ([]entity.Car, error) {
	var carsIDs []string

	err := p.db.Model(&entity.Car{}).Pluck("id", &carsIDs).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	var cars []entity.Car
	err = p.db.Model(&entity.Rent{}).Where("car_id NOT IN ?", carsIDs).Find(&cars).Error
	if err != nil {
		return nil, errors.ErrorDBWrapper(err)
	}

	utils.ExtractLogger(ctx).Info("repo layers successfully passed", nil)
	return cars, nil
}
