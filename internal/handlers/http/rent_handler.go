package http

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/metrics"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type RentHandler struct {
	uc service.Rent
	m  metrics.Metrics
}

func RegisterRentHandler(uc service.Rent, m metrics.Metrics) *RentHandler {
	return &RentHandler{
		uc: uc,
		m:  m,
	}
}

// StartRent godoc
// @Summary StartRent
// @Description Starts new rent.
// @Tags Rent
// @Produce json
// @Param carID path string true "car ID"
// @Param Authorization header string true "token"
// @Success 201 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/rent/{carID} [post]
func (r *RentHandler) StartRent(c *fiber.Ctx) error {
	token := extractAuthHeader(c)
	carID := c.Params("carID")

	if err := r.uc.StartRent(c.UserContext(), token, carID); err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		r.m.Increment(st, c.Method())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	r.m.Increment(http.StatusCreated, c.Method())
	c.Status(http.StatusCreated)
	return nil
}

// FinishRent godoc
// @Summary FinishRent
// @Description Finishes rent.
// @Tags Rent
// @Produce json
// @Param rentID path string true "rent ID"
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/rent/{rentID} [delete]
func (r *RentHandler) FinishRent(c *fiber.Ctx) error {
	token := extractAuthHeader(c)
	rentID := c.Params("rentID")

	bill, err := r.uc.FinishRent(c.UserContext(), token, rentID)
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		r.m.Increment(st, c.Method())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	r.m.Increment(http.StatusOK, c.Method())
	_ = c.Status(http.StatusOK).JSON(fiber.Map{"bill": bill})
	return nil
}

// GetRentHistory godoc
// @Summary GetRentHistory
// @Description Gets all user`s rents.
// @Tags Rent
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/rent/history [get]
func (r *RentHandler) GetRentHistory(c *fiber.Ctx) error {
	token := extractAuthHeader(c)

	rents, err := r.uc.GetRentHistory(c.UserContext(), token)
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		r.m.Increment(st, c.Method())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	r.m.Increment(http.StatusOK, c.Method())
	_ = c.Status(http.StatusOK).JSON(fiber.Map{"rents": rents})
	return nil
}

// GetAvailableCars godoc
// @Summary GetAvailableCars
// @Description Gets all available cars at the moment.
// @Tags Rent
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/rent/available [get]
func (r *RentHandler) GetAvailableCars(c *fiber.Ctx) error {
	cars, err := r.uc.GetAvailableCars(c.UserContext())
	if err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		r.m.Increment(st, c.Method())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	r.m.Increment(http.StatusOK, c.Method())
	_ = c.Status(http.StatusOK).JSON(fiber.Map{"cars": cars})
	return nil
}
