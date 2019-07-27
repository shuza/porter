docker_build:
	@echo "Building binary ..."
	GOOS=linux GOARCH=amd64 go build

	@echo "Building docker image ..."
	docker build -t shuzasa/user-cli:v1.0 .

docker_run: docker_build
	@echo "Running docker image ..."
	docker run --net="host" \
	shuzasa/user-cli

deploy: docker_build
	kubectl apply -f cli-deploy.yaml

delete:
	kubectl delete deployment cli