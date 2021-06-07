package main

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	influxToken string = "l8oHiNT7pvFBkw8g6qw4zWla5G8Ox7exI7oqSGpZJqqRk102Ad5lO1sJWvFaDFTulriKtC-E_tyMx9QMAJxnqQ=="
	influxUrl   string = "http://localhost:8086"
)

func connect() {
	client := influxdb2.NewClient(influxUrl, influxToken)
	fmt.Println(client)
}
