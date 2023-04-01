{{- $data := delete . "content" -}}
{{- $data = set . "draft" (not .publishedAt) -}}

---
{{ toYaml $data }}
---

{{ .content }}
