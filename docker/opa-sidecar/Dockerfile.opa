FROM openpolicyagent/opa:0.55.0-istio

FROM ubuntu:rolling
RUN mkdir /app
COPY --from=0 /app/opa_envoy_linux_amd64 /app/opa_envoy_linux_amd64
COPY cfg.yaml /app/cfg.yaml

COPY entrypoint.sh /app/entrypoint.sh

ENTRYPOINT [ "/app/entrypoint.sh" ]