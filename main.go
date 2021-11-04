package main

import (
    "fmt"
    "log"
    "net/http"
		"time"
)

var VIDEO_STATE string = "UPLOADING"
var WAIT_SECOND time.Duration = 10

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Video state : %s", VIDEO_STATE)
}

func handlerAsAuthor(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "(As Author) Video state : %s", VIDEO_STATE)
}

func manageVideoState() {
	fmt.Println(VIDEO_STATE)
	time.Sleep(WAIT_SECOND * time.Second)
	VIDEO_STATE = "READY"
	fmt.Println(VIDEO_STATE)
}

func main() {
		go manageVideoState()
    http.HandleFunc("/", handler)
    http.HandleFunc("/author", handlerAsAuthor)
    log.Fatal(http.ListenAndServe(":2375", nil))
}