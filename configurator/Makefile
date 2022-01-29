
docker_tag=configurator
container=configurator
PORT=5555

build: Dockerfile
	docker build -t $(docker_tag) .

prep:
	git clone https://github.com/ginkgobioworks/react-json-schema-form-builder.git

run: stop build
	docker run -d --rm  --name $(container) -v app/src/data:/app/src/data -p $(PORT):3000 $(docker_tag)

stop:
	-docker stop $(container)
	-docker rm $(container)

health:
	curl -k http://localhost:$(PORT)/health && echo

version:
	curl -k http://localhost:$(PORT)/version && echo

json:
	python -m json.tool app/src/schema.json
