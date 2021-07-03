package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	lastEntryIdDb  int = 1
	influxDbClient influxdb2.Client
	wg             sync.WaitGroup
	influxToken    string
	influxUrl      string = "http://influxdb:8086"
	influxOrg      string
	influxBuc      string
)

const (
	DelayTime          time.Duration = 15
	DelayTimeWriteData time.Duration = 5
)

func init() {
	loadEnv()

	influxToken = getEnvByKey("DOCKER_INFLUXDB_TOKEN")
	influxOrg = getEnvByKey("DOCKER_INFLUXDB_INIT_ORG")
	influxBuc = getEnvByKey("DOCKER_INFLUXDB_INIT_BUCKET")

	// delay for db setting
	time.Sleep(time.Second * DelayTime)

	influxDbClient = influxdb2.NewClient(influxUrl, influxToken)
}

func main() {
	wg.Add(1)

	// go routine for getting data thingspeak
	go func() {
		for {
			tsLastData, err := getRequest(getURLString())
			if err != nil {
				log.Fatal(err)
			}
			lastEntryIdTs := tsLastData.EntryId
			if lastEntryIdTs != lastEntryIdDb || lastEntryIdTs == 1 {
				// update in db
				fmt.Println("update db, new entryId: ", lastEntryIdTs)

				//goroutine for sending TS point to channel
				go func() {
					sendHumidityPointToChannel(tsLastData)
					sendAirQualityPointToChannel(tsLastData.AirQuality)
				}()

				//check and send mail
				go sendMailProcess(tsLastData)

				lastEntryIdDb = lastEntryIdTs
			} else {
				fmt.Println("no update db")
			}
			time.Sleep(time.Second * DelayTime)
		}
	}()

	// go routine for writing points from channel to influxdb
	go writeDataInfluxDb()

	wg.Wait()
	influxDbClient.Close()
}
