package main

import "github.com/gorilla/mux"
import "7_mysql_db/handlers"
import "7_mysql_db/common/datastore"
import "log"
import (
	"7_mysql_db/common"
	"net/http"
)

func main() {

	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := datastore.NewDatastore(datastore.MYSQL, "zrd:123456@tcp(127.0.0.1:3306)/gopherfacedb")
	//db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	//db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")

	defer db.Close()
	env := common.Env{DB: db}

	if err != nil {
		log.Print(err)
	}
	r := mux.NewRouter()
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	http.ListenAndServe("127.0.0.1:2315", r)

}
