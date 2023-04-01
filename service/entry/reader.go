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
	sectionModel   = "section"
	sectionPathKey = "path"
)

// Parses an entry for prepare to write to an YAML file.
func GetEntry(payload *EntryPayload) *Entry {
	model := payload.Model
	entryType := getCollectionType(model)
	entry := payload.Entry

	locale := ""
	if entry["locale"] != nil {
		locale = entry["locale"].(string)
	}

	sectionPath := ""
	filename := ""

	if entryType == entryTypeSingle {
		// A data file
		filename = fmt.Sprintf("%s.yaml", model)
	} else {
		if model == sectionModel {
			// Is a index page of a section
			if v, ok := payload.Entry[sectionPathKey].(string); ok {
				sectionPath = v
			}
			filename = "_index.md"
		} else {
			// Or a content page
			if section, ok := payload.Entry[sectionModel].(map[string]any); ok {
				if v, ok := section[sectionPathKey].(string); ok {
					sectionPath = v
				}
			}

			filename = fmt.Sprintf("%s.md", getUniqueFilename(entry))
		}
	}

	// Include the section dir, if exists
	if sectionPath != "" {
		filename = path.Join(sectionPath, filename)
	} else {
		filename = path.Join(model, filename)
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
		Filename: filename,
		Data:     payload.Entry,
	}
}

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
