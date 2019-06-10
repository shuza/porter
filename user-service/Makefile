protobuf:
	protoc -I . --go_out=plugins=grpc:. ./proto/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/user-service .

run:
	docker run -p 8082:8082 -e PORT=:8082 shuzasa/user-service
