#!/usr/bin/env bash

# Kill, rebuild and run
# shellcheck disable=SC2046
# shellcheck disable=SC2009
kill $(ps -e | grep _build/whisper | awk '{print $1}'| head -n 1)
make clean-client && make client
