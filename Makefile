build:
	@go build -o bin/petpal-go main.go 

run: build
	@./bin/petpal-go
