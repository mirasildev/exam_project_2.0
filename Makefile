#.SILENT:
DB_URL=postgresql://postgres:1105@localhost:5432/exam_2?sslmode=disable


swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

#migrateup1:
#	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

#migratedown1:
#	migrate -path migrations -database "$(DB_URL)" -verbose down 1

local-up:
	docker compose --env-file ./.env.docker up -d

.PHONY: start migrateup migratedown
