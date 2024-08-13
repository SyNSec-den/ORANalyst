#!/bin/bash

# Copyright 2019-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

# Assume that SGI network gateway knows routes to UE pool
{{- if .Values.config.sriov.enabled }}
ip route add {{ .Values.networks.ue.subnet }} via {{ .Values.networks.sgi.gateway }}
{{- end }}

ip link set {{ .Values.config.nginx.sgiInterface.name }} mtu {{ .Values.config.nginx.sgiInterface.mtu }}

cp /conf/nginx.conf  /etc/nginx/
nginx -g "daemon off;"
