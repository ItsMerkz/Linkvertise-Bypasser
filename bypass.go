package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type bypassresponse struct {
	BypassedUrl string `json:"destination"`
}

func main() {
	fmt.Print("URL To Bypass > ")
	var bypassurl string
	fmt.Scanln(&bypassurl)
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://bypass.bot.nu/bypass2?url=%v", bypassurl), nil)

	client := http.Client{}
	client.Timeout = time.Second * 10
	resp, err := client.Do(req)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Bypassing")
	fmt.Println("")
	fmt.Println("Getting Link...")
	var x bypassresponse
	err = json.Unmarshal(b, &x)
	if err != nil {
		fmt.Printf("Error %v", err)
	}
	fmt.Printf("Bypass Url | %v", x.BypassedUrl)

}
