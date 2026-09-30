LOCAL_BIN=$(CURDIR)/.bin
PROJECT_NAME=enrichment

include .make-deps/*.mk

args=`arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

env=./.env
composefile=./docker/docker-compose.yml

.PHONY: fumpt
fumpt: $(GOFUMPT_BIN)
	$(GO_ENV) $(GOFUMPT_BIN) -w ./cmd ./internal

.PHONY: lint
lint: $(GOLANG_LINT_BIN) fumpt
	$(GO_ENV) $(GOLANG_LINT_BIN) run -v --fix ./...

server:
	go run ./cmd/main.go

## Запуск в docker-compose с ребилдом контейнера detached
dc-up: dc-down
	docker compose -f $(composefile) --env-file $(env) up --build -d

dc-down:
	docker compose -f $(composefile) --env-file $(env) down
