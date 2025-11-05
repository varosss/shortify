package links

import (
	"context"
	"fmt"
	"shortify/internal/db"
	"shortify/internal/utils"

	"gorm.io/gorm"
)

type LinksService struct {
	shortLinksRepo *db.ShortLinksRepo
}

func NewLinksService(conn *gorm.DB) *LinksService {
	return &LinksService{shortLinksRepo: db.NewShortLinksRepo(conn)}
}

func (s *LinksService) AddLinks(ctx context.Context, userId int, links []LinkData) []LinkData {
	shortLinks := []db.ShortLink{}

	for _, link := range links {
		shortLinks = append(
			shortLinks,
			db.ShortLink{
				Code:        utils.GenerateShortULID(),
				OriginalUrl: link.Url,
				UserID:      uint(userId),
			},
		)
	}

	s.shortLinksRepo.InsertBatch(ctx, &shortLinks)

	linksForResult := []LinkData{}
	for _, link := range shortLinks {
		linksForResult = append(
			linksForResult,
			MakeLinkFromShortLink(&link),
		)
	}

	return linksForResult
}

func (s *LinksService) GetLink(ctx context.Context, code string, userId int) (*LinkData, error) {
	shortLink, err := s.shortLinksRepo.FindOneByCodeAndUserId(ctx, code, userId)
	if err != nil {
		return nil, fmt.Errorf("short link not found: user_id=%d, code=%s", userId, code)
	}

	link := MakeLinkFromShortLink(shortLink)

	return &link, nil
}

func (s *LinksService) DeleteLink(ctx context.Context, code string, userId int) error {
	err := s.shortLinksRepo.DeleteOneByCodeAndUserId(ctx, code, userId)
	if err != nil {
		return fmt.Errorf("couldn't delete link: %s", err.Error())
	}

	return nil
}

func (s *LinksService) GetLinks(ctx context.Context, userId int) ([]LinkData, error) {
	shortLinks, err := s.shortLinksRepo.FindManyByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	linksData := []LinkData{}
	for _, shortLink := range shortLinks {
		linksData = append(linksData, MakeLinkFromShortLink(&shortLink))
	}

	return linksData, nil
}
