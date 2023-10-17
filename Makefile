SHELL=/bin/sh
GO := $(shell command -v go 2> /dev/null)
TERN := $(shell command -v tern 2> /dev/null)

check-go:
ifndef GO
	$(error "go is not installed! Aborting")
endif

check-tern:
ifndef TERN
	go install github.com/jackc/tern@latest
endif

migrate: check-go check-tern
	tern migrate -m "sql/migration"