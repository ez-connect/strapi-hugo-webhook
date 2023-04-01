package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"strapiwebhook/helper/zlog"
	"strapiwebhook/service/entry"
)

var (
	testToken = "19aacb860cf1282312cbd670dd459e4b2f770fcebdb0795aa6ae5b432a2cee2dae3462dae657cc254e716bab02a1133dbb825da9b07d02ebb86bd8ef2940ee11df66305f263ea22d0f0313eefffbb881f30a78a641862a4ac9c9bf8fa185547f3dfb017e86e14cbab214744c265c16e1a749acb6a8a40d18cb1db9eccf85ce76"
)

func TestListEntry(t *testing.T) {
	data, err := list(
		"http://localhost:1337/api/documents?populate=*&pagination[page]=1&pagination[pageSize]=2",
		testToken,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, data.Data)
}

func TestGetEntry(t *testing.T) {
	data, err := get(
		"http://localhost:1337/api/documents/115?populate=*",
		testToken,
	)
	assert.NoError(t, err)

	payload := getEntryPayload("document", data.Data)
	e := entry.GetEntry(payload)
	assert.NoError(t, entry.WriteEntry("../../data/test", "../../helper/template", e))
}

func TestFetchWriteEntryList(t *testing.T) {
	err := FetchAndWriteEntryList(
		"../../../hugo-theme",
		"../../helper/template",
		"document",
		"http://localhost:1337/api/documents?populate=*&pagination[pageSize]=100",
		testToken,
	)

	assert.NoError(t, err)

	err = FetchAndWriteEntryList(
		"../../../hugo-theme",
		"../../helper/template",
		"article",
		"http://localhost:1337/api/articles?populate=*&pagination[pageSize]=100",
		testToken,
	)

	assert.NoError(t, err)
}

func TestFetchWriteEntry(t *testing.T) {
	err := FetchAndWriteEntry(
		"../../../hugo-theme",
		"../../helper/template",
		"document",
		"http://localhost:1337/api/documents/115?populate=*",
		testToken,
	)

	assert.NoError(t, err)
}

func init() {
	zlog.InitLogger(false)
}
