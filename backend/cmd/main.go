package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func getItems(w http.ResponseWriter, r *http.Request) {
	items := []Item{
		{ID: 1, Name: "Docker", Img: "https://static-00.iconduck.com/assets.00/docker-icon-2048x2048-5mc7mvtn.png"},
		{ID: 2, Name: "Nginx", Img: "https://www.svgrepo.com/show/373924/nginx.svg"},
		{ID: 3, Name: "GitHub", Img: "https://cdn-icons-png.flaticon.com/512/25/25231.png"},
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(items)
	if err != nil {
		fmt.Fprintf(w, "err json Encode")
	}
}

func main() {
	http.HandleFunc("/items", getItems)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
