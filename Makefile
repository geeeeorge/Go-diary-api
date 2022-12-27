ENV_LOCAL=\
	DIARY_APP_PORT=8080 \
	DIARY_APP_HOST=127.0.0.1 \
	DIARY_DB_HOST=localhost \
	DIARY_DB_PORT=5432 \
	DIARY_DB_NAME=daily_app \
	DIARY_DB_USER=postgres \
	DIARY_DB_PASSWORD=postgres

ENV_INTEGRATION=\
	DIARY_APP_PORT=8080 \
	DIARY_APP_HOST=127.0.0.1 \
	DIARY_DB_HOST=localhost \
	DIARY_DB_PORT=5433 \
	DIARY_DB_NAME=daily_app_test \
	DIARY_DB_USER=postgres \
	DIARY_DB_PASSWORD=postgres

.PHONY: app-start
app-start:
	$(ENV_LOCAL) \
	go run cmd/go-diary-api/main.go

.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/geeeeorge/Go-diary-api"
