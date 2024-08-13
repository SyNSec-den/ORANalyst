#!/bin/bash
# Copyright 2021-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

set -ex

{{ if .Values.config.rosCore.debug }}
while true; do sleep 3600; done;
{{ else }}
cd /root/TrajBridge-PX4;
source devel/setup.bash; roslaunch bridge_px4 gcs.launch
{{ end }}
