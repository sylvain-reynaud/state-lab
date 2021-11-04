package main

import (
    "fmt"
    "log"
    "net/http"
		"time"
)

var VIDEO_STATE string = "UPLOADING"
var WAIT_SECOND int = 10

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Video state : %s", VIDEO_STATE)
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
    log.Fatal(http.ListenAndServe(":2375", nil))
}