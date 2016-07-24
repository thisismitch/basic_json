package main

import (
	"encoding/json"
	"net/http"
)

type Path struct {
	Path string
	Authors []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	body := Path{r.URL.Path[1:], []string{"manicas", "bsinclair"}}

	response, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
