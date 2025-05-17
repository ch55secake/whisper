#!/usr/bin/env bash

# Kill, rebuild and run
kill $(ps -e | grep _build/whisper | awk '{print $1}'| head -n 1)
go build -o "$(pwd)/_build/" && cd "$(pwd)/_build/" && ./whisper
