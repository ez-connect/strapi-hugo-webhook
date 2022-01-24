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
	var (
		entry *pb.EntryContent
		err   error
	)

	// Validate
	if req.Entry == nil {
		return nil, errors.New("no entry found")
	}

	// Writes entry to file
	if isSingleType(req.Model) {
		if entry, err = getSingleTypeEntry(req); err == nil {
			err = writeSingleTypeEntry(entry)
		}

	} else {
		if entry, err = getCollectionTypeEntry(req); err == nil {
			err = writeCollectionTypeEntry(req.Model, entry)
		}
	}

	// Debug
	fmt.Println("event:", req.Event, "model:", req.Model, "locale:", entry.Locale, "name:", entry.Filename)

	// Write entry

	// Build

	// Commit

	// Push

	return &pb.EntryResponse{Request: req, Response: entry}, err
}
