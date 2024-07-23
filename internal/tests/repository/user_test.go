package repository

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	myPostgres "github.com/MaksKazantsev/DriverGO/internal/repositories/postgres"
	mock_log "github.com/MaksKazantsev/DriverGO/internal/tests/mocks/logger"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

var (
	preparedUserData          []entity.User
	preparedUserProfileData   []entity.UserProfile
	preparedNotificationsData []entity.Notification
)

type postgresUserSuite struct {
	suite.Suite

	ctx context.Context

	container          testcontainers.Container
	user, password, db string

	ctrl *gomock.Controller
	repo repositories.User
	conn *gorm.DB
}

func TestPostgresUserRepo(t *testing.T) {
	suite.Run(t, new(postgresUserSuite))
}

func (pu *postgresUserSuite) SetupTest() {
	// Setting up vars
	pu.user, pu.password, pu.db = "postgres", "postgres", "driverGO"
	pu.ctrl = gomock.NewController(pu.T())

	// connecting to a container
	ct, err := pu.newPostgresUserInstance()
	pu.Require().NoError(err)
	port, err := ct.MappedPort(context.Background(), "5432")
	pu.Require().NoError(err)

	// opening database
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s dbname=%s password=%s sslmode=disable", port.Port(), pu.user, pu.db, pu.password)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pu.Require().NoError(err)

	// migrations
	err = conn.AutoMigrate(
		&entity.User{},
		&entity.UserProfile{},
		&entity.Notification{},
	)
	pu.Require().NoError(err)

	// Setting up vars
	pu.conn = conn
	pu.container = ct
	pu.repo = myPostgres.NewRepository(conn)

	// Logger to ctx
	logger := mock_log.NewMockLogger(pu.ctrl)
	logger.EXPECT().Trace(gomock.Any(), gomock.Any()).AnyTimes()

	pu.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	pu.ctx = context.WithValue(pu.ctx, utils.LoggerKey, logger)

	preparedUserData = []entity.User{
		{
			ID:            "1",
			Username:      "Test",
			Email:         "testing@gmail.com",
			Notifications: 1,
		},
		{
			ID:            "3",
			Username:      "Test",
			Email:         "testing3@gmail.com",
			Notifications: 1,
		},
	}
	preparedUserProfileData = []entity.UserProfile{
		{
			ID:        "1",
			Joined:    time.Now(),
			RentHours: 5 * time.Hour,
			Email:     "testing@gmail.com",
		},
		{
			ID:        "3",
			Joined:    time.Now(),
			RentHours: 5 * time.Hour,
			Email:     "testing3@gmail.com",
		},
	}
	preparedNotificationsData = []entity.Notification{
		{
			UserID:    "1",
			Topic:     "test",
			Title:     "test",
			Body:      "test for test",
			CreatedAt: time.Now(),
		},
		{
			UserID:    "3",
			Topic:     "test",
			Title:     "test",
			Body:      "test for test",
			CreatedAt: time.Now(),
		},
	}

	pu.conn.Save(preparedUserData)
	pu.conn.Save(preparedUserProfileData)
	pu.conn.Save(preparedNotificationsData)
}

func (pu *postgresUserSuite) TestAboutMe() {
	// Correct request
	info, err := pu.repo.AboutMe(pu.ctx, "1")
	pu.Require().NoError(err)
	pu.Require().Equal(preparedUserProfileData[0].RentHours, info.RentHours)
	pu.Require().Equal(preparedUserData[0].Notifications, info.Notifications)

	// User does not exist
	info, err = pu.repo.AboutMe(pu.ctx, "2")
	pu.Require().Error(err)
	pu.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "entity not found"), err)
	pu.Require().Equal(entity.UserInfo{}, info)
}

func (pu *postgresUserSuite) TestGetProfile() {
	// Correct request
	profile, err := pu.repo.GetProfile(pu.ctx, "1")
	pu.Require().NoError(err)
	pu.Require().Equal(preparedUserProfileData[0].Email, profile.Email)

	// User does not exist
	profile, err = pu.repo.GetProfile(pu.ctx, "2")
	pu.Require().Error(err)
	pu.Require().Equal(entity.UserProfile{}, profile)
	pu.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "entity not found"), err)
}

func (pu *postgresUserSuite) TestGetNotifications() {
	// Correct request
	notis, err := pu.repo.GetNotifications(pu.ctx, "1")
	pu.Require().NoError(err)
	pu.Require().NotNil(notis)
	for _, v := range notis {
		pu.Require().Equal("1", v.UserID)
	}

	info, err := pu.repo.AboutMe(pu.ctx, "1")
	pu.Require().NoError(err)
	pu.Require().Equal(0, info.Notifications)

	// User does not exist
	notis, err = pu.repo.GetNotifications(pu.ctx, "2")
	pu.Require().Error(err)
	pu.Require().Nil(notis)
	pu.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "entity not found"), err)
}

func (pu *postgresUserSuite) newPostgresUserInstance() (testcontainers.Container, error) {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     pu.user,
				"POSTGRES_PASSWORD": pu.password,
				"POSTGRES_DB":       pu.db,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("database system is ready to accept connections"),
				wait.ForListeningPort("5432/tcp"),
			),
		},
		Started: true,
	})
	if err != nil {
		return nil, err
	}
	return container, nil
}
