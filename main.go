package main

import (
    "fmt"
    "log"
    "net/http"
		"time"
)

var VIDEO_STATE string

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Video state : %s", VIDEO_STATE)
}

func manageVideoState() {
	VIDEO_STATE = "UPLOADING"
	fmt.Println(VIDEO_STATE)
	time.Sleep(10 * time.Second)
	VIDEO_STATE = "READY"
	fmt.Println(VIDEO_STATE)
}

func main() {
		go manageVideoState()
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":2375", nil))
}