docker_run:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/user-cli .
	docker run --net="host" \
	shuzasa/user-cli