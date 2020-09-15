IMAGE_TAG ?= latest

build:
	docker build . -t ls-proto-deps:${IMAGE_TAG}

