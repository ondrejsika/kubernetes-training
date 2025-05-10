build-and-push:
	@make build
	@make push

build:
	docker build --platform=linux/amd64 -f Dockerfile.php -t ondrejsika/php-example-php .
	docker build --platform=linux/amd64 -f Dockerfile.nginx -t ondrejsika/php-example-nginx .

push:
	docker push ondrejsika/php-example-php
	docker push ondrejsika/php-example-nginx
