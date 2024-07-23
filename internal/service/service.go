package service

import (
	"github.com/MaksKazantsev/DriverGO/internal/notifications"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
)

type Service struct {
	Authorization
	Rent
	CarManagement
	User
}

func NewService(repo repositories.Repository, notifier notifications.Notifier) *Service {
	return &Service{Authorization: NewAuth(repo), Rent: NewRent(repo, notifier), CarManagement: NewCarManagement(repo), User: NewUser(repo)}
}
