TARGET_ARCH ?= amd64
BUILD_DIR ?= build
BIN_DIR := $(BUILD_DIR)/learn_cobra

$(BIN_DIR):
	mkdir -p $@

# build learn_cobra
learn_cobra:
	go build -o $(BIN_DIR)/$@ ./cmd/learn_cobra
	
# remove artifact file
clean:
	rm -rf $(BUILD_DIR)

# .PHONY: 