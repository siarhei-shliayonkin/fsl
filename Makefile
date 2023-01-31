RUN	    = go run
BUILD   = go build
TEST    = go test
BIN_DIR = ./bin

PROJECT_NAME = fsl
PROJECT_DIR	 = ./cmd

default: run

.PHONY: run build bindir test
run:
	@$(RUN) $(PROJECT_DIR)

build: bindir
	@$(BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) $(PROJECT_DIR)/main.go

bindir:
	@if [ ! -d $(BIN_DIR) ]; then echo "binary dir does not exist, creating.."; mkdir -p $(BIN_DIR); fi

test:
	$(TEST) -v ./...
