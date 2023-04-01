{{- $data := set . "draft" (not .publishedAt) -}}

{{- with $data -}}
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
locale: {{ .locale }}
draft: {{ $draft }}
---

{{ .content }}
{{- end }}
