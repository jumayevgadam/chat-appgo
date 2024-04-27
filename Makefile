postgresinit:
	docker run --name postgres-db -p 7701:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -d postgres:latest

postgres:
	docker exec -it postgres-db psql -U postgres

createdb:
	docker exec -it postgres-db createdb --username=postgres -O gadamus go-chatnewnew

dropdb:
	docker exec -it postgres-db dropdb g-chatnewnew

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:7701/go-chatnewnew?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:7701/go-chatnewnew?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown