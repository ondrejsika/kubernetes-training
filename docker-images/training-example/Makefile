build-push-all:
	@make build-push-default
	@make build-push-ondrejsika
	@make build-push-argocd
	@make build-push-mavenir
	@make build-push-ogsoft
	@make build-push-thermofisher
	@make build-push-aardwark
	@make build-push-ttc
	@make build-push-tietoevry
	@make build-push-spcss
	@make build-push-daktela
	@make build-push-heidelbergmaterials
	@make build-push-eon
	@make build-push-ceps
	@make build-push-generali
	@make build-push-elm
	@make build-push-t-mobile
	@make build-push-rozhlas
	@make build-push-era
	@make build-push-gpe
	@make build-push-exainfra
	@make build-push-cloud-native-bratislava
	@make build-push-continental
	@make build-push-cez
	@make build-push-vitkovice
	@make build-push-ostrava
	@make build-push-ataccama
	@make build-push-ica

_build-push:
	docker build --platform linux/amd64 \
		--build-arg IMAGE=sikalabs/hello-world-server:${NAME} \
		--tag ondrejsika/training-example:${NAME} \
		- < Dockerfile
	docker push ondrejsika/training-example:${NAME}
	# ghcr.io
	docker tag ondrejsika/training-example:${NAME} ghcr.io/ondrejsika/training-example:${NAME}
	docker push ghcr.io/ondrejsika/training-example:${NAME}

_use:
	docker pull ondrejsika/training-example:${NAME}
	docker tag ondrejsika/training-example:${NAME} ondrejsika/training-example
	docker push ondrejsika/training-example
	# ghcr.io
	docker tag ondrejsika/training-example ghcr.io/ondrejsika/training-example
	docker push ghcr.io/ondrejsika/training-example

# === Build and Push ===

build-push-default:
	@make _build-push NAME=latest

build-push-ondrejsika:
	@make _build-push NAME=ondrejsika

build-push-argocd:
	@make _build-push NAME=argocd

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

build-push-frequentis:
	@make _build-push NAME=frequentis

build-push-spcss:
	@make _build-push NAME=spcss

build-push-daktela:
	@make _build-push NAME=daktela

build-push-heidelbergmaterials:
	@make _build-push NAME=heidelbergmaterials

build-push-eon:
	@make _build-push NAME=eon

build-push-ceps:
	@make _build-push NAME=ceps

build-push-generali:
	@make _build-push NAME=generali

build-push-elm:
	@make _build-push NAME=elm

build-push-t-mobile:
	@make _build-push NAME=t-mobile

build-push-rozhlas:
	@make _build-push NAME=rozhlas

build-push-era:
	@make _build-push NAME=era

build-push-gpe:
	@make _build-push NAME=gpe

build-push-exainfra:
	@make _build-push NAME=exainfra

build-push-cloud-native-bratislava:
	@make _build-push NAME=cloud-native-bratislava

build-push-continental:
	@make _build-push NAME=continental

build-push-cez:
	@make _build-push NAME=cez

build-push-vitkovice:
	@make _build-push NAME=vitkovice

build-push-ostrava:
	@make _build-push NAME=ostrava

build-push-ataccama:
	@make _build-push NAME=ataccama

build-push-ica:
	@make _build-push NAME=ica

# === Use ===

use-default:
	@make _use NAME=default

use-ondrejsika:
	@make _use NAME=ondrejsika

use-argocd:
	@make _use NAME=argocd

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

use-frequentis:
	@make _use NAME=frequentis

use-spcss:
	@make _use NAME=spcss

use-daktela:
	@make _use NAME=daktela

use-heidelbergmaterials:
	@make _use NAME=heidelbergmaterials

use-eon:
	@make _use NAME=eon

use-ceps:
	@make _use NAME=ceps

use-generali:
	@make _use NAME=generali

use-elm:
	@make _use NAME=elm

use-t-mobile:
	@make _use NAME=t-mobile

use-rozhlas:
	@make _use NAME=rozhlas

use-era:
	@make _use NAME=era

use-gpe:
	@make _use NAME=gpe

use-exainfra:
	@make _use NAME=exainfra

use-cloud-native-bratislava:
	@make _use NAME=cloud-native-bratislava

use-continental:
	@make _use NAME=continental

use-cez:
	@make _use NAME=cez

use-vitkovice:
	@make _use NAME=vitkovice

use-ostrava:
	@make _use NAME=ostrava

use-ataccama:
	@make _use NAME=ataccama

use-ica:
	@make _use NAME=ica
