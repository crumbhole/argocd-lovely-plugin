{{- define "service-name" -}}
{{- if semverCompare "> 1.21-0" .Capabilities.KubeVersion.GitVersion -}}
service-for-kubernetes-versions-newer-than-1-21
{{- else -}}
service-for-kubernetes-versions-older-than-1-22
{{- end -}}
{{- end -}}
