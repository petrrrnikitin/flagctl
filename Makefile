COMPOSE = docker compose -f infra/docker-compose.yml --env-file .env

.PHONY: up down logs ps migrate \
        api-gateway-test api-gateway-lint api-gateway-deptrac \
        auth-test auth-lint \
        evaluator-test evaluator-lint \
        audit-test audit-lint \
        frontend-test frontend-lint

# ─── Infra ────────────────────────────────────────────────────────────────────

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logs -f $(service)

ps:
	$(COMPOSE) ps

migrate:
	$(COMPOSE) exec api-gateway php bin/console doctrine:migrations:migrate --no-interaction

# ─── API Gateway ──────────────────────────────────────────────────────────────

api-gateway-test:
	$(COMPOSE) exec api-gateway php bin/phpunit

api-gateway-lint:
	$(COMPOSE) exec api-gateway vendor/bin/phpstan analyse

api-gateway-deptrac:
	$(COMPOSE) exec api-gateway vendor/bin/deptrac analyse

# ─── Auth Service ─────────────────────────────────────────────────────────────

auth-test:
	cd services/auth && go test ./...

auth-lint:
	cd services/auth && golangci-lint run ./...

# ─── Evaluator Service ────────────────────────────────────────────────────────

evaluator-test:
	cd services/evaluator && go test ./...

evaluator-lint:
	cd services/evaluator && golangci-lint run ./...

# ─── Audit Service ────────────────────────────────────────────────────────────

audit-test:
	cd services/audit && go test ./...

audit-lint:
	cd services/audit && golangci-lint run ./...

# ─── Frontend ─────────────────────────────────────────────────────────────────

frontend-test:
	cd frontend && npm run test

frontend-lint:
	cd frontend && npm run lint