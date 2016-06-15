package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var port int

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "spammer -p <port>",
	Short: "Start the spammer server",
	Run: func(cmd *cobra.Command, args []string) {
		port := fmt.Sprintf(":%d", port)
		fmt.Printf("Server listening on port %s\n", port)

		http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
			resp.Write([]byte("Hello idiot"))
		})
		http.HandleFunc("/register", registerHandler)
		http.ListenAndServe(port, nil)
	},
}

// Register handler for a new connection
func registerHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Received a new request")

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

func init() {
	// Bind the port flag
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8000, "The port to listen for register requests")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
