build-inject:
	cd admission && make build 	&& make push
build-opa:
	cd docker/opa-sidecar && make build 	
build-init:
	cd docker/init-container && make build 	
build-proxy:
	cd docker/proxy-sidecar && make build 	
 