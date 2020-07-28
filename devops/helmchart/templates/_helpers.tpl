{{/* Expand the name of the chart. */}}
{{- define "mychart.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{/* Shared by all deployments */}
{{- define "common.shared" -}}
- name: LOG_LEVEL
  value: {{ .Values.common.logLevel | quote }}
- name: LOG_FORMAT
  value: {{ .Values.common.logFormat | quote }}
- name: COMMON_CACHE_BUSTER
  value: {{ .Values.common.cacheBuster | quote }}
{{- end -}}

{/* Shared by services requiring auth0 */}
{{- define "common.auth0" -}}
- name: AUTH0_CLIENT_ID
  value: {{ required "common.auth0ClientId required" .Values.common.auth0ClientId | quote }}
- name: AUTH0_CLIENT_SECRET
  value: {{ required "common.auth0ClientSecret required" .Values.common.auth0ClientSecret | quote }}
- name: AUTH0_PEMFILE_PATH
  value: {{ required "common.auth0PemfilePath required" .Values.common.auth0PemfilePath | quote }}
- name: AUTH0_REDIRECT_URI
  value: {{ required "common.auth0RedirectUri required" .Values.common.auth0RedirectUri | quote }}
{{- end -}}

{/* Service specific shared by all deployments */}
{{- define "service.shared" -}}
- name: APP_NAME
  value: {{ .name | quote }}
- name: LISTEN_ADDRESS
  value: {{ .listenAddress | quote }}
- name: CACHE_BUSTER
  value: {{ .cacheBuster | default "1" | quote }}
{{- end -}}

{/* Service specific env for rabbitmq */}
{{- define "service.rabbit" -}}
- name: AMQP_DSN
  value: {{ required "<service>.amqpDsn required" .amqpDsn | quote }}
{{- end -}}
