init:
	echo "Init"

build-docker-image:
	docker build -t go-stateservice -f Dockerfile .

push:
	docker tag go-stateservice gcr.io/golden-tenure-231009/go-stateservice:2.0
	docker push gcr.io/golden-tenure-231009/go-stateservice:2.0

all: init  build-docker-image push
