package repository

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	myPostgres "github.com/MaksKazantsev/DriverGO/internal/repositories/postgres"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
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
)

var preparedData []entity.Car

type postgresCarSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	repo               repositories.CarManagement
	conn               *gorm.DB
	container          testcontainers.Container
	user, password, db string

	ctx context.Context
}

func TestPostgresCarRepo(t *testing.T) {
	suite.Run(t, new(postgresCarSuite))
}

func (pc *postgresCarSuite) SetupTest() {
	// Setting up vars
	pc.user, pc.password, pc.db = "postgres", "postgres", "driverGO"
	pc.ctrl = gomock.NewController(pc.T())

	// connection to container
	ct, err := pc.newPostgresCarInstance()
	pc.Require().NoError(err)
	port, err := ct.MappedPort(context.Background(), "5432")
	pc.Require().NoError(err)

	// opening database
	dsn := fmt.Sprintf("host=127.0.0.1 port=%s user=%s dbname=%s password=%s sslmode=disable", port.Port(), pc.user, pc.db, pc.password)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pc.Require().NoError(err)

	// migrations
	err = conn.AutoMigrate(
		&entity.User{},
		&entity.UserProfile{},
		&entity.Car{},
		&entity.Rent{},
		&entity.RentHistory{},
		&entity.Bill{},
	)
	pc.Require().NoError(err)

	// Setting up vars
	pc.conn = conn
	pc.container = ct
	pc.repo = myPostgres.NewRepository(conn)

	// Logger to ctx
	logger := mock_log.NewMockLogger(pc.ctrl)
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	pc.ctx = context.WithValue(context.Background(), utils.IdempotencyKey, uuid.NewString())
	pc.ctx = context.WithValue(pc.ctx, utils.LoggerKey, logger)

	// Prepared data
	preparedData = []entity.Car{
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

	pc.conn.Save(preparedData)
}

func (pc *postgresCarSuite) TestAddCar() {
	var req entity.Car
	req.ID = "3"
	req.Class = "Standard"
	req.Brand = "Audi"

	err := pc.repo.AddCar(pc.ctx, req)
	pc.Require().NoError(err)
}

func (pc *postgresCarSuite) TestRemoveCar() {
	// Correct request
	err := pc.repo.RemoveCar(pc.ctx, preparedData[0].ID)
	pc.Require().NoError(err)

	// Request to delete a not existing entity
	err = pc.repo.RemoveCar(pc.ctx, "154")
	pc.Require().Error(err)
}

func (pc *postgresCarSuite) TestEditCar() {
	// Correct request
	err := pc.repo.EditCar(pc.ctx, models.CarReq{Brand: "Mercedes-Benz", Class: "Premium"}, "1")
	pc.Require().NoError(err)

	// Request to edit a not existing entity
	err = pc.repo.EditCar(pc.ctx, models.CarReq{Brand: "BMW", Class: "Premium"}, "421")
	pc.Require().Error(err)
	pc.Require().Equal(errors.NewError(errors.ERR_NOT_FOUND, "car does not exist"), err)
}

func (p *postgresCarSuite) newPostgresCarInstance() (testcontainers.Container, error) {
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_USER":     p.user,
				"POSTGRES_PASSWORD": p.password,
				"POSTGRES_DB":       p.db,
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
