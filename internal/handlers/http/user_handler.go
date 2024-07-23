package http

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	uc service.User
}

// AboutMe godoc
// @Summary AboutMe
// @Description Gets main information about yourself.
// @Tags User
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/user/me [get]
func (u *UserHandler) AboutMe(c *fiber.Ctx) error {
	token := extractAuthHeader(c)

	info, err := u.uc.AboutMe(c.UserContext(), token)
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"info": info})
	return nil
}

// GetProfile godoc
// @Summary GetProfile
// @Description Gets user`s profile.
// @Tags User
// @Produce json
// @Param userID path string true "userID"
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/{userID} [get]
func (u *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Params("userID")

	profile, err := u.uc.GetProfile(c.UserContext(), userID)
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"profile": profile})
	return nil
}

// GetNotifications godoc
// @Summary GetNotifications
// @Description Gets user`s profile.
// @Tags User
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/notifications [get]
func (u *UserHandler) GetNotifications(c *fiber.Ctx) error {
	token := extractAuthHeader(c)

	notifications, err := u.uc.GetNotifications(c.UserContext(), token)
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	_ = c.Status(http.StatusOK).JSON(fiber.Map{"notifications": notifications})
	return nil
}
