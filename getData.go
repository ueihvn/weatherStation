package main

import (
	"context"
	"fmt"
	"log"
)

func getAirQualityData() {

	queryAPI := influxDbClient.QueryAPI(influxOrg)

	result, err := queryAPI.Query(context.Background(), `from(bucket:"nhom3-bucker")|> range(start: -10h) |> filter(fn: (r) => r._measurement == "airQualitySensor")`)

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}

		fmt.Printf("row: %s\n", result.Record().String())

		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	}
}
