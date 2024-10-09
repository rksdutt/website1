package main

import (
    "net/http"
    "log"
)

func main() {
    // Serve the React app
    fs := http.FileServer(http.Dir("../frontend/build"))
    http.Handle("/", fs)

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
