package impl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"strapiwebhook/base/pb"
)

var testService = NewService()

func TestEntry(t *testing.T) {
	req := &pb.EntryRequest{}
	_, err := testService.Entry(context.TODO(), req)
	assert.NoError(t, err)
}

// func TestMedia(t *testing.T) {
// 	req := &pb.MediaRequest{}
// 	_, err := testService.Media(context.TODO(), req)
// 	assert.NoError(t, err)
// }
