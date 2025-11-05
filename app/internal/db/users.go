package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex;size:100;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ShortLinks   []ShortLink `gorm:"foreignKey:UserID"`
}

type UsersRepo struct {
	conn *gorm.DB
}

func NewUsersRepo(conn *gorm.DB) *UsersRepo {
	return &UsersRepo{conn: conn}
}

func (r *UsersRepo) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	user, err := gorm.G[User](r.conn).Where("email = ?", email).First(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepo) InsertOne(ctx context.Context, user *User) (*User, error) {
	err := gorm.G[User](r.conn).Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}
