package db

import (
	"context"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
)

type Repository interface {
	Auth
}

type Auth interface {
	Register(ctx context.Context, data models.RegisterReq) (models.AuthResponse, error)
	Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error)
}
