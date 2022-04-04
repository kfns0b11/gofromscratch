SHELL := /bin/bash

run:
	go run main.go

VERSION := 1.0.0

all: service

service:
	docker build \
		-f build/docker/Dockerfile \
		-t service:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		.


KIND_CLUSTER := test-k8s

kind-up:
	kind create cluster \
		--name $(KIND_CLUSTER) \
		--config deployment/kind/kind-config.yaml

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)
