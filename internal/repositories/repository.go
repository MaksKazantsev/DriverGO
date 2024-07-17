package repositories

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
)

type Repository interface {
	Auth
}

type Auth interface {
	Register(ctx context.Context, data entity.User) (models.AuthResponse, error)
	Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error)
	Refresh(ctx context.Context, uuid, token string) (models.AuthResponse, error)
	GetPasswordAndID(ctx context.Context, email string) (string, string, error)
}
