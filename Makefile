GO111MODULE=on
all:
	go build -o bin/mallserver ./main.go

doc:
	swag init

clean:
	go clean -i ./...
	@rm -rf ./bin