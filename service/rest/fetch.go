package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"strapiwebhook/helper"
	"strapiwebhook/service/entry"
)

// Fetch & write a list of entries
func FetchAndWriteEntryList(siteDir, templateDir, model, uri, token string) error {
	data, err := list(uri, token)
	if err != nil {
		return err
	}

	for _, v := range data.Data {
		if err := writeEntry(siteDir, templateDir, model, v); err != nil {
			return err
		}
	}

	return nil
}

// Fetch & write an entry
func FetchAndWriteEntry(siteDir, templateDir, model, uri, token string) error {
	data, err := get(uri, token)
	if err != nil {
		return err
	}

	return writeEntry(siteDir, templateDir, model, data.Data)
}

//---------------------------------------------------------

func httpGet(uri, token string, v any) error {
	data, err := helper.HttpGet(
		uri,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token),
		},
	)

	if err != nil {
		return err
	}

	if data.StatusCode != http.StatusOK {
		return errors.New(string(*data.Body))
	}

	return json.Unmarshal(*data.Body, &v)
}

// List all entries - Strapi Rest API
// http://localhost:1337/api/documents?populate=*&pagination[page]=1&pagination[pageSize]=100
func list(uri, token string) (*ListResponse, error) {
	res := &ListResponse{}
	err := httpGet(uri, token, &res)
	return res, err
}

func get(uri, token string) (*GetResponse, error) {
	res := &GetResponse{}
	err := httpGet(uri, token, &res)
	return res, err
}

func getEntryPayload(model string, data map[string]any) *entry.EntryPayload {
	res := &entry.EntryPayload{
		Event: entry.EventEntryCreate,
		Model: model,
	}

	entry := map[string]any{
		"id": data["id"],
	}

	if attributes, ok := data["attributes"].(map[string]any); ok {
		for k, v := range attributes {
			entry[k] = v
		}
		delete(data, "attributes")
	}

	for k, v := range entry {
		if val, ok := v.(map[string]any); ok {
			if data, ok := val["data"].(map[string]any); ok {
				if attributes, ok := data["attributes"].(map[string]any); ok {
					for k1, v1 := range attributes {
						data[k1] = v1
					}

					delete(data, "attributes")
				}

				entry[k] = data
			}
		}
	}

	res.Entry = entry
	return res
}

func writeEntry(siteDir, templateDir, model string, data map[string]any) error {
	payload := getEntryPayload(model, data)
	e := entry.GetEntry(payload)
	return entry.WriteEntry(siteDir, templateDir, e)
}
