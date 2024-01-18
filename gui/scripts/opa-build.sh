#!/bin/sh

cd $1
/app/opa_envoy_linux_amd64 build -b src --output $2