{{- define "name" -}}
  {{- printf "%s-%s" .Release.Name .Chart.Name }}
{{- end }}

{{- define "image" -}}
  {{- if .Values.image }}
    {{- .Values.image }}
  {{- else }}
    {{- .Values.global.image }}
  {{- end }}
{{- end }}

{{- define "replicas" -}}
  {{- if .Values.replicas }}
    {{- .Values.replicas }}
  {{- else }}
    {{- .Values.global.replicas }}
  {{- end }}
{{- end }}
