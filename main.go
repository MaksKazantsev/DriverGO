package main

import (
	_ "github.com/MaksKazantsev/DriverGO/docs"
	"github.com/MaksKazantsev/DriverGO/internal/app"
	"github.com/MaksKazantsev/DriverGO/internal/config"
)

//go:generate mockgen -source=./internal/service/auth.go -destination=./internal/tests/mocks/mock_auth.go

// @title DriverGO server API
// @version 1.0

func main() {
	app.MustStart(config.MustLoad())
}
