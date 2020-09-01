test:
	go test -mod=vendor -v -race -failfast -cover -coverprofile=coverage.out ./services ./handlers || exit 1

e2e-test:
	docker-compose -f docker/e2e/docker-compose.yml down
	docker-compose -f docker/e2e/docker-compose.yml build
	docker-compose -f docker/e2e/docker-compose.yml up --abort-on-container-exit --remove-orphans
	docker-compose -f docker/e2e/docker-compose.yml stop
	docker-compose -f docker/e2e/docker-compose.yml rm -f

coverage-html: test
	go tool cover -html=coverage.out