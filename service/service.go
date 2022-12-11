package service

import (
	"errors"
	"net/http"
	"strapiwebhook/helper"
	"strapiwebhook/service/entry"
	"strapiwebhook/service/media"
)

type Service struct {
	http.Handler
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/entry":
		entry.HandleEntry(w, r)
	case "/media":
		media.HandleMedia(w, r)
	default:
		helper.WriteHttpError(w, http.StatusNotFound, errors.New("not found"))
	}
}
