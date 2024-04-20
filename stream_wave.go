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

func addMonster(w http.ResponseWriter, r *http.Request) {
	var monster = r.URL.Path[len("/add/"):]
	spawns = append(spawns, monster)
	fmt.Printf("Added %s to queue\n", monster)
}
func fetchMonsters(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetched queue next wave soon\n")
	io.WriteString(w, strings.Join(spawns[:], "|"))
	// empty the queue
	spawns = []string{}
}

func main() {
	http.HandleFunc("/add/", addMonster)
	http.HandleFunc("/fetch", fetchMonsters)

	fmt.Printf("Starting server on port 51525\n")
	err := http.ListenAndServe(":51525", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
