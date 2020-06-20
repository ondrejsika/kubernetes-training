IMAGE = ondrejsika/liveness

all: build push

build:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)
