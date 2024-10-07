test:
	go test -v all
.PHONY: test

format:
	@find $(WORKSPACE_DIR) -name "*.go" | while read -r gofile; do \
		gofmt -w "$$gofile"; \
		goimports -w "$$gofile"; \
		gopls format "$$gofile"; \
		echo "Formatted: $$gofile"; \
	done
.PHONY: format
