package main

import (
	"fmt"
	"log"
	"time"
)

var lastEntryIdDb int

const DelayTime time.Duration = 15

func init() {
	// get last entry db

	lastEntryIdDb = 111
}

func main() {
	//getRequest(getURLString())
	for {
		tsLastData, err := getRequest(getURLString())
		if err != nil {
			log.Fatal(err)
		}
		lastEntryIdTs := tsLastData.EntryId
		if lastEntryIdTs != lastEntryIdDb {
			// update in db
			fmt.Println("update db, new entryId: ", lastEntryIdTs)
			lastEntryIdDb = lastEntryIdTs
		} else {
			fmt.Println("no update db")
		}
		time.Sleep(time.Second * DelayTime)
	}
}
