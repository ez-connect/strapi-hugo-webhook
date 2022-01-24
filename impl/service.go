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
	data := parseEntry(req.Entry)
	fmt.Println("event:", req.Event, "model:", req.Model, "id:", data["id"], "locale:", data["locale"])

	if isCollectionType(req) {
		md := getMarkdown(data)
		fmt.Println(md.Name)
		fmt.Println(md.Text)
	} else {
		fmt.Println(getYAML(data))
	}

	// Get text
	if isCollectionType(req) {
		md := getMarkdown(data)
		fmt.Println(md.Name)
		fmt.Println(md.Text)
	} else {
		fmt.Println(getYAML(data))
	}

	// Write entry

	// Build

	// Commit

	// Push

	return &pb.EntryResponse{Request: req}, nil
}
