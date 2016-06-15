package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// The response expected by the server
type Response struct {
	Result interface{} `json:"result"`
}

// Main loop of the game
func clientGame(user *User) {

	fmt.Printf("User %s starts a new game, good luck to him\n", user)

	// Wait until a user has failed to answer 5 times
	for user.Failure < 5 {

		// Create a new random question
		question := NewQuestion()
		reader := strings.NewReader(question.Content())
		url := user.Addr

		// Create a new client for the connection
		client := http.Client{
			Timeout: time.Duration(5 * time.Second),
		}
		resp, err := client.Post("http://"+url+"/answer", "application/json", reader)

		// The player has crashed or didn't answer in time, lose points and make a new attempt
		if err != nil {
			fmt.Printf("User %s failed to respond in time\n", user)
			user.Points -= 5
			user.Failure += 1
			continue
		}

		// Unmarshal the response from the player
		content, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		response := &Response{}
		err = json.Unmarshal(content, response)
		fmt.Println(string(content))

		// The answer is not formatted properly, lose points and make a new attempt
		if err != nil {
			fmt.Printf("User %s failed to format the response\n", user)
			user.Points -= 2
			continue
		}

		// Check if the response is correct
		fmt.Printf("User %s answered %v, expected %v\n", user, response.Result, question.Answer())
		if response.Result == question.Answer() {
			fmt.Printf("User %s answered successfully\n", user)
			user.Points += 1
		} else {
			fmt.Printf("User %s failed to provide the right answer\n", user)
			user.Points -= 1
		}
	}

	// The player has failed to much, delete him
	fmt.Printf("User %s stopped the game and reached %d points\n", user, user.Points)
	game.deletePlayer(user)

}
