package handlers

import (
	appHandlers "github.com/MaksKazantsev/DriverGO/internal/handlers/http"
	"github.com/MaksKazantsev/DriverGO/internal/log"
	"github.com/MaksKazantsev/DriverGO/internal/metrics"
	"github.com/MaksKazantsev/DriverGO/internal/middleware"
	"github.com/MaksKazantsev/DriverGO/internal/middleware/wrappers"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Controller interface {
	SetupRoutes(app *fiber.App, log log.Logger, m metrics.Metrics)
}

type controller struct {
	srvc *service.Service

	UserHandler *appHandlers.UserHandler
	AuthHandler *appHandlers.AuthHandler
	RentHandler *appHandlers.RentHandler
	CarHandler  *appHandlers.CarHandler
}

func NewController(srvc *service.Service, m metrics.Metrics) Controller {
	return &controller{srvc: srvc, UserHandler: appHandlers.RegisterUserHandler(srvc.User, m), AuthHandler: appHandlers.RegisterAuthHandler(srvc.Authorization, validator.NewValidator(), m), RentHandler: appHandlers.RegisterRentHandler(srvc, m), CarHandler: appHandlers.RegisterCarHandler(srvc, m)}
}

func (c controller) SetupRoutes(app *fiber.App, log log.Logger, m metrics.Metrics) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	auth := app.Group("/auth").Use(wrappers.EmbedLogger(log), wrappers.WithIdempotencyKey(), wrappers.WithRequestLatency(m))
	auth.Post("/register", c.AuthHandler.Register)
	auth.Put("/login", c.AuthHandler.Login)
	auth.Get("/refresh", c.AuthHandler.Refresh)

	v1 := app.Group("/v1").Use(wrappers.EmbedLogger(log), wrappers.WithIdempotencyKey(), wrappers.WithRequestLatency(m))
	rent := v1.Group("/rent").Use(middleware.CheckAuth())
	rent.Post("/:carID", c.RentHandler.StartRent)
	rent.Delete("/:rentID", c.RentHandler.FinishRent)
	rent.Get("/history", c.RentHandler.GetRentHistory)
	rent.Get("/available", c.RentHandler.GetAvailableCars)

	admin := v1.Group("/admin").Use(middleware.RejectNotAdmin(), wrappers.EmbedLogger(log), wrappers.WithIdempotencyKey())
	admin.Post("/", c.CarHandler.AddCar)
	admin.Delete("/:carID", c.CarHandler.RemoveCar)
	admin.Put("/:carID", c.CarHandler.EditCar)

	user := v1.Group("/user").Use(middleware.CheckAuth(), wrappers.EmbedLogger(log), wrappers.WithIdempotencyKey())
	user.Get("/me", c.UserHandler.AboutMe)
	user.Get("/:userID", c.UserHandler.GetProfile)
	user.Get("/notifications", c.UserHandler.GetNotifications)

}
