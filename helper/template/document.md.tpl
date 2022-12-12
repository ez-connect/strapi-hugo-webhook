{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
title: {{ .title }}
description: {{ .description }}
path: {{ .path }}
weight: {{ .weight }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
draft: {{ $draft }}
---

{{ .content }}
