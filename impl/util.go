package impl

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gosimple/slug"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"

	"strapi-webhook/base/pb"
)

var (
	// Single types
	singleTypeModels = []string{"site", "home"}

	// // Collection types
	// collectionTypeModels = []string{"article", "career", "category", "contributor", "document", "page", "resume", "tag", "user"}

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

// Marshals with indent 2
func marshalYAML(v interface{}) (string, error) {
	// // Index 4, default
	// buf, _ := yaml.Marshal(data)
	// return string(buf)

	// Customize encoder
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(&v)
	return buf.String(), err
}

// Is collection type?
func isSingleType(model string) bool {
	for _, e := range singleTypeModels {
		if model == e {
			return true
		}
	}

	return false
}

// Converts a gRPC struct to a map
func parseEntry(entry *structpb.Struct) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	buf, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &res)
	return res, err
}

// Returns an unique file name
func getFileName(entry map[string]interface{}) string {
	id := entry["id"]
	slug := slug.Make(fmt.Sprintf("%v", entry["title"]))
	return fmt.Sprintf("%s-%v", slug, id) // page-title-slug-id
}

// Returns Hugo's front matter
func getFrontMatter(entry map[string]interface{}) (string, error) {
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

	return marshalYAML(data)
}

// Returns a single type entry to a YAML file
func getSingleTypeEntry(req *pb.EntryRequest) (*pb.EntryContent, error) {
	entry, err := parseEntry(req.Entry)
	if err != nil {
		return nil, err
	}

	frontMatter, err := getFrontMatter(entry)
	if err != nil {
		return nil, err
	}

	res := pb.EntryContent{
		Locale:   fmt.Sprintf("%s", entry["locale"]),
		Filename: fmt.Sprintf("%s.yaml", req.Model),
		Text:     frontMatter,
	}

	return &res, nil
}

// Returns a collection type entry to a markdown file
func getCollectionTypeEntry(req *pb.EntryRequest) (*pb.EntryContent, error) {
	entry, err := parseEntry(req.Entry)
	if err != nil {
		return nil, err
	}

	// Content
	frontMatter, err := getFrontMatter(entry)
	if err != nil {
		return nil, err
	}

	content := fmt.Sprintf("%s", entry["content"])

	res := pb.EntryContent{
		Locale:   fmt.Sprintf("%s", entry["locale"]),
		Filename: fmt.Sprintf("%s.md", getFileName(entry)),
		Text:     fmt.Sprintf("%s---\n\n%s\n", frontMatter, content),
	}

	return &res, nil
}
