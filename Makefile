build :
	@echo "Building..."
	@go build -o bin/$(APP_NAME) 

run :
	@go run cmd/$(APP_NAME)/main.go

test :
	@echo "Testing..."		
	@go test -v ./...