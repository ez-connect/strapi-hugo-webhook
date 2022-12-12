package entry

import (
	"encoding/json"
	"errors"
	"net/http"

	"strapiwebhook/helper"
	"strapiwebhook/helper/zlog"
	"strapiwebhook/service/config"
)

// Handle `service.EntryPath` endpoint
func HandleEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helper.WriteHttpError(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	req := &EntryPayload{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteHttpError(w, http.StatusBadRequest, errors.New("unable decode payload"))
		return
	}

	// Write the entry to a file, or delete a markdown file
	var (
		res    *Entry
		status = http.StatusBadRequest
		err    error
	)

	switch req.Event {
	case eventEntryCreate, eventEntryUpdate:
		res, err = writeEntry(config.SiteDir, config.TemplateDir, req)

	case eventEntryDelete:
		res, err = deleteEntry(config.SiteDir, req)

	default:
		err = errors.New("events ignored")
	}

	// Has any error?
	if err != nil {
		zlog.Warnw("entry", "status", status, "request", req, "error", err)
		helper.WriteHttpError(w, status, err)
		return
	}

	// OK
	zlog.Infow("entry", "status", http.StatusOK, "request", req)
	helper.WriteHttpResponse(w, http.StatusOK, res)

	// Post commands
	go func() {
		if config.PostCmd != "" {
			helper.RunCommand(config.SiteDir, config.PostCmd)
		}

		if config.PostDebouncedCmd != "" {
			helper.RunDebouncedCommand(config.SiteDir, config.PostDebouncedCmd)
		}
	}()
}
