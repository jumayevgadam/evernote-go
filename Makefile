run:	
	go run cmd/main.go

tidy:
	go mod tidy

migrate_create:
	migrate create -ext sql -dir internal/database/migrations/ -seq tables	

migrate_up:
	migrate	-path internal/database/migrations/ -database "postgresql://postgres:123456@localhost:5432/evernote?sslmode=disable"	-verbose up

migrate_force:
	migrate -path internal/database/migrations/ -database "postgresql://postgres:123456@localhost:5432/evernote?sslmode=disable" force 1

.PHONY: run tidy migrate_create migrate_run migrate_force swag-init