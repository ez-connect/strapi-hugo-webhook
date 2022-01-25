package impl

import (
	"fmt"
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

// Writes a YAML data file
func writeSingleTypeEntry(entry *pb.EntryContent) error {
	filename := path.Join("data", entry.Locale, entry.Filename)
	return writeFile(filename, entry.Text)
}

// Writes a markdown content file
func writeCollectionTypeEntry(model string, entry *pb.EntryContent) error {
	filename := path.Join("content", entry.Locale, model, entry.Filename)
	return writeFile(filename, entry.Text)
}

func writeEntry(req *pb.EntryRequest, entry *pb.EntryContent) error {
	if isSingleType(req.Model) {
		return writeSingleTypeEntry(entry)
	}

	return writeCollectionTypeEntry(req.Model, entry)
}
