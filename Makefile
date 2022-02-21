
BUILDARCH=x86_64
TARGET=trapmux
image = alpine:3.15
docker_tag_trapmux = kkearne/trapmux
container_trapmux = trapmux
configuration_path_trapmux = /Users/kellskearney/go/src/trapmux/tools
docker_tag_clickhouse = kkearne/clickhouse
container_clickhouse = clickhouse
#configuration_path_clickhouse = /Users/kellskearney/go/src/trapmux/tools


build:
	cd cmds && ./buildall.sh

build_all: plugins build

plugins:
	cd txPlugins && ./build_plugins.sh

deps:
	go get ./...

test: build
	go test
	cd txPlugins && go test

fmt:
	gofmt -w txPlugins/*.go
	gofmt -w txPlugins/*/*/*.go
	gofmt -w cmds/*/*.go
	git commit -m "gofmt" -a

rpm: build
	rpmbuild -ba tools/rpm.spec

gosec:
	gosec -exclude=G107 ./...

clean: clean_plugins
	rm -rf ~/rpmbuild/BUILD/${TARGET} ~/rpmbuild/BUILD/${BUILDARCH}/*
	go clean

clean_plugins:
	find txPlugins -name \*.so -delete

install:
	cd ~/rpmbuild/RPMS/${BUILDARCH} && sudo yum install -y `ls -1rt | tail -1`

push:
	git push -u origin $(shell git symbolic-ref --short HEAD)

