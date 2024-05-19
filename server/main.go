package main

import (
	_ "embed"
	"net/http"
)

//go:embed f.json
var data []byte

func main() {
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}),
	)

	http.ListenAndServe(":8000", m)
}
