build-push-all:
	@make build-push-ondrejsika
	@make build-push-mavenir
	@make build-push-ogsoft
	@make build-push-thermofisher
	@make build-push-aardwark
	@make build-push-ttc
	@make build-push-tietoevry

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

build-push-ondrejsika:
	@make _build-push NAME=ondrejsika

build-push-mavenir:
	@make _build-push NAME=mavenir

build-push-ogsoft:
	@make _build-push NAME=ogsoft

build-push-thermofisher:
	@make _build-push NAME=thermofisher

build-push-aardwark:
	@make _build-push NAME=aardwark

build-push-ttc:
	@make _build-push NAME=ttc

build-push-tietoevry:
	@make _build-push NAME=tietoevry

use-ondrejsika:
	@make _use NAME=ondrejsika

use-mavenir:
	@make _use NAME=mavenir

use-ogsoft:
	@make _use NAME=ogsoft

use-thermofisher:
	@make _use NAME=thermofisher

use-aardwark:
	@make _use NAME=aardwark

use-ttc:
	@make _use NAME=ttc

use-tietoevry:
	@make _use NAME=tietoevry
