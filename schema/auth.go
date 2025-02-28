package schema

type ILoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ILoginRes struct {
	AccessToken string `json:"accessToken"`
}
