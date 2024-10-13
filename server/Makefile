build:
	go build github.com/Prokopevs/GoLaniakea/server/cmd/server

migrateup:
	migrate -path db/migrations -database "postgres://postgres:postgres@192.168.48.2:5432/laniakeadb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:postgres@192.168.48.2:5432/laniakeadb?sslmode=disable" -verbose down