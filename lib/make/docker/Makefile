.PHONY: docker/build/ubuntu/lastest
docker/build/ubuntu/latest: docker/build/ubuntu

.PHONY: docker/build/ubuntu
docker/build/ubuntu:
	docker build --target ubuntu -t tecli/ubuntu:latest .

.PHONY: docker/build/ubuntu/bionic
docker/build/ubuntu/bionic:
	docker build --target bionic -t tecli/ubuntu:bionic .