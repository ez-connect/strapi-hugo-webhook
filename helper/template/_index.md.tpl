{{- $data := delete . "content" -}}
{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
{{ toYaml $data }}
draft: {{ $draft }}
---

{{ .content }}
