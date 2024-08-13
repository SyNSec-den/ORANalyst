#!/bin/sh

# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

set -xe

{{- if .Values.config.coreDump.enabled }}
cp /free5gc/webconsole/webui /tmp/coredump/
{{- end }}

cd /free5gc

cat config/webuicfg.conf

GOTRACEBACK=crash ./webconsole/webconsole -webuicfg config/webuicfg.conf
