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
		return entryError(errors.New("no entry found"))
	}

	// Write entry to file
	entry, err := getEntry(req)
	if err != nil {
		return entryError(err)
	}

	if err := writeEntry(req, entry); err != nil {
		return entryError(err)
	}

	// Debug
	fmt.Println("event:", req.Event, "model:", req.Model, "locale:", entry.Locale, "name:", entry.Filename)

	// Build
	if err = hugoBuild(); err != nil {
		return entryError(err)
	}

	// Commit + Push
	if gitCommitMsg != "" {
		if err := gitCommit(gitCommitMsg); err != nil {
			return entryError(err)
		}

		if err := gitPush(); err != nil {
			return entryError(err)
		}
	}

	return &pb.EntryResponse{Request: req, Response: entry}, err
}

func entryError(err error) (*pb.EntryResponse, error) {
	fmt.Println("error:", err)
	return nil, err
}
