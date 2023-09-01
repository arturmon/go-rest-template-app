.PHONY: build
build:
	@$(MAKE) swagger
	@$(MAKE) info-go
	go build cmd

.PHONY: run
run:
	@$(MAKE) swagger
	go run cmd/main.go

.PHONY: up
up:
	@$(MAKE) swagger
	docker-compose up -d --force-recreate

.PHONY: restart
restart:
	@$(MAKE) down
	@$(MAKE) swagger
	@$(MAKE) up

halt:
	docker-compose stop

.PHONY: down
down:
	docker-compose down --rmi local -v --remove-orphans

init: ##- Runs make modules, tools and tidy.
	@$(MAKE) modules
	@$(MAKE) install-tools
	@$(MAKE) tidy

tidy:
	go mod tidy

install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest

modules:
	go mod download

lint:
	golangci-lint run --timeout 5m

test:
	$(MAKE) up
	go test -v ./...
	${MAKE} down

swagger:
	cd cmd && swag init -d ./,../web

info-go: ##- (opt) Prints go.mod updates, module-name and current go version.
	@echo "[go.mod]" > .info-go
	@$(MAKE) get-go-outdated-modules >> .info-go
	@$(MAKE) info-module-name >> .info-go
	@go version >> .info-go
	@cat .info-go

# TODO: switch to "-m direct" after go 1.17 hits: https://github.com/golang/go/issues/40364
get-go-outdated-modules: ##- (opt) Prints outdated (direct) go modules (from go.mod).
	@((go list -u -m -f '{{if and .Update (not .Indirect)}}{{.}}{{end}}' all) 2>/dev/null | grep " ") || echo "go modules are up-to-date."

info-module-name: ##- (opt) Prints current go module-name.
	@echo "go module-name: '${GO_MODULE_NAME}'"