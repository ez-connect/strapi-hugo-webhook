package impl

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/structpb"

	"strapi-webhook/base/pb"
)

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

// Write a collection type entry to a markdown file
func writeCollectionEntry(entry *structpb.Struct) error {
	return nil
}

// Write a single type entry to a YAML file
func writeSingleEntry(entry *structpb.Struct) error {
	return nil
}
