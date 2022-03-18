# Gin-Budget-Controll

This project is a improvment made from my [Alura Challenge Backend](https://github.com/RaphaSalomao/alura-challenge-backend)
## Run Locally

What you will need: Golang v1.18 or greater, Docker, Linux operating system and [Swagger](https://github.com/swaggo/swag#getting-started)

Clone the project

```bash
git clone git@github.com:RaphaSalomao/alura-challenge-backend.git
```

Run postgres docker container

```bash
make db-up
```

Start the server
```bash
make run
```
If you don't have swag binary on the project root, run this instead
```bash
go run application.go
```

Access the documentation at http://localhost:5000/swagger/index.html

## Running Tests

To run tests, run the following commands

Run postgres test database docker container 
```bash
make test-db-up
```

Run tests
```bash
make run-test
```
