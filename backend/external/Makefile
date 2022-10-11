build:
	GOARCH=amd64 CGO_ENABLED=0 go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o external cmd/main.go
test:
	go test -v ./...
up: 
	docker-compose up --build
