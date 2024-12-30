#!/usr/bin/env bash

# Kill, rebuild and run
kill $(ps -e | grep _build/whisper | awk '{print $1}'| head -n 1)
go build -o ~/Projects/whisper/_build/ && cd ~/Projects/whisper/_build/ && ./whisper

