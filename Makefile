run:
	go run src/main.go
dev:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run src/main.go
test:
	go test -v ./...
generateMigration:
	migrate create -seq -ext=.sql -dir=./src/db/psql/migrations $(name)
migrateUp:
	migrate -path=./src/db/psql/migrations -database=$(TWEET_DSN) up
migrateDown:
	migrate -path=./src/db/psql/migrations -database=$(TWEET_DSN) down 1