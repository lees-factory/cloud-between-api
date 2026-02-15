.PHONY: build run deploy-build clean

APP_NAME := cloud-between-api

build:
	go build -o $(APP_NAME) .

run:
	go run .

deploy-build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(APP_NAME) .

clean:
	rm -f $(APP_NAME)
