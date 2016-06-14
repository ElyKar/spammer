package main

import (
	"net/http"
	"fmt"
	"strings"
	"encoding/json"
	"io/ioutil"
)

type Register struct {
	Port int
	Name string
}

var listenPort int = 8001
var masterUrl string = "http://127.0.0.1:8000"

func subscribe()  {
	register := &Register{
		Port: listenPort,
		Name: "anthony",
	}
	content, err := json.Marshal(register)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := strings.NewReader(string(content))
	resp, err := http.Post(masterUrl+"/register", "application/json", buff)
	if err != nil {
		fmt.Println("Subscribe post error:" + err.Error())
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response body parse error:" + err.Error())
	}
	fmt.Println(string(respBody))
}

func handleHttp(res http.ResponseWriter, req *http.Request){
	fmt.Printf("Method:%s\n", req.Method)
	fmt.Printf("Body:%s\n", req.Body)
}

func startServer(){
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}

func main(){
	subscribe()
	startServer()
}