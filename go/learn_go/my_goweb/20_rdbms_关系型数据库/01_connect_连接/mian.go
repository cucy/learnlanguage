package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "zr2d:123456@test")
	check(err)

	defer db.Close()
	check(err)
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}
func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	check(err)
}
func check(e error) {
	if e != nil {
		fmt.Println(err)
	}

}
