#!/usr/bin/env bash

# Watch changes, kill the process and then rebuild
fswatch -o ~/Projects/whisper/client/ | xargs -n1 -I{} ./kbuild.sh