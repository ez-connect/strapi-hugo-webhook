package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"

	"strapi-webhook/base/pb"
)

var (
	// Single types
	singleTypeModels = []string{"site", "home", "nav"}

	// // Collection types
	// collectionTypeModels = []string{"article", "career", "category", "contributor", "document", "page", "resume", "tag", "user"}

	// Ignore these fields
	implicitFields = []string{"localizations"}

	// Enable `populateCreatorFields` in `schema.json`
	// to allows `createdBy`, `updatedBy` in webook payloads
	relationFields = map[string]string{
		"createdBy": "username",
		"updatedBy": "username",
	}
)

// Marshals with indent 2
func marshalYAML(v interface{}) (string, error) {
	// // Indent 4, default
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
func grpcStruct2Map(entry *structpb.Struct) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	buf, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &res)
	return res, err
}

// Returns an unique file name
func getUniqueFilename(entry map[string]interface{}) string {
	id := entry["id"].(float64)
	var title string
	if entry["title"] != nil {
		title = entry["title"].(string)
	}

	slug := slug.Make(title)
	return fmt.Sprintf("%s-%v", slug, id) // page-title-slug-id
}

// Returns Hugo's front matter
func getFrontMatter(entry map[string]interface{}) (string, error) {
	// Ingore fields
	for _, field := range implicitFields {
		delete(entry, field)
	}

	// Ingore nil & `content` fields
	data := map[string]interface{}{}
	for k, v := range entry {
		if v != nil && k != "content" {
			data[k] = v
		}
	}

	// Is draft?
	if _, ok := data["publishedAt"]; !ok {
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

	// Tags: tag-a, tag-b
	if data["tags"] != nil {
		tags := strings.Split(data["tags"].(string), ",")
		data["tags"] = tags
	}

	return marshalYAML(data)
}

// Gets an entry
func getEntry(req *pb.EntryRequest) (*pb.EntryContent, error) {
	entry, err := grpcStruct2Map(req.Entry)
	if err != nil {
		return nil, err
	}

	frontMatter, err := getFrontMatter(entry)
	if err != nil {
		return nil, err
	}

	model := req.Model
	isSingle := isSingleType(model)
	res := pb.EntryContent{
		Id:           int64(entry["id"].(float64)),
		Model:        model,
		IsSingleType: isSingle,
	}

	if entry["locale"] != nil {
		res.Locale = entry["locale"].(string)
	} else {
		res.Locale = localeDefault
	}

	if entry["parent"] != nil {
		res.Parent = entry["parent"].(string)
	}

	if isSingle {
		res.Filename = fmt.Sprintf("%s.yaml", req.Model)
		res.Text = frontMatter
	} else {
		// Is a index page of a section
		if entry["isIndex"] != nil {
			res.Filename = "_index.md"
			// Or a content page
		} else {
			res.Filename = fmt.Sprintf("%s.md", getUniqueFilename(entry))
		}

		if entry["content"] != nil {
			res.Text = fmt.Sprintf("---\n%s---\n\n%s", frontMatter, entry["content"])
		} else {
			res.Text = fmt.Sprintf("---\n%s---\n\n", frontMatter)
		}
	}

	return &res, nil
}

// Gets a media urls
func getMedia(req *pb.MediaRequest) (*pb.MediaContent, error) {
	media, err := grpcStruct2Map(req.Media)
	if err != nil {
		return nil, err
	}

	res := pb.MediaContent{Url: media["url"].(string)}

	// Responsive files
	if media["formats"] != nil {
		formats := media["formats"].(map[string]interface{})

		// Thumbnail
		if v, ok := formats["thumbnail"]; ok {
			thumbnail := v.(map[string]interface{})
			res.Thumbnail = thumbnail["url"].(string)

		}

		// Small
		if v, ok := formats["small"]; ok {
			small := v.(map[string]interface{})
			res.Thumbnail = small["url"].(string)

		}
	}

	return &res, nil
}
