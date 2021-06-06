package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	apiKey    string = "FBKWQ3H7RBLR9GDZ"
	channelId string = "1377314"
)

type TSLastEntry struct {
	CreatedAt  string  `json:"created_at"`
	EntryId    int     `json:"entry_id"`
	Tempature  float32 `json:"field1,string"`
	Humidity   float32 `json:"field2,string"`
	Pressure   float32 `json:"field3,string"`
	Altitude   float32 `json:"field4,string"`
	AirQuality float32 `json:"field5,string"`
}

func getURLString() string {
	strUrl := "http://api.thingspeak.com/channels/" + channelId + "/feeds/last.json"

	urlThingSpeak, err := url.Parse(strUrl)
	if err != nil {
		log.Fatal(err)
	}

	urlValues := urlThingSpeak.Query()
	urlValues.Add("api_key", apiKey)

	urlThingSpeak.RawQuery = urlValues.Encode()

	fmt.Println(urlThingSpeak.String())

	return urlThingSpeak.String()
}

func decodeTsJSON(data []byte) (tsLastEntry TSLastEntry, err error) {
	err = json.Unmarshal(data, &tsLastEntry)
	if err != nil {
		return
	}

	return
}

func getRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	ts, err := decodeTsJSON(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ts)
}
