package entity

import "time"

type User struct {
	ID            string `json:"id" gorm:"id;primaryKey"`
	Email         string `json:"email" gorm:"email;unique"`
	Username      string `json:"username" gorm:"username"`
	Password      string `json:"password" gorm:"password"`
	RFToken       string `json:"rf_token" gorm:"rf_token"`
	FBToken       string `json:"fb_token" gorm:"fb_token"`
	Notifications int    `json:"notifications" gorm:"notifications"`
}

type UserProfile struct {
	ID        string        `json:"id" gorm:"primaryKey"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Joined    time.Time     `json:"joined" gorm:"default:CURRENT_TIMESTAMP"`
	Bio       string        `json:"bio" gorm:"default:not stated"`
	Sex       string        `json:"sex" gorm:"default:not stated"`
	RentHours time.Duration `json:"rent_hours"`
}

type Car struct {
	ID    string `json:"ID" gorm:"primaryKey"`
	Brand string `json:"brand" gorm:"brand"`
	Class string `json:"class" gorm:"class"`
}

type Rent struct {
	ID        string    `json:"ID" gorm:"primaryKey"`
	CarID     string    `json:"car_id" gorm:"car_id;unique"`
	UserID    string    `json:"user_id" gorm:"user_id"`
	CarClass  string    `json:"car_class" gorm:"car_class"`
	StartTime time.Time `json:"start_time" gorm:"start_time"`
}

type RentHistory struct {
	ID           string        `json:"ID" gorm:"primaryKey"`
	CarID        string        `json:"car_id" gorm:"car_id"`
	UserID       string        `json:"user_id" gorm:"user_id"`
	CarClass     string        `json:"car_class" gorm:"car_class"`
	RentDuration time.Duration `json:"rent_duration" gorm:"rent_duration"`
}

type Bill struct {
	ID     string  `json:"id" gorm:"primaryKey"`
	UserID string  `json:"user_id" gorm:"user_id"`
	Price  float64 `json:"price" gorm:"price"`
}

type UserInfo struct {
	Username      string        `json:"username"`
	Email         string        `json:"email"`
	Joined        time.Time     `json:"joined"`
	Bio           string        `json:"bio"`
	Sex           string        `json:"sex"`
	Notifications int           `json:"notifications"`
	RentHours     time.Duration `json:"rent_hours"`
}

type Notification struct {
	UserID    string    `json:"userID"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
