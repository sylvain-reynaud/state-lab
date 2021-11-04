package main

import (
    "fmt"
    "log"
    "net/http"
		"time"
)

var VIDEO_STATE string = "UPLOADING"
var WAIT_SECOND time.Duration = 10


var video Video

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Video : %s", video)
}

func handlerAsAuthor(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "(As Author) Video : %s", video)
}

func manageVideoState() {
	fmt.Println(VIDEO_STATE)
	time.Sleep(WAIT_SECOND * time.Second)
	VIDEO_STATE = "READY"
	fmt.Println(VIDEO_STATE)
}

func main() {
	video.name = "Le Bitcoin n'apporte rien de nouveau, change my mind"
	video.author = "Léo"

	go manageVideoState()
  http.HandleFunc("/", handler)
  http.HandleFunc("/iamleo", handlerAsAuthor)
  log.Fatal(http.ListenAndServe(":2375", nil))
}

/*
 TODOs :
 - Add video struct : name, author
	*/