package app

import (
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/config"
	"github.com/MaksKazantsev/DriverGO/internal/handlers"
	"github.com/MaksKazantsev/DriverGO/internal/log"
	"github.com/MaksKazantsev/DriverGO/internal/notifications"
	"github.com/MaksKazantsev/DriverGO/internal/repositories/postgres"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// Loading .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("failed to load .env file: " + err.Error())
	}

	// New logger
	l := log.MustInit(cfg.Env)

	// New repository
	repo := postgres.NewRepository(postgres.MustConnect(cfg.DB.Postgres.GetDSN()))
	l.Info("repository layer init success", nil)

	// New service
	srvc := service.NewService(repo, notifications.NewNotifier(repo))
	l.Info("service layer init success", nil)

	// New controller
	ctrl := handlers.NewController(srvc)
	l.Info("controller init success", nil)

	// New fiber app
	app := fiber.New()
	ctrl.SetupRoutes(app, l)

	run(func() {
		l.Info("starting server...", log.WithData("port", cfg.Port))
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			panic("failed to listen to TCP: " + err.Error())
		}
	}, app)
}

func run(fn func(), app *fiber.App) {
	go fn()

	chDone := make(chan os.Signal, 1)
	signal.Notify(chDone, syscall.SIGINT|syscall.SIGTERM)
	<-chDone

	if err := app.Shutdown(); err != nil {
		panic("failed to shutdown app: " + err.Error())
	}
}
