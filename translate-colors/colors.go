package main

import (
	"encoding/json"
	"net/http"
)

type Color struct {
	Names map[string]string
	Hex   string
}

var colors = []Color{
	{Names: map[string]string{"pl": "czerwony", "en": "red", "de": "rot"}, Hex: "#ff0000"},
	{Names: map[string]string{"pl": "zielony", "en": "green", "de": "grun"}, Hex: "#00ff00"},
	{Names: map[string]string{"pl": "niebieski", "en": "blue", "de": "blau"}, Hex: "#0000ff"},
	{Names: map[string]string{"pl": "zolty", "en": "yellow", "de": "gelb"}, Hex: "#ffff00"},
	{Names: map[string]string{"pl": "czarny", "en": "black", "de": "schwarz"}, Hex: "#000000"},
}

var langs = map[string]bool{"pl": true, "en": true, "de": true}

func colorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	lng := r.URL.Query().Get("lng")

	if name == "" || lng == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "brak parametru 'name' lub 'lng'"})
		return
	}

	if !langs[lng] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "brak języka: " + lng})
		return
	}

	for _, c := range colors {
		for _, n := range c.Names {
			if n == name {
				json.NewEncoder(w).Encode(map[string]string{
					"color": name,
					"lng":   lng,
					"name":  c.Names[lng],
					"value": c.Hex,
				})
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "nie znaleziono koloru: " + name})
}

func main() {
	http.HandleFunc("/color", colorHandler)
	http.ListenAndServe(":8080", nil)
}
