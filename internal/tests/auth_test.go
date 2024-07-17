package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	appHandlers "github.com/MaksKazantsev/DriverGO/internal/handlers/http"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	mock_service "github.com/MaksKazantsev/DriverGO/internal/tests/mocks"
	"github.com/MaksKazantsev/DriverGO/internal/utils/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	s := AuthSuite{srvr: fiber.New()}

	go func() {
		require.NoError(t, s.srvr.Listen("localhost:3001"))
	}()

	suite.Run(t, &s)
}

type AuthSuite struct {
	suite.Suite

	cl   *http.Client
	ctrl *gomock.Controller
	srvr *fiber.App
}

func (at *AuthSuite) SetupTest() {
	at.cl = http.DefaultClient
	at.ctrl = gomock.NewController(at.T())
}

func (at *AuthSuite) TearDownTest() {
	at.ctrl.Finish()
}

func (at *AuthSuite) TestRegister() {
	clientMock := mock_service.NewMockAuthorization(at.ctrl)
	clientMock.EXPECT().
		Register(gomock.Any(), gomock.Eq(models.RegisterReq{Email: "string@gmail.com", Username: "string", Password: "stringpassword"})).
		Times(1).
		Return(models.AuthResponse{
			AccessToken:  "aToken",
			RefreshToken: "rToken",
			UUID:         "id5414134",
		}, nil)

	handler := appHandlers.RegisterAuthHandler(clientMock, validator.NewValidator())
	at.srvr.Post("/auth/register", handler.Register)

	body := models.RegisterReq{Email: "string@gmail.com", Username: "string", Password: "stringpassword"}
	b, err := json.Marshal(body)
	at.Require().NoError(err)

	url := "http://localhost:3001/auth/register"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	at.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("Sending request to /auth/register")
	res, err := at.cl.Do(req)
	at.Require().NoError(err)

	defer func() { _ = res.Body.Close() }()

	b, err = io.ReadAll(res.Body)
	at.Require().NoError(err)

	fmt.Printf("Response body: %s\n", string(b))
	fmt.Printf("Response status: %d\n", res.StatusCode)

	at.Require().Equal(http.StatusCreated, res.StatusCode, "Expected status code 200 Created")

	at.Require().Equal("{\"result\":{\"accessToken\":\"aToken\",\"refreshToken\":\"rToken\",\"UUID\":\"id5414134\"}}", string(b), "AccessToken should not be empty")
}
