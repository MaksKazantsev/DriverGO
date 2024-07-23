package service

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/notifications"
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

type userTest struct {
	suite.Suite

	ctrl *gomock.Controller
	ctx  context.Context
	repo *mock_service.MockRepository
	noti notifications.Notifier
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userTest))
}

func (u *userTest) SetupTest() {
	u.ctrl = gomock.NewController(u.T())

	logger := mock_log.NewMockLogger(u.ctrl)
	logger.EXPECT().Trace(gomock.Any(), gomock.Any()).AnyTimes()
	u.noti = notifications.NewNotifier(u.repo)

	u.repo = mock_service.NewMockRepository(u.ctrl)

	u.ctx = context.WithValue(context.Background(), utils.LoggerKey, logger)
	u.ctx = context.WithValue(u.ctx, utils.IdempotencyKey, uuid.NewString())
}

func (u *userTest) TestAboutMe() {
	// Correct request
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})

	info := entity.UserInfo{
		Username: "Test",
		Email:    "testing@gmail.com",
		Joined:   time.Now(),
	}

	u.repo.EXPECT().AboutMe(gomock.Any(), "1").Times(1).Return(info, nil)

	infoRes, err := service.NewService(u.repo, u.noti).AboutMe(u.ctx, token)
	u.Require().NoError(err)
	u.Require().NotEqual(entity.UserInfo{}, infoRes)

	u.repo.EXPECT().AboutMe(gomock.Any(), "1").Times(1).Return(entity.UserInfo{}, errors.NewError(errors.ERR_NOT_FOUND, "entity does not exist"))

	infoRes, err = service.NewService(u.repo, u.noti).AboutMe(u.ctx, token)
	u.Require().Error(err)
	u.Require().Equal(entity.UserInfo{}, infoRes)
}

func (u *userTest) TestGetProfile() {
	// Correct request
	profile := entity.UserProfile{
		ID:       "1",
		Username: "Test",
		Email:    "testing@gmail.com",
	}

	u.repo.EXPECT().GetProfile(gomock.Any(), "1").Times(1).Return(profile, nil)

	profileRes, err := service.NewService(u.repo, u.noti).GetProfile(u.ctx, "1")
	u.Require().NoError(err)
	u.Require().NotEqual(entity.UserProfile{}, profileRes)

	// User does not exist
	u.repo.EXPECT().GetProfile(gomock.Any(), "1").Times(1).Return(entity.UserProfile{}, errors.NewError(errors.ERR_NOT_FOUND, "entity does not exist"))

	profileRes, err = service.NewService(u.repo, u.noti).GetProfile(u.ctx, "1")
	u.Require().Error(err)
	u.Require().Equal(entity.UserProfile{}, profileRes)
}

func (u *userTest) TestGetNotifications() {
	// Correct request
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})

	notifs := []entity.Notification{
		{
			UserID:    "1",
			Title:     "test",
			Topic:     "test",
			Body:      "test for test",
			CreatedAt: time.Now(),
		},
		{
			UserID:    "1",
			Title:     "test",
			Topic:     "test",
			Body:      "test for test",
			CreatedAt: time.Now(),
		},
	}

	u.repo.EXPECT().GetNotifications(gomock.Any(), "1").Times(1).Return(notifs, nil)

	notificationsRes, err := service.NewService(u.repo, u.noti).GetNotifications(u.ctx, token)
	u.Require().NoError(err)
	u.Require().NotEqual(notificationsRes, nil)
	for _, v := range notificationsRes {
		u.Require().Equal("1", v.UserID)
	}

	// Invalid token
	notificationsRes, err = service.NewService(u.repo, u.noti).GetNotifications(u.ctx, "invalidtoken")
	u.Require().Error(err)
	u.Require().Nil(notificationsRes)
}
