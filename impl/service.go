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

	// Parse the entry
	entry, err := getEntry(req)
	if err != nil {
		return entryError(err)
	}

	// Debug
	fmt.Println("event:", req.Event, "model:", req.Model, "locale:", entry.Locale, "name:", entry.Filename)

	// Ingore media file
	if entry.Model == "file" {
		fmt.Println("ignore model `file`")
		return nil, nil
	}

	// Write the entry to a file
	switch req.Event {
	case EventEntryCreate, EventEntryUpdate:
		// Delete the old one
		if err := deleteEntry(entry); err != nil {
			return entryError(err)
		}

		if err := writeEntry(entry); err != nil {
			return entryError(err)
		}

	case EventEntryDelete:
		if err := writeEntry(entry); err != nil {
			return entryError(err)
		}

	default:
		return entryError(errors.New("unsupported event"))
	}

	// Build + sync
	if err := buildAndSync(gitCommitMsg); err != nil {
		return nil, err
	}

	return &pb.EntryResponse{Request: req, Response: entry}, err
}

func (s serviceImpl) Media(ctx context.Context, req *pb.MediaRequest) (*pb.MediaResponse, error) {
	// Parse the media
	media, err := getMedia(req)
	if err != nil {
		return mediaError(err)
	}

	fmt.Println("event:", req.Event, "url:", media.Url)

	// Download the all media formats
	switch req.Event {
	case EventMediaCreate, EventMediaUpdate:
		err = writeMedia(media)
	case EventMediaDelete:
		err = deleteMedia(media)
	default:
		err = errors.New("unknow event")
	}

	if err != nil {
		return mediaError(err)
	}

	// Build + sync
	if err := buildAndSync(gitCommitMsg); err != nil {
		return nil, err
	}

	return &pb.MediaResponse{Request: req, Response: media}, nil
}

func entryError(err error) (*pb.EntryResponse, error) {
	fmt.Println("error:", err)
	return nil, err
}

func mediaError(err error) (*pb.MediaResponse, error) {
	fmt.Println("error:", err)
	return nil, err
}
