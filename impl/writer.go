package impl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"strapiwebhook/base/pb"
)

// Writes a file
func writeFile(filename, text string) error {
	GetLogger().Infow("write file", "filename", filename)
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	return err
}

// Deletes a file
func deleteFile(filename string) error {
	GetLogger().Infow("delete file", "filename", filename)
	if err := os.Remove(filename); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

// Writes a YAML data file or a markdown content file
func writeEntry(entry *pb.EntryContent) error {
	var filename string
	if entry.IsSingleType {
		filename = path.Join(siteDir, "data", entry.Locale, entry.Filename)
	} else {
		filename = path.Join(siteDir, "content", entry.Locale, entry.Model, entry.Parent, entry.Filename)
	}
	return writeFile(filename, entry.Text)
}

// Deletes a markdown content file
func deleteEntry(entry *pb.EntryContent) error {
	// Delete single type file
	if entry.IsSingleType {
		return deleteFile(entry.Filename)
	}

	// Delete a file name with the last name is the entry id
	// pattern := path.Join(siteDir, "content", entry.Locale, entry.Model, fmt.Sprintf("**/*-%v.md", entry.Id))
	// files, err := filepath.Glob(pattern)
	// DEV: Glob doesn't support `**`
	// https://github.com/golang/go/issues/11862
	files := []string{}
	err := filepath.Walk(
		path.Join(siteDir, "content", entry.Locale, entry.Model),
		func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, fmt.Sprintf("-%v.md", entry.Id)) {
				files = append(files, path)
			}
			return nil
		},
	)

	if err != nil {
		return err
	}

	for _, f := range files {
		if err := deleteFile(f); err != nil {
			return err
		}
	}

	return nil
}

// Downloads and write the file from http
func downloadMedia(url string) error {
	filename := path.Join(siteDir, "static", url)
	GetLogger().Infow("download media", "url", url, "filename", filename)

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

// Writes media files
func writeMedia(media *pb.MediaContent) error {
	urls := []string{media.Url, media.Thumbnail, media.Small}
	for _, url := range urls {
		if url != "" {
			if err := downloadMedia(url); err != nil {
				return err
			}
		}
	}

	return nil
}

// Delete media files
func deleteMedia(media *pb.MediaContent) error {
	urls := []string{media.Url, media.Thumbnail, media.Small}
	for _, url := range urls {
		if url != "" {
			filename := path.Join(siteDir, "static", url)
			GetLogger().Infow("delete media", "url", media.Url, "filename", filename)
			if err := deleteFile(filename); err != nil {
				return err
			}
		}
	}

	return nil
}
