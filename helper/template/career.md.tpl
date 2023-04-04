{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYamlByFields . "title" "description" "status" "recommended" "location" }}
{{- with .categories }}
categories:
{{ indent (toYaml (split . ",")) 2 }}
{{- end }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .Date }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
{{ toYamlByFields . "locale" "draft" }}
---

{{ .content }}
{{- end }}
