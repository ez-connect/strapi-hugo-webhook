{{- $data := delete . "content" -}}
{{- $data = set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYaml . }}
---
{{- end }}

{{ .content }}
