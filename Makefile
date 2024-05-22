migrate-up:
	migrate -path database/migration/ -database "postgresql://postgres:root@localhost:5432/mangdb?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration/ -database "postgresql://postgres:root@localhost:5432/mangdb?sslmode=disable" -verbose down