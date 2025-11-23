build-push:
	@make _build-push VERSION=v0
	@make _build-push VERSION=v1
	@make _build-push VERSION=v2
	@make _build-push VERSION=v3
	@make _build-push VERSION=v4
	@make _build-push VERSION=v5

_build-push:
	docker build --platform linux/amd64 \
		--build-arg IMAGE=ghcr.io/sikalabs/hello-world-server \
		--build-arg VERSION=${VERSION} \
		--tag ondrejsika/training-example:${VERSION} \
		- < Dockerfile
	docker push ondrejsika/training-example:${VERSION}
	# ghcr.io
	docker tag ondrejsika/training-example:${VERSION} ghcr.io/ondrejsika/training-example:${VERSION}
	docker push ghcr.io/ondrejsika/training-example:${VERSION}
