.PHONY: docker/build/ubuntu/lastest
docker/build/ubuntu/latest: docker/build/ubuntu

.PHONY: docker/build
docker/build: docker/delete
	rm -rf docker/ dist/ && mkdir -p docker
	docker build -f Dockerfile.Golang . -t tecli-go
	docker create -ti --name tecli-go tecli-go
	docker cp tecli-go:/dist/tecli docker/tecli
	docker container stop tecli-go
	docker container rm tecli-go
	docker build -f Dockerfile.Ubuntu.20.04 -t tecli-ubuntu:20.04 .
	# docker cp tecli:/dist/tecli tecli-linux-amd64 

.PHONY: docker/delete
docker/delete:
	docker stop tecli-go || true && docker rm tecli-go || true

# docker tag tecli-ubuntu:20.04 valtersilvaau/tecli-ubuntu:20.04
# docker push tecli-ubuntu:20.04 valtersilvaau/tecli-ubuntu:20.04

# don't forget to login: docker login
.PHONY: docker/publish
docker/publish:
	docker tag tecli-go valtersilvaau/tecli-go:latest
	docker push valtersilvaau/tecli-go:latest

.PHONY: docker/build/ubuntu
docker/build/ubuntu:
	docker build --target ubuntu -t tecli/ubuntu:latest .

.PHONY: docker/build/ubuntu/bionic
docker/build/ubuntu/bionic:
	docker build --target bionic -t tecli/ubuntu:bionic .