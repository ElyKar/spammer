package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Operation struct {
	Type     string  `json:"type"`
	OperandA float64 `json:"opA"`
	OperandB float64 `json:"opB"`
}

var operations [4]string

func NewRandomOp() *Operation {
	a := rand.Float64() + 1
	b := rand.Float64() + 1
	op := operations[rand.Int()%4]
	return &Operation{
		Type:     op,
		OperandA: a,
		OperandB: b,
	}
}

func (o *Operation) AddOne() {
	o.OperandA += 1
}

func init() {
	operations = [4]string{"ADD", "SUB", "MUL", "DIV"}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	operation := NewRandomOp()
	encoded, err := json.Marshal(operation)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.Write([]byte(req.RemoteAddr + "\n"))
	resp.Write(encoded)
}

func main() {
	http.HandleFunc("/question", handler)
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8000", nil)
}
