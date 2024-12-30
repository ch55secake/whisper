#!/usr/bin/env bash

# Watch changes, call the other script to kill and rebuild the client 
while true; do 
	fswatch -o ~/Projects/whisper/client/ | xargs -n1 -I{} ./kbuild.sh && echo "Change detected... rebuilding"
done
