package auth

type ErrorAuthResponse struct {
	Error string `json:"error"`
}

type SuccessAuthResponse struct {
	AuthToken string `json:"auth_token"`
}
