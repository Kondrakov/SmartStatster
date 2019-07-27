// SmartStatsterCore project main.go
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
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

func main() {

	flag.Parse()

	s := fasthttp.Server{
		Handler: handler,
	}
	errFH := s.ListenAndServe(*addr)
	if errFH != nil {
		log.Fatalf("error in ListenAndServe: %s", errFH)
	}
}

func handler(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html;charset=utf-8")

	ctx.WriteString("<html><head></head><body>" +
		"<button id=\"get_updates\">GET updates</button>" +
		"<br/>" +
		"<button id=\"get_sendmessage\">GET sendmessage</button>" +
		"<input id=\"message_input\"/>" +
		"<input id=\"chat_id_input\"/>" +

		"<script type=\"text/javascript\">" +
		"var getUpdates = document.getElementById(\"get_updates\");" +
		"var getUpdatesMethod = (event) => {" +
		"var request = new XMLHttpRequest();" +
		"request.open(\"GET\", \"/?action=getupdates\", true);" +
		"request.setRequestHeader(\"Content-type\", \"application/x-www-form-urlencoded\");" +
		//"var params = \"action=getupdates\";" +
		//"request.send(params);" +
		"request.send();" +
		"};" +
		"getUpdates.addEventListener(\"click\", getUpdatesMethod);" +

		"var getSendmessage = document.getElementById(\"get_sendmessage\");" +
		"var bodySendmessage = document.getElementById(\"message_input\");" +
		"var chatIDInput = document.getElementById(\"chat_id_input\");" +
		"var getSendmessageMethod = (event) => {" +
		"var request = new XMLHttpRequest();" +
		"request.open(\"GET\", \"/?action=sendmessage&value=\" + bodySendmessage.value + \"&chatid=\" + chatIDInput.value, true);" +
		"request.setRequestHeader(\"Content-type\", \"application/x-www-form-urlencoded\");" +
		"request.send();" +
		"};" +
		"getSendmessage.addEventListener(\"click\", getSendmessageMethod);" +
		"</script></body></html>")

	fmt.Println("Query string is %q\n", string(ctx.QueryArgs().Peek("action")))
	fmt.Println(ctx.QueryArgs())

	actionHandler(ctx)
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

	var actionNext string = ""
	switch action {
	case "getupdates":
		actionNext = "?action=getupdates"
	case "sendmessage":
		if value != "" {
			actionNext = "?action=sendmessage" + "&value=" + value + "&chatid=" + chatid
		}
	case "":
		actionNext = ""
	}
	if actionNext != "" {
		resp, err := http.Get("http://127.0.0.1:8085/bot_get" + actionNext)
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
