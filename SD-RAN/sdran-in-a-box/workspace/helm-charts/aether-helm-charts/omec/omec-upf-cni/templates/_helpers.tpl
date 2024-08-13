{{- /*

# Copyright 2019-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

*/ -}}

{{/*
Renders a set of standardised labels.
*/}}
{{- define "omec-upf-cni.metadata_labels" -}}
{{- $application := index . 0 -}}
{{- $context := index . 1 -}}
tier: node
release: {{ $context.Release.Name }}
app: {{ $application }}
{{- end -}}
