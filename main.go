package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	t := time.Now()
	time := t.Format("02.01.2006, 15:04:05 Uhr")

	data := make(map[string]string)
	data["pageTitle"] = "Jenkins X Test"
	data["paragraph1"] = "This is a test of Jenkins X."
	data["paragraph2"] = "And it seems to be working."
	data["time"] = time

	tmpl, err := template.New("index").Parse(`{{define "I"}} 
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>{{ .pageTitle }}</title>
	
			<style>
				body {
					font-family: sans-serif;
				}
	
				.box {
					margin: 60px auto;
					width: 400px;
					padding: 30px;
					border: 1px solid #666666;
					border-radius: 5px;
				}
	
				.time {
					font-size: 80%;
					color: #666666;
				}
			</style>
		</head>
		<body>
			<div class="box">
				<h1>{{ .pageTitle }}</h1>
		
				<p>{{ .paragraph1 }}</p>
		
				<p>{{ .paragraph2 }}</p>
		
				<p class="time">{{ .time }}</p>
			</div>
		</body>
	</html>
	{{end}}`)
	err = tmpl.ExecuteTemplate(w, "I", data)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
