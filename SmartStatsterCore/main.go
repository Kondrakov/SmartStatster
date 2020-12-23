// SmartStatsterCore project main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

const (
	pi float64 = 3.1415
)

var addr = flag.String("addr", "127.0.0.1:8080",
	"TCP address to listen to for incoming connections")

//var bottoken string
//var params map = make(map[string]string)
var params Params = Params{"", ""}

func main() {

	ParseCsv()

	flag.Parse()

	s := fasthttp.Server{
		Handler: actionHandler,
	}
	errFH := s.ListenAndServe(*addr)
	if errFH != nil {
		log.Fatalf("error in ListenAndServe: %s", errFH)
	}
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
		fmt.Println(record)
	}
}

func actionHandler(ctx *fasthttp.RequestCtx) {
	var jsonBlob = []byte("")

	var action string = string(ctx.QueryArgs().Peek("action"))
	var value string = string(ctx.QueryArgs().Peek("value"))
	var chatid string = string(ctx.QueryArgs().Peek("chatid"))

	fmt.Println("one")
	fmt.Println(value)
	fmt.Println(chatid)
	fmt.Println("one::" + action)

	var actionNext string = fmt.Sprintf(params.BaseUrl, params.Token)
	switch action {
	case "getupdates":
		actionNext += "/getUpdates"
	case "sendmessage":
		if value != "" {
			actionNext += "/sendMessage?chat_id=570051893&text=Greetings"
		}
	case "":
		actionNext = ""
	}
	if actionNext != "" {
		resp, err := http.Get(actionNext)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		for true {
			bs := make([]byte, 1014)
			n, err := resp.Body.Read(bs)
			jsonBlob = []byte(string(bs[:n]))
			fmt.Println(string(bs[:n]))

			if n == 0 || err != nil {
				break
			}
		}
		fmt.Println(actionNext)
		if actionNext == "?action=getupdates" {
			var updates Updates
			errUnm := json.Unmarshal(jsonBlob, &updates)
			if errUnm != nil {
				fmt.Println("error:", errUnm)
			}
			fmt.Println("%+v", updates)

			if len(updates.Result) > 0 {
			}
			for i := 0; i < len(updates.Result); i++ {
				var responseReactionWords = [3]string{"Hello", "Hi", "Привет"}
				rand.Shuffle(len(responseReactionWords), func(ind, jnd int) {
					responseReactionWords[ind], responseReactionWords[jnd] = responseReactionWords[jnd], responseReactionWords[ind]
				})
				for j := 0; j < len(responseReactionWords); j++ {
					fmt.Println("-- " + responseReactionWords[j] + " -- " + updates.Result[i].Message.Text)
					if responseReactionWords[j] == updates.Result[i].Message.Text {
						fmt.Println("%+v", updates.Result[i].Message.Text)
						actionNext = "?action=sendmessage" + "&value=" + responseReactionWords[0] + "&chatid=" + strconv.Itoa(updates.Result[i].Message.Chat.Id)

						fmt.Println(updates.Result[i].Message.Chat.Id)

						resp, err := http.Get("http://127.0.0.1:8085/bot_get" + actionNext)
						if err != nil {
							fmt.Println(err)
							return
						}
						defer resp.Body.Close()
					}
				}
			}
		}
	}
}
