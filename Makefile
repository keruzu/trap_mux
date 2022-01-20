
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
	go build

build_all: plugins build

plugins:
	cd txPlugins && ./build_plugins.sh

deps:
	go get ./...

test: build
	go test
	cd txPlugins && go test

fmt:
	gofmt -w *.go
	gofmt -w txPlugins/*.go
	gofmt -w txPlugins/actions/*/*.go
	gofmt -w txPlugins/generators/*/*.go
	gofmt -w txPlugins/metrics/*/*.go
	gofmt -w cmds/*/*.go
	git commit -m "gofmt" -a

rpm: build
	rpmbuild -ba tools/rpm.spec

clean: clean_plugins
	rm -rf ~/rpmbuild/BUILD/${TARGET} ~/rpmbuild/BUILD/${BUILDARCH}/*
	go clean

clean_plugins:
	find txPlugins -name \*.so -delete

install:
	cd ~/rpmbuild/RPMS/${BUILDARCH} && sudo yum install -y `ls -1rt | tail -1`

push:
	git push -u origin $(shell git symbolic-ref --short HEAD)

# ----  Docker: trapmux  ----------------------------
.PHONY: trapmux
trapmux:
	DOCKER_BUILD=0 docker build -t $(docker_tag_trapmux) -f tools/docker/Dockerfile .

trapmux_aws:
	DOCKER_BUILD=0 docker build -t $(docker_tag_trapmux) -f tools/docker/Dockerfile.amazonlinux .

run:
	docker run --name $(container_trapmux) -v $(configuration_path):/opt/trapmux/etc -p 162:162 -p 5080:80 $(docker_tag_trapmux)

stop:
	docker stop $(container_trapmux)
	docker rm $(container_trapmux)

pull:
	docker pull $(image)

# ----  Docker: clickhouse  ----------------------------
clickhouse:
	DOCKER_BUILD=0 docker build -t $(docker_tag_clickhouse) -f tools/docker/Dockerfile.clickhouse .

run_click:
	docker run --name $(container_clickhouse) -v $(configuration_path):/opt/trapmux/etc -p 162:162 -p 5080:80 $(docker_tag_clickhouse)

stop_click:
	docker stop $(container_trapmux)
	docker rm $(container_trapmux)

# ----  AWS  ----------------------------
codebuild:
# Need to run the following first
# aws configure
	#aws cloudformation deploy --template-file tools/aws/codebuild_cfn.yml --stack-name trapmuxrpm --capabilities CAPABILITY_IAM
	aws cloudformation deploy --template-file tools/aws/codebuild_docker.yml --stack-name trapmuxdocker --capabilities CAPABILITY_IAM
	#aws cloudformation deploy --template-file tools/aws/codebuild_batch_cfn.yml --stack-name trapmuxbatchrpm --capabilities CAPABILITY_IAM --parameter-overrides StreamId=rpm BuildSpec=tools/aws/buildspec_batch_rpm.yml
	#aws cloudformation deploy --template-file tools/aws/codebuild_batch_cfn.yml --stack-name trapmuxbatchnopkg --capabilities CAPABILITY_IAM --parameter-overrides StreamId=nopkg BuildSpec=tools/aws/buildspec_batch_nopkg.yml CodeBuildImage=aws/codebuild/standard:5.0 

