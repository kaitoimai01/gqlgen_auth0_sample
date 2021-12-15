.PHONY: up
up:
	@docker compose up

.PHONY: upd
upd:
	@docker compose up -d

.PHONY: tidy
tidy:
	@docker compose exec app go mod tidy

.PHONY: fmt
fmt:
	@docker compose exec app go fmt ./...

.PHONY: gqlgen
gqlgen:
	@docker compose exec app go run -mod=mod github.com/99designs/gqlgen generate
