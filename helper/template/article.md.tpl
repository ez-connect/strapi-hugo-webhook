{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
---
{{ toYamlByFields . "title" "description" }}
tags:
{{ indent (toYaml (split .tags ",")) 2 }}
{{- with .recommended }}
recommended: {{ . }}
{{- end }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
{{- with .thumbnail.data }}
thumbnail: {{ toYaml . }}
{{- end }}
{{ toYamlByFields . "locale" "draft" }}
---

{{ .content }}
{{- end }}
