MIGRATIONS_DIR=./migration/scripts
CURRENT_TIME=$(shell date +"%Y%m%d")

.PHONY: create-migration

create-migration:
ifndef name
	$(error name is undefined)
endif
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(CURRENT_TIME)_$(name)
	@echo "Migration files created in $(MIGRATIONS_DIR)"

build-di:
	wire ./internal/di