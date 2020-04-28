.PHONY: dependency unit-test integration-test docker-up docker-down clear 

dependency:
	@go get -v ./...

integration-test:
	@docker-compose -f docker-test/docker-compose.yml up -d
	sleep 5
	-go test -v ./...
	-@docker-compose -f docker-test/docker-compose.yml down

unit-test: dependency
	@go test -v -short ./...

docker-up:
	@docker-compose -f docker/docker-compose.yml up -d

docker-down:
	@docker-compose -f docker/docker-compose.yml down

test-docker-up:
	@docker-compose -f docker-test/docker-compose.yml up -d

test-docker-down:
	@docker-compose -f docker-test/docker-compose.yml down

clear: docker-down
