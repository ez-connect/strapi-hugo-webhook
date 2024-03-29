package media

import (
	"encoding/json"
	"errors"
	"net/http"

	"strapiwebhook/helper"
	"strapiwebhook/helper/zlog"
	"strapiwebhook/service/config"
)

func HandleMedia(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.WriteHttpError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	req := &MediaPayload{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteHttpError(w, http.StatusBadRequest, errors.New("unable decode payload"))
		return
	}

	// Write the entry to a file, or delete a markdown file
	var (
		res    *Media
		status = http.StatusBadRequest
		err    error
	)

	// Download the all media formats
	switch req.Event {
	case EventMediaCreate, EventMediaUpdate:
		err = writeMedia(config.SiteDir, config.StrapiAddr, req)
	case EventMediaDelete:
		err = deleteMedia(config.SiteDir, req)
	default:
		err = errors.New("unknow event")
	}

	// Has any error?
	if err != nil {
		zlog.Warnw("entry", "status", status, "request", req, "error", err)
		helper.WriteHttpError(w, status, err)
		return
	}

	// OK
	zlog.Errorw("media", "status", http.StatusOK, "request", req)
	helper.WriteHttpResponse(w, http.StatusOK, res)

	// Post commands
	go func() {
		if config.Cmd != "" {
			helper.RunCommand(config.SiteDir, config.Cmd)
		}

		if config.DebouncedCmd != "" {
			helper.RunDebouncedCommand(config.SiteDir, config.DebouncedCmd)
		}
	}()
}
