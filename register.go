package main

import "fmt"

// Global state of the game
var game *Game = NewGame()

// Struct representing a player
type User struct {
	// Ip of the player
	Ip string
	// Port the player listens to
	Port int
	// Combination of the address and port
	Addr string
	// Name of the player
	Name string
	// Number of points the player has
	Points int
	// Number of failures the player has done
	Failure int
}

// String representation of a user
func (u *User) String() string {
	return fmt.Sprintf("%s [%s]", u.Name, u.Addr)
}

// Represent the state of the game
type Game struct {
	players map[string]*User
}

// Creates a new empty game
func NewGame() *Game {
	return &Game{make(map[string]*User)}
}

// Add a player to the game (if not exist already
func (g *Game) addPlayer(u *User) {
	if _, ok := g.players[u.Addr]; !ok {
		g.players[u.Addr] = u
	}
}

// Remove a player from the game
func (g *Game) deletePlayer(u *User) {
	if _, ok := g.players[u.Addr]; ok {
		delete(g.players, u.Addr)
	}
}
