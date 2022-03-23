db-up:
	docker-compose -f docker-compose.db.yml -p gin-budget-control_dev up -d

db-start:
	docker start gin-budget-control_dev

test-db-up:
	docker-compose -f docker-compose.db.test.yml -p gin-budget-control_test up -d

test-db-start:
	docker start gin-budget-control_test

swagger:
	swag init --exclude router/ -g src/application.go 

run:
	make swagger
	go run application.go

run-test:
	go test ./test/...

generate-deploy-package:
	make swagger
	zip -r deploy.zip . -x '*.git*' -x '.env'

build:
	go build -o application application.go