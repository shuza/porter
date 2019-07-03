include .env

protobuf:
	protoc -I . --go_out=plugins=micro:. ./proto/*.proto

docker_build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/user-service:$(APP_VERSION) .

docker_run: docker_build
	docker run -p 8082:8082 \
	-e DB_HOST:$(DB_HOST) \
	-e DB_USER:$(DB_USER) \
	-e DB_NAME=$(DB_NAME) \
	-e DB_PASSWORD=$(DB_PASSWORD) \
	shuzasa/user-service:$(APP_VERSION)

run:
	MICRO_SERVER_ADDRESS=$(MICRO_SERVER_ADDRESS) \
	DB_HOST=$(DB_HOST) \
	DB_USER=$(DB_USER) \
	DB_NAME=$(DB_NAME) \
	DB_PASSWORD=$(DB_PASSWORD) \
	go run main.go