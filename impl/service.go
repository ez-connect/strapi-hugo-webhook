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
		return nil, errors.New("no entry found")
	}

	// Write entry to file
	entry, err := getEntry(req)
	if err != nil {
		return nil, err
	}

	if err := writeEntry(req, entry); err != nil {
		return nil, err
	}

	// Debug
	fmt.Println("event:", req.Event, "model:", req.Model, "locale:", entry.Locale, "name:", entry.Filename)

	if err != nil {
		return nil, err
	}

	// Build
	err = hugoBuild()

	// Commit

	// Push

	return &pb.EntryResponse{Request: req, Response: entry}, err
}
