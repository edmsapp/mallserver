GO111MODULE=on
all:
	go build -o bin/mallserver ./main.go

clean:
	go clean -i ./...
	@rm -rf ./bin