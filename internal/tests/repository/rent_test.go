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
	preparedRentData    []entity.Rent
	preparedCarData     []entity.Car
	prepatedRentHistory []entity.RentHistory
)

type postgresRentSuite struct {
	suite.Suite

	ctx context.Context

	container          testcontainers.Container
	user, password, db string

	ctrl *gomock.Controller
	repo repositories.Rent
	conn *gorm.DB
}

func TestPostgresRentRepo(t *testing.T) {
	suite.Run(t, new(postgresRentSuite))
}

func (pr *postgresRentSuite) SetupTest() {
	// Setting up vars
	pr.user, pr.password, pr.db = "postgres", "postgres", "driverGO"
	pr.ctrl = gomock.NewController(pr.T())

	// connecting to a container
	ct, err := pr.newPostgresRentInstance()
	pr.Require().NoError(err)
	port, err := ct.MappedPort(context.Background(), "5432")
	pr.Require().NoError(err)

	// opening database
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s dbname=%s password=%s sslmode=disable", port.Port(), pr.user, pr.db, pr.password)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pr.Require().NoError(err)

	// migrations
	err = conn.AutoMigrate(
		&entity.User{},
		&entity.UserProfile{},
		&entity.Car{},
		&entity.Rent{},
		&entity.RentHistory{},
		&entity.Bill{},
	)
	pr.Require().NoError(err)

	// Setting up vars
	pr.conn = conn
	pr.container = ct
	pr.repo = myPostgres.NewRepository(conn)

	// Logger to ctx
	logger := mock_log.NewMockLogger(pr.ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	pr.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	pr.ctx = context.WithValue(pr.ctx, utils.LoggerKey, logger)

	preparedRentData = []entity.Rent{
		{
			ID:        "1",
			UserID:    "1",
			CarID:     "1",
			CarClass:  "Standard",
			StartTime: time.Now(),
		},
	}
	preparedCarData = []entity.Car{
		{
			ID:    "1",
			Class: "Standard",
			Brand: "Ford",
		},
		{
			ID:    "2",
			Class: "Premium",
			Brand: "BMW",
		},
	}
	prepatedRentHistory = []entity.RentHistory{
		{
			ID:           "1",
			CarID:        "1",
			UserID:       "1",
			CarClass:     "Standard",
			RentDuration: time.Minute * 30,
		},
	}

	pr.conn.Save(preparedCarData)
	pr.conn.Save(preparedRentData)
	pr.conn.Save(prepatedRentHistory)
}

func (pr *postgresRentSuite) TestStartRent() {
	// Correct request
	err := pr.repo.StartRent(pr.ctx, "1", "2")
	pr.Require().NoError(err)

	// Car already in rent
	err = pr.repo.StartRent(pr.ctx, "1", "1")
	pr.Require().Error(err)
	pr.Require().Equal(errors.NewError(errors.ERR_BAD_REQUEST, "car already in rent"), err)

	// Car does not exist
	err = pr.repo.StartRent(pr.ctx, "1", "123")
	pr.Require().Error(err)
	pr.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "car does not exist"), err)
}

func (pr *postgresRentSuite) TestFinishRent() {
	// Correct request
	bill, err := pr.repo.FinishRent(pr.ctx, "1", "1")
	pr.Require().NoError(err)
	pr.Require().NotEqual(entity.Bill{}, bill)
	pr.Require().Equal(bill.UserID, "1")

	// Rent does not exist
	bill, err = pr.repo.FinishRent(pr.ctx, "1", "2")
	pr.Require().Error(err)
	pr.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "entity not found"), err)
	pr.Require().Equal(entity.Bill{}, bill)
}

func (pr *postgresRentSuite) TestGetRentHistory() {
	// Correct request
	history, err := pr.repo.GetRentHistory(pr.ctx, "1")
	pr.Require().NoError(err)
	pr.Require().NotNil(history)
	pr.Require().GreaterOrEqual(len(history), 1)

	// Empty history
	history, err = pr.repo.GetRentHistory(pr.ctx, "2")
	pr.Require().NoError(err)
	pr.Require().NotNil(history)
	pr.Require().Equal(len(history), 0)
}

func (pr *postgresRentSuite) TestGetAvailableCars() {
	// Correct request
	cars, err := pr.repo.GetAvailableCars(pr.ctx)
	pr.Require().NoError(err)
	pr.Require().Equal(len(cars), len(preparedCarData)-len(preparedRentData))
	pr.Require().NotNil(cars)
}

func (pr *postgresRentSuite) newPostgresRentInstance() (testcontainers.Container, error) {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     pr.user,
				"POSTGRES_PASSWORD": pr.password,
				"POSTGRES_DB":       pr.db,
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
