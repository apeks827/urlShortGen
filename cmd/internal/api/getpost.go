package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"urlShortGen/cmd/internal/service"
)

type Api struct {
	services *service.Service
	errorLog *log.Logger
	infoLog  *log.Logger
}

func NewRequest(services *service.Service, errorLog *log.Logger, infoLog *log.Logger) *Api {
	return &Api{
		services: services,
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func (handle *Api) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle.shortUrlsHandler)
	return mux
}

func (handle *Api) shortUrlsHandler(write http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/" {
		handle.notFound(write)
		return
	}

	if request.Method == http.MethodPost {
		if request.Header.Get("Content-Type") != "text/plain; charset=utf-8" {
			handle.clientError(write, http.StatusUnsupportedMediaType)
			return
		}

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			handle.serverError(write)
			return
		}

		url := string(body)
		if url == "" {
			handle.clientError(write, http.StatusBadRequest)
			return
		}

		shortUrl, err := handle.services.PostUrl(url)

		if err != nil {
			if shortUrl != "" {
				handle.clientError(write, http.StatusBadRequest)
				write.Write([]byte(fmt.Sprintf("For URL: %s has already been generated: %s\n", url, shortUrl)))
				return
			}
			handle.serverError(write)
			return
		}

		write.Write([]byte(shortUrl))
		return
	}

	if request.Method == http.MethodGet {
		shortUrl := request.URL.Query().Get("url")
		if shortUrl == "" {
			handle.clientError(write, http.StatusBadRequest)
			return
		}

		url, err := handle.services.GetUrl(shortUrl)
		if err != nil {
			handle.serverError(write)
			return
		}
		if url == "" {
			handle.clientError(write, http.StatusBadRequest)
			return
		}

		write.Write([]byte(url))
		return
	}

	write.Header().Set("Allow", http.MethodPost+"; "+http.MethodGet)
	handle.clientError(write, http.StatusMethodNotAllowed)
}
