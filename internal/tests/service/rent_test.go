package service

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	mock_service "github.com/MaksKazantsev/DriverGO/internal/tests/mocks"
	mock_log "github.com/MaksKazantsev/DriverGO/internal/tests/mocks/logger"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestRentSuite(t *testing.T) {
	suite.Run(t, new(RentTest))
}

type RentTest struct {
	suite.Suite

	ctx context.Context

	repo *mock_service.MockRepository
	ctrl *gomock.Controller
}

func (rt *RentTest) SetupTest() {
	rt.ctrl = gomock.NewController(rt.T())
	rt.repo = mock_service.NewMockRepository(rt.ctrl)

	logger := mock_log.NewMockLogger(rt.ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	rt.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	rt.ctx = context.WithValue(rt.ctx, utils.LoggerKey, logger)
}

func (rt *RentTest) TearDownTest() {
	rt.ctrl.Finish()
}

func (rt *RentTest) TestStartRent() {
	// Correct Request
	rt.repo.EXPECT().StartRent(gomock.Any(), "1", "1").Times(1).Return(nil)

	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})
	err := service.NewService(rt.repo).StartRent(rt.ctx, token, "1")
	rt.Require().NoError(err)

	// Car does not exist
	rt.repo.EXPECT().StartRent(gomock.Any(), "2", "1").Times(1).Return(errors.NewError(errors.ERR_NOT_FOUND, "car not found"))

	token, _ = utils.NewToken(utils.ACCESS, utils.TokenData{ID: "2", Role: "user"})
	err = service.NewService(rt.repo).StartRent(rt.ctx, token, "1")
	rt.Require().Error(err)

	// Wrong token
	err = service.NewService(rt.repo).StartRent(rt.ctx, "12342141", "1")
	rt.Require().Error(err)
}

func (rt *RentTest) TestFinishRent() {
	// Correct Request
	rentID := "1"
	userID := "1"
	res := entity.Bill{
		ID:     "1",
		UserID: userID,
		Price:  67.2,
	}

	rt.repo.EXPECT().FinishRent(gomock.Any(), "1", "1").Times(1).Return(res, nil)
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: userID, Role: "user"})
	bill, err := service.NewService(rt.repo).FinishRent(rt.ctx, token, rentID)
	rt.Require().NoError(err)
	rt.Require().NotEqual(entity.Bill{}, bill)
	rt.Require().Equal(bill.ID, userID)

	// Rent does not exist
	rt.repo.EXPECT().FinishRent(gomock.Any(), "1", "1").Times(1).Return(entity.Bill{}, errors.NewError(errors.ERR_NOT_FOUND, "rent does not exist"))
	bill, err = service.NewService(rt.repo).FinishRent(rt.ctx, token, rentID)
	rt.Require().Error(err)
	rt.Require().Equal(entity.Bill{}, bill)

	// Wrong token
	bill, err = service.NewService(rt.repo).FinishRent(rt.ctx, "12342141", "1")
	rt.Require().Error(err)
	rt.Require().Equal(entity.Bill{}, bill)
}

func (rt *RentTest) TestGetRentHistory() {
	// Correct Request
	res := []entity.RentHistory{
		{
			ID:           "1",
			CarID:        "1",
			UserID:       "1",
			CarClass:     "Premium",
			RentDuration: time.Minute * 1,
		},
		{
			ID:           "2",
			CarID:        "2",
			UserID:       "1",
			CarClass:     "Standard",
			RentDuration: time.Minute * 7,
		},
	}

	rt.repo.EXPECT().GetRentHistory(gomock.Any(), "1").Times(1).Return(res, nil)
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})
	res, err := service.NewService(rt.repo).GetRentHistory(rt.ctx, token)
	rt.Require().NoError(err)
	rt.Require().NotEqual(nil, res)
	for _, v := range res {
		rt.Require().Equal(v.UserID, "1")
	}

	// No rent history
	rt.repo.EXPECT().GetRentHistory(gomock.Any(), "1").Times(1).Return([]entity.RentHistory{}, nil)
	res, err = service.NewService(rt.repo).GetRentHistory(rt.ctx, token)
	rt.Require().NoError(err)
	rt.Require().Equal(res, []entity.RentHistory{})

	// Wrong token
	res, err = service.NewService(rt.repo).GetRentHistory(rt.ctx, "12342141")
	rt.Require().Error(err)
	rt.Require().Nil(res)
}

func (rt *RentTest) TestGetAvailableCars() {
	// Correct Request
	res := []entity.Car{
		{
			ID:    "1",
			Brand: "BMW",
			Class: "Premium",
		},
		{
			ID:    "2",
			Brand: "Mercedes-Benz",
			Class: "Premium",
		},
		{
			ID:    "3",
			Brand: "Ford",
			Class: "Standard",
		},
	}

	rt.repo.EXPECT().GetAvailableCars(gomock.Any()).Times(1).Return(res, nil)
	res, err := service.NewService(rt.repo).GetAvailableCars(rt.ctx)
	rt.Require().NoError(err)
	rt.Require().NotEqual(nil, res)

	// No available cars
	rt.repo.EXPECT().GetAvailableCars(gomock.Any()).Times(1).Return([]entity.Car{}, nil)
	res, err = service.NewService(rt.repo).GetAvailableCars(rt.ctx)
	rt.Require().NoError(err)
	rt.Require().Equal(res, []entity.Car{})

}
