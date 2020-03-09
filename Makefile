.DEFAULT_GOAL := docker_run

client_dev:
	cd client && yarn run dev

server_local:
	go run ./server

docker_build:
	docker build -t talang-land .

docker_run:
	docker run -p 9876:9876 talang-land
