package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/Todo", TodoList)
	router.HandleFunc("/Todo/{Id}", TodoDetail)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index rout handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - We had modified our code!  %q", html.EscapeString(r.URL.Path))
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo List")
}

func TodoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	fmt.Fprintf(w, "The requested ID is %q", id)
}
