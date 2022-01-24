package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testEntry = map[string]interface{}{
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

func TestGetFrontMatter(t *testing.T) {
	res := getFrontMatter(testEntry)
	t.Error(res)
}

func TestGetMardown(t *testing.T) {
	res, err := getMarkdown(testEntry)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Text)
	t.Error(res.Text)
}
