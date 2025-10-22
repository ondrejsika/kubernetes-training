all:
	@make _build_and_push IMAGE=ghcr.io/ondrejsika/counter
	@make _build_and_push IMAGE=ghcr.io/ondrejsika/counter-frontend
	@make _build_and_push IMAGE=ghcr.io/ondrejsika/counter-frontend-go
	@make _build_and_push IMAGE=redis

_build_and_push:
	docker build --platform linux/amd64 \
		-t ttl.sh/docker-training/dev/${IMAGE} \
		--build-arg IMAGE=${IMAGE} \
		. && \
	docker push ttl.sh/docker-training/dev/${IMAGE}
