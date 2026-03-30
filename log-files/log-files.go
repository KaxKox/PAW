package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Błąd pliku", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		record := []string{
			time.Now().Format(time.RFC3339),
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		}

		if err := writer.Write(record); err != nil {
			http.Error(w, "Błąd zapisu", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "zapisano w log.csv")
	})

	fmt.Println("Serwer działa na porcie :8080")
	http.ListenAndServe(":8080", nil)
}