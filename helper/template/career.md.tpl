{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
title: {{ .title }}
status: {{ .status }}
categories:
{{ indent (toYaml (split .categories ",")) 2 }}
recommended: {{ .recommended }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .Date }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
draft: {{ $draft }}
---

{{ .content }}
