package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var db *sql.DB

func main() {
	var err error
	connStr := "postgres://user:password@localhost:5432/library?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS books (id SERIAL PRIMARY KEY, title TEXT, author TEXT)")

	http.HandleFunc("GET /books", getBooks)
	http.HandleFunc("POST /books", createBook)
	http.HandleFunc("DELETE /books/{id}", deleteBook)

	fmt.Println("Serwer dziala na :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT id, title, author FROM books")
	var books []Book
	for rows.Next() {
		var b Book
		rows.Scan(&b.ID, &b.Title, &b.Author)
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	json.NewDecoder(r.Body).Decode(&b)
	err := db.QueryRow("INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id", b.Title, b.Author).Scan(&b.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(b)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	db.Exec("DELETE FROM books WHERE id = $1", id)
	w.WriteHeader(http.StatusNoContent)
}