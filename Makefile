
MICROSERVICE_NAME=pablogod/sherlockgo:latest

gen-docs:
	swag init

run-local: gen-docs
	swag init && go run main.go

build: gen-docs
	docker build -t $(MICROSERVICE_NAME) .

release:
	docker push $(MICROSERVICE_NAME) 

build-and-release: build release
