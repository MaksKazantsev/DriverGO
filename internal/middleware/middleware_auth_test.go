package middleware

import (
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

type MiddleWareSuitTest struct {
	suite.Suite
	srvr *fiber.App
}

func TestMiddleWareAuth(t *testing.T) {
	suite.Run(t, new(MiddleWareSuitTest))
}

func (m *MiddleWareSuitTest) SetupTest() {
	m.srvr = fiber.New()

	checkAuth := m.srvr.Group("/check")
	checkAuth.Use(CheckAuth())
	checkAuth.Get("/auth", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).SendString("Authorized")
	})

	rejectNotAdmin := m.srvr.Group("/reject")
	rejectNotAdmin.Use(RejectNotAdmin())
	rejectNotAdmin.Get("/not_admin", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).SendString("Authorized")
	})
}

func (m *MiddleWareSuitTest) TestRejectNotAdmin() {
	// Correct token
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "admin"})

	// Correct token request
	req := httptest.NewRequest(http.MethodGet, "/reject/not_admin", nil)
	req.Header.Set("Authorization", "Bearer: "+token)

	resp, err := m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusOK, resp.StatusCode)

	// Incorrect token request
	req = httptest.NewRequest(http.MethodGet, "/reject/not_admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusBadRequest, resp.StatusCode)

	// Incorrect permission level
	token, _ = utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})

	// Correct token request
	req = httptest.NewRequest(http.MethodGet, "/reject/not_admin", nil)
	req.Header.Set("Authorization", "Bearer: "+token)

	resp, err = m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusMethodNotAllowed, resp.StatusCode)

}

func (m *MiddleWareSuitTest) TestCheckAuth() {
	// Correct token
	token, _ := utils.NewToken(utils.ACCESS, utils.TokenData{ID: "1", Role: "user"})

	// Correct token request
	req := httptest.NewRequest(http.MethodGet, "/check/auth", nil)
	req.Header.Set("Authorization", "Bearer: "+token)

	resp, err := m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusOK, resp.StatusCode)

	// Wrong token request
	req = httptest.NewRequest(http.MethodGet, "/check/auth", nil)
	req.Header.Set("Authorization", "Bearer: wrongtoken")

	resp, err = m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusMethodNotAllowed, resp.StatusCode)

	// Wrong "Bearer: "
	req = httptest.NewRequest(http.MethodGet, "/check/auth", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusBadRequest, resp.StatusCode)

	// No token provided
	req = httptest.NewRequest(http.MethodGet, "/check/auth", nil)
	req.Header.Set("Authorization", "Bearer: "+"")

	resp, err = m.srvr.Test(req)
	m.Require().NoError(err)
	m.Require().Equal(http.StatusMethodNotAllowed, resp.StatusCode)
}
