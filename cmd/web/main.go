package main

import (
    "log"
    "net/http"
)

// A handler function named 'home'
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from newsletter"))
}



func main() {

    // mux is our router (multiplexer)
    mux := http.NewServeMux()

    // the route pattern/endpoint/URL path
    mux.HandleFunc("/", home)

    
    // start a local web server
    log.Print("starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)

}
