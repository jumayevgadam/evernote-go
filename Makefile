local_run:	
	go run cmd/main.go -config=local

swag_init:
	swag init -g ./cmd/main.go -o ./docs --pd

tidy:
	go mod tidy

migrate_create:
	migrate create -ext sql -dir internal/database/migrations/ -seq tables	

migrate_up:
	migrate	-path internal/database/migrations/ -database "postgresql://postgres:123456@localhost:5432/evernote?sslmode=disable"	-verbose up

migrate_force:
	migrate -path internal/database/migrations/ -database "postgresql://postgres:123456@localhost:5432/evernote?sslmode=disable" force 2

.PHONY: local_run tidy migrate_create migrate_run migrate_force