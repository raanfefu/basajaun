namespace:
	kubectl apply -f kubernetes/authz-system-namespace.yaml
build-inject:
	cd admission && make build 	
push-inject:
	cd admission && make push