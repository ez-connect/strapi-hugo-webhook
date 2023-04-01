{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYaml $data }}
---

{{ .content }}
{{- end }}
