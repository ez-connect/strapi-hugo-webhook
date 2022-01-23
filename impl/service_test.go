package impl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"strapi-webhook/base/pb"
)

var testService = NewService()

func TestEntry(t *testing.T) {
	req := &pb.EntryRequest{}
	_, err := testService.Entry(context.TODO(), req)
	assert.NoError(t, err)
}
