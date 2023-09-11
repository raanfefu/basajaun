FROM ubuntu:mantic

RUN apt update && apt install curl jq python3-pip -y
RUN  pip install check-jsonschema --break-system-packages
RUN curl -L -o opa https://github.com/open-policy-agent/opa/releases/download/v0.55.0/opa_linux_arm64_static 
RUN chmod 755 ./opa && cp opa /bin/opa
RUN  pip install jinja2-cli --break-system-packages
RUN mkdir /scripts /template  

COPY azcopy_linux_arm64_10.20.1/ /usr/bin
COPY template/   /template/
COPY jsonschema/ /jsonschema/
COPY scripts/   /scripts/

RUN chmod a+x /scripts/entrypoint.sh
ENV PATH="$PATH:/scripts"

WORKDIR /workspace

#ENTRYPOINT [ "entrypoint.sh" ]

