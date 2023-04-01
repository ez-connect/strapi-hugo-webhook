{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
title: {{ .title }}
description: {{ .description }}
{{- with .section }}
path: {{ .path }}
{{- end }}
weight: {{ .weight }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
locale: {{ .locale }}
draft: {{ $draft }}
---

{{ .content }}
