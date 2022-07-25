MAIN_ROOT ?= ./toolkit
PKG = "github.com/saitofun/qkit/gen/cmd"

install: download
	cd $(MAIN_ROOT) && go install

build:
	cd $(MAIN_ROOT) && go build

download:
	go mod download -x

dep:
	go get -u ./...
