package impl

import (
	"fmt"
	"os"
	"path"
	"strapi-webhook/base/pb"
)

// Writes a file
func writeFile(filename, text string) error {
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

// Writes a markdown content file
func writeCollectionTypeEntry(model string, entry *pb.EntryContent) error {
	filename := path.Join("content", entry.Locale, model, entry.Filename)
	return writeFile(filename, entry.Text)
}

// Writes a YAML data file
func writeSingleTypeEntry(entry *pb.EntryContent) error {
	filename := path.Join("data", entry.Locale, fmt.Sprintf("%s.yaml", entry.Filename))
	return writeFile(filename, entry.Text)
}
