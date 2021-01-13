// SmartStatsterCore project main.go
package main

import (
	"encoding/json"
	"fmt"

	"flag"

	"time"
)

var addr = flag.String("addr", "127.0.0.1:8080",
	"TCP address to listen to for incoming connections")

var params Params = parseCsvParams()

var lastq map[int][2]string

func main() {

	lastq = make(map[int][2]string)

	fmt.Println("started")

	go workerBotResp(time.NewTicker(10000 * time.Millisecond))

	select {}
}

func workerBotResp(ticker *time.Ticker) {
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

		msgLen := len(updates.Result)
		for i := 0; i < msgLen; i++ {
			if convint(lastq[updates.Result[i].Message.Chat.Id][0]) > 0 {
				if convint(lastq[updates.Result[i].Message.Chat.Id][0]) < updates.Result[i].Message.MessageId {
					currq := [2]string{"", ""}
					currq[0] = fmt.Sprintf("%d", updates.Result[i].Message.MessageId)
					currq[1] = updates.Result[i].Message.Text
					lastq[updates.Result[i].Message.Chat.Id] = currq
				} else {
					exchcurrq := lastq[updates.Result[i].Message.Chat.Id]
					currq := [2]string{exchcurrq[0], ""}
					lastq[updates.Result[i].Message.Chat.Id] = currq
				}
			} else {
				currq := [2]string{"", ""}
				currq[0] = fmt.Sprintf("%d", updates.Result[i].Message.MessageId)
				currq[1] = updates.Result[i].Message.Text
				lastq[updates.Result[i].Message.Chat.Id] = currq
			}
		}

		for k, v := range lastq {
			if string(v[1]) != "" {
				body, err := get(fmt.Sprintf(params.BaseUrl, params.Token) +
					fmt.Sprintf("/sendMessage?chat_id=%d", k) +
					"&text=" + answer(string(v[1])))
				fmt.Println(body)
				fmt.Println(err)
			} else {
				fmt.Println("answered")
			}

		}
		timeSince = time.Since(lastUpdate).Seconds()
		fmt.Println("Time to work ", fmt.Sprintf("%g", timeSince))
		lastUpdate = time.Now()
	}
}
