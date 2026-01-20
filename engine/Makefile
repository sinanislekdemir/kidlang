# Makefile

DIST_DIR := dist

.PHONY: default
default: build

install:
	@go install

build:
	@mkdir -p $(DIST_DIR)
	@go build -o $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH)$(shell test "$(shell go env GOOS)" = "windows" && echo ".exe" || echo "")

linux-amd64:
	@mkdir -p $(DIST_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/kidlang-linux-amd64

linux-arm64:
	@mkdir -p $(DIST_DIR)
	@GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/kidlang-linux-arm64

darwin-amd64:
	@mkdir -p $(DIST_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/kidlang-darwin-amd64

darwin-arm64:
	@mkdir -p $(DIST_DIR)
	@GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/kidlang-darwin-arm64

win64:
	@mkdir -p $(DIST_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/kidlang-windows-amd64.exe

windows-amd64:
	@mkdir -p $(DIST_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/kidlang-windows-amd64.exe

windows-arm64:
	@mkdir -p $(DIST_DIR)
	@GOOS=windows GOARCH=arm64 go build -o $(DIST_DIR)/kidlang-windows-arm64.exe

all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64 windows-arm64

release:
	@mkdir -p $(DIST_DIR)
	@echo "Building release binaries with optimizations..."
	@echo "Note: Darwin builds require macOS due to CGO/ncurses dependencies"
	@echo "Note: Cross-compiling ARM64 from AMD64 may fail with CGO"
	@echo ""
	@echo "Building Linux AMD64..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-linux-amd64
	@echo "Building Windows AMD64..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-windows-amd64.exe
	@echo "Building Windows ARM64..."
	@GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-windows-arm64.exe
	@echo ""
	@echo "Creating compressed archives..."
	@cd $(DIST_DIR) && tar -czf kidlang-linux-amd64.tar.gz kidlang-linux-amd64 && echo "  ✓ kidlang-linux-amd64.tar.gz"
	@cd $(DIST_DIR) && zip -q kidlang-windows-amd64.zip kidlang-windows-amd64.exe && echo "  ✓ kidlang-windows-amd64.zip"
	@cd $(DIST_DIR) && zip -q kidlang-windows-arm64.zip kidlang-windows-arm64.exe && echo "  ✓ kidlang-windows-arm64.zip"
	@echo ""
	@echo "Release packages created in $(DIST_DIR)/:"
	@ls -lh $(DIST_DIR)/*.tar.gz $(DIST_DIR)/*.zip 2>/dev/null
	@echo ""
	@echo "Cleaning up uncompressed binaries..."
	@rm -f $(DIST_DIR)/kidlang-linux-amd64 $(DIST_DIR)/kidlang-windows-*.exe
	@echo "Done! Release files are in $(DIST_DIR)/"

release-darwin:
	@mkdir -p $(DIST_DIR)
	@echo "Building Darwin releases (must be run on macOS)..."
	@echo "Building Darwin AMD64..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-darwin-amd64
	@echo "Building Darwin ARM64..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-darwin-arm64
	@echo "Creating compressed archives..."
	@cd $(DIST_DIR) && tar -czf kidlang-darwin-amd64.tar.gz kidlang-darwin-amd64 && echo "  ✓ kidlang-darwin-amd64.tar.gz"
	@cd $(DIST_DIR) && tar -czf kidlang-darwin-arm64.tar.gz kidlang-darwin-arm64 && echo "  ✓ kidlang-darwin-arm64.tar.gz"
	@echo "Cleaning up uncompressed binaries..."
	@rm -f $(DIST_DIR)/kidlang-darwin-amd64 $(DIST_DIR)/kidlang-darwin-arm64
	@echo "Done! Darwin releases are in $(DIST_DIR)/"

release-current:
	@mkdir -p $(DIST_DIR)
	@echo "Building optimized release for current platform..."
	@go build -ldflags="-s -w" -trimpath -o $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH)$(shell test "$(shell go env GOOS)" = "windows" && echo ".exe" || echo "")
	@echo "Creating archive..."
	@if [ "$(shell go env GOOS)" = "windows" ]; then \
		cd $(DIST_DIR) && zip -q kidlang-$(shell go env GOOS)-$(shell go env GOARCH).zip kidlang-$(shell go env GOOS)-$(shell go env GOARCH).exe; \
		echo "Created: $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH).zip"; \
	else \
		cd $(DIST_DIR) && tar -czf kidlang-$(shell go env GOOS)-$(shell go env GOARCH).tar.gz kidlang-$(shell go env GOOS)-$(shell go env GOARCH); \
		echo "Created: $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH).tar.gz"; \
	fi
	@rm -f $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH) $(DIST_DIR)/kidlang-$(shell go env GOOS)-$(shell go env GOARCH).exe

clean:
	@rm -rf $(DIST_DIR)

debug-test:
	dlv test github.com/sinanislekdemir/kidlang/$(pkg) -- -test.v

test:
	@go test -v ./...

fmt:
	@go fmt ./...
