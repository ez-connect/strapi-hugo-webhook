package impl

import (
	"context"
	"errors"
	"fmt"

	"strapi-webhook/base"
	"strapi-webhook/base/pb"
)

type serviceImpl struct {
	base.Service
}

func NewService() base.Service {
	return serviceImpl{}
}

func (s serviceImpl) Entry(ctx context.Context, req *pb.EntryRequest) (*pb.EntryResponse, error) {
	// Validate
	if req.Entry == nil {
		handleEntryError(errors.New("no entry found"))
	}

	// Write entry to file
	entry, err := getEntry(req)
	if err != nil {
		handleEntryError(err)
	}

	if err := writeEntry(req, entry); err != nil {
		handleEntryError(err)
	}

	// Debug
	fmt.Println("event:", req.Event, "model:", req.Model, "locale:", entry.Locale, "name:", entry.Filename)

	// Build
	if err = hugoBuild(); err != nil {
		handleEntryError(err)
	}

	// Commit + Push
	if gitCommitMsg != "" {
		if err := gitCommit(gitCommitMsg); err != nil {
			handleEntryError(err)
		}

		if err := gitPush(); err != nil {
			handleEntryError(err)
		}
	}

	return &pb.EntryResponse{Request: req, Response: entry}, err
}

func handleEntryError(err error) (*pb.EntryResponse, error) {
	fmt.Println("error:", err)
	return nil, err
}
