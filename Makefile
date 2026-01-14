dep:
	GO111MODULE=on go mod download

test:
	go test ./...

run:
	go run main.go

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	go build -tags mock -o playground .

local/docker-build:
	docker build -t go-playground:local .

start:
	docker-compose up -d

ratelimiter:
	go run cmd/ratelimiter/main.go

logparser:
	go run cmd/logparser/main.go -path=logs -query="service=auth level=warning"
