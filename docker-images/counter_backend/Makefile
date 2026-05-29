IMAGE=ttl.sh/counter-backend
BUILD_ID=0

build-and-push:
	docker build . \
		--build-arg BUILD_ID=${BUILD_ID} \
		-t ${IMAGE}:v${BUILD_ID}
	docker push ${IMAGE}:v${BUILD_ID}
