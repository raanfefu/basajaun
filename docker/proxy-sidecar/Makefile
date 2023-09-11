DOCKER_PLATFORM=linux/amd64
build:	

	@printf "********************************************\n" 
	docker build  --platform=${DOCKER_PLATFORM} \
	  -t `cat  ../../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.registry`/`cat  ../../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.image_proxy` \
	  -f Dockerfile.proxy  . 
	@printf "********************************************\n\n\n" 

	@printf "********************************************\n" 
	docker push `cat  ../../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.registry`/`cat  ../../charts/admission-controller/values.yaml | yq | jq -r .json_values.deployment.image_proxy` 
	@printf "********************************************\n\n\n" 