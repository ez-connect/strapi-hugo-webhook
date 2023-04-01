package entry

import (
	"fmt"
	"path"
	"strapiwebhook/service/config"

	"github.com/gosimple/slug"
)

const (
	entryTypeSingle     = "single"
	entryTypeCollection = "collection"
	entryTypeIngore     = "ingore" // ingore entries

	// Specific model to create a nested section with a `_index.md`
	nestedSectionModel   = "nested-section"
	nestedSectionPathKey = "path"
)

// The webook doesn't have the type of an entry?
// So, we will based on `collectionModels` to get its type.
func getCollectionType(model string) string {
	for _, e := range config.SingleTypes {
		if model == e {
			return entryTypeSingle
		}
	}

	for _, e := range config.CollectionTypes {
		if model == e {
			return entryTypeCollection
		}
	}

	return entryTypeIngore
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

	nestedSectionPath := ""
	filename := ""

	if entryType == entryTypeSingle {
		// A data file
		filename = fmt.Sprintf("%s.yaml", model)
	} else {
		if model == nestedSectionModel {
			// Is a index page of a section
			if v, ok := payload.Entry[nestedSectionPath].(string); ok {
				nestedSectionPath = v
			}
			filename = "_index.md"
		} else {
			// Or a content page
			if section, ok := payload.Entry["section"].(map[string]any); ok {
				if v, ok := section[nestedSectionPathKey].(string); ok {
					nestedSectionPath = v
				}
			}

			filename = fmt.Sprintf("%s.md", getUniqueFilename(entry))
		}
	}

	// Include the nested section dir, if exists
	filename = path.Join(nestedSectionPath, filename)

	if entryType == entryTypeCollection {
		filename = path.Join("content", locale, model, filename)
	} else {
		filename = path.Join("data", locale, filename)
	}

	return &Entry{
		Id:       int64(entry["id"].(float64)),
		Model:    model,
		Type:     entryType,
		Locale:   locale,
		Filename: filename,
		Data:     payload.Entry,
	}
}
