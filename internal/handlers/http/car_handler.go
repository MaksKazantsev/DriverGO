package http

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/service"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CarHandler struct {
	uc service.CarManagement
}

func RegisterCarHandler(uc service.CarManagement) *CarHandler {
	return &CarHandler{
		uc: uc,
	}
}

// AddCar godoc
// @Summary AddCar
// @Description Adds new car to the pool of available cars. Can be executed only by admin.
// @Tags CarManagement
// @Produce json
// @Param input body models.CarReq true "car request"
// @Param Authorization header string true "token"
// @Success 201 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/admin/add [post]
func (cr *CarHandler) AddCar(c *fiber.Ctx) error {
	var req models.CarReq

	err := c.BodyParser(&req)
	if err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err = cr.uc.AddCar(c.UserContext(), req); err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	c.Status(http.StatusCreated)
	return nil
}

// RemoveCar godoc
// @Summary RemoveCar
// @Description Removes car from the pool of available cars. Can be executed only by admin.
// @Tags CarManagement
// @Produce json
// @Param carID path string true "car`s ID"
// @Param Authorization header string true "token"
// @Success 201 {object} int
// @Failure 400 {object} errors.HTTPError
// @Failure 404 {object} errors.HTTPError
// @Failure 405 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /v1/admin/remove/{carID} [delete]
func (cr *CarHandler) RemoveCar(c *fiber.Ctx) error {
	carID := c.Params("carID")

	if err := cr.uc.RemoveCar(c.UserContext(), carID); err != nil {
		st, msg := errors.FromError(err, c.UserContext())
		_ = c.Status(st).JSON(errors.HTTPError{ErrorCode: st, ErrorMsg: msg})
		return nil
	}

	c.Status(http.StatusOK)
	return nil
}