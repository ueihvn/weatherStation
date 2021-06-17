build:
	sudo docker build . -f Dockerfile -t weather-station-alpine:1
start:
	sudo docker-compose -f production.yml --env-file ./dev.env up -d
clean:
	sudo docker-compose -f production.yml down
