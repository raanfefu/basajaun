build:
	@printf "Starting build admission controller ...\n" 
	export `cat ./scripts/build.env | xargs`  && ./scripts/build.sh
	@printf "Finish admission controller ...\n" 
	
push:
	@printf "push docker image..." 
	export `cat ./scripts/build.env | xargs`  && ./scripts/push.sh	