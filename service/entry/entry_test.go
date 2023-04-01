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
			"model": "home",
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
	assert.NoError(t, json.Unmarshal([]byte(testSinglePayload), &req))
	res := GetEntry(req)
	assert.Equal(t, "home", res.Model)

	assert.Equal(t, entryTypeSingle, res.Type)
	assert.Equal(t, "data/home.yaml", res.Filename)
}

func TestWriteEntry(t *testing.T) {
	payload := &EntryPayload{}
	assert.NoError(t, json.Unmarshal([]byte(testSinglePayload), &payload))
	entry := GetEntry(payload)
	assert.NoError(t, WriteEntry("example", "../../helper/template", entry))

	assert.NoError(t, json.Unmarshal([]byte(testCollectionPayload), &payload))
	entry = GetEntry(payload)
	assert.NoError(t, WriteEntry("example", "../../helper/template", entry))
}
