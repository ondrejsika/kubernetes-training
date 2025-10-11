help:
	@echo see Makefile content: cat Makefile

fmt:
	yarn run prettier-write

fmt-check:
	yarn run prettier-check

setup-git-hooks:
	rm -rf .git/hooks
	(cd .git && ln -s ../.git-hooks hooks)
