package helper

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed all:template
	templateEmbed embed.FS

	// Template funcs
	templateFuns = template.FuncMap{
		"add":            func(a int, b int) int { return a + b },
		"upperCase":      strings.ToUpper,
		"lowerCase":      strings.ToLower,
		"titleCase":      cases.Title(language.English, cases.NoLower).String,
		"join":           strings.Join,
		"split":          strings.Split,
		"toYaml":         toYaml,
		"toYamlByFields": toYamlByFields,
		"indent":         indent,
		"set":            setKey,
		"delete":         deleteKeys,
	}
)

// Reads an embeded file
func ReadEmbed(filename string) (string, error) {
	buf, err := templateEmbed.ReadFile(filename)
	return string(buf), err
}

func WriteAllEmbed(dir string) error {
	fs, err := templateEmbed.ReadDir("template")
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for _, v := range fs {
		buf, err := templateEmbed.ReadFile(path.Join("template", v.Name()))
		if err != nil {
			return err
		}

		err = os.WriteFile(path.Join(dir, v.Name()), buf, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// Executes a template
func ExecuteTemplate(filename string, data any) (string, error) {
	t, err := newTemplate(filename)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	return buf.String(), err
}

// Reads an embeded file
func readTemplate(filename string) (string, error) {
	buf, err := os.ReadFile(filename)
	return string(buf), err
}

// Creates a Go template from an embeded file
func newTemplate(filename string) (*template.Template, error) {
	// t, err := template.ParseFiles(filename)
	// t = t.Funcs(templateFuns)
	// return t, err

	// buf, err := ioutil.ReadFile(filename)
	buf, err := readTemplate(filename)
	if err != nil {
		return nil, err
	}

	return template.New("app").Option("missingkey=zero").Funcs(templateFuns).Parse(buf)
}

// --------------------------------------------------------
// Template functions
// --------------------------------------------------------

func toYaml(v any) string {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(v); err != nil {
		return "{}"
	}

	return strings.TrimSpace(buf.String())
}

func toYamlByFields(v any, fields ...string) string {
	res := []string{}
	for _, f := range fields {
		fieldValue, _ := v.(map[string]any)[f]
		res = append(res, toYaml(map[string]any{f: fieldValue}))
	}

	return strings.Join(res, "\n")
}

func indent(text string, value int) string {
	buf := []string{}
	lines := strings.Split(strings.TrimSpace(text), "\n")
	space := ""
	for i := 0; i < value; i++ {
		space += " "
	}

	for _, v := range lines {
		buf = append(buf, fmt.Sprintf("%s%s", space, v))
	}

	return strings.Join(buf, "\n")
}

func setKey(source map[string]any, k string, v any) map[string]any {
	source[k] = v
	return source
}

func deleteKeys(v map[string]any, keys ...string) map[string]any {
	res := map[string]any{}
	for k, v := range v {
		res[k] = v
	}

	for _, k := range keys {
		delete(res, k)
	}

	return res
}
