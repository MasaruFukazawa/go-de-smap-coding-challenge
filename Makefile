.PHONY: help up down build restart logs shell clean ps

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

up: ## Start containers
	docker-compose up -d

down: ## Stop and remove containers
	docker-compose down

build: ## Build containers
	docker-compose build

rebuild: ## Rebuild containers
	docker-compose build --no-cache

restart: ## Restart containers
	docker-compose restart

logs: ## Show logs (use CTRL+C to exit)
	docker-compose logs -f

logs-app: ## Show app logs
	docker-compose logs -f app

logs-nginx: ## Show nginx logs
	docker-compose logs -f nginx

shell: ## Enter app container shell
	docker-compose exec app bash

ps: ## Show running containers
	docker-compose ps

clean: ## Stop containers and remove volumes
	docker-compose down -v

dev: ## Start containers and show logs
	docker-compose up

stop: ## Stop containers without removing
	docker-compose stop

start: ## Start existing containers
	docker-compose start
