{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
id: {{ .id }}
title: {{ .title }}
description: {{ .description }}
tags:
{{ indent (toYaml (split .tags ",")) 2 }}
recommended: {{ .recommended }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
locale: {{ .locale }}
draft: {{ $draft }}
---

{{ .content }}
