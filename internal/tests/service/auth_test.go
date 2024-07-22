package service

import (
	"context"
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

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTest))
}
func (a *AuthTest) SetupTest() {
	a.ctrl = gomock.NewController(a.T())
	a.repo = mock_service.NewMockRepository(a.ctrl)

	logger := mock_log.NewMockLogger(a.ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	a.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	a.ctx = context.WithValue(a.ctx, utils.LoggerKey, logger)
}

func (a *AuthTest) TearDownTest() {
	a.ctrl.Finish()
}

type AuthTest struct {
	suite.Suite

	ctx context.Context

	ctrl *gomock.Controller
	repo *mock_service.MockRepository
}

func (a *AuthTest) TestRefreshToken() {
	res := models.AuthResponse{
		RefreshToken: "somerefresh",
		AccessToken:  "someaccess",
		UUID:         "1",
	}

	token, _ := utils.NewToken(utils.REFRESH, utils.TokenData{ID: "1", Role: "user"})
	a.repo.EXPECT().Refresh(gomock.Any(), "1", gomock.AssignableToTypeOf(token)).Times(1).Return(res, nil)

	res, err := service.NewService(a.repo).Refresh(a.ctx, token)
	a.Require().NoError(err)
	a.Require().Equal("1", res.UUID)
	a.Require().NotEqual(token, res.RefreshToken)
	a.Require().Equal(res.RefreshToken, "somerefresh")
}
