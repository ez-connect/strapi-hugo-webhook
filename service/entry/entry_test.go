package entry

import (
	"encoding/json"
	"strapiwebhook/helper/zlog"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSinglePayload = `
		{
			"event": "entry.create",
			"createdAt": "2020-01-10T08:47:36.649Z",
			"model": "address",
			"entry": {
				"id": 1,
				"geolocation": {},
				"city": "Paris",
				"postal_code": null,
				"category": null,
				"full_name": "Paris",
				"createdAt": "2020-01-10T08:47:36.264Z",
				"updatedAt": "2020-01-10T08:47:36.264Z",
				"cover": null,
				"images": []
			}
		}
	`

	testCollectionPayload = `
		{
			"event": "entry.create",
			"createdAt": "2020-01-10T08:47:36.649Z",
			"model": "article",
			"entry": {
				"id": 1,
				"title": "Sample title",
				"description": "Article description",
				"tags": "tag-1,tag-2,tag-3",
				"content": "TBD",
				"createdAt": "2020-01-10T08:47:36.264Z",
				"updatedAt": "2020-01-10T08:47:36.264Z"
			}
		}
	`
)

func init() {
	zlog.InitLogger(false)
}

func TestGetEntryContent(t *testing.T) {
	req := &EntryPayload{}
	json.Unmarshal([]byte(testSinglePayload), &req)
	res := getEntry(req)
	assert.Equal(t, "address", res.Model)

	assert.Equal(t, entryTypeSingle, res.Type)
	assert.Equal(t, "data/address.yaml", res.Filename)
}

func TestWriteEntry(t *testing.T) {
	entry := &EntryPayload{}
	assert.NoError(t, json.Unmarshal([]byte(testSinglePayload), &entry))
	_, err := writeEntry("example", "../../data", entry)
	assert.NoError(t, err)

	assert.NoError(t, json.Unmarshal([]byte(testCollectionPayload), &entry))
	_, err = writeEntry("example", "../../data", entry)
	assert.NoError(t, err)
}
