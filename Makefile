db-up:
	docker-compose -f docker-compose.yml -p gin-budget-control_dev up -d

db-start:
	docker start gin-rest-api_postgres_1

test-db-up:
	docker-compose -f docker-compose.test.yml -p gin-budget-control_test up -d

run:
	./swag init -g application.go 
	go run application.go

run-test:
	go test ./test/...

generate-deploy-package:
	./swag init -g application.go 
	zip -r deploy.zip . -x '*.git*' -x '.env'