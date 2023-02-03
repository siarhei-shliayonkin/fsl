RUN	    = go run
BUILD   = CGO_ENABLED=0 go build
TEST    = go test
BIN_DIR = ./bin

GIT_COMMIT := $(shell git describe --tags --dirty=-unsupported --always || echo pre-commit)
IMAGE_VERSION ?= $(GIT_COMMIT)

PROJECT_NAME = fsl
PROJECT_DIR	 = ./cmd
IMAGE_NAME   = $(PROJECT_NAME):$(IMAGE_VERSION)

default: run

.PHONY: run build bindir test
run:
	@$(RUN) $(PROJECT_DIR)

build: bindir
	@$(BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) $(PROJECT_DIR)/main.go

bindir:
	@if [ ! -d $(BIN_DIR) ]; then echo "binary dir does not exist, creating.."; mkdir -p $(BIN_DIR); fi

test:
	$(TEST) -cover ./...

.PHONY: image image-run image-stop
image: build
	@docker build -t $(IMAGE_NAME) .
	docker tag $(IMAGE_NAME) $(PROJECT_NAME):latest

image-run:
	docker run -d --rm -p 8081:8081 --name $(PROJECT_NAME) $(IMAGE_NAME)

image-stop:
	docker stop $(PROJECT_NAME)
