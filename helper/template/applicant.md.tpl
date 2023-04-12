{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data  -}}
---
title: {{ .title }}
phone: {{ .phone }}
email: {{ .email }}
status: {{ .status }}
{{- toYaml .job }}
createdAt: {{ .createdAt }}
updatedAt: {{ .updatedAt }}
locale: {{ .locale }}
draft: {{ $draft }}
---

{{ .content }}
{{- end }}
