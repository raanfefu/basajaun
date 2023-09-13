#!/bin/sh
cat /app/env/sa.env 
export $(cat /app/env/sa.env | xargs)

/app/opa_envoy_linux_amd64 run \
	--server \
	--addr=localhost:8181 \
	--diagnostic-addr=0.0.0.0:8282 \
	--set=plugins.envoy_ext_authz_grpc.addr=:9191 \
	--set=plugins.envoy_ext_authz_grpc.query=data.envoy.authz.allow \
	--set=decision_logs.console=true \
	--ignore=.* \
	-c \
	/app/cfg.yaml