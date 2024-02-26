postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine3.19

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root swole_goal

dropdb:
	docker exec -it postgres16 dropdb swole_goal

# https://community.grafana.com/t/dial-tcp-127-0-0-1-connect-connection-refused/13071/6
migrateup:
	migrate -path backend/db/migration -database "postgresql://root:secret@host.docker.internal:5432/swole_goal?sslmode=disable" -verbose up

migratedown:
	migrate -path backend/db/migration -database "postgresql://root:secret@host.docker.internal:5432/swole_goal?sslmode=disable" -verbose down

.PHONY: createdb postgres dropdb migrateup migratedown