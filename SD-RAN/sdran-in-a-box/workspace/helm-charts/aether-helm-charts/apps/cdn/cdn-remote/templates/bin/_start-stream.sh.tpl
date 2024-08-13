#!/bin/bash

# Copyright 2019-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

QUALITY=$1

while true; do
    ffmpeg -re -i /opt/cdn/movies/$QUALITY.mp4 -c copy -f flv rtmp://{{ tuple "ant-media" . | include "cdn-remote.get_domain" }}:1935/LiveApp/$QUALITY;
done
