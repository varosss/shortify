package auth

import (
	"context"
	"fmt"
	"shortify/internal/db"
	"shortify/internal/user"
	"time"

	"gorm.io/gorm"
)

type AuthService struct {
	jwtManager  *JWTManager
	userService *user.UserService
}

func NewAuthService(conn *gorm.DB) *AuthService {
	return &AuthService{
		jwtManager:  NewJWTManager(time.Hour),
		userService: user.NewUserService(conn),
	}
}

func (s *AuthService) AuthUser(
	ctx context.Context,
	email string,
	password string,
	action AuthAction,
) (string, error) {
	var (
		user *db.User
		err  error
	)

	switch action {
	case RegisterAction:
		user, err = s.userService.Register(ctx, email, password)
	case LoginAction:
		user, err = s.userService.Login(ctx, email, password)
	default:
		return "", fmt.Errorf("unknown action: %s", string(action))
	}

	if err != nil {
		return "", err
	}

	auth_token, err := s.jwtManager.Generate(int64(user.ID))
	if err != nil {
		return "", err
	}

	return auth_token, nil
}
