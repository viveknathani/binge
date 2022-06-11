#!/bin/bash

# A script to process a video for MPEG-DASH.
# $1 is the path to your video file.
# $2 is your video id.

mkdir -p ./content/$2
ffmpeg -re -i $1 -map 0 -map 0 -c:a libfdk_aac -c:v libx264 \
-b:v:0 800k -b:v:1 300k -s:v:1 320x170 -profile:v:1 baseline \
-profile:v:0 main -bf 1 -keyint_min 120 -g 120 -sc_threshold 0 \
-b_strategy 0 -ar:a:1 22050 -use_timeline 1 -use_template 1 \
-window_size 5 -adaptation_sets "id=0,streams=v id=1,streams=a" \
-seg_duration 10 -f dash ./content/$2/$2.mpd