{{- /*

# Copyright 2019-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

*/ -}}

{{/*
Render a set of standardised label
*/}}
{{- define "accelleran-cbrs-common.metadata_labels" -}}
{{- $application := index . 0 -}}
{{- $context := index . 1 -}}
release: {{ $context.Release.Name }}
app: {{ $application }}
{{- end -}}

{{/*
Render ServiceAccount, Role, and RoleBinding required for kubernetes-entrypoint.
*/}}
{{- define "accelleran-cbrs-common.service_account" -}}
{{- $context := index . 1 -}}
{{- $saName := index . 0 -}}
{{- $saNamespace := $context.Release.Namespace }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ $saName }}
  namespace: {{ $saNamespace }}
  labels:
{{ tuple $saName $context | include "accelleran-cbrs-common.metadata_labels" | indent 4 }}
---
{{- if semverCompare ">=1.16-0" $context.Capabilities.KubeVersion.GitVersion }}
apiVersion: rbac.authorization.k8s.io/v1
{{- else }}
apiVersion: rbac.authorization.k8s.io/v1beta1
{{- end }}
kind: RoleBinding
metadata:
  name: {{ $saName }}
  namespace: {{ $saNamespace }}
  labels:
{{ tuple $saName $context | include "accelleran-cbrs-common.metadata_labels" | indent 4 }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ $saName }}
subjects:
  - kind: ServiceAccount
    name: {{ $saName }}
    namespace: {{ $saNamespace }}
---
{{- if semverCompare ">=1.16-0" $context.Capabilities.KubeVersion.GitVersion }}
apiVersion: rbac.authorization.k8s.io/v1
{{- else }}
apiVersion: rbac.authorization.k8s.io/v1beta1
{{- end }}
kind: Role
metadata:
  name: {{ $saName }}
  namespace: {{ $saNamespace }}
  labels:
{{ tuple $saName $context | include "accelleran-cbrs-common.metadata_labels" | indent 4 }}
rules:
  - apiGroups:
      - ""
      - extensions
      - batch
      - apps
    verbs:
      - get
      - list
      - patch
    resources:
      - statefulsets
      - daemonsets
      - jobs
      - pods
      - services
      - endpoints
      - configmaps
{{- end -}}
