install-deps:
	go mod tidy

app: cmd/layoffs/main.go
	go build -o app cmd/layoffs/main.go

run: build
	./app

clean: app
	rm -f app


### Container section ###
DEFAULT_IMAGE_NAME := layoffs-api
DEFAULT_IMAGE_TAG := latest
DEFAULT_FULL_IMAGE_NAME := ${DEFAULT_IMAGE_NAME}:${DEFAULT_IMAGE_TAG}
DEFAULT_CONTAINER_NAME := ${DEFAULT_IMAGE_NAME}

image: cmd/layoffs/main.go
	docker build -f Dockerfile -t ${DEFAULT_FULL_IMAGE_NAME} .

image-clean:
	docker rmi ${DEFAULT_FULL_IMAGE_NAME}

container: image
	docker run \
					--name ${DEFAULT_CONTAINER_NAME} \
					-p 3000:3000 \
					${DEFAULT_FULL_IMAGE_NAME}
stop:
	docker stop ${DEFAULT_CONTAINER_NAME}

kill: 
	docker container kill ${DEFAULT_CONTAINER_NAME}



