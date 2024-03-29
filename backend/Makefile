# NOTE for dev env
# DB_URL=postgresql://root:secret@host.docker.internal:5432/swole_goal?sslmode=disable
DB_URL=postgresql://root:secret@localhost:5432/swole_goal?sslmode=disable
DB_URL_AWS=postgresql://root:swole-goal.cllc73kmmpqz.us-east-2.rds.amazonaws.com@swole-goal.cllc73kmmpqz.us-east-2.rds.amazonaws.com:5432/swole_goal
# QUICKFIX: if build is somehow not connecting to external DB extensions, and you are using WSL, check if you have a postgres port being listened on by using powershell
# this stack overflow quesiton will guide you well -> https://stackoverflow.com/a/63007311/19919302
# PERMANENT FIX: check to see if postgres runs on start on windows. if it does, disable that, and see if the problemn persists.
build_postgres: # builds the postgres container, with these settings
	docker run --name postgres16 --network swole-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine3.19

start_postgres: # starts the postgres container
	docker start postgres16

stop_postgres: # kills the postgres container
	docker container stop postgres16

createdb: # creates the database 'swole-goal' within the postgres container
	docker exec -it postgres16 createdb --username=root --owner=root swole_goal

dropdb: # drops the database 'swole-goal' within the postgres container
	docker exec -it postgres16 dropdb swole_goal

# steps into our db using psql from the terminal
# to see the correct DB, use \c swole_goal
step_db:
	docker exec -it postgres16 psql -U root

# with the wsl specifications that I have while developing this, this answer below made the go-migration work!
# https://community.grafana.com/t/dial-tcp-127-0-0-1-connect-connection-refused/13071/6
migrateup: # migrates the database 'swole-goal' up to the current migration version
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown: # migrates the database 'swole-goal' back to the previous migration version
	migrate -path db/migration -database "${DB_URL}" -verbose down

# for AWS RDS instance only
migrateupaws: # migrates the database 'swole-goal' up to the current migration version
	migrate -path db/migration -database "${DB_URL_AWS}" -verbose up

# generates a new sqlc compiled sql statement to run
sqlc:
	sqlc generate

# runs the tests that we have for our crud operations
test:
	go clean -testcache
	go test -v -cover ./...

# lints the files found in our go files!
# see https://golangci-lint.run/usage/install/ for more documentation
golint:
	golangci-lint run ./.../...

# lints the files in our backend repo, aside from our migration files (premade)
# see https://docs.sqlfluff.com/en/stable/gettingstarted.html
sqlint:
	sqlfluff lint ./
	@ echo 'always run your sql against this website to make it prettier: https://shorturl.at/xyNZ3'
	@ echo 'this will format your sql to get the passing ci'

# makes the documentation for our golang code easier to read
# see https://pkg.go.dev/golang.org/x/tools/cmd/godoc
godoc:
	@ echo 'http://localhost:9090/pkg/'
	godoc -http :9090

# runs the server
server:
	@ go run main.go

# this tells the makefile that these commands are not related to any files in particular - so makefile does not get confused
# https://stackoverflow.com/a/2145605/19919302
.PHONY: createdb postgres dropdb migrateup migratedown sqlc golint sqlint