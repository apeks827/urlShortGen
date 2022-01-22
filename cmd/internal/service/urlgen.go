package service

import (
	"math/rand"
	"urlShortGen/cmd/internal/database"
)

type UrlListService struct {
	repo         database.UrlList
	stringLength int
	characters   []rune
}

func NewUrlListService(repo database.UrlList, stringLength int, characters []rune) *UrlListService {
	chars := make([]rune, 0, len(characters))
	chars = append(chars, characters...)
	return &UrlListService{
		repo:         repo,
		stringLength: stringLength,
		characters:   chars,
	}
}

func (s *UrlListService) PostUrl(longUrl string) (string, error) {
	var shortUrl string
	var err error
	shortUrlAvailable := false
	for !shortUrlAvailable {
		shortUrl = s.getUniqueString()
		shortUrlAvailable, err = s.repo.Available(shortUrl)
		if err != nil {
			return "", err
		}
	}
	err = s.repo.PostUrl(shortUrl, longUrl)
	if err != nil {
		shortUrl, errQuery := s.repo.GetShortUrl(longUrl)
		if errQuery != nil {
			return "", errQuery
		}
		return shortUrl, err
	}
	return shortUrl, nil
}

func (s *UrlListService) getUniqueString() string {
	rand.Seed(rand.Int63())
	uniqueRuneArray := make([]rune, s.stringLength)
	for i := range uniqueRuneArray {
		uniqueRuneArray[i] = s.characters[rand.Intn(len(s.characters))]
	}
	return string(uniqueRuneArray)
}

func (s *UrlListService) GetUrl(shortUrl string) (string, error) {
	return s.repo.GetUrl(shortUrl)
}
