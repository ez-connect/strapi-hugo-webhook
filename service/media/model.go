package media

type MediaPayload struct {
	Event string         `json:"event,omitempty"`
	Media map[string]any `json:"media,omitempty"`
	Url   string         `json:"url,omitempty"`
}

type Media struct {
	Url       string `json:"url,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Small     string `json:"small,omitempty"`
}

const (
	EventMediaCreate = "media.create"
	EventMediaUpdate = "media.update"
	EventMediaDelete = "media.delete"
)
