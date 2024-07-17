package service

import (
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
)

type Service struct {
	Authorization
}

func NewService(repo repositories.Repository) *Service {
	return &Service{Authorization: NewAuth(repo)}
}
