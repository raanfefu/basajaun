build-inject:
	cd admission && make build 	&& make push
build-opa:
	cd docker/opa-sidecar && make build 	
build-init:
	cd docker/init-container && make build 	
build-proxy:
	cd docker/proxy-sidecar && make build 

run-redis:
	docker run -p 6379:6379 --rm -v ${PWD}/redis-conf:/usr/local/etc/redis --name myredis redis redis-server /usr/local/etc/redis/redis.conf	
 