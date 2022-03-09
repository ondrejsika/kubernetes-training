help:
	@echo see Makefile content: cat Makefile

fmt:
	yarn run prettier-write

fmt-check:
	yarn run prettier-check

minikube-image-ls:
	minikube image ls

minikube-image-save:
	minikube image ls > minikube-images.txt

minikube-image-load:
	cat minikube-images.txt | xargs minikube image load

minikube-image-commit:
	git add minikube-images.txt
	git commit -m "[generated] chore(minikube-images.txt): Update minikube-images.txt"
