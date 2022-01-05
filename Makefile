.PHONY: run_local
run_local:
	sudo docker-compose up -d --force-recreate --remove-orphans

.PHONY: run_local_tests
run_local_tests:
	sudo docker-compose up -d --force-recreate --remove-orphans server_db redis

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: stop_local
stop_local:
	docker-compose down

.PHONY: remove_containers
remove_containers:
	-docker stop $$(docker ps -aq)
	-docker rm $$(docker ps -aq)

.PHONY: integration_test
integration_test:
	go test -tags=integration ./integration_tests -count=1 -run=$(INTEGRATION_TEST_SUITE_PATH) 

.PHONY: armageddon
armageddon:
	-make remove_containers
	-docker builder prune -f
	-docker network prune -f
	-docker volume rm $$(docker volume ls --filter dangling=true -q)
	-docker rmi $$(docker images -a -q) -f


.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage1.out -coverpkg=./... -cover ./...
	cat coverage1.out | grep -v mock | grep -v proto | grep -v cmd | grep -v models > cover.out
	go tool cover -func cover.out && go tool cover -html cover.out

.DEFAULT_GOAL := run_local
