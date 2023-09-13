#!/bin/sh

CFG_FILE=/envoy/cfg.yaml
CFG_DEFAULT=/envoy/default/cfg.yaml

if test -f "$FILE"; then
    envoy --config-path $CFG_FILE
else
    envoy --config-path $CFG_DEFAULT
fi


 
