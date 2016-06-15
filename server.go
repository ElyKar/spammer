package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Register handler for a new connection
func registerHandler(resp http.ResponseWriter, req *http.Request) {

	content, _ := ioutil.ReadAll(req.Body)
	_ = req.Body.Close()
	var user *User = &User{}

	_ = json.Unmarshal(content, user)
	user.Ip = strings.Split(req.RemoteAddr, ":")[0]
	user.Addr = fmt.Sprintf("%s:%d", user.Ip, user.Port)
	game.addPlayer(user)

	fmt.Printf("Added another user: %v\n", *user)
	go clientGame(user)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8000", nil)
}
