package main

import (
	"html/template"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	time := t.Format("02.01.2006, 15:04:05 Uhr")

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	data := make(map[string]string)
	data["pageTitle"] = "Jenkins X Test"
	data["paragraph1"] = "This is a test of Jenkins X."
	data["paragraph2"] = "And it seems to be working."
	data["time"] = time

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
