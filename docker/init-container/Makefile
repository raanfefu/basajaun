DOCKER_PLATFORM=linux/amd64
IMAGE_NAME=acrbackstagedev1.azurecr.io/proxy_init_base:v5

build:	

	@printf "********************************************\n" 
	docker build  --platform=${DOCKER_PLATFORM} -t ${IMAGE_NAME} -f Dockerfile.init  . 
	@printf "********************************************\n\n\n" 

	@printf "********************************************\n" 
	docker push ${IMAGE_NAME} 
	@printf "********************************************\n\n\n" 