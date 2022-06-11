
MICROSERVICE_NAME=pablogod/sherlockgo:latest

build:
	docker build -t $(MICROSERVICE_NAME) .

release:
	docker push $(MICROSERVICE_NAME) 

build-and-release: build release
