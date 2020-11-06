package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/slowmanchan/moneyapi/models"
)

// App holds all major dependancies for the app such
// as http router and the db connection
type App struct {
	router *httprouter.Router
	db     *sqlx.DB
}

// New instantiates an app with the appropriate major
// app dependencies
func New() *App {
	db, err := sqlx.Open("postgres", "dbname=gomoney sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return &App{
		db:     db,
		router: httprouter.New(),
	}
}

// Start will start the app and listen on a port. It will
// have all the http routes
func (a *App) Start() {
	a.router.GET("/", a.Index)
	fmt.Println("Now listening on localhost:4100")
	log.Fatal(http.ListenAndServe(":4100", a.router))
}

// Index is the handler for the index route of the api
func (a *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rawBankStatements, err := models.AllRawBankStatements(a.db)
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
