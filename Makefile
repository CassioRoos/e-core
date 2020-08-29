
e2e-test:
	rm -rf docker/e2e/.data
	mkdir -p docker/e2e/.data
	docker-compose -f docker/e2e/docker-compose.yml down
	docker-compose -f docker/e2e/docker-compose.yml build
	docker-compose -f docker/e2e/docker-compose.yml up --abort-on-container-exit --remove-orphans
	docker-compose -f docker/e2e/docker-compose.yml stop
	docker-compose -f docker/e2e/docker-compose.yml rm -f