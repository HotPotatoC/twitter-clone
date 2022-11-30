.DEFAULT_GOAL := help

# Source: https://gist.github.com/prwhite/8168133?permalink_comment_id=3624253#gistcomment-3624253
help: ## Show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nAvailable Commands: \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

start: ## Start containers
	docker-compose up -d

fresh-start: docker-reset start _trigger_delay db-setup-logical-replication ## Start containers with fresh data

stop: ## Stops the containers
	docker-compose down

docker-reset: ## Reset the containers (deletes containers and volumes)
	./reset_containers.sh

db-sync-replica-schema: ## Sync replica schema (pg_dump)
	./pg_sync_replica_schema.sh

db-setup-logical-replication: ## Setup postgresql logical replication
	./pg_setup_logical_replication.sh

_trigger_delay:
	sleep 1