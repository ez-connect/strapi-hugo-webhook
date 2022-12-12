package entry

import (
	"fmt"
	"path"

	"github.com/gosimple/slug"
)

const (
	entryTypeSingle     = "single"
	entryTypeCollection = "collection"
	indexModel          = "index" // specific model to create a `_index.md`
)

var (
	collectionModels = []string{
		"contributor",
		"article",
		"document",
		"career",
		"project",
		"page",
		"resume",
	}
)

// The webook doesn't have the type of an entry?
// So, we will based on `collectionModels` to get its type.
func getCollectionType(model string) string {
	for _, e := range collectionModels {
		if model == e {
			return entryTypeCollection
		}
	}

	return entryTypeSingle
}

// Returns an unique file name: `${page-title-slug}-${id}`
func getUniqueFilename(entry map[string]any) string {
	id := entry["id"].(float64)
	var title string
	if entry["title"] != nil {
		title = entry["title"].(string)
	}

	slug := slug.Make(title)
	return fmt.Sprintf("%s-%v", slug, id)
}

// Parses an entry for prepare to write to an YAML file.
func getEntry(payload *EntryPayload) *Entry {
	model := payload.Model
	entryType := getCollectionType(model)
	entry := payload.Entry

	locale := ""
	if entry["locale"] != nil {
		locale = entry["locale"].(string)
	}

	parent := ""
	if entry["path"] != nil {
		parent = entry["path"].(string)
	}

	filename := ""
	if entryType == entryTypeSingle {
		// A data file
		filename = fmt.Sprintf("%s.yaml", payload.Model)
	} else {
		if payload.Model == indexModel {
			// Is a index page of a section
			filename = "_index.md"
		} else {
			// Or a content page
			filename = fmt.Sprintf("%s.md", getUniqueFilename(entry))
		}

	}

	if entryType == entryTypeCollection {
		filename = path.Join("content", locale, filename)
	} else {
		filename = path.Join("data", locale, filename)
	}

	return &Entry{
		Id:       int64(entry["id"].(float64)),
		Model:    model,
		Type:     entryType,
		Locale:   locale,
		Parent:   parent,
		Filename: filename,
		Data:     payload.Entry,
	}
}
