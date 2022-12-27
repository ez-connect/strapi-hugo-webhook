package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEntryPayload(t *testing.T) {
	data, err := list("http://localhost:1337/api/documents?pagination[page]=1&populate=*")
	assert.NoError(t, err)

	doc := data.Data[1]
	payload := toEntryPayload("document", doc)
	t.Error(payload)
}
