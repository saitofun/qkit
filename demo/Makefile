
install_toolkit:
	go install github.com/saitofun/qkit/gen/cmd/...

generate: install_toolkit
	go generate ./...

migrate:
	go run cmd/demo/main.go migrate
