{{- $data := delete . "title" "position" "tel" "email" "webpage" "address" }}
{{- $data = delete $data "createdBy" "createdAt", "updatedBy" "updatedAt"
{{- $data = delete $data "draft" "content" }}

{{- $draft := true -}}
{{- if .publishedAt }}
	{{- $draft = false -}}
{{- end -}}

---
title: {{ .title }}
avatar:
  url: {{ index .avatar url }}
position: {{ .position }}
tel: {{ .tel }}
email: {{ .email }}
webpage: {{ .webpage }}
address: {{ .address }}
createdBy: {{ .createdBy.username }}
createdAt: {{ .createdAt }}
updatedBy: {{ .updatedBy.username }}
updatedAt: {{ .updatedAt }}
locale: {{ .locale }}
draft: {{ $draft }}

{{- toYaml $data }}
---

{{ .content }}
