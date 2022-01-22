package inmemory

import (
	"errors"
)

type InMemoryDB struct {
	longUrlsMap map[string]string
}

func NewInMemoryDB() *InMemoryDB {
	longUrlsMap := make(map[string]string)
	return &InMemoryDB{
		longUrlsMap: longUrlsMap,
	}
}

func (db *InMemoryDB) GetUrl(shortUrl string) (string, error) {
	if shortUrl == "" {
		return "", errors.New("InMemoryDB: wrong initial data to read")
	}
	longUrl, ok := db.longUrlsMap[shortUrl]
	if ok {
		return longUrl, nil
	}
	return "", nil
}

func (db *InMemoryDB) GetShortUrl(longUrl string) (string, error) {
	if longUrl == "" {
		return "", errors.New("InMemoryDB: wrong initial data to read")
	}
	for key, value := range db.longUrlsMap {
		if longUrl == value {
			return key, nil
		}
	}
	return "", nil
}

func (db *InMemoryDB) PostUrl(shortUrl string, longUrl string) error {
	if shortUrl == "" || longUrl == "" {
		return errors.New("InMemoryDB: wrong initial data to write")
	}

	isExistShortUrl, _ := db.GetShortUrl(longUrl)
	if isExistShortUrl != "" {
		return errors.New("InMemoryDB: already have this link")
	}

	isExistLongUrl, _ := db.GetUrl(shortUrl)
	if isExistLongUrl != "" {
		return errors.New("InMemoryDB: already have this shortlink")
	}

	db.longUrlsMap[shortUrl] = longUrl
	return nil
}

func (db *InMemoryDB) Available(shortUrl string) (bool, error) {
	if shortUrl == "" {
		return false, errors.New("InMemoryDB: wrong initial data to write")
	}
	_, ok := db.longUrlsMap[shortUrl]
	return !ok, nil
}
