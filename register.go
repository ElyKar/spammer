package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "strings"
)

type User struct {
	Ip   string  
	Port int 
	Name string 
}

var users []User

func registerHandler(resp http.ResponseWriter, req *http.Request) {
	
    content, _ := ioutil.ReadAll(req.Body)

    var user *User = &User{}

    _ = json.Unmarshal(content, user)

    user.Ip = strings.Split(req.RemoteAddr, ":")[0]

    users = append(users, *user)

    fmt.Println(*user)
}