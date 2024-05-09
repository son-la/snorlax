DOCKER_IMAGE_NAME:=sonlax/snorlax
VERSION:=0.0.1

.PHONY: build
build:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .


.PHONY: push
push:
	docker push -t $(DOCKER_IMAGE_NAME):$(VERSION)

.PHONY: run
run:
	docker run --rm $(DOCKER_IMAGE_NAME):$(VERSION) $(ARG)
