// SmartStatsterCore project main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	type Update struct {
		UpdateId int `json:"update_id"`
		Message  struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id           int    `json:"id"`
				IsBot        bool   "is_bot"
				FirstName    string "first_name"
				LastName     string "last_name"
				LanguageCode string "language_code"
			} `json:"from"`
			Chat struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date int    `json:"date"`
			Text string `json:"text"`
		} `json:"message"`
	}
	type Updates struct {
		Ok     bool `json:"ok"`
		Result []Update
	}

	var jsonBlob = []byte("")

	resp, err := http.Get("")
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
