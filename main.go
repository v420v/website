package main

import (
	"net/http"
	"os"
	"text/template"
)

func mainHandler(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func main() {
	dir, _ := os.Getwd() // 追記
	http.HandleFunc("/", mainHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/")))) //追記
	http.ListenAndServe(":8000", nil)
}
