// Package main provides ...
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

type Route struct {
    Name            string
    Pattern         string
    Method          string
    HandlerFunc     http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Test",
        "/test",
        "GET",
        Test,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}

func main() {
    ip := GetIp()
    help := fmt.Sprintf("Simple Server running on %s%s", ip, config.Port)
    fmt.Println(help)
    router := NewRouter()
    log.Fatal(http.ListenAndServe(config.Port, router))
}
