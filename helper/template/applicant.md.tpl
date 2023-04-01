{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data  -}}
---
title: {{ .title }}
phone: {{ .phone }}
email: {{ .email }}
status: {{ .status }}
{{- toYaml .job }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
locale: {{ .locale }}
draft: {{ $draft }}
---

{{ .content }}
{{- end }}
