TARGET_ARCH ?= amd64
BUILD_DIR ?= build
BIN_DIR := $(BUILD_DIR)/learn_bbgo

$(BIN_DIR):
	mkdir -p $@

# build learn_bbgo
learn_bbgo:
	go build -o $(BIN_DIR)/$@ ./cmd/learn_bbgo
	
# remove artifact file
clean:
	rm -rf $(BUILD_DIR)

# .PHONY: 