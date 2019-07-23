docker_run:
	@echo "Building binary ..."
	GOOS=linux GOARCH=amd64 go build

	@echo "Building docker image ..."
	docker build -t shuzasa/user-cli .

	@echo "Running docker image ..."
	docker run --net="host" \
	shuzasa/user-cli