package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jipachec/grb"
	"github.com/jipachec/grb/database"
	"github.com/jipachec/grb/server"
	"fmt" //
)

func main() {
	fmt.Println("hello world") //
	db := database.New()
	db.AddEntry(grb.NewFakeEntry())
	db.AddEntry(grb.NewFakeEntry())
	db.AddEntry(grb.NewFakeEntry())

	myServer := server.New(db)

	router := mux.NewRouter()
	router.NewRoute().
		Path("/").
		HandlerFunc(myServer.EntryList).
		Methods("GET").
		Name("EntryList")

	router.NewRoute().
		Path("/{key}").
		HandlerFunc(myServer.EntryGet).
		Methods("GET").
		Name("EntryGet")

	http.ListenAndServe(":8000", router) // http://localhost:8000/
}
