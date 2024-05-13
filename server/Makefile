migrateup:
	migrate -path db/migrations -database "postgres://postgres:2409@localhost:5432/laniakeadb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:2409@localhost:5432/laniakeadb?sslmode=disable" -verbose down

build_images:
	docker build -t server/post-api:latest -f Dockerfile .