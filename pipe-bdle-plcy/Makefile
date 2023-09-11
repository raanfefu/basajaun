IMAGE_NAME=opa-builder

build:
	docker build  -t opa-builder .
run:
	make build 
	cd ${P} && docker run -v ${P}:/workspace  -v ${PWD}/jsonschema/:/jsonschema/ -v ${PWD}/template/:/template/ --rm -it opa-builder bash
 