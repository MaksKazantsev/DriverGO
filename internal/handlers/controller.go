package http

import "github.com/gofiber/fiber/v2"

type Controller interface {
	SetupRoutes(app *fiber.App)
}

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c controller) SetupRoutes(app *fiber.App) {
	//TODO implement me
	panic("implement me")
}
