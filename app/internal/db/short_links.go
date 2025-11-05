package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ShortLink struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"uniqueIndex"`
	OriginalUrl string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ShortLinksRepo struct {
	db *gorm.DB
}

func NewShortLinksRepo(db *gorm.DB) *ShortLinksRepo {
	return &ShortLinksRepo{db: db}
}

func (r *ShortLinksRepo) FindManyByUserID(ctx context.Context, userId int) ([]ShortLink, error) {
	shortLinks, err := gorm.G[ShortLink](r.db).Where("user_id = ?", userId).Find(ctx)
	if err != nil {
		return nil, err
	}

	return shortLinks, nil
}

func (r *ShortLinksRepo) FindOneByCode(ctx context.Context, code string) (*ShortLink, error) {
	shortLink, err := gorm.G[ShortLink](r.db).Where("code = ?", code).First(ctx)
	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

func (r *ShortLinksRepo) FindOneByCodeAndUserId(ctx context.Context, code string, userId int) (*ShortLink, error) {
	shortLink, err := gorm.G[ShortLink](r.db).Where("code = ? AND user_id = ?", code, userId).First(ctx)
	if err != nil {
		return nil, err
	}

	return &shortLink, nil
}

func (r *ShortLinksRepo) DeleteOneByCodeAndUserId(ctx context.Context, code string, userId int) error {
	_, err := gorm.G[ShortLink](r.db).Where("code = ? AND user_id = ?", code, userId).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShortLinksRepo) InsertBatch(ctx context.Context, shortLinks *[]ShortLink) {
	r.db.Create(shortLinks)
}
