#!/bin/bash

# Copyright 2020-present Open Networking Foundation
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/google/video_analytics_server_main \
        --mediapipe_graph_path=/google/demo_graph.pbtxt \
        --person_detection_tf_saved_model_dir=/google/saved_model/ \
        --camera_scene_geometry_path=/google/camera_scene_geometry.pbtxt \
        --grpc_port=50051 \
        --mq_address=localhost:5672 \
        --mediapipe_detection_topic_name=phylo.mediapipe_detection \
        --person_detection_topic_name=phylo.person_detection \
        --bbox_decoded_video_frame_topic_name=phylo.bbox_decoded_video_frame \
        --publish_to_mq=true \
        --publish_to_log=true
