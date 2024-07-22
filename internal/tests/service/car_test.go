package service

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	mock_service "github.com/MaksKazantsev/DriverGO/internal/tests/mocks"
	mock_log "github.com/MaksKazantsev/DriverGO/internal/tests/mocks/logger"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCarSuite(t *testing.T) {
	suite.Run(t, new(CarTest))
}

type CarTest struct {
	suite.Suite

	ctx context.Context

	ctrl *gomock.Controller
	repo *mock_service.MockRepository
}

func (r *CarTest) SetupTest() {
	r.ctrl = gomock.NewController(r.T())
	r.repo = mock_service.NewMockRepository(r.ctrl)

	logger := mock_log.NewMockLogger(r.ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	r.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	r.ctx = context.WithValue(r.ctx, utils.LoggerKey, logger)
}

func (r *CarTest) TearDownTest() {
	r.ctrl.Finish()
}

func (r *CarTest) TestAddCar() {
	// Correct Request
	car := entity.Car{
		ID:    "1",
		Class: "Premium",
		Brand: "BMW",
	}

	r.repo.EXPECT().AddCar(gomock.Any(), gomock.AssignableToTypeOf(car)).Times(1).Return(nil)

	req := models.CarReq{
		Class: "Premium",
		Brand: "BMW",
	}

	err := service.NewService(r.repo).AddCar(r.ctx, req)
	r.Require().NoError(err)

	// Invalid class
	car = entity.Car{
		ID:    "1",
		Class: "UnknownClassBugaga",
		Brand: "BMW",
	}

	r.repo.EXPECT().AddCar(gomock.Any(), gomock.AssignableToTypeOf(car)).Times(1).Return(errors.NewError(errors.ERR_BAD_REQUEST, "invalid class type"))

	req = models.CarReq{
		Class: "UnknownClassBugaga",
		Brand: "BMW",
	}

	err = service.NewService(r.repo).AddCar(r.ctx, req)
	r.Require().Error(err)
}

func (r *CarTest) TestRemoveCar() {
	// Correct Request
	carID := "1"

	r.repo.EXPECT().RemoveCar(gomock.Any(), carID).Times(1).Return(nil)

	err := service.NewService(r.repo).RemoveCar(r.ctx, carID)
	r.Require().NoError(err)
}

func (r *CarTest) TestEditCar() {
	// Correct Request
	req := models.CarReq{
		Class: "Premium",
		Brand: "BMW",
	}
	carID := "1"

	r.repo.EXPECT().EditCar(gomock.Any(), req, carID).Times(1).Return(nil)

	err := service.NewService(r.repo).EditCar(r.ctx, req, carID)
	r.Require().NoError(err)
}
