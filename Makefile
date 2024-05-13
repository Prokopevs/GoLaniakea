migrateup:
	migrate -path db/migrations -database "postgres://postgres:2409@localhost:5432/laniakeadb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:2409@localhost:5432/laniakeadb?sslmode=disable" -verbose down

build_image:
	docker build -t server/post-api:latest -f ./server/Dockerfile ./server