MIGRATION_DIR = ./migration
DB_URL = postgres://st-user:st-pass@localhost:6000/chat-db
migrate-up:
	@echo "Running migration up on: $(DB_URL)"
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up 

migrate-down:
	@echo "Rolling back latest migration on: $(DB_URL)" 
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down 

migrate-status:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" status 

migrate-version:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" version 

