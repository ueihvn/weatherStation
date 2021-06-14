build:
	docker build . -f Dockerfile -t weather-station-alpine:1
start:
	docker-compose -f production.yml --env-file ./dev.env up -d
clean:
	docker-compose -f production.yml down
