.PHONY: migrate migrate_down migrate_up migrate_version local swaggo test

# ==============================================================================
# Go migrate postgresql

force:
	migrate -database sqlite3://c/GstData/auth_db -path migrations force 1

version:
	migrate -database sqlite3://c/GstData/auth_db -path migrations version

migrate_up:
	migrate -database sqlite3://C/GstData/auth_db -path migrations up 1

migrate_down:
	migrate -database sqlite3://c:\\GstData\\auth_db -path migrations down 1

# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g cmd/api/main.go

# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...


# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache