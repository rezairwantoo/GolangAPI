.PHONY: build test cover-report bindata mock-generate
BINDATA := $(shell command -v go-bindata 2> /dev/null)

build:
	@echo ">> Building ..."
	@go build -o case2 ./
	@echo ">> Finished"

build-migrate:
	@GOOS=linux GOARCH=amd64
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o ./bin/case2 ./
	@echo "build Migration CMD"
	@go mod tidy && go build -o case2-migration ./
	@echo "Finished"

test:
	@go clean -testcache;
	@echo ">> running test for all folders"
	@go test -race ./... -cover -coverpkg=./... -coverprofile=cover.out 
	@go tool cover -func=cover.out | tail -n 1 | awk '{ print $3 }'

lint:
	@golangci-lint run --issues-exit-code 0 --out-format code-climate | jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'

cover-report:
	@go tool cover -html=cover.out

mock-generate:
	@mockgen -source=repository/interface.go -destination=mocks/mock_repository.go -package=mocks
	@mockgen -source=usecase/interface.go -destination=mocks/mock_usecase.go -package=mocks -mock_names Usecase=MockUsecase

postgres-migrate-up:
	@ migrate -path resources/postgres/migrations/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

postgres-migrate-create:
	@migrate create -ext .sql -dir resources/postgres/migrations $(name)
