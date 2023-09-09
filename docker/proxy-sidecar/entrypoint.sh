#!/bin/sh

CFG_FILE=/envoy/cfg.yaml

if test -f "$FILE"; then
    envoy --config-path $CFG_FILE
else
    exit -1
fi


 
