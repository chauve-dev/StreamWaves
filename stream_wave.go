package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var spawns = []string{}
var afflictions = []string{}
var characters = []string{}

func addMonster(w http.ResponseWriter, r *http.Request) {
	var monster = r.URL.Path[len("/add/"):]
	spawns = append(spawns, monster)
	fmt.Printf("Added %s to queue\n", monster)
}
func fetchMonsters(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetched queue next wave soon\n")
	io.WriteString(w, strings.Join(spawns[:], "|"))
	spawns = []string{}
}

func addAffliction(w http.ResponseWriter, r *http.Request) {
	var affliction = r.URL.Path[len("/add_affliction/"):]
	afflictions = append(afflictions, affliction)
	fmt.Printf("Added %s to queue\n", affliction)
}

func fetchAffliction(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetched afflictions next afflictions wave soon\n")
	io.WriteString(w, strings.Join(afflictions[:], "|"))
	afflictions = []string{}
}

func addCharacter(w http.ResponseWriter, r *http.Request) {
	var character = r.URL.Path[len("/add_character/"):]
	characters = append(characters, character)
	fmt.Printf("Added %s to queue\n", character)
}

func fetchCharacter(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetched characters next characters wave soon\n")
	io.WriteString(w, strings.Join(characters[:], "|"))
	characters = []string{}
}

func main() {
	http.HandleFunc("/add/", addMonster)
	http.HandleFunc("/fetch", fetchMonsters)

	http.HandleFunc("/add_affliction/", addAffliction)
	http.HandleFunc("/fetch_affliction", fetchAffliction)

	http.HandleFunc("/add_character/", addCharacter)
	http.HandleFunc("/fetch_character", fetchCharacter)

	fmt.Printf("Starting server on port 51525\n")
	err := http.ListenAndServe(":51525", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
