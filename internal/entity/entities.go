package entity

import "time"

type User struct {
	ID       string `json:"id" gorm:"id;primaryKey"`
	Email    string `json:"email" gorm:"email;unique"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	RFToken  string `json:"rf_token" gorm:"rf_token"`
	FBToken  string `json:"fb_token" gorm:"fb_token"`
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

type UserInfo struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Joined    string        `json:"joined"`
	Bio       string        `json:"bio"`
	Sex       string        `json:"sex"`
	RentHours time.Duration `json:"rent_hours"`
}
