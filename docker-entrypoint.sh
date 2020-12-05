#! /bin/bash

if [[ "$#" -gt 0 && ("$1" == --yaml=* || "$1" == -yaml=*) ]]; then
    # https://tldp.org/LDP/abs/html/string-manipulation.html for substring removal.
    export YAML_FILE="${1#*=}"
    ./urlshort --yaml="$YAML_FILE"
else
    ./urlshort
fi