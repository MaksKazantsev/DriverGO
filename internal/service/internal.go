package service

import (
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/entity"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"time"
)

func UserReqToEntity(data models.RegisterReq) entity.User {
	return entity.User{ID: data.ID, Username: data.Username, Password: data.Password, RFToken: data.RToken, FBToken: data.FBToken, Email: data.Email}
}

func FormatDuration(ms int64) string {
	duration := time.Duration(ms) * time.Millisecond

	days := duration / (24 * time.Hour)
	duration -= days * 24 * time.Hour

	hours := duration / time.Hour
	duration -= hours * time.Hour

	minutes := duration / time.Minute
	duration -= minutes * time.Minute

	seconds := duration / time.Second

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}
