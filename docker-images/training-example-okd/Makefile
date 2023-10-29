build-push-all:
	@make build-push-ondrejsika
	@make build-push-digiteq
	@make build-push-ceps

_build-push:
	docker build --platform linux/amd64 \
		--build-arg IMAGE=sikalabs/hello-world-server:${NAME} \
		--tag ondrejsika/training-example-okd:${NAME} \
		--tag ghcr.io/ondrejsika/training-example-okd:${NAME} \
		- < Dockerfile
	docker push ondrejsika/training-example-okd:${NAME}
	docker push ghcr.io/ondrejsika/training-example-okd:${NAME}

_use:
	docker pull ondrejsika/training-example-okd:${NAME}
	docker tag ondrejsika/training-example-okd:${NAME} ondrejsika/training-example-okd
	docker tag ondrejsika/training-example-okd:${NAME} ghcr.io/ondrejsika/training-example-okd
	docker push ondrejsika/training-example-okd
	docker push ghcr.io/ondrejsika/training-example-okd

build-push-ondrejsika:
	@make _build-push NAME=ondrejsika

build-push-digiteq:
	@make _build-push NAME=digiteq

build-push-ceps:
	@make _build-push NAME=ceps

use-ondrejsika:
	@make _use NAME=ondrejsika

use-digiteq:
	@make _use NAME=digiteq

use-ceps:
	@make _use NAME=ceps
