package main

import (
    "fmt"
    "log"
    "net/http"
		"time"
)

var WAIT_SECOND time.Duration = 10

var video Video

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, video.GetDownloadMessage("random"))
}

func handlerAsAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, video.GetDownloadMessage("Léo"))
}

func manageVideoState() {
	fmt.Println(video.currentState)
	time.Sleep(WAIT_SECOND * time.Second)
	ready := &Ready{
		video: &video,
	}
	video.setState(ready)
	fmt.Println(video.currentState)
}

func main() {
	video.name = "Le Bitcoin n'apporte rien de nouveau, change my mind"
	video.author = "Léo"

	uploading := &Uploading{
		video: &video,
	}
	video.setState(uploading)

	go manageVideoState()
  http.HandleFunc("/", handler)
  http.HandleFunc("/iamleo", handlerAsAuthor)
  log.Fatal(http.ListenAndServe(":2375", nil))
}