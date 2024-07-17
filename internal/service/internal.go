package service

import (
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
)

func UserReqToEntity(data models.RegisterReq) entity.User {
	return entity.User{ID: data.ID, Username: data.Username, Password: data.Password, RFToken: data.RToken, FBToken: data.FBToken, Email: data.Email}
}
