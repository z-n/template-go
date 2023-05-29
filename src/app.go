package main

import (
	"html/template"
    "net/http"
	"os"
)

func main() {
	port := "3000"
	if(os.Getenv("PORT") != "") {
		port = os.Getenv("PORT")
	}
	tmpl := template.Must(template.ParseFiles("src/templates/index.html"))

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, "")
    })

    fs := http.FileServer(http.Dir("src/static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":" + port, nil)
}