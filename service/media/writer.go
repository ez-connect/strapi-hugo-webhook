package media

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strapiwebhook/helper/zlog"
)

// Writes media files
func writeMedia(siteDir, strapiAddr string, payload *MediaPayload) error {
	media := getMedia(payload)
	urls := []string{payload.Url, media.Thumbnail, media.Small}
	for _, url := range urls {
		if url != "" {
			if err := downloadMedia(siteDir, strapiAddr, url); err != nil {
				return err
			}
		}
	}

	return nil
}

// Delete media files
func deleteMedia(siteDir string, payload *MediaPayload) error {
	media := getMedia(payload)
	urls := []string{media.Url, media.Thumbnail, media.Small}
	for _, url := range urls {
		if url != "" {
			filename := path.Join(siteDir, "static", url)
			zlog.Infow("delete media", "url", media.Url, "filename", filename)
			if err := os.Remove(filename); err != nil {
				zlog.Warnw("delete file", "error", err)
			}
		}
	}

	return nil
}

// Downloads and write the file from http
func downloadMedia(siteDir, strapiAddr, url string) error {
	filename := path.Join(siteDir, "static", url)
	zlog.Infow("download media", "url", url, "filename", filename)

	// Download the file
	res, err := http.Get(fmt.Sprintf("%s%s", strapiAddr, url))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Write the file
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		return err
	}

	return nil
}
