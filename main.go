package main

import (
	"bookstore/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models.NewDB("postgres://bookstore:bookstore@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	http.Handle("/books", booksIndex(env))
	http.Handle("/books/show", booksShow(env))
	http.Handle("/books/create", booksCreate(env))
	http.ListenAndServe(":3000", nil)
}

func booksIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := models.AllBooks(env.db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, bk := range bks {
			fmt.Fprintf(w, "%s, %s, %s, %.2f rub.\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
		}
	})
}

func booksShow(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		isbn := r.FormValue("isbn")
		if isbn == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		bk, err := models.BookShow(env.db, &isbn)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintf(w, "%s, %s, %s, %.2f rub.\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	})
}

func booksCreate(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")

		if isbn == "" || title == "" || author == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		price, err := strconv.ParseFloat(r.FormValue("price"), 32)

		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		bk := models.Book{
			Isbn:   isbn,
			Title:  title,
			Author: author,
			Price:  float32(price),
		}

		rowsAffected, err := models.CreateBook(env.db, &bk)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
	})
}
