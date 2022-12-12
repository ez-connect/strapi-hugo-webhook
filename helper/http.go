package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strapiwebhook/helper/zlog"
)

func WriteHttpError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if _, err := w.Write([]byte(err.Error())); err != nil {
		zlog.Errorw("write http error", "error", err)
	}
}

func WriteHttpResponse(w http.ResponseWriter, status int, v any) {
	buf, err := json.Marshal(v)
	if err != nil {
		WriteHttpError(w, http.StatusBadGateway, errors.New("marshal response error"))
		return
	}

	if _, err = w.Write(buf); err != nil {
		zlog.Errorw("write http response", "error", err)
	}
}
