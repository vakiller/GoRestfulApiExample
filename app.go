// app.go
// create by Aiden 03/04/2021

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbName string) {
	fmt.Printf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

}

func (a *App) Run(addr string) {}
