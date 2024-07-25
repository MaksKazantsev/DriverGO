package http

import (
	"net/http"

	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/metrics"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthHandler(uc service.Authorization, v validator.Validator, m metrics.Metrics) *AuthHandler {
	return &AuthHandler{
		uc:        uc,
		validator: v,
		m:         m,
	}
}

type AuthHandler struct {
	uc        service.Authorization
	validator validator.Validator
	m         metrics.Metrics
}

// Register godoc
// @Summary Register
// @Description Registers new user.
// @Tags Auth
// @Produce json
// @Param input body models.RegisterReq true "register request"
// @Success 201 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
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
		a.m.Increment(st, "POST")
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	a.m.Increment(http.StatusCreated, "POST")
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
//	@Failure        400 {object} errors.HTTPError
//	@Failure        404 {object} errors.HTTPError
//	@Failure        405 {object} errors.HTTPError
//	@Failure        500 {object} errors.HTTPError
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
		a.m.Increment(st, "PUT")
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	a.m.Increment(http.StatusOK, "PUT")
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
//	@Failure        400 {object} errors.HTTPError
//	@Failure        404 {object} errors.HTTPError
//	@Failure        405 {object} errors.HTTPError
//	@Failure        500 {object} errors.HTTPError
//	@Router         /auth/refresh [get]
func (a *AuthHandler) Refresh(c *fiber.Ctx) error {
	token := extractAuthHeader(c)

	data, err := a.uc.Refresh(c.UserContext(), token)
	if err != nil {
		st, msg := errors.FromError(err, c.Context())
		a.m.Increment(st, "GET")
		_ = c.Status(st).JSON(errors.HTTPError{
			ErrorCode: st, ErrorMsg: msg,
		})
		return nil
	}

	a.m.Increment(http.StatusOK, "GET")
	_ = c.Status(http.StatusOK).JSON(fiber.Map{"result": data})
	return nil
}
