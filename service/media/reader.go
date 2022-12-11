package media

// Gets a media urls
func getMedia(payload *MediaPayload) *Media {
	res := &Media{Url: payload.Media["url"].(string)}

	// Responsive files
	formats, ok := payload.Media["formats"].(map[string]any)
	if ok {
		// return nil, errors.New("not found any `formats`")
		// Thumbnail
		if v, ok := formats["thumbnail"]; ok {
			thumbnail := v.(map[string]any)
			res.Thumbnail = thumbnail["url"].(string)

		}

		// Small
		if v, ok := formats["small"]; ok {
			small := v.(map[string]any)
			res.Small = small["url"].(string)
		}
	}

	return res
}
