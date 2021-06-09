package main

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

const (
	influxToken string = "l8oHiNT7pvFBkw8g6qw4zWla5G8Ox7exI7oqSGpZJqqRk102Ad5lO1sJWvFaDFTulriKtC-E_tyMx9QMAJxnqQ=="
	influxUrl   string = "http://localhost:8086"
	influxOrg   string = "nhom3"
	influxBuc   string = "nhom3-bucker"
)

var pointsCh = make(chan *write.Point, 200)

func writeDataInfluxDb() {
	writeApi := influxDbClient.WriteAPI(influxOrg, influxBuc)
	maxFailValue := 0
	for {
		point, ok := <-pointsCh
		if !ok {
			maxFailValue++
			if maxFailValue > 3 {
				time.Sleep(time.Second * DelayTimeWriteData)
			}
			continue
		} else {
			maxFailValue = 0
			writeApi.WritePoint(point)
		}
	}

}

func sendHumidityPointToChannel(tsLastEntry TSLastEntry) {
	//huminity sensor
	hs := influxdb2.NewPointWithMeasurement("humiditySensor").
		AddTag("source", "BME280").
		AddTag("location", "Ho Chi Minh").
		AddField("temperature", tsLastEntry.Tempature).
		AddField("humidity", tsLastEntry.Humidity).
		AddField("pressure", tsLastEntry.Pressure).
		AddField("altitude", tsLastEntry.Altitude).
		SortFields().SetTime(time.Now())
	fmt.Println(hs)
	pointsCh <- hs
}

func sendAirQualityPointToChannel(avgPpm float32) {
	//air quality sensor
	aqs := influxdb2.NewPointWithMeasurement("airQualitySensor").
		AddTag("source", "MQ135").
		AddTag("location", "Ho Chi Minh").
		AddField("avgPpm", avgPpm).
		SetTime(time.Now())
	fmt.Println(aqs)
	pointsCh <- aqs
}
