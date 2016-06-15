package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Struct sent on a register request
type Register struct {
	// Port the user listens to
	Port int
	// Name of the player
	Name string
}

// The response expected by the server
type Response struct {
	Result interface{} `json:"result"`
}

var (
	// Variable holding the name of the player
	name string
	// Variable holding the port to listen too
	listenPort int
	// Variable holding the url of the master server
	masterUrl string
)

// The cobra command for lauching the client
var rootCmd *cobra.Command = &cobra.Command{
	Use:   "client -p <port> -n <name> -s <server>",
	Short: "Runs a client to play our awesome spammer game",
	Run: func(cmp *cobra.Command, args []string) {
		subscribe()
		startServer()
	},
}

// Registers a new client to the game
func subscribe() {
	register := &Register{
		Port: listenPort,
		Name: name,
	}

	content, err := json.Marshal(register)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := strings.NewReader(string(content))
	resp, err := http.Post("http://"+masterUrl+"/register", "application/json", buff)
	if err != nil {
		fmt.Println("Subscribe post error:" + err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response body parse error:" + err.Error())
	}
	fmt.Println(string(respBody))
}

// A dummy handler that log the request and always return 42
func dummyHandler(resp http.ResponseWriter, req *http.Request) {
	// Do not answer too fast
	time.Sleep(time.Duration(2 * time.Second))

	// Log the request
	content, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	fmt.Printf("Received a new request: '%s'\n", content)

	// Send the response
	result := &Response{42}
	data, _ := json.Marshal(result)
	resp.Write(data)

}

// Start the server
func startServer() {
	http.HandleFunc("/answer", dummyHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}

// Bind the command flags to the global variable
func init() {
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "user", "Name of the player")
	rootCmd.PersistentFlags().StringVarP(&masterUrl, "server", "s", "127.0.0.1:8000", "Server hosting the game")
	rootCmd.PersistentFlags().IntVarP(&listenPort, "port", "p", 8001, "Port to listen to")
}

// Launch the game
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
