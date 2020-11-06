package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/slowmanchan/moneyapi/models"
)

type Env struct {
	db *sqlx.DB
}

func main() {
	db, err := sqlx.Open("postgres", "dbname=gomoney sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db: db}

	router := httprouter.New()

	router.GET("/", env.Index)

	fmt.Println("Now listening on localhost:4100")
	log.Fatal(http.ListenAndServe(":4100", router))
}

func (e *Env) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rawBankStatements, err := models.AllRawBankStatements(e.db)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	data, err := json.Marshal(rawBankStatements)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(data)
	fmt.Println("Index")
}
