# Taco Book Store Example

This is a simple api using taco framework to manage and buy books from a bookstore

### To run:
    $ docker compose up # starts database
    $ go run cmd/api/main.go # starts server

### To ping with DB connection check:
    $ curl http://localhost:8000/api/health/ping

### To check docs:
    http://localhost:8000/docs

    