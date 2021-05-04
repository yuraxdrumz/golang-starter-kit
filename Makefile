default: init mod
run:
	@go run main.go
init:
	@git config core.hooksPath .githooks
	@echo "successfully added git hooks"
mod:
	@go mod tidy