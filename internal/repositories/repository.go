package repositories

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
)

type Repository interface {
	Auth
	Rent
	User
	CarManagement
	NotifierRepo
}

type Auth interface {
	Register(ctx context.Context, data entity.User) (models.AuthResponse, error)
	Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error)
	Refresh(ctx context.Context, uuid, token string) (models.AuthResponse, error)
	GetPasswordAndID(ctx context.Context, email string) (string, string, error)
}
type Rent interface {
	StartRent(ctx context.Context, userID, carID string) error
	FinishRent(ctx context.Context, userID, rentID string) (entity.Bill, error)
	GetRentHistory(ctx context.Context, userID string) ([]entity.RentHistory, error)
	GetAvailableCars(ctx context.Context) ([]entity.Car, error)
}
type CarManagement interface {
	AddCar(ctx context.Context, car entity.Car) error
	RemoveCar(ctx context.Context, carID string) error
	EditCar(ctx context.Context, data models.CarReq, carID string) error
}

type User interface {
	AboutMe(ctx context.Context, userID string) (entity.UserInfo, error)
	GetProfile(ctx context.Context, userID string) (entity.UserProfile, error)
	GetNotifications(ctx context.Context, userID string) ([]entity.Notification, error)
}

type NotifierRepo interface {
	GetFBToken(ctx context.Context, userID string) (string, error)
	SaveNotification(ctx context.Context, notification entity.Notification) error
}
