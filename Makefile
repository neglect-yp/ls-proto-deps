IMAGE_TAG ?= latest
IMAGE_NAME ?= ghcr.io/neglect-yp/ls-proto-deps:${IMAGE_TAG}

build:
	docker build . -t ${IMAGE_NAME}

push: build
	docker push ${IMAGE_NAME}
