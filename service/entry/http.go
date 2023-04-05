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

	payload := &EntryPayload{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		helper.WriteHttpError(w, http.StatusBadRequest, errors.New("unable decode payload"))
		return
	}

	entry := GetEntry(payload)
	if entry.Type == entryTypeIngore {
		zlog.Infow("entry", "ingore", entry.Model)
		helper.WriteHttpError(w, http.StatusNotImplemented, errors.New("ingore entry"))
		return
	}

	// Write the entry to a file, or delete a markdown file
	var (
		status = http.StatusBadRequest
		err    error
	)

	switch payload.Event {
	case EventEntryCreate, EventEntryUpdate:
		err = WriteEntry(config.SiteDir, config.TemplateDir, entry)

	case EventEntryDelete:
		err = deleteEntry(config.SiteDir, entry)

	default:
		err = errors.New("events ignored")
	}

	// Has any error?
	if err != nil {
		zlog.Warnw("entry", "status", status, "request", payload, "error", err)
		helper.WriteHttpError(w, status, err)
		return
	}

	// OK
	// zlog.Infow("entry", "status", http.StatusOK, "request", payload)
	zlog.Infow("entry", "status", http.StatusOK, "event", payload.Event)
	helper.WriteHttpResponse(w, http.StatusOK, entry)

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
