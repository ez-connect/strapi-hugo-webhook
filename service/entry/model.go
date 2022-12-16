package entry

// Example payload
//
//	{
//		"event": "entry.create",
//		"createdAt": "2020-01-10T08:47:36.649Z",
//		"model": "address",
//		"entry": {
//		  "id": 1,
//		  "geolocation": {},
//		  "city": "Paris",
//		  "postal_code": null,
//		  "category": null,
//		  "full_name": "Paris",
//		  "createdAt": "2020-01-10T08:47:36.264Z",
//		  "updatedAt": "2020-01-10T08:47:36.264Z",
//		  "cover": null,
//		  "images": []
//		}
//	}
type EntryPayload struct {
	Event     string         `json:"event,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	Model     string         `json:"model,omitempty"`
	Entry     map[string]any `json:"entry,omitempty"`
}

type Entry struct {
	Id       int64          `json:"id,omitempty"`
	Locale   string         `json:"locale,omitempty"`
	Model    string         `json:"model,omitempty"`
	Type     string         `json:"type,omitempty"`
	Filename string         `json:"filename,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
}

const (
	eventEntryCreate = "entry.create"
	eventEntryUpdate = "entry.update"
	eventEntryDelete = "entry.delete"
	// eventEntryPublish   = "entry.publish"
	// eventEntryUnpublish = "entry.unpublish"
)
