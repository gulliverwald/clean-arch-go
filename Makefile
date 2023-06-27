SHELL:=/bin/bash
ROOT_DIR:=$(shell dirname $(shell pwd))

.PHONY: up
up: ## Start the app in docker 
	cd docker && docker compose up -d

.PHONY: migrate-db
migrate-db: ## Migrate the database.sql into the mysql container
	cd docker && docker exec -i clean_arch_go_mysql mysql -uuser -ppassword example < database.sql

.PHONY: clean
clean: ## Clean up the app in docker
	docker rmi clean-arch

.PHONY: update-app
update-app: ## Update the app in docker
	cd docker &&  docker rm --force clean_arch_app && docker rmi clean-arch && docker compose up -d app

.PHONY: down
down: ## Remove the app in docker 
	cd docker && docker compose down && rm -r example.sql

.PHONY: help
help: ## Print the targets and their descriptions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
