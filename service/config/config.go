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

	SiteDir       = "../hugo-theme" // hugo site dir
	LocaleDefault = "en"            // default locale

	// Collection type models
	CollectionTypes = []string{
		"nested-section",
		"contributor",
		"article",
		"document",
		"career",
		"project",
		"page",
		"resume",
	}

	// Single type models
	SingleTypes = []string{"site",
		"home",
		"nav",
		"about",
	}

	TemplateDir = "helper/template"
	PostCmd     = "hugo --gc --minify"

	DebouncedTimeout = int64(300) // git timeout in seconds, 5m
	PostDebouncedCmd = ""         //"git add .; git commit -m 'Sync cms'; git push;"
)
