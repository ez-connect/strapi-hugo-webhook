{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
id: {{ .id }}
title: {{ .title }}
phone: {{ .phone }}
email: {{ .email }}
status: {{ .status }}
{{- toYaml .job }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
draft: {{ $draft }}
---

{{ .content }}
