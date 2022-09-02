IMAGE = ondrejsika/metrics-example
CONTAINER = metrics-example

all: build push

build:
	docker build --platform linux/amd64 -t $(IMAGE) .

push:
	docker push $(IMAGE)

run:
	docker run --name $(CONTAINER) -d -p 80:80 $(IMAGE)

stop:
	docker rm -f $(CONTAINER)
