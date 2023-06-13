# Go AUTH

Go REST API with JWT authentication

## Requirements

* Go v1.19
* Swagger v0.30.3

## Tasks (Completed)

* Implement authentication API
* Write tests for handlers
* Generate docs
* Enable CORS

## Server

```bash
go run main.go
```

## Tests

```bash
cd tests
go test -run .
```

## Documentation

### Generate docs
```bash
swagger generate spec -o ./swagger.yaml 
```

### Serve docs
```bash
swagger serve swagger.yaml
```