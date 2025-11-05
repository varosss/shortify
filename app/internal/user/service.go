package user

import (
	"context"
	"errors"
	"shortify/internal/db"
	"shortify/internal/utils"

	"gorm.io/gorm"
)

type UserService struct {
	usersRepo *db.UsersRepo
}

func NewUserService(conn *gorm.DB) *UserService {
	return &UserService{usersRepo: db.NewUsersRepo(conn)}
}

func (s *UserService) Register(ctx context.Context, email string, password string) (*db.User, error) {
	_, err := s.usersRepo.FindOneByEmail(ctx, email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := s.usersRepo.InsertOne(ctx, &db.User{Email: email, PasswordHash: passwordHash})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, email string, password string) (*db.User, error) {
	user, err := s.usersRepo.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("user is not exist")
	}

	err = utils.ComparePassword(user.PasswordHash, password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
