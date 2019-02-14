// Package main provides ...
package main

import (
    "fmt"
    "net/http"
)

//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//w.WriteHeader(http.StatusOK)
//vars := mux.Vars(r) // "github.com/gorilla/mux"
//todoId := vars["todoId"]

func Test(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello\n")
}
