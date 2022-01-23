package impl

import (
	"context"
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
	// Debug
	fmt.Println(req.Event, isCollectionType(req), req.Model)
	fmt.Println(parseEntry(req.Entry))

	// Write entry

	// Build

	// Commit

	// Push

	return &pb.EntryResponse{Request: req}, nil
}
