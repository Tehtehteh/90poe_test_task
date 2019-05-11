build:
	export GO111MODULE="on"; \
	go mod download; \
	go mod vendor; \
	GO111MODULE="on" CGO_ENABLED=0 go build -mod=vendor -a -ldflags '-s' -installsuffix cgo -o main cmd/server/main.go

install:
	export GO111MODULE="on"; \
	go mod download;

run:
	export GO111MODULE="on"; \
	go run cmd/server/main.go

test:
	go test -race ./...
