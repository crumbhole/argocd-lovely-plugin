{{- define "rbac.apiVersion" -}}
{{- if .Capabilities.APIVersions.Has "rbac.authorization.k8s.io/v2" -}}
"v2"
{{- else -}}
"v1"
{{- end -}}
{{- end -}}
