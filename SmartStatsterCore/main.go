// SmartStatsterCore project main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"

	"os"

	"flag"

	"time"

	"github.com/valyala/fasthttp"
)

var addr = flag.String("addr", "127.0.0.1:8080",
	"TCP address to listen to for incoming connections")

var params Params = Params{"", ""}

func main() {
	ParseCsv()

	fmt.Println("test stream")
	fmt.Println("started")

	go worker(time.NewTicker(10000 * time.Millisecond))
	for {
	}

}

func worker(ticker *time.Ticker) {
	var lastUpdate = time.Now()
	var timeSince float64
	for range ticker.C {

		body, err := get(fmt.Sprintf(params.BaseUrl, params.Token) + "/getUpdates")
		fmt.Println(body)
		fmt.Println(err)

		var updates Updates
		errUnm := json.Unmarshal(body, &updates)
		if errUnm != nil {
			fmt.Println("error:", errUnm)
		}
		fmt.Println("%+v", updates)

		if len(updates.Result) > 0 {
			fmt.Println(updates.Result[0].Message.Text)
			body, err := get(fmt.Sprintf(params.BaseUrl, params.Token) +
				"/sendMessage?chat_id=570051893&text=" +
				fmt.Sprintf("%g", average()))
			fmt.Println(body)
			fmt.Println(err)
		}

		timeSince = time.Since(lastUpdate).Seconds()
		fmt.Println("Time to work ", fmt.Sprintf("%g", timeSince))
		lastUpdate = time.Now()
	}
}

func get(uri string) ([]byte, error) {
	statusCode, body, err := fasthttp.Get(nil, uri)
	if err != nil {
		return nil, err
	}
	fmt.Println(statusCode)
	return body, nil
}

func ParseCsv() {
	file, err := os.Open("params.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		if record[0] == "baseurl" {
			params.BaseUrl = record[1]
		}
		if record[0] == "token" {
			params.Token = record[1]
		}
	}
}
