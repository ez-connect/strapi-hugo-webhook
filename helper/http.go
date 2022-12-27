package helper

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"strapiwebhook/helper/zlog"
)

type HttpResponse struct {
	Request    *http.Request
	Response   *http.Response
	StatusCode int
	Body       *[]byte
}

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

func HttpSendRequest(method, uri string, headers map[string]string, data io.Reader) (*HttpResponse, error) {
	req, err := http.NewRequest(method, uri, data)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Method = method

	client := &http.Client{
		// x509: certificate signed by unknown authority error
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &HttpResponse{
		Request:    req,
		Response:   resp,
		StatusCode: resp.StatusCode,
		Body:       &body,
	}

	return res, nil
}

func HttpGet(uri string, headers map[string]string) (*HttpResponse, error) {
	return HttpSendRequest(http.MethodGet, uri, headers, nil)
}

func HttpPost(uri string, headers map[string]string, data io.Reader) (*HttpResponse, error) {
	return HttpSendRequest(http.MethodPost, uri, headers, data)
}
