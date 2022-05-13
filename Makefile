build:
	go build -o ./bin/binge main.go

test:
	go test -v ./...

run:
	export PORT=8080 && export DATABASE_URL=postgres://viveknathani:root@localhost:5432/binge && ./bin/binge