dep:
	GO111MODULE=on go mod download

test:
	go test ./...

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	go build -tags mock -o playground .

local/docker-build:
	docker build -t go-playground:local .

start:
	docker-compose up -d
