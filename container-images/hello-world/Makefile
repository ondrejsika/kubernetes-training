all:
	make build-push \
		SRC=sikalabs/hello-world-server:ondrejsika \
		DST=ondrejsika/training-example
	make build-push \
		SRC=sikalabs/hello-world-server:o2 \
		DST=ondrejsika/training-example-o2
	make build-push \
		SRC=sikalabs/hello-world-server:cra \
		DST=ondrejsika/training-example-cra
	make build-push \
		SRC=sikalabs/hello-world-server:trask \
		DST=ondrejsika/training-example-trask
	make build-push \
		SRC=sikalabs/hello-world-server:mstp \
		DST=ondrejsika/training-example-mstp
	make build-push \
		SRC=sikalabs/hello-world-server:atos \
		DST=ondrejsika/training-example-atos
	make build-push \
		SRC=sikalabs/hello-world-server:red \
		DST=ondrejsika/training-example:red
	make build-push \
		SRC=sikalabs/hello-world-server:blue \
		DST=ondrejsika/training-example:blue
	make build-push \
		SRC=sikalabs/hello-world-server:green \
		DST=ondrejsika/training-example:green
	make build-push \
		SRC=sikalabs/hello-world-server:yellow \
		DST=ondrejsika/training-example:yellow


use-o2:
	make pull-tag-push \
		SRC=ondrejsika/training-example-o2 \
		DST=ondrejsika/training-example

use-cra:
	make pull-tag-push \
		SRC=ondrejsika/training-example-cra \
		DST=ondrejsika/training-example

use-trask:
	make pull-tag-push \
		SRC=ondrejsika/training-example-trask \
		DST=ondrejsika/training-example

use-mstp:
	make pull-tag-push \
		SRC=ondrejsika/training-example-mstp \
		DST=ondrejsika/training-example

use-atos:
	make pull-tag-push \
		SRC=ondrejsika/training-example-atos \
		DST=ondrejsika/training-example

build-push:
	docker pull ${SRC}
	docker build \
	--build-arg IMAGE=${SRC} \
	--tag ${DST} \
	- < Dockerfile
	docker push ${DST}

pull-tag-push:
	docker pull ${SRC}
	docker tag ${SRC} ${DST}
	docker push ${DST}
