

build:
	rm -rf docker/dist
	mkdir docker/dist docker/dist/angular  docker/dist/template docker/dist/scripts/
	cp template/* docker/dist/template/
	cp scripts/opa-build.sh scripts/sa.env docker/dist/scripts
	cd angular && ng build && mv dist/gui-bdle-plcy/* ../docker/dist/angular/
	env GOOS=linux GOARCH=amd64  go build -o docker/dist/main ./cmd/main.go 
	cd docker && docker build --platform=linux/amd64 -t temporalgui -f Dockerfile.gui .
	docker run -it --rm -p 8000:8090 temporalgui bash
	

	