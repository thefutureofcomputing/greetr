package main

import (
    "fmt"
    "log"
    "net/http"
    "flag"
)

var bindAddr = flag.String("bind-addr", ":8080", "Address/port to bind listening port to")

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s! I love you.", r.URL.Path[1:])
}

func main() {
	flag.Parse()
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(*bindAddr, nil))
}