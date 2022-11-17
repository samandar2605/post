swag:
	swag init -g ./api/api.go -o api/docs 

run:
	go run cmd/main.go


DB_URL=postgresql://postgres:12345@localhost:5432/blog?sslmode=disable

migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

.PHONY: start migrateup migratedown