package handlers

import (
	appHandlers "github.com/MaksKazantsev/DriverGO/internal/handlers/http"
	"github.com/MaksKazantsev/DriverGO/internal/log"
	"github.com/MaksKazantsev/DriverGO/internal/middleware/wrappers"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Controller interface {
	SetupRoutes(app *fiber.App, log log.Logger)
}

type controller struct {
	srvc *service.Service

	UserHandler *appHandlers.UserHandler
	AuthHandler *appHandlers.AuthHandler
}

func NewController(srvc *service.Service) Controller {
	return &controller{srvc: srvc, AuthHandler: appHandlers.RegisterAuthHandler(srvc.Authorization, validator.NewValidator())}
}

func (c controller) SetupRoutes(app *fiber.App, log log.Logger) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	auth := app.Group("/auth").Use(wrappers.EmbedLogger(log), wrappers.WithIdempotencyKey())
	auth.Post("/register", c.AuthHandler.Register)
	auth.Put("/login", c.AuthHandler.Login)
	auth.Get("/refresh", c.AuthHandler.Refresh)
}
