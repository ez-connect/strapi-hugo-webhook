package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"strapiwebhook/helper"
	"strapiwebhook/service/entry"
)

// pagination[page]

func list(uri string) (*ListResponse, error) {
	data, _ := helper.HttpGet(
		uri,
		map[string]string{
			"Authorization": "Bearer b8a9d39212af31ee0cb3ad429b873873e8985f355b135658a5ac42bd1f9ad30a5d6b69c06e872561c3e99c3d4abb1ff7194bfbf147f3a14588d4f76491b0fc66fad92527edbcf0d53217d7ad0cfba94939e0c84d2f4187fdc2489f283ef54641f35c2c60b627fd96d7d091849bfbe5db788d6285090b78551074ce7d5ef46e62",
		},
	)

	if data.StatusCode != http.StatusOK {
		return nil, errors.New(string(*data.Body))
	}

	res := &ListResponse{}
	if err := json.Unmarshal(*data.Body, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func toEntryPayload(model string, data map[string]any) *entry.EntryPayload {
	res := &entry.EntryPayload{
		Event: entry.EventEntryCreate,
		Model: model,
	}

	if attributes, ok := data["attributes"].(map[string]any); ok {
		for k, v := range attributes {
			data[k] = v
		}
		delete(data, "attributes")
	}

	entry := map[string]any{}
	for k, v := range data {
		if val, ok := v.(map[string]any); ok {
			if data, ok := val["data"].(map[string]any); ok {
				if attributes, ok := data["attributes"].(map[string]any); ok {
					for k1, v1 := range attributes {
						data[k1] = v1
					}

					delete(data, "attributes")
				}

				entry[k] = data
			} else {
				entry[k] = v
			}
		} else {
			entry[k] = v
		}
	}

	res.Entry = entry
	return res
}
