## version
SERVER_VERSION = v1.0.5
## command
GO           = go
GO_VENDOR    = go mod
MKDIR_P      = mkdir -p

## build 
.PHONY: build
build:
	GO111MODULE=on $(GO) build -v -o _output/balancer ./


## dockerfile

.PHONY: docker.build
docker.build: 
	docker build --no-cache --rm --tag charlie1380/balancer:$(SERVER_VERSION) -f ./build/Dockerfile .
