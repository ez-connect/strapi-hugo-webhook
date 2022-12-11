package config

// Args passed via build in the `Makefile`
// -ldflags="-X 'pkg/service.BuildDate=$(name)' -X 'pkg/service.Branch=$(version)...'"
var (
	Name        string
	Description string
	Version     string
	BuildDate   string
	Branch      string
	Hash        string
	BuildMode   string
)

// Command args
var (
	StrapiAddr  = "http://localhost:1337" // strapi server
	StrapiToken = ""                      // strapi token

	SiteDir         = "example" // hugo site dir
	LocaleDefault   = "en"      // default locale
	CollectionTypes = []string{"contributor", "article", "document", "career", "project", "page", "resume"}

	TemplateDir = "helper/template"
	PostCmd     = "hugo --gc --minify"

	DebouncedTimeout = int64(300) // git timeout in seconds, 5m
	PostDebouncedCmd = "git add .; git commit -m 'Sync cms'; git push;"
)
