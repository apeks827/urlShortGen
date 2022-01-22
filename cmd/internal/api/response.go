package api

import (
	"net/http"
)

func (handle *Api) serverError(write http.ResponseWriter) {
	http.Error(write, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (handle *Api) clientError(write http.ResponseWriter, status int) {
	http.Error(write, http.StatusText(status), status)
}

func (handle *Api) notFound(write http.ResponseWriter) {
	handle.clientError(write, http.StatusNotFound)
}
