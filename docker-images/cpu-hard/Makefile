IMAGE = ondrejsika/cpu-hard

all: build push

build:
	docker build --platform linux/amd64 -t $(IMAGE) .

push:
	docker push $(IMAGE)
