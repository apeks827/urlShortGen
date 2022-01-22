package database

type UrlArchive struct {
	UrlList
}

type UrlList interface {
	GetUrl(shortUrl string) (string, error)
	GetShortUrl(longUrl string) (string, error)
	PostUrl(shortUrl string, longUrl string) error
	Available(shortUrl string) (bool, error)
}

func NewArchive(db UrlList) *UrlArchive {
	return &UrlArchive{
		UrlList: db,
	}
}
