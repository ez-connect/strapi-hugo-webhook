package impl

import (
	"encoding/json"
	"fmt"

	"github.com/gosimple/slug"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"

	"strapi-webhook/base/pb"
)

var (
	// Ignore these fields
	implicitFields = []string{"localizations"}

	// Enable `populateCreatorFields` in `schema.json`
	// to allows `createdBy`, `updatedBy` in webook payloads
	relationFields = map[string]string{
		"tags":      "name",
		"createdBy": "username",
		"updatedBy": "username",
	}
)

// An entry content
type EntryContent struct {
	Name string
	Text string // file content
}

// Is collection type?
func isCollectionType(event *pb.EntryRequest) bool {
	return event.Model != ""
}

// Converts a gRPC struct to a map
func parseEntry(entry *structpb.Struct) map[string]interface{} {
	res := map[string]interface{}{}
	buf, _ := json.Marshal(entry)
	json.Unmarshal(buf, &res)

	return res
}

// Gets Hugo's front matter
func getFrontMatter(entry map[string]interface{}) string {
	// Ingore fields
	for _, field := range implicitFields {
		delete(entry, field)
	}

	// Ingore content
	data := map[string]interface{}{}
	for k, v := range entry {
		if k != "content" {
			data[k] = v
		}
	}

	// Is draft?
	if _, ok := entry["publishedAt"]; !ok {
		data["draft"] = true
	}

	// Relationships
	for k, v := range relationFields {
		// fmt.Printf("%T", data[k])
		if relation, ok := data[k].([]interface{}); ok {
			names := []string{}
			for _, e := range relation {
				names = append(names, fmt.Sprintf("%s", e.(map[string]interface{})[v]))
			}

			data[k] = names
			continue
		}

		if relation, ok := data[k]; ok {
			data[k] = fmt.Sprintf("%s", relation.(map[string]interface{})[v])
		}
	}

	buf, _ := yaml.Marshal(data)
	return string(buf)
}

// Write a collection type entry to a markdown file
func getMarkdown(entry map[string]interface{}) *EntryContent {
	// Name
	id := entry["id"]
	slug := slug.Make(fmt.Sprintf("%v", entry["title"]))
	name := fmt.Sprintf("%s-%v", slug, id) // page-title-slug-id

	// Content
	frontMatter := getFrontMatter(entry)
	content := fmt.Sprintf("%s", entry["content"])

	res := EntryContent{
		Name: name,
		Text: fmt.Sprintf("%s---\n\n%s\n", frontMatter, content),
	}

	return &res
}

// Write a single type entry to a YAML file
func getYAML(entry map[string]interface{}) string {
	return getFrontMatter(entry)
}
