{{- $data := delete . "content" -}}
{{- $data = set $data "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYaml . }}
---
{{- end }}

{{ .content }}
