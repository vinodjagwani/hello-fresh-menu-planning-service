BINARY_NAME=hello-fresh-menu-planning-service
export PROFILE=$(ENV_PROFILE)

docker.start:
	docker-compose -f ./docker/docker-compose.yml up --build

docker.stop:
	docker-compose -f ./docker/docker-compose.yml down;

test.integration:
	 GO111MODULE=on go test -tags=integration ./internal/adaptors/e2e_test -v -count=1

generate.swagger:
	swag init