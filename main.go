package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/LaMonF/FDJ_SLACK/parser"
	l "github.com/LaMonF/FDJ_SLACK/log"
)

var win = []int{7, 14, 22, 28, 42}
var luckWin = 5

var slackURL = os.Getenv("SLACK_HOOK_URL")
var localhost = "http://localhost:8888"

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/lotoResult", lotoResult) // set router
	err := http.ListenAndServe(":9090", nil)   // set listen port
	if err != nil {
		l.Err("ListenAndServe: ", err)
	}
}

func lotoResult(w http.ResponseWriter, r *http.Request) {
	//Leaving this for dev purpose, we will remove it later
	// =================
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Fprintf(w, "Hello astaxie!") // send data to client side
	// =================

	if slackURL == "" {
		slackURL = localhost
	}

	// sendImmediateResponseToSlack(w, "Fetching latest loto result")

	p := parser.NewParser()
	data := p.FetchData()
	result := p.ParseData(data)
	for index, result := range result {
		if index == 0 { // only first result
			l.Info(result)
			//We can improve this post by using the URL from the POST request
			//See (https://api.slack.com/slash-commands -> Sending delayed responses)
			sendResponseToSlack(w, result.String())
			if result.IsWinning(win, luckWin) {
				postToSlack(slackURL, "ON A GAGNÉ !!!")
			}
		}
	}
}

func sendImmediateResponseToSlack(w http.ResponseWriter, post string) {
	l.Info("Sending back immediate response")
	// fmt.Fprintf(w, post) // send data to client side

	message := `{"text" : "` + post + `"}`
	l.Info(message)

	var jsonStr = []byte(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(jsonStr)
}

func sendResponseToSlack(w http.ResponseWriter, post string) {
	l.Info("Sending back immediate response")
	// fmt.Fprintf(w, post) // send data to client side

	//We set the response_type to in_channel (everyone can see it) instead of ephemeral (only you) by default
	message := `{"response_type": "in_channel","text" : "` + post + `"}`
	l.Info(message)

	var jsonStr = []byte(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(jsonStr)
}

func postToSlack(slackURL string, post string) {
	message := `{"text" : "` + post + `"}`

	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//TODO: Manage response, handle errors
	l.Info("response Body:", string(body))
}
