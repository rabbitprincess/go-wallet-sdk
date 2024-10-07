build:
	@mkdir -p ./bin; \
	for moddir in $$(find . -name 'go.mod' -exec dirname {} \;); do \
		basename=$$(basename $$moddir); \
		echo "Building $$basename..."; \
		cd $$moddir && go build -o ../bin/$$basename; \
	done
.PHONY: build

test:
	@p=$(shell pwd); \
	for module in $$(grep 'use' go.work | awk '{print $$2}'); do \
		echo "Running tests in $$module..."; \
		cd $$p/$$module && go test ./...; \
	done
.PHONY: test

format:
	@find $(WORKSPACE_DIR) -name "*.go" | while read -r gofile; do \
		gofmt -w "$$gofile"; \
		goimports -w "$$gofile"; \
		gopls format "$$gofile"; \
		echo "Formatted: $$gofile"; \
	done
.PHONY: format
