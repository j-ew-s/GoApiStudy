package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow!")
}

func TodoList(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Request Project"},
		Todo{Name: "Develop project"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	fmt.Fprintf(w, "The requested ID is : %q", id)
}
