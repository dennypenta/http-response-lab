package main

import (
	_ "embed"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"server/models"
)

//go:embed f.json
var data []byte

func main() {
	gob.Register(models.AutoGenerated{})
	m := http.NewServeMux()

	var mm []models.AutoGenerated
	if err := json.Unmarshal(data, &mm); err != nil {
		panic(err)
	}

	m.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}),
	)

	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gob.NewEncoder(w).Encode(mm)
	}),
	)

	http.ListenAndServe(":8000", m)
}
