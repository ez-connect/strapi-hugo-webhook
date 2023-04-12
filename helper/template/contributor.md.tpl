{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYamlByFields . "user" }}
{{- with .avatar }}
avatar:
{{ indent (toYamlByFields . "alternativeText" "caption" "url") 2 }}
{{- end }}
{{ toYamlByFields . "locale" "draft" }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
{{ toYamlByFields . "locale" "draft" }}
---

{{ .content }}
{{- end }}
