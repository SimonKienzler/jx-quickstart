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
	time := t.Format("02.01.2006, 15:04:05")

	data := make(map[string]string)
	data["pageTitle"] = "Jenkins X Test (Dev-Branch)"
	data["paragraph1"] = `This is a test of Jenkins X using devpods. They provide 
	a simple way of developing a cloud based application.`
	data["paragraph2"] = "And it is most definetely working. 😃"
	data["time"] = time

	tmpl, err := template.New("index").Parse(`{{define "I"}} 
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
			<title>{{ .pageTitle }}</title>
			
			<link href="https://fonts.googleapis.com/css?family=Montserrat:400,700" rel="stylesheet">
			<style>
				body {
					font-family: 'Montserrat',sans-serif;
					background: #00b0f0;
					background-attachment: fixed;
					background-image: linear-gradient(-45deg, #00b0f0, #003c7e);
					background-position: top left;
					background-size: 400%;
				}

				h1 {
					font-weight: 700;
					color: #ffffff;
					margin: 0px 0px 30px 0px;
				}

				p {
					font-weight: 300;
					color: #ffffff;
				}
	
				.box {
					margin: 100px auto;
					width: 400px;
					padding: 30px;
					border: 1px solid rgba(255,255,255,0.4);
					border-radius: 5px;
					background: rgba(255,255,255,0.1);
					box-shadow: 0px 0px 10px rgba(55,55,55,0.7);
				}

				hr {
					color: #ffffff;
					margin: 10px 0px;
					border-style: solid;
					border-width: .5px;
				}
	
				.time {
					font-size: 80%;
					color: #ffffff;
					text-align: center;
				}
			</style>
		</head>
		<body>
			<div class="box">
				<h1>{{ .pageTitle }}</h1>
		
				<p>{{ .paragraph1 }}</p>
		
				<p>{{ .paragraph2 }}</p>

				<hr>
		
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
