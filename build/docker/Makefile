
BUILDARCH=x86_64
TARGET=trapmux
image = almalinux:minimal
docker_tag_trapmux = kkearne/trapmux
build_image = kkearne/trapmux_build
container_trapmux = trapmux
configuration_path_trapmux = /Users/kellskearney/go/src/trapmux/tools
docker_tag_clickhouse = kkearne/clickhouse
container_clickhouse = clickhouse
#configuration_path_clickhouse = /Users/kellskearney/go/src/trapmux/tools


# ----  Docker: trapmux  ----------------------------
.PHONY: trapmux
trapmux:
	DOCKER_BUILD=0 docker build -t $(build_image) -f Dockerfile .

trapmux_aws:
	DOCKER_BUILD=0 docker build -t $(docker_tag_trapmux) -f tools/docker/Dockerfile.amazonlinux .

run:
	#docker run --platform linux/amd64 --name trapmux_build $(build_image)
	docker run --name trapmux_build $(build_image)

stop:
	-docker stop trapmux_build
	-docker rm trapmux_build

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

