docker_build:
	docker image build --rm -t neversitup -f _cicd/build/dockers/Dockerfile .
	
docker_up:

	docker-compose up -d api

docker_down:
	docker-compose down --volumes
	docker container prune -f
	docker volume prune -f
