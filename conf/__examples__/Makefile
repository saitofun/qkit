
gen_models: install_tools
	@echo "generate models"
	go generate github.com/iotexproject/w3bstream/base/examples/pkg/models
	@echo ""

install_tools:
	@echo "install tools"
	go install github.com/iotexproject/w3bstream/base/tools/cmd/cmd/tools
	@echo ""

start_postgres:
	docker-compose -f docker-compose-pg.yaml up

init_db:
	@echo "init database"
	psql 'postgres://pguser:pgpassword@127.0.0.1:5432' -c 'create database demo' && psql 'postgres://pguser:pgpassword@127.0.0.1:5432' -c 'create schema demo'
	@echo ""

download:
	@echo "download go modules"
	go mod download
	@echo ""

dep:
	go get -u ./...

build: download install_tools gen_models
	@echo "build demo"
	cd cmd/demo && go build -v -o demo
	mkdir -pv bin
	mv cmd/demo/demo bin/
	@echo ""

run: build
	cp cmd/demo/config bin/ -r
	cd bin && ./demo

clean:
	docker-compose -f docker-compose-pg.yaml down; rm -rf data sqlxdemo
	rm -rf data log
