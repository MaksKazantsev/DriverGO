package service

import (
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
)

type Service struct {
	Authorization
	Rent
	CarManagement
}

func NewService(repo repositories.Repository) *Service {
	return &Service{Authorization: NewAuth(repo), Rent: NewRent(repo), CarManagement: NewCarManagement(repo)}
}
