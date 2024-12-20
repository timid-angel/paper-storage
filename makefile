MAIN_PKG_DIR = ./server
SRC = ./server/...
BIN_DIR = ./bin
BIN_NAME = runner

.PHONY: build clean

# builds the project, entry file at MAIN_PKG_DIR
build:
	go build -o $(BIN_DIR)/$(BIN_NAME) ${MAIN_PKG_DIR}

# deletes built binaries
clean:
	rm -rf $(BIN_DIR)