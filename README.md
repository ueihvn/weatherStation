# Weather Station for IOT subject:
NT532.L21 Internet Of Things

## overview
Get data from private channel on ThingSpeak using RESTAPI
Store data into InfluxDB
Send Gmail with custom threshold
Running with Docker
## using Makefile for run
### make run
Using Dockerfile for build image
### make start
run production.yml for start running project
### make clean 
for cleaning container down project

## Dockerfile
WORKDIR: set working directory for any RUN,CMD,..
using go install for build and put binary into $GOBIN

##writeDataInfluxDb
Using point to store data base on line protocol
### line protocol

humiditySensor,source= "BME280",location="Ho Chi Minh" temperature=30.1,humidity=62.33,pressure=1007.11,altitude=54.3 timestamp

airQualitySensor,source= "MQ135",location="Ho Chi Minh" avgPpm=6.2 timestamp

## production.yml
running 3 contaier
### influxdb
influxdb: for running influxdb contaier
### influxdb_cli
influxdb_cli: for running initial setup for influxdb
links: for share enviroment variables between influxdb and influxdb_cli
### server
Go application for get data ThingSpeak from private channel, write data ThingSpeak and Email when match condition
