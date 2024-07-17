package models

type RegisterReq struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RToken   string `json:"-"`
	FBToken  string `json:"fb_token"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RToken   string `json:"-"`
	FBToken  string `json:"fb_token"`
}
