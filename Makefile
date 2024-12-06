################################################################################
# Generate gRPC code
# Set PROTO_PATH variable before to use
################################################################################

.DEFAULT_GOAL := help

.PHONY: proto
proto: clean
	cd proto && buf generate --path $(PROTO_PATH)
	@echo "> Successfully generated gRPC code"

.PHONY: clean
clean: check ## Clean code generated folder for proto
	@echo "> Cleaning up previous generated files for $(PROTO_PATH)"
	@rm -rf '_gen_go/$(PROTO_PATH)'

.PHONY: check
check:
	@[ -z "$(PROTO_PATH)" ] && echo "PROTO_PATH is not set" && exit 1 || true

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
