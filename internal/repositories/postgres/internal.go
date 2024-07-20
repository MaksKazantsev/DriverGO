package postgres

import (
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"time"
)

func calculateRentPrice(class string, startTime time.Time) (float64, error) {
	rentTime := time.Since(startTime)
	minutes := int(rentTime.Minutes())
	hours := int(rentTime.Hours())

	switch class {
	case "Standard":
		if minutes < 60 {
			return float64(minutes) * 0.1, nil
		}
		return float64(hours) * 5, nil
	case "Premium":
		if minutes < 60 {
			return float64(minutes) * 0.3, nil
		}
		return float64(hours) * 15, nil
	}
	return 0, errors.NewError(errors.ERR_BAD_REQUEST, "invalid class type")
}
