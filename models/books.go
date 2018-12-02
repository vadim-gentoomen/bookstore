package models

import (
	"database/sql"
)

type Book struct {
	Isbn   string `sql:"primary_key"`
	Title  string
	Author string
	Price  float32
}

func AllBooks(db *sql.DB) ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func BookShow(db *sql.DB, isbn *string) (*Book, error) {

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", *isbn)

	bk := new(Book)
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return nil, err
	}

	return bk, nil
}

func CreateBook(db *sql.DB, book *Book) (int64, error) {
	result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
