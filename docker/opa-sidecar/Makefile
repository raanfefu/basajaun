DOCKER_PLATFORM=linux/amd64
IMAGE_NAME=acrbackstagedev1.azurecr.io/opa_azure_sa:0.55.0-istio

build:	

	@printf "********************************************\n" 
	docker build  --platform=${DOCKER_PLATFORM} -t ${IMAGE_NAME} -f Dockerfile.opa  . 
	@printf "********************************************\n\n\n" 

	@printf "********************************************\n" 
	docker push ${IMAGE_NAME} 
	@printf "********************************************\n\n\n" 