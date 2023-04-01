package entry

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"strapiwebhook/helper"
	"strapiwebhook/helper/zlog"
)

// Writes a file
func WriteEntry(siteDir, templateDir string, entry *Entry) error {
	// Ingore media file due to wrong trigger in case mistake in the Webhook settings
	// TODO: Automaticaly switch to media???
	if entry.Model == "file" {
		return errors.New("Entry endpoint isn't used for media")
	}

	filename := path.Join(siteDir, entry.Filename)
	zlog.Infow("write file", "filename", filename)
	if err := os.MkdirAll(path.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	// Execute template
	template := getTemplate(templateDir, entry)
	buf, err := helper.ExecuteTemplate(template, entry.Data)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// TODO: Write use a template
	_, err = f.WriteString(buf)
	return err
}

// Deletes a markdown file
func deleteEntry(siteDir string, entry *Entry) error {
	// Delete single type file
	filename := path.Join(siteDir, entry.Filename)
	if entry.Type == entryTypeSingle {
		if err := os.Remove(filename); err != nil {
			zlog.Warnw("delete file", "error", err)
		}

		return nil
	}

	// Delete a file name with the last name is the entry id
	// pattern := path.Join(siteDir, "content", entry.Locale, entry.Model, fmt.Sprintf("**/*-%v.md", entry.Id))
	// files, err := filepath.Glob(pattern)
	// DEV: Glob doesn't support `**`
	// https://github.com/golang/go/issues/11862
	files := []string{}
	err := filepath.Walk(
		path.Dir(filename),
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
		zlog.Infow("delete file", "filename", f)
		if err := os.Remove(filename); err != nil {
			zlog.Warnw("delete file", "error", err)
		}

	}

	return nil
}

// Gets the template file for a `model`.
// Returns default template if not found.
func getTemplate(templateDir string, entry *Entry) string {
	if entry.Type == entryTypeSingle {
		filename := path.Join(templateDir, fmt.Sprintf("%s.yaml", entry.Model))
		if _, err := os.Stat(filename); err != nil {
			return path.Join(templateDir, "_single.yaml.tpl")
		}

		return filename
	}

	if entry.Model == sectionModel {
		return path.Join(templateDir, "_index.md.tpl")
	}

	filename := path.Join(templateDir, fmt.Sprintf("%s.md.tpl", entry.Model))
	if _, err := os.Stat(filename); err != nil {
		return path.Join(templateDir, "_collection.md.tpl")
	}

	return filename
}
