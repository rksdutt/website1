package main

import (
    "html/template"
    "net/http"
    "log"
)

type PageVariables struct {
    Title string
}

func main() {
    http.HandleFunc("/", HomePage)

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    pageVariables := PageVariables{
        Title: "My Dynamic Go Website",
    }

    t, err := template.ParseFiles("static/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, pageVariables)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
