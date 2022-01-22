package service

import (
	"urlShortGen/cmd/internal/database"
)

type UrlList interface {
	GetUrl(shortUrl string) (string, error)
	PostUrl(longUrl string) (string, error)
}

type Service struct {
	UrlList
}

func NewService(repo *database.UrlArchive, stringLength int, characters []rune) *Service {
	return &Service{
		UrlList: NewUrlListService(repo.UrlList, stringLength, characters),
	}
}
