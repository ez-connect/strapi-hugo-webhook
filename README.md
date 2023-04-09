# strapi-hugo-webhook

Takes webhooks from Strapi then generate Hugo's data or content files and rebuild a site.

## Webhook

- `http://localhost:8080/entry`
- `http://localhost:8080/media`

## Usage

```bash
Usage:
  strapiwebhook [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fetch       Fetch entries from Strapi Rest API
  help        Help about any command
  serve       Start the server
  template    Write sample templates

Flags:
  -C, --collections strings   collection type models (default [section,contributor,article,document,career,project,page,resume])
  -h, --help                  help for strapiwebhook
  -l, --locale string         default locale (default "en")
  -S, --singles strings       single type models (default [site,home,nav,about])
  -d, --site-dir string       website site root dir (default "../hugo-theme")
  -s, --strapi-host string    strapi listen address (default "http://localhost:1337")
  -t, --strapi-token string   strapi api token
  -T, --template-dir string   template dir (default "helper/template")
  -v, --version               version for strapiwebhook

Use "strapiwebhook [command] --help" for more information about a command.
```

### Serve

```bash
Usage:
  strapiwebhook serve [flags]

Flags:
  -c, --cmd string              commands to run after trigger (default "echo 'Build the site'; hugo --gc --minify;")
      --debounced-cmd string    post debounced commands to run
      --debounced-timeout int   debounced timeout in second (default 300)
  -h, --help                    help for serve

Global Flags:
  -C, --collections strings   collection type models (default [section,contributor,article,document,career,project,page,resume])
  -l, --locale string         default locale (default "en")
  -S, --singles strings       single type models (default [site,home,nav,about])
  -d, --site-dir string       website site root dir (default "../hugo-theme")
  -s, --strapi-host string    strapi listen address (default "http://localhost:1337")
  -t, --strapi-token string   strapi api token
  -T, --template-dir string   template dir (default "helper/template")
```

### Fetch

```bash
Usage:
  strapiwebhook fetch [command]

Available Commands:
  get         Fetch an entry from Strapi Rest API
  list        fetch entries from Strapi Rest API

Flags:
  -e, --endpoint string   entry enpoint (default "sections")
  -h, --help              help for fetch
  -m, --model string      entry model (default "section")

Global Flags:
  -C, --collections strings   collection type models (default [section,contributor,article,document,career,project,page,resume])
  -l, --locale string         default locale (default "en")
  -S, --singles strings       single type models (default [site,home,nav,about])
  -d, --site-dir string       website site root dir (default "../hugo-theme")
  -s, --strapi-host string    strapi listen address (default "http://localhost:1337")
  -t, --strapi-token string   strapi api token
  -T, --template-dir string   template dir (default "helper/template")

Use "strapiwebhook fetch [command] --help" for more information about a command.
```

## Template

TBD
