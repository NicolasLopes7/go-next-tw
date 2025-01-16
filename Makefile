.DEFAULT_GOAL=run
.PHONY: server

server:
	@go run cmd/server/*.go

gen-pages:
	@cd test-next-app && pnpm build && cd ..

deps:
	@go mod tidy && cd test-next-app && pnpm install

run: gen-pages server