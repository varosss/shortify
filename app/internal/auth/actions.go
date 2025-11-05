package auth

type AuthAction string

const (
	RegisterAction AuthAction = "register"
	LoginAction    AuthAction = "login"
)
