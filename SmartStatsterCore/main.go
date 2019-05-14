// SmartStatsterCore project main.go
package main

//"strconv"
import (
	"encoding/json"
	"fmt"
	"net/http"

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
		"<button id=\"get_request\">GET request</button>" +
		"<input type=\"password\" id=\"action\"/>" +
		"<script type=\"text/javascript\">" +
		"var getRequest = document.getElementById(\"get_request\");" +
		"var getMethod = (event) => {" +
		"var request = new XMLHttpRequest();" +
		"request.open(\"GET\", \"http://127.0.0.1:8080/bot_get?action=getupdates\", true);" +
		"request.setRequestHeader(\"Content-type\", \"application/x-www-form-urlencoded\");" +
		"var params = \"action=getupdates\";" +
		"request.send(params);" +
		"console.log(\"hhhh\");};" +
		"getRequest.addEventListener(\"click\", getMethod);" +
		"</script></body></html>")

	fmt.Println("Query string is %q\n", string(ctx.QueryArgs().Peek("action")))

	//var d bool = "getupdates" == string(ctx.QueryArgs().Peek("action"))
	//fmt.Println(strconv.FormatBool(d))

	actionHandler(string(ctx.QueryArgs().Peek("action")))
}

func actionHandler(action string) {

	var jsonBlob = []byte("")

	var actionNext string = ""
	switch action {
	case "getupdates":
		fmt.Println("one")
		actionNext = "?action=getupdates"
	case "sendmessage":
		actionNext = "?action=sendmessage"
	}
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

	var updates Updates
	errUnm := json.Unmarshal(jsonBlob, &updates)
	if errUnm != nil {
		fmt.Println("error:", errUnm)
	}
	fmt.Println("%+v", updates)
}
