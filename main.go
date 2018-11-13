package main

import (
    "html/template"
    "log"
    "net/http"
)

type Page struct {
    Title    string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

    p := Page{"Titre de ma page"}

    t := template.New("Label de ma template")

    t = template.Must(t.ParseFiles("tmpl/layout.tmpl"))

    err := t.ExecuteTemplate(w, "layout", p)

    if err != nil {
        log.Fatalf("Template execution: %s", err)
    }
}

func main() {
    
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":9090", nil)
}