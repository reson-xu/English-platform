package user

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserResp struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type RegisterResp struct {
	User        AuthUserResp `json:"user"`
	AccessToken string       `json:"accessToken"`
}

type LoginResp = RegisterResp
