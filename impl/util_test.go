package impl

import (
	"strapi-webhook/base/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	testEntry = map[string]interface{}{
		"id":    1,
		"title": "page-title",
		// "publishedAt": "000",
		"tags": []interface{}{
			map[string]interface{}{"id": 1, "name": "tag-name"},
		},
		"createdBy": map[string]interface{}{
			"id":        1,
			"firstname": "Vinh",
			"lastname":  "Nguyen",
			"username":  "vinh@hotmail.com",
			"createdAt": "2022-01-23T13:59:43.490Z",
			"updatedAt": "2022-01-24T10:00:36.441Z",
		},
		"updatedBy": map[string]interface{}{
			"id":        1,
			"firstname": "Vinh",
			"lastname":  "Nguyen",
			"username":  "vinh2@hotmail.com",
			"createdAt": "2022-01-23T13:59:43.490Z",
			"updatedAt": "2022-01-24T10:00:36.441Z",
		},
		"content": "page-content",
	}

	testEntryRequest *pb.EntryRequest
)

func init() {
	entry, _ := structpb.NewStruct(testEntry)
	testEntryRequest = &pb.EntryRequest{
		Model: "test-model",
		Entry: entry,
	}

}

func TestGetFrontMatter(t *testing.T) {
	res, err := getFrontMatter(testEntry)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetSingleTypeEntry(t *testing.T) {
	res, err := getSingleTypeEntry(testEntryRequest)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Text)
}
