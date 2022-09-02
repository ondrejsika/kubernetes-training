build-push-all:
	@make build-push-mavenir
	@make build-push-ogsoft

_build-push:
	docker build --platform linux/amd64 \
		--build-arg IMAGE=sikalabs/hello-world-server:${NAME} \
		--tag ondrejsika/training-example:${NAME} \
		- < Dockerfile
	docker push ondrejsika/training-example:${NAME}

_use:
	docker pull ondrejsika/training-example:${NAME}
	docker tag ondrejsika/training-example:${NAME} ondrejsika/training-example
	docker push ondrejsika/training-example

build-push-mavenir:
	@make _build-push NAME=mavenir

build-push-ogsoft:
	@make _build-push NAME=ogsoft

use-mavenir:
	@make _use NAME=mavenir

use-ogsoft:
	@make _use NAME=ogsoft
