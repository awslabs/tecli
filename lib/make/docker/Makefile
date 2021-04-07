.PHONY: docker/build/ubuntu/lastest
docker/build/ubuntu/latest: docker/build/ubuntu

.PHONY: docker/build
docker/build: docker/delete
	rm -rf dist/ && mkdir -p dist/docker
	docker build -f Dockerfile.Golang . -t tecli-go
	docker create -ti --name tecli-go tecli-go
	docker cp tecli-go:/dist/tecli dist/docker/tecli
	docker container stop tecli-go
	docker container rm tecli-go
	docker build -f Dockerfile.Ubuntu -t tecli-ubuntu .
	docker build -f Dockerfile.Alpine -t tecli-alpine .

.PHONY: docker/delete
docker/delete:
	docker stop tecli-go || true && docker rm tecli-go || true

# don't forget to login: docker login
.PHONY: docker/publish/go
docker/publish/go:
	docker tag tecli-go ${DOCKERHUB_USERNAME}/tecli-go:$(TAG)
	docker push ${DOCKERHUB_USERNAME}/tecli-go:$(TAG)

.PHONY: docker/publish/alpine
docker/publish/alpine: docker/build
	docker tag tecli-alpine ${DOCKERHUB_USERNAME}/tecli-alpine:$(TAG)
	docker push ${DOCKERHUB_USERNAME}/tecli-alpine:$(TAG)

.PHONY: docker/publish/ubuntu
docker/publish/ubuntu:
	docker tag tecli-ubuntu ${DOCKERHUB_USERNAME}/tecli-ubuntu:$(TAG)
	docker push ${DOCKERHUB_USERNAME}/tecli-ubuntu:$(TAG)

.PHONY: docker/publish
docker/publish: docker/publish/go docker/publish/alpine docker/publish/ubuntu

.PHONY: docker/build/ubuntu
docker/build/ubuntu:
	docker build --target ubuntu -t tecli/ubuntu:latest .

.PHONY: docker/image/size
docker/image/size:
	docker image inspect $(IMAGE) --format='{{.Size}}' 

.PHONY:docker/cleanup/containers
docker/cleanup/containers:
	docker container stop $(shell docker container ls -aq)
	docker container rm $(shell docker container ls -aq)

.PHONY:docker/cleanup/images
docker/cleanup/images:
	docker image prune --all --force
	docker images --quiet | xargs docker rmi

.PHONY: docker/build/ubuntu/bionic
docker/build/ubuntu/bionic:
	docker build --target bionic -t tecli/ubuntu:bionic .