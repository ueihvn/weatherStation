package main

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
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
	pointsCh <- hs
}

func sendAirQualityPointToChannel(avgPpm float32) {
	//air quality sensor
	aqs := influxdb2.NewPointWithMeasurement("airQualitySensor").
		AddTag("source", "MQ135").
		AddTag("location", "Ho Chi Minh").
		AddField("avgPpm", avgPpm).
		SetTime(time.Now())
	pointsCh <- aqs
}
