# Weather Station for IOT subject:

NT532.L21 Internet Of Things

## Overeview

Get data from private channel on ThingSpeak using RESTAPI
Store data into InfluxDB
Send Gmail with custom threshold, email body subject
Running with Docker

## Using Makefile for run

### make run

Using Dockerfile for build image

### make start

run production.yml for start running project

### make clean

for cleaning container down project

## Dockerfile

WORKDIR: set working directory for any RUN,CMD,..
using go install for build and put binary into $GOBIN

## writeDataInfluxDb

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

## dev.env

DOCKER_INFLUXDB_INIT_USERNAME=user_name_influxdb
DOCKER_INFLUXDB_INIT_PASSWORD=password_influxdb
DOCKER_INFLUXDB_INIT_ORG=org_influxdb
DOCKER_INFLUXDB_INIT_BUCKET=bucket_influxdb
DOCKER_INFLUXDB_TOKEN=token_influxdb
GMAIL_USERNAME=gmail_user_send
GMAIL_PASSWORD=gmail_password_send
