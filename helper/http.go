package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteHttpError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func WriteHttpResponse(w http.ResponseWriter, v any) {
	buf, err := json.Marshal(v)
	if err != nil {
		WriteHttpError(w, http.StatusBadGateway, errors.New("marshal response error"))
		return
	}

	w.Write(buf)
}
