{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYamlByFields . "title" "description" "recommended" }}
tags:
{{ indent (toYaml (split .tags ",")) 2 }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
{{- with .thumbnail }}
thumbnail: {{ indent (toYaml .)  2 }}
{{- end }}
{{ toYamlByFields . "locale" "draft" }}
---

{{ .content }}
{{- end }}
