package impl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"strapi-webhook/base/pb"
)

// Writes a file
func writeFile(filename, text string) error {
	output := path.Join(siteDir, filename)
	fmt.Println("Write:", output)

	if err := os.MkdirAll(path.Dir(output), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	return err
}

// Writes a YAML data file or a markdown content file
func writeEntry(entry *pb.EntryContent) error {
	var filename string
	if entry.IsSingleType {
		filename = path.Join("data", entry.Locale, entry.Filename)
	} else {
		filename = path.Join("content", entry.Locale, entry.Model, entry.Filename)
	}
	return writeFile(filename, entry.Text)
}

// Deletes a markdown content file
func deleteEntry(entry *pb.EntryContent) error {
	var filename string
	if entry.IsSingleType {
		filename = entry.Filename
	} else {
		filename = path.Join("content", entry.Locale, entry.Model, fmt.Sprintf("*-%v.md", entry.Id))
	}

	if err := os.Remove(filename); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

// Downloads and write the file from http
func downloadMedia(url string) error {
	filename := path.Join(siteDir, "static", url)
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
			if err := os.Remove(filename); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
			}
		}
	}

	return nil
}
