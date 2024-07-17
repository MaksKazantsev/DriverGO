package http

import (
	"net/http"

	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthHandler(uc service.Authorization, v validator.Validator) *AuthHandler {
	return &AuthHandler{
		uc:        uc,
		validator: v,
	}
}

type AuthHandler struct {
	uc        service.Authorization
	validator validator.Validator
}

// Register godoc
// @Summary Register
// @Description Registers new user.
// @Tags Auth
// @Produce json
// @Param input body models.RegisterReq true "register request"
// @Success 201 {object} int
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 405 {object} string
// @Failure 500 {object} string
// @Router /auth/register [post]
func (a *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := a.validator.ValidateRegistration(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	data, err := a.uc.Register(c.UserContext(), req)
	if err != nil {
		st, msg := errors.FromError(err, c.Context())
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusCreated).JSON(fiber.Map{"result": data})
	return nil
}

// Login godoc
// @Summary Login
// @Description Logins user to the system.
// @Tags Auth
// @Produce json
// @Param input body models.LoginReq true "login request"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/login [put]
func (a *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginReq

	if err := c.BodyParser(&req); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := a.validator.ValidateLogin(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	data, err := a.uc.Login(c.UserContext(), req)
	if err != nil {
		st, msg := errors.FromError(err, c.Context())
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"result": data})
	return nil
}

// Refresh godoc
// @Summary Refresh
// @Description Updates token pair.
// @Tags Auth
// @Produce json
// @Param Authorization header string true "refresh token"
//
//	@Success        200 {object} int
//	@Failure        400 {object} string
//	@Failure        404 {object} string
//	@Failure        405 {object} string
//	@Failure        500 {object} string
//	@Router         /auth/refresh [get]
func (a *AuthHandler) Refresh(c *fiber.Ctx) error {
	token := extractAuthHeader(c)

	data, err := a.uc.Refresh(c.UserContext(), token)
	if err != nil {
		st, msg := errors.FromError(err, c.Context())
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"result": data})
	return nil
}
