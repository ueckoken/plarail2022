all: fetch reload

fetch:
	git pull

up:
	docker-compose up -d

reload: build
	docker-compose up -d

build: down
	docker-compose build

down:
	docker-compose down
test:
	go test -v ./...
buildGoBin:
	CGO_ENABLED=0 go build -o app ./cmd/main.go
