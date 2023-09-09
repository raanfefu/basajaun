DOCKER_PLATFORM=linux/amd64
IMAGE_NAME=acrbackstagedev1.azurecr.io/envoy_base:v1.10.0

build:	
	@printf "********************************************\n" 
	docker build  --platform=${DOCKER_PLATFORM} -t ${IMAGE_NAME} -f Dockerfile.proxy . 
	@printf "********************************************\n\n\n" 
	@printf "********************************************\n" 
	docker push ${IMAGE_NAME} 
	@printf "********************************************\n\n\n" 